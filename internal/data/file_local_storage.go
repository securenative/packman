package data

import (
	"encoding/json"
	"fmt"
	"github.com/securenative/packman/internal/etc"
)

type fileLocalStorage struct {
	storage  map[string]string
	filePath string
}

func NewFileLocalStorage(filePath string) (LocalStorage, error) {
	m, err := file2Map(filePath)
	if err != nil {
		return nil, err
	}
	return &fileLocalStorage{filePath: filePath, storage: m}, nil
}

func (this *fileLocalStorage) Put(key, value string) error {
	this.storage[key] = value
	return this.flush()
}

func (this *fileLocalStorage) Get(key string) (string, error) {
	value, ok := this.storage[key]
	if !ok {
		return "", fmt.Errorf("failed to find key: %s in local storage", key)
	}
	return value, nil
}

func (this *fileLocalStorage) flush() error {
	err := etc.WriteFile(this.filePath, this.storage, etc.JsonEncoder)
	return err
}

func file2Map(filePath string) (map[string]string, error) {
	if !etc.FileExists(filePath) {
		return make(map[string]string), nil
	}

	content, err := etc.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var out map[string]string
	err = json.Unmarshal([]byte(content), &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
