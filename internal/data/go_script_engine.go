package data

import (
	"fmt"
	"os/exec"
)

type GoScriptEngine struct {
}

func NewGoScriptEngine() *GoScriptEngine {
	return &GoScriptEngine{}
}

func (this *GoScriptEngine) Run(scriptFile string, args []string) error {
	var cmdArgs []string
	cmdArgs = append(cmdArgs, "run")
	cmdArgs = append(cmdArgs, scriptFile)
	cmdArgs = append(cmdArgs, "--")
	cmdArgs = append(cmdArgs, args...)

	fmt.Println(fmt.Sprintf("Runnin main.go with %v", cmdArgs))

	cmd := exec.Command("go", cmdArgs...)
	result, err := cmd.Output()
	if err != nil {
		return err
	}

	fmt.Println(string(result))

	return nil
}
