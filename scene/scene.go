package scene

import (
	"encoding/json"
	"errors"
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

// Scene is an object that manages the saving of game states on an
// object-to-object basis into a JSON file.
type Scene struct {
	loaders []Loader
	savers  []Saver
}

// SetLoadSceners sets the list of Loader objects that will be prompted during
// a call to Load.
func (scene *Scene) SetLoadSceners(loaders []Loader) {

	scene.loaders = loaders
}

// SetSaveSceners sets the list of Saver objects that will be prompted during
// a call to Save.
func (scene *Scene) SetSaveSceners(savers []Saver) {

	scene.savers = savers
}

// Load reads the information from the given JSON file and gives it to the
// Loader objects.
func (scene *Scene) Load(r io.Reader) error {

	decoder := json.NewDecoder(r)

	data := make([]map[string]interface{}, 0)

	for i := 0; ; i++ {
		data = append(data, make(map[string]interface{}))
		err := decoder.Decode(&data[i])
		if err != nil && err.Error() == "EOF" {
			data = data[:len(data)-1]
			break
		} else if err != nil {
			panic(err)
		}
	}

	used := make([]bool, len(data))
	for _, obj := range scene.loaders {
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
	if usedCnt < len(scene.loaders) {
		return errors.New("some objects did not have corresponding data")
	}

	return nil
}

// Save reads the information to the given location using data recieved from
// the Saver objects.
func (scene *Scene) Save(w io.Writer) error {

	encoder := json.NewEncoder(w)

	for _, val := range scene.savers {
		err := encoder.Encode(val.SaveScene())
		if err != nil {
			return err
		}
	}

	return nil
}
