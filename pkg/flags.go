package pkg

import "fmt"

func ParseFlags(cmdArgs []string) map[string]string {

	if len(cmdArgs)%2 != 0 {
		panic(fmt.Sprintf("incoming flags must be an even array but got %v", cmdArgs))
	}

	out := make(map[string]string)
	for idx, item := range cmdArgs {
		if idx%2 == 0 {
			out[item] = cmdArgs[idx+1]
		}
	}

	return out
}
