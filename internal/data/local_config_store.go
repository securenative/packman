package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type LocalConfigStore struct {
	dataPath string
}

func NewLocalConfigStore(dataPath string) *LocalConfigStore {
	return &LocalConfigStore{dataPath: dataPath}
}

func (this *LocalConfigStore) Put(key string, value interface{}) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	path := this.key2path(key)
	return ioutil.WriteFile(path, bytes, os.ModePerm)
}

func (this *LocalConfigStore) Get(key string, valueOut interface{}) bool {
	bytes, err := ioutil.ReadFile(this.key2path(key))
	if err != nil {
		return false
	}

	return json.Unmarshal(bytes, valueOut) == nil
}

func (this *LocalConfigStore) key2path(key string) string {
	path := fmt.Sprintf(this.dataPath, os.PathSeparator, key)
	return path
}
