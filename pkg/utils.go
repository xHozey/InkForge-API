package pkg

import "strings"

func validateLength(str string) bool {
	if len(str) > 30 || len(strings.TrimSpace(str)) == 0 {
		return false
	}
	return true
}
