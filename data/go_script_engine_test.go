package data

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func setup() {
	path := tempPath()

	status0 := `
	package main
	
	import "os"
	
	func main() {
		os.Exit(0)
	}
	`

	status1 := `
	package main
	
	import "os"
	
	func main() {
		os.Exit(1)
	}
	`
	err := ioutil.WriteFile(filepath.Join(path, string(os.PathSeparator), "status_0.go"), []byte(status0), os.ModePerm)
	if err != nil {
		panic(err.Error())
	}

	err = ioutil.WriteFile(filepath.Join(path, string(os.PathSeparator), "status_1.go"), []byte(status1), os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
}

func tempPath() string {
	path := fmt.Sprintf("%s%c%s", os.TempDir(), os.PathSeparator, "packtest")
	return path
}

func TestGoScriptEngine_Run(t *testing.T) {
	setup()
	runner := NewGoScriptEngine()

	err := runner.Run(filepath.Join(tempPath(), string(os.PathSeparator), "status_0.go"), []string{})
	assert.Nil(t, err)

	err = runner.Run(filepath.Join(tempPath(), string(os.PathSeparator), "status_1.go"), []string{})
	assert.NotNil(t, err)
}
