package scene

import (
	"encoding/json"
	"errors"
	"io"
)

type LoadScener interface {
	LoadScene(map[string]interface{}) bool
}

type Loader struct {
	objects []LoadScener
}

func (loader *Loader) SetLoadSceners(objects []LoadScener) {

	loader.objects = objects
}

func (loader *Loader) Load(r io.Reader) error {

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
	for _, obj := range loader.objects {
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
	if usedCnt < len(loader.objects) {
		return errors.New("some objects did not have corresponding data")
	}

	return nil
}
