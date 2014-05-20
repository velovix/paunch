package scene

import (
	"encoding/json"
	"errors"
	"io"
)

type Loader interface {
	LoadScene(map[string]interface{}) bool
}

type Saver interface {
	SaveScene() map[string]interface{}
}

type Scene struct {
	loaders []Loader
	savers  []Saver
}

func (scene *Scene) SetLoadSceners(loaders []Loader) {

	scene.loaders = loaders
}

func (scene *Scene) SetSaveSceners(savers []Saver) {

	scene.savers = savers
}

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
