package paunch

import (
	"errors"
	al "github.com/vova616/go-openal/openal"
)

// Audio is an object that deals with Paunch's sound system, as a whole. An
// Audio object must be created before any other sound operations can take
// place.
type Audio struct {
	device  *al.Device
	context *al.Context
}

// NewAudio creates a new Audio object. This must be done before any other
// audio operations happen.
func NewAudio() (Audio, error) {

	var audio Audio

	audio.device = al.OpenDevice("")
	if audio.device == nil {
		return audio, errors.New("failed to open device")
	}

	audio.context = audio.device.CreateContext()
	if ok := audio.context.Activate(); !ok {
		return audio, errors.New("failed to make context current")
	}

	return audio, nil
}

// Destroy cleans up the Audio object. After this method, audio opperations
// will no longer work until NewAudio is called again. This should be done
// before the program exits.
func (audio Audio) Destroy() {
	audio.device.CloseDevice()
	audio.context.Destroy()
}
