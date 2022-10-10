package functions

import "testing"

func TestRatio(t *testing.T) {
	test := []struct {
		input    []string
		expected float64
	}{
		{[]string{
			"DUHBHB",
			"DUBUHD",
			"UBUUHU",
			"BHBDHH",
			"DDDDUB",
			"UDBDUH",
			"OOOOOO",
			"MMMMMN",
			"NNNNNN",
			"HHHHHH",
			"YYYYYY",
			"XXXXXX"},
			0.5},
	}

	for _, tt := range test {
		got := CountRatio(tt.input)
		if got != tt.expected {
			t.Errorf("got %v, want %v", got, tt.expected)
		}
	}
}
