package pkg

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

const replyFile = "reply.json"

func Reply(data interface{}) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	destPath := filepath.Join(pwd, replyFile)
	if err = ioutil.WriteFile(destPath, bytes, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func ReadReply(projectPath string) (interface{}, error) {

	destPath := filepath.Join(projectPath, replyFile)
	bytes, err := ioutil.ReadFile(destPath)
	if err != nil {
		return nil, err
	}

	var data = new(interface{})
	if err = json.Unmarshal(bytes, data); err != nil {
		return nil, err
	}

	return data, nil
}
