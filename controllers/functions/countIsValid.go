package functions

func CountIsValid(s []string) int {
	count := 0
	for _, v := range s {
		if IsValid(v) {
			count++
		}
	}
	return count
}
