package data

import (
	"github.com/securenative/packman/internal/etc"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestGenericScriptEngine_Run(t *testing.T) {
	filePath := filepath.Join(os.TempDir(), "packman_script_engine.go")
	defer os.Remove(filePath)
	err := etc.WriteFile(filePath, script, etc.StringEncoder)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	flags := map[string]string{
		"flag1": "flag1",
		"flag2": "flag2",
		"flag3": "flag3",
	}

	s := NewGenericScriptEngine("go run")
	reply, err := s.Run(filePath, flags)
	assert.Nil(t, err)
	for k, v := range flags {
		assert.EqualValues(t, v, reply[k])
	}
}

const script = `
package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		panic("the script requires exactly 3 arguments")
	}

	flagsPath := os.Args[1]
	replyPath := os.Args[2]

	bytes, err := ioutil.ReadFile(flagsPath)
	if err != nil {
		panic(err)
	}

	var m map[string]interface{}
	err = json.Unmarshal(bytes, &m)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(replyPath, bytes, os.ModePerm)
	if err != nil {
		panic(err)
	}

	os.Exit(0)
}
`
