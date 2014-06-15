package paunch

import (
	"bytes"
	"encoding/binary"
	"errors"
	vorbis "github.com/velovix/vorbis"
	al "github.com/vova616/go-openal/openal"
	"io/ioutil"
	"os"
	"strings"
)

// Sound represents a playable sound clip.
type Sound struct {
	source *al.Source
	buffer *al.Buffer

	SampleRate  int32
	NumChannels int32
}

type riffHeaderObj struct {
	chunkID   []byte
	chunkSize int32
	format    []byte
}

type fmtHeaderObj struct {
	chunkID       []byte
	chunkSize     int32
	audioFormat   int16
	numChannels   int16
	sampleRate    int32
	byteRate      int32
	blockAlign    int16
	bitsPerSample int16
}

type dataHeaderObj struct {
	chunkID   []byte
	chunkSize int32
}

type soundInfo struct {
	sampleRate  int32
	bitRate     int16
	numChannels int16
	data8       []byte
	data16      []int16
}

func loadOGG(filename string) (soundInfo, error) {

	file, err := os.Open(filename)
	if err != nil {
		return soundInfo{}, err
	}

	data, readErr := ioutil.ReadAll(file)
	if readErr != nil {
		return soundInfo{}, readErr
	}

	sound, numChannels, decodeErr := vorbis.Decode(data)
	if decodeErr != nil {
		return soundInfo{}, decodeErr
	}
	return soundInfo{sampleRate: 44100, bitRate: 16, numChannels: int16(numChannels), data16: sound}, nil
}

func loadWAV(filename string) (soundInfo, error) {

	file, err := os.Open(filename)
	if err != nil {
		return soundInfo{}, err
	}
	defer file.Close()

	var length int

	var riffHeader riffHeaderObj
	riffHeaderBytes := make([]byte, 12)
	length, err = file.Read(riffHeaderBytes)
	if err != nil {
		return soundInfo{}, err
	}
	if length != len(riffHeaderBytes) {
		return soundInfo{}, errors.New("unexpected end of file")
	}
	riffHeader.chunkID = riffHeaderBytes[0:4]
	if string(riffHeader.chunkID) != "RIFF" {
		return soundInfo{}, errors.New("no RIFF header ID in first WAV chunk")
	}
	stream := bytes.NewBuffer(riffHeaderBytes[4:8])
	binary.Read(stream, binary.LittleEndian, &riffHeader.chunkSize)
	riffHeader.format = riffHeaderBytes[8:12]
	if string(riffHeader.format) != "WAVE" {
		return soundInfo{}, errors.New("no WAVE header ID in first WAV chunk")
	}

	var fmtHeader fmtHeaderObj
	fmtHeaderBytes := make([]byte, 24)
	length, err = file.ReadAt(fmtHeaderBytes, int64(len(riffHeaderBytes)))
	if err != nil {
		return soundInfo{}, err
	}
	if length != len(fmtHeaderBytes) {
		return soundInfo{}, errors.New("unexpected end of file")
	}
	fmtHeader.chunkID = fmtHeaderBytes[0:4]
	if string(fmtHeader.chunkID) != "fmt " {
		return soundInfo{}, errors.New("no fmt header ID in second WAV chunk")
	}
	stream = bytes.NewBuffer(fmtHeaderBytes[4:8])
	binary.Read(stream, binary.LittleEndian, &fmtHeader.chunkSize)
	stream = bytes.NewBuffer(fmtHeaderBytes[8:10])
	binary.Read(stream, binary.LittleEndian, &fmtHeader.audioFormat)
	stream = bytes.NewBuffer(fmtHeaderBytes[10:12])
	binary.Read(stream, binary.LittleEndian, &fmtHeader.numChannels)
	stream = bytes.NewBuffer(fmtHeaderBytes[12:16])
	binary.Read(stream, binary.LittleEndian, &fmtHeader.sampleRate)
	stream = bytes.NewBuffer(fmtHeaderBytes[16:20])
	binary.Read(stream, binary.LittleEndian, &fmtHeader.byteRate)
	stream = bytes.NewBuffer(fmtHeaderBytes[20:22])
	binary.Read(stream, binary.LittleEndian, &fmtHeader.blockAlign)
	stream = bytes.NewBuffer(fmtHeaderBytes[22:24])
	binary.Read(stream, binary.LittleEndian, &fmtHeader.bitsPerSample)

	var dataHeader dataHeaderObj
	dataHeaderBytes := make([]byte, 8)
	length, err = file.ReadAt(dataHeaderBytes, int64(len(riffHeaderBytes)+int(8+fmtHeader.chunkSize)))
	if err != nil {
		return soundInfo{}, err
	}
	if length != len(dataHeaderBytes) {
		return soundInfo{}, errors.New("unexpected end of file")
	}
	dataHeader.chunkID = dataHeaderBytes[0:4]
	if string(dataHeader.chunkID) != "data" {
		return soundInfo{}, errors.New("no data header ID in third WAV chunk")
	}
	stream = bytes.NewBuffer(dataHeaderBytes[4:8])
	binary.Read(stream, binary.LittleEndian, &dataHeader.chunkSize)

	if fmtHeader.bitsPerSample == 8 {
		data := make([]byte, dataHeader.chunkSize)
		length, err = file.ReadAt(data, int64(len(riffHeaderBytes)+int(8+fmtHeader.chunkSize)+len(dataHeaderBytes)))
		if length != len(data) {
			return soundInfo{}, errors.New("unexpected end of file")
		}

		return soundInfo{sampleRate: fmtHeader.sampleRate, bitRate: fmtHeader.bitsPerSample, numChannels: fmtHeader.numChannels, data8: data}, nil
	} else if fmtHeader.bitsPerSample == 16 {
		data := make([]int16, dataHeader.chunkSize/2)
		dataBytes := make([]byte, dataHeader.chunkSize)
		length, err = file.ReadAt(dataBytes, int64(len(riffHeaderBytes)+int(8+fmtHeader.chunkSize)+len(dataHeaderBytes)))
		if length != len(dataBytes) {
			return soundInfo{}, errors.New("unexpected end of file")
		}

		for i := 0; i < len(dataBytes); i += 2 {
			stream = bytes.NewBuffer(dataBytes[i : i+2])
			binary.Read(stream, binary.LittleEndian, &data[i/2])
		}
		return soundInfo{sampleRate: fmtHeader.sampleRate, bitRate: fmtHeader.bitsPerSample, numChannels: fmtHeader.numChannels, data16: data}, nil
	}

	return soundInfo{}, errors.New("invalid bitrate")
}

