package jsonparser

import (
	"encoding/json"
	"io/ioutil"
)

func ParseFromJsonFile(filePath string, target interface{}) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, target)
	return err
}
