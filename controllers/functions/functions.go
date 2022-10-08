package functions

import (
	"strings"
)

// função que verifica se a string é valida
func IsValid(s string) bool {
	if len(s) <= 4 {
		return false
	}

	//String letters can only be: (B, U, D, H)
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

// função que calcula o numero de strings validas
func CountIsValid(s []string) int {
	count := 0
	for _, v := range s {
		if IsValid(v) {
			count++
		}
	}
	return count
}

// função que calcula o numero de strings invalidas
func CountInValid(s []string) int {
	count := 0
	for _, v := range s {
		if !IsValid(v) {
			count++
		}
	}
	return count
}

// função que calcula o ratio
func CountRatio(s []string) float64 {
	return float64(CountIsValid(s)) / float64(len(s))
}
