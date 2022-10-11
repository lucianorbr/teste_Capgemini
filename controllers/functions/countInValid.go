package functions

// CountInValid counts the number of invalid blocks
func CountInValid(s []string) int {
	count := 0
	for _, v := range s {
		if !IsValid(v) {
			count++
		}
	}
	return count
}
