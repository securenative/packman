package etc

import (
	"encoding/json"
	"github.com/fatih/color"
)

func PrintInfo(message string, args ...interface{}) {
	c := color.New(color.FgCyan)
	_, _ = c.Printf(message, args...)
}

func PrintResponse(message string, args ...interface{}) {
	c := color.New(color.FgYellow).Add(color.Italic)
	_, _ = c.Printf(message, args...)
}

func PrettyPrintJson(m map[string]interface{}) {
	bytes, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		PrintError(err.Error())
		return
	}

	PrintResponse("%s\n", string(bytes))
}

func PrintError(message string, args ...interface{}) {
	c := color.New(color.FgRed).Add(color.Bold)
	_, _ = c.Printf(message, args...)
}

func PrintSuccess(message string, args ...interface{}) {
	c := color.New(color.FgGreen).Add(color.Bold)
	_, _ = c.Printf(message, args...)
}
