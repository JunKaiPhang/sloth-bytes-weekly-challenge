package main

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Empty input",
			input:    "",
			expected: []string{},
		},
		{
			name:     "Single digit - 2",
			input:    "2",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "Single digit - 7 (4 letters)",
			input:    "7",
			expected: []string{"p", "q", "r", "s"},
		},
		{
			name:  "Two digits - 23",
			input: "23",
			expected: []string{
				"ad", "ae", "af",
				"bd", "be", "bf",
				"cd", "ce", "cf",
			},
		},
		{
			name:  "Two digits - 79 (both have 4 letters)",
			input: "79",
			expected: []string{
				"pw", "px", "py", "pz",
				"qw", "qx", "qy", "qz",
				"rw", "rx", "ry", "rz",
				"sw", "sx", "sy", "sz",
			},
		},
		{
			name:  "Three digits - 234",
			input: "234",
			expected: []string{
				"adg", "adh", "adi",
				"aeg", "aeh", "aei",
				"afg", "afh", "afi",
				"bdg", "bdh", "bdi",
				"beg", "beh", "bei",
				"bfg", "bfh", "bfi",
				"cdg", "cdh", "cdi",
				"ceg", "ceh", "cei",
				"cfg", "cfh", "cfi",
			},
		},
		{
			name:  "Three digits - 568",
			input: "568",
			expected: []string{
				"jmt", "jmu", "jmv",
				"jnt", "jnu", "jnv",
				"jot", "jou", "jov",
				"kmt", "kmu", "kmv",
				"knt", "knu", "knv",
				"kot", "kou", "kov",
				"lmt", "lmu", "lmv",
				"lnt", "lnu", "lnv",
				"lot", "lou", "lov",
			},
		},
		{
			name:  "Four digits - 2345",
			input: "2345",
			expected: []string{
				"adgj", "adgk", "adgl", "adhj", "adhk", "adhl", "adij", "adik", "adil",
				"aegj", "aegk", "aegl", "aehj", "aehk", "aehl", "aeij", "aeik", "aeil",
				"afgj", "afgk", "afgl", "afhj", "afhk", "afhl", "afij", "afik", "afil",
				"bdgj", "bdgk", "bdgl", "bdhj", "bdhk", "bdhl", "bdij", "bdik", "bdil",
				"begj", "begk", "begl", "behj", "behk", "behl", "beij", "beik", "beil",
				"bfgj", "bfgk", "bfgl", "bfhj", "bfhk", "bfhl", "bfij", "bfik", "bfil",
				"cdgj", "cdgk", "cdgl", "cdhj", "cdhk", "cdhl", "cdij", "cdik", "cdil",
				"cegj", "cegk", "cegl", "cehj", "cehk", "cehl", "ceij", "ceik", "ceil",
				"cfgj", "cfgk", "cfgl", "cfhj", "cfhk", "cfhl", "cfij", "cfik", "cfil",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := letterCombinations(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("Input %s → got %v, want %v", tc.input, res, tc.expected)
			}
		})
	}
}
