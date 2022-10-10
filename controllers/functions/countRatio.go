package functions

func CountRatio(s []string) float64 {
	return float64(CountIsValid(s)) / float64(len(s))
}
