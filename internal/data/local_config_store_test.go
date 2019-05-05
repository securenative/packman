package data

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var configStore *LocalConfigStore

func TestMain(m *testing.M) {

	path := fmt.Sprintf("%s%c%s", os.TempDir(), os.PathSeparator, "packtest")

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}

	configStore = NewLocalConfigStore(path)

	status := m.Run()

	err = os.RemoveAll(path)
	if err != nil {
		panic(err)
	}

	os.Exit(status)
}

func TestLocalConfigStore_IntegrationTest(t *testing.T) {
	const key = "mykey"
	type data struct {
		Name string
		Age  int
	}

	var nilData *data
	notFound := configStore.Get(key, nilData)
	assert.Nil(t, nilData)
	assert.False(t, notFound)

	obj := &data{Name: "matan", Age: 31}
	err := configStore.Put(key, obj)
	assert.Nil(t, err)

	var objData data
	found := configStore.Get(key, &objData)
	assert.True(t, found)
	assert.NotNil(t, obj)
	assert.EqualValues(t, *obj, objData)
}
