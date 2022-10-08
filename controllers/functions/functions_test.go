package functions

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	test := []struct {
		input    string
		expected bool
	}{
		{"DUHBHB", true},
		{"DUBUHD", true},
		{"UBUUHU", true},
		{"BHBDHH", true},
		{"DDDDUB", true},
		{"UDBDUH", true},
		{"NNNNNN", false},
		{"MMMMMN", false},
		{"NNNNNN", false},
		{"HHHHHH", false},
		{"YYYYYY", false},
		{"XXXXXX", false},
	}

	for _, tt := range test {
		got := IsValid(tt.input)
		if got != tt.expected {
			t.Errorf("got %v, want %v", got, tt.expected)
		}
	}
}

func TestInvalid(t *testing.T) {
	test := []struct {
		input    string
		expected bool
	}{
		{"DUHBHB", false},
		{"DUBUHD", false},
		{"UBUUHU", false},
		{"BHBDHH", false},
		{"DDDDUB", false},
		{"UDBDUH", false},
		{"NNNNNN", true},
		{"MMMMMN", true},
		{"NNNNNN", true},
		{"HHHHHH", true},
		{"YYYYYY", true},
		{"XXXXXX", true},
	}

	for _, tt := range test {
		got := !IsValid(tt.input)
		if got != tt.expected {
			t.Errorf("got %v, want %v", got, tt.expected)
		}
	}
}

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
			"NNNNNN",
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

func TestCountInValid(t *testing.T) {
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
			"NNNNNN",
			"MMMMMN",
			"NNNNNN",
			"HHHHHH",
			"YYYYYY",
			"XXXXXX"},
			6},
	}

	for _, tt := range test {
		got := CountInValid(tt.input)
		if got != tt.expected {
			t.Errorf("got %v, want %v", got, tt.expected)
		}
	}
}

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
			"NNNNNN",
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
