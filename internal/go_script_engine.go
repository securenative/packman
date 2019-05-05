package internal

import "os/exec"

type GoScriptEngine struct {
}

func NewGoScriptEngine() *GoScriptEngine {
	return &GoScriptEngine{}
}

func (this *GoScriptEngine) Run(scriptFile string, args []string) error {
	var cmdArgs []string
	cmdArgs = append(cmdArgs, "run")
	cmdArgs = append(cmdArgs, scriptFile)
	cmdArgs = append(cmdArgs, args...)

	cmd := exec.Command("go", cmdArgs...)
	return cmd.Run()
}
