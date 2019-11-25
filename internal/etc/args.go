package etc

import "strings"

func ArgsToFlagsMap(args []string) map[string]string {
	out := make(map[string]string)
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
