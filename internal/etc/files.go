package etc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func ReadFile(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	return string(bytes), err
}

func WriteFile(path string, data interface{}, encoder func(interface{}) ([]byte, error)) error {
	dataBytes, err := encoder(data)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, dataBytes, os.ModePerm)
}

var StringEncoder = func(input interface{}) ([]byte, error) {
	return []byte(fmt.Sprintf("%s", input)), nil
}

var JsonEncoder = func(input interface{}) ([]byte, error) {
	return json.Marshal(input)
}
