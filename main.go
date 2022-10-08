package main

import (
	"fmt"
	"github.com/lucianorbr/teste_Capgemini/controllers/functions"
)

var letters = []string{
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
	"XXXXXX",
}

var resultisvalid = functions.CountIsValid(letters)
var resultratio = functions.CountRatio(letters)
var resultinvalid = functions.CountInValid(letters)

func main() {

	for i := 0; i < len(letters); i++ {
		fmt.Printf("{\"is_valid\": %v}\n", functions.IsValid(letters[i]))
	}

	fmt.Printf("{\"count_valid\": %d}\n", resultisvalid)
	fmt.Printf("{\"count_invalid\": %d}\n", resultinvalid)
	fmt.Printf("{\"ratio\": %.1f}\n", resultratio)

}
