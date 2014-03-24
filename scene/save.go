package scene

import (
	"encoding/json"
	"io"
)

type SaveScener interface {
	SaveScene() map[string]interface{}
}

type Saver struct {
	objects []SaveScener
}

func (saver *Saver) SetSaveSceners(objects []SaveScener) {

	saver.objects = objects
}

func (saver *Saver) Save(w io.Writer) error {

	encoder := json.NewEncoder(w)

	for _, val := range saver.objects {
		err := encoder.Encode(val.SaveScene())
		if err != nil {
			return err
		}
	}

	return nil
}
