package functions

import (
	"strings"
)

func IsValid(s string) bool {
	if len(s) <= 4 {
		return false
	}

	if strings.ContainsAny(s, "BUDH") == false {
		return false
	}

	for i := 0; i < len(s)-3; i++ {
		if strings.Contains(s[i+1:], s[i:i+4]) {
			return false
		}
	}
	return true
}
