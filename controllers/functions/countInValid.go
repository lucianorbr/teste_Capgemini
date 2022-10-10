package functions

func CountInValid(s []string) int {
	count := 0
	for _, v := range s {
		if !IsValid(v) {
			count++
		}
	}
	return count
}