// NewSound returns a new Sound object based on the provided file. As of right
// now, NewSound only supports standard WAV files.
func NewSound(filename string) (*Sound, error) {

	sound := &Sound{}

	tmpSource := al.NewSource()
	sound.source = &tmpSource
	tmpBuffer := al.NewBuffer()
	sound.buffer = &tmpBuffer

	var err error
	var info soundInfo
	split := strings.Split(filename, ".")
	if split[len(split)-1] == "wav" {
		info, err = loadWAV(filename)
	} else if split[len(split)-1] == "ogg" {
		info, err = loadOGG(filename)
	}
	if err != nil {
		return sound, err
	}

	if info.bitRate == 8 {
		if info.numChannels == 1 {
			sound.buffer.SetData(al.FormatMono8, info.data8, info.sampleRate)
		} else if info.numChannels == 2 {
			sound.buffer.SetData(al.FormatStereo8, info.data8, info.sampleRate)
		} else {
			return sound, errors.New("invalid number of channels")
		}
	} else if info.bitRate == 16 {
		if info.numChannels == 1 {
			sound.buffer.SetDataInt(al.FormatMono16, info.data16, info.sampleRate)
		} else if info.numChannels == 2 {
			sound.buffer.SetDataInt(al.FormatStereo16, info.data16, info.sampleRate)
		} else {
			return sound, errors.New("invalid number of channels")
		}
	}

	sound.source.Seti(al.AlBuffer, int32(*sound.buffer))

	return sound, nil
}

// Play plays the Sound object from it's current time. If the Sound has never
// been played, that time will be at the beginning.
func (sound *Sound) Play() {

	sound.source.Play()
}

// Pause pauses the Sound object at it's current time. The Play method will
// resume from where the Pause method left off.
func (sound *Sound) Pause() {

	sound.source.Pause()
}

// Stop pauses the Sound object and brings it's time to the beginning.
func (sound *Sound) Stop() {

	sound.source.Stop()
}

// SetLoop sets whether or not the Sound object should loop at the end of it's
// sample. The default value is false.
func (sound *Sound) SetLoop(willLoop bool) {

	sound.source.SetLooping(willLoop)
}

// GetPlaying returns a SoundState value reflecting the playing state of the
// Sound object.
func (sound *Sound) GetPlaying() SoundState {

	return SoundState(sound.source.State())
}

// GetGain returns the gain (volume) of the Sound object.
func (sound *Sound) GetGain() float32 {

	return sound.source.GetGain()
}

// SetGain sets the gain (volume) of the Sound object. 1 is the default value.
func (sound *Sound) SetGain(gain float32) {

	sound.source.SetGain(gain)
}

// Destroy cleans up the Sound object, which will no longer be playable. This
// should be done after the Sound object is no longer needed.
func (sound *Sound) Destroy() {

	tmpSource := *sound.source
	tmpBuffer := *sound.buffer
	al.DeleteSource(tmpSource)
	al.DeleteBuffer(tmpBuffer)
	sound.source = &tmpSource
	sound.buffer = &tmpBuffer
}
