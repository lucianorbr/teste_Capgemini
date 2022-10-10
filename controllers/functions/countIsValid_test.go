package functions

import "testing"

func TestCountIsValid(t *testing.T) {
	test := []struct {
		input    []string
		expected int
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
			6},
	}

	for _, tt := range test {
		got := CountIsValid(tt.input)
		if got != tt.expected {
			t.Errorf("got %v, want %v", got, tt.expected)
		}
	}
}
