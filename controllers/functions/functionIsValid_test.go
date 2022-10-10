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
		{"OOOOOO", false},
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
		{"OOOOOO", true},
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
