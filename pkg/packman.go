package packman

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadFlags() map[string]string {
	var out map[string]string
	path := os.Args[1]
	flagsContent, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(flagsContent, &out)
	if err != nil {
		panic(err)
	}

	return out
}

func WriteReply(model interface{}) {
	bytes, err := json.Marshal(model)
	if err != nil {
		panic(err)
	}

	path := os.Args[2]
	err = ioutil.WriteFile(path, bytes, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
