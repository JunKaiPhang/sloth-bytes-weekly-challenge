package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTraceWordPath(t *testing.T) {
	assert := assert.New(t)

	t.Run("it should return expected data", func(t *testing.T) {
		result1 := traceWordPath("BISCUIT", [][]string{
			{"B", "I", "T", "R"},
			{"I", "U", "A", "S"},
			{"S", "C", "V", "W"},
			{"D", "O", "N", "E"},
		})
		assert.Equal([][]int{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {1, 1}, {0, 1}, {0, 2}}, result1)

		result2 := traceWordPath("HELPFUL", [][]string{
			{"L", "I", "T", "R"},
			{"U", "U", "A", "S"},
			{"L", "U", "P", "O"},
			{"E", "F", "E", "H"},
		})
		assert.Nil(result2)

		result3 := traceWordPath("UKULELE", [][]string{
			{"N", "H", "B", "W"},
			{"E", "X", "A", "D"},
			{"L", "A", "U", "U"},
			{"E", "L", "U", "K"},
		})
		assert.Equal([][]int{{2, 3}, {3, 3}, {3, 2}, {3, 1}, {3, 0}, {2, 0}, {1, 0}}, result3)

		result4 := traceWordPath("SURVIVAL", [][]string{
			{"V", "L", "R", "L"},
			{"V", "A", "I", "V"},
			{"I", "O", "S", "C"},
			{"V", "R", "U", "F"},
		})
		assert.Equal([][]int{{2, 2}, {3, 2}, {3, 1}, {3, 0}, {2, 0}, {1, 0}, {1, 1}, {0, 1}}, result4)
	})
}
