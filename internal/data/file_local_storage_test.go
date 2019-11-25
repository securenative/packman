package data

import (
	"github.com/securenative/packman/internal/etc"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestFileLocalStorage_Init_NewFile(t *testing.T) {
	filePath := filepath.Join(os.TempDir(), "packman_local_storage_test.json")
	key := "somekey"
	value := "somevalue"

	defer os.Remove(filePath)
	s, err := NewFileLocalStorage(filePath)
	assert.NotNil(t, s)
	assert.Nil(t, err)
	assert.False(t, etc.FileExists(filePath))

	_, err = s.Get(key)
	assert.NotNil(t, err)

	err = s.Put(key, value)
	assert.Nil(t, err)
	assert.FileExists(t, filePath)

	storedValue, err := s.Get(key)
	assert.Nil(t, err)
	assert.EqualValues(t, value, storedValue)

	_, err = s.Get(key + "123")
	assert.NotNil(t, err)
}

func TestFileLocalStorage_Init_ExistingValidFile(t *testing.T) {
	filePath := filepath.Join(os.TempDir(), "packman_local_storage_test_valid.json")
	key := "somekey"
	value := "somevalue"

	defer os.Remove(filePath)
	err := etc.WriteFile(filePath, map[string]interface{}{key: value}, etc.JsonEncoder)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	s, err := NewFileLocalStorage(filePath)
	assert.NotNil(t, s)
	assert.Nil(t, err)
	assert.True(t, etc.FileExists(filePath))

	storedValue, err := s.Get(key)
	assert.Nil(t, err)
	assert.EqualValues(t, value, storedValue)

	newKey := "new key"
	newValue := "new value"
	err = s.Put(newKey, newValue)
	assert.Nil(t, err)

	storedValue, err = s.Get(newKey)
	assert.Nil(t, err)
	assert.EqualValues(t, newValue, storedValue)
}

func TestFileLocalStorage_Init_ExistingInvalidFile(t *testing.T) {
	filePath := filepath.Join(os.TempDir(), "packman_local_storage_test_invalid.json")

	defer os.Remove(filePath)
	err := etc.WriteFile(filePath, "not json", etc.StringEncoder)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	s, err := NewFileLocalStorage(filePath)
	assert.Nil(t, s)
	assert.NotNil(t, err)
	assert.True(t, etc.FileExists(filePath))
}
