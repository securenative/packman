package data

import (
	"encoding/json"
	"fmt"
	"github.com/securenative/packman/internal/etc"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type genericScriptEngine struct {
	command string
}

func NewGenericScriptEngine(command string) *genericScriptEngine {
	return &genericScriptEngine{command: command}
}

func (this *genericScriptEngine) Run(scriptPath string, flags map[string]string) (map[string]interface{}, error) {
	flagsFile := pwdPath(scriptPath, "~flags.json")
	replyFile := pwdPath(scriptPath, "~reply.json")
	err := etc.WriteFile(flagsFile, flags, etc.JsonEncoder)

	mainCommand, args, err := splitCommand(this.command)
	if err != nil {
		return nil, err
	}

	var cmdArgs []string
	cmdArgs = append(cmdArgs, args...)
	cmdArgs = append(cmdArgs, filepath.Base(scriptPath))
	cmdArgs = append(cmdArgs, panicOrString(filepath.Abs, flagsFile))
	cmdArgs = append(cmdArgs, panicOrString(filepath.Abs, replyFile))

	cmd := exec.Command(mainCommand, cmdArgs...)
	cmd.Dir = filepath.Dir(scriptPath)

	etc.PrintInfo(fmt.Sprintf("Running '%s %v'", mainCommand, cmdArgs))
	result, err := cmd.CombinedOutput()
	if err != nil {
		if result != nil {
			etc.PrintError(" FAILED\n")
			etc.PrintError(string(result) + "\n")
		}
		return nil, err
	}
	etc.PrintSuccess(" OK\n")
	etc.PrintResponse(string(result))

	etc.PrintInfo("Trying to read reply file: %s...", replyFile)
	content, err := etc.ReadFile(replyFile)
	if err != nil {
		etc.PrintError(" FAILED\n")
		etc.PrintError("Unable to read reply file from: %s\n", replyFile)
		return nil, err
	}

	var out map[string]interface{}
	err = json.Unmarshal([]byte(content), &out)
	if err != nil {
		return nil, err
	}

	os.Remove(flagsFile)
	os.Remove(replyFile)

	etc.PrintSuccess(" OK\n")
	etc.PrettyPrintJson(out)

	return out, nil
}

func pwdPath(scriptPath string, newName string) string {
	fileName := filepath.Base(scriptPath)
	scriptFolder := strings.Replace(scriptPath, fileName, "", 1)
	return filepath.Join(scriptFolder, newName)
}

func splitCommand(command string) (string, []string, error) {
	parts := strings.Split(command, " ")
	if parts != nil && len(parts) > 0 {
		return parts[0], parts[1:], nil
	}
	return "", nil, fmt.Errorf("cannot parse command %s, the command syntax should be as follows: 'commnad arg1 arg2 arg3 ...'", command)
}

func panicOrString(f func(s string) (string, error), s string) string {
	str, err := f(s)
	if err != nil {
		panic(err)
	}
	return str
}
