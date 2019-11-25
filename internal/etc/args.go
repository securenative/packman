package etc

import (
	"fmt"
	"strings"
)

func ArgsToFlagsMap(args []string) map[string]string {
	out := make(map[string]string)

	if len(args)%2 != 0 {
		panic(fmt.Errorf("num of flags must be an even number but got %d. flags must be k-v pairs such as -port 4000 -host 127.0.0.1", len(args)))
	}

	for idx, arg := range args {
		if isFlagKey(arg) {
			key := strings.TrimPrefix(arg, "-")
			out[key] = args[idx+1]
		}
	}
	return out
}

func isFlagKey(arg string) bool {
	if strings.HasPrefix(arg, "-") {
		return true
	}
	return false
}
