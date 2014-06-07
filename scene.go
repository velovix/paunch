package paunch

import (
	"encoding/json"
	"fmt"
	"io"
)

// Loader is an interface for objects that are able to respond to data from a
// Scene object. An implementation of LoadScene is expected to return true if
// the data it recieved should no longer be sent to other Loader objects.
type Loader interface {
	LoadScene(map[string]interface{}) bool
}

// Saver is an interface for objects that are able to supply data for a Scene
// object.
type Saver interface {
	SaveScene() map[string]interface{}
}

// SceneEncoder creates a JSON encoded file containing header information and
// the data given to the SceneEncoder by its Saver objects.
type SceneEncoder struct {
	encoder *json.Encoder
	header  map[string]interface{}
	savers  []Saver
}

// SceneDecoder reads from a JSON file as created by a SceneEncoder object.
type SceneDecoder struct {
	decoder         *json.Decoder
	hasHeaderLoaded bool
	loaders         []Loader
}

// NewSceneEncoder creates a new SceneEncoder object that will save to the
// given io.Writer.
func NewSceneEncoder(w io.Writer) SceneEncoder {

	var sceneEncoder SceneEncoder

	sceneEncoder.encoder = json.NewEncoder(w)
	sceneEncoder.header = make(map[string]interface{})
	sceneEncoder.savers = make([]Saver, 0)

	return sceneEncoder
}

// NewSceneDecoder creates a new SceneDecoder object that will read from the
// given io.Reader.
func NewSceneDecoder(r io.Reader) SceneDecoder {

	var sceneDecoder SceneDecoder

	sceneDecoder.decoder = json.NewDecoder(r)
	sceneDecoder.loaders = make([]Loader, 0)

	return sceneDecoder
}

// SetLoaders sets the list of Loader objects that will be prompted during
// a call to Load.
func (sceneDecoder *SceneDecoder) SetLoaders(loaders []Loader) {

	sceneDecoder.loaders = loaders
}

// SetSavers sets the list of Saver objects that will be prompted during
// a call to Save.
func (sceneEncoder *SceneEncoder) SetSavers(savers []Saver) {

	sceneEncoder.savers = savers
}

// GetHeader returns the header information from the file.
func (sceneDecoder *SceneDecoder) GetHeader() (map[string]interface{}, error) {

	defer func() { sceneDecoder.hasHeaderLoaded = true }()

	info := make(map[string]bool)
	err := sceneDecoder.decoder.Decode(&info)
	if err != nil {
		return make(map[string]interface{}), err
	}

	if info["hasHeader"] {
		header := make(map[string]interface{})
		err = sceneDecoder.decoder.Decode(&header)
		if err != nil {
			return make(map[string]interface{}), err
		}
		return header, nil
	}

	return make(map[string]interface{}), nil
}

// Load reads the information from the given JSON file and gives it to the
// Loader objects.
func (sceneDecoder *SceneDecoder) Load() error {

	if !sceneDecoder.hasHeaderLoaded {
		_, err := sceneDecoder.GetHeader()
		if err != nil {
			return err
		}
	}

	data := make([]map[string]interface{}, 0)

	for i := 0; ; i++ {
		data = append(data, make(map[string]interface{}))
		err := sceneDecoder.decoder.Decode(&data[i])
		if err != nil && err.Error() == "EOF" {
			data = data[:len(data)-1]
			break
		} else if err != nil {
			return err
		}
	}

	used := make([]bool, len(data))
	for _, obj := range sceneDecoder.loaders {
		for j, val := range data {
			if !used[j] && obj.LoadScene(val) {
				used[j] = true
				break
			}
		}
	}

	usedCnt := 0
	for _, val := range used {
		if val {
			usedCnt++
		}
	}
	if usedCnt < len(sceneDecoder.loaders) {
		return fmt.Errorf("some objects did not have corresponding data: %v used, %v total", usedCnt, len(sceneDecoder.loaders))
	}

	return nil
}

// SetHeader sets the header information of the file.
func (sceneEncoder *SceneEncoder) SetHeader(data map[string]interface{}) {

	sceneEncoder.header = data
}

// Save reads the information to the given location using data recieved from
// the Saver objects.
func (sceneEncoder *SceneEncoder) Save() error {

	info := make(map[string]bool)
	info["hasHeader"] = len(sceneEncoder.header) > 0
	err := sceneEncoder.encoder.Encode(info)
	if err != nil {
		return err
	}

	if info["hasHeader"] {
		err = sceneEncoder.encoder.Encode(sceneEncoder.header)
		if err != nil {
			return err
		}
	}

	for _, val := range sceneEncoder.savers {
		err = sceneEncoder.encoder.Encode(val.SaveScene())
		if err != nil {
			return err
		}
	}

	return nil
}
