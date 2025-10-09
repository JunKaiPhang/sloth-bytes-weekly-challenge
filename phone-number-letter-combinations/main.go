package main

import (
	"fmt"
)

var numLetterMap = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func main() {
	res := letterCombinations("23")
	fmt.Println(res)
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var result []string

	recurFunc(0, "", digits, &result)

	return result
}

func recurFunc(index int, temp string, digits string, result *[]string) {
	if index == len(digits) {
		*result = append(*result, temp)
		return
	}

	letters := numLetterMap[digits[index]]
	for _, letter := range letters {
		recurFunc(index+1, temp+letter, digits, result)
	}
}
