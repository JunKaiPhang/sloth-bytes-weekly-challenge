package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvalAlgebra(t *testing.T) {
	assert := assert.New(t)

	t.Run("it should return expected data", func(t *testing.T) {
		assert.Equal(17, evalAlgebra("2 + x = 19"))

		assert.Equal(3, evalAlgebra("4 - x = 1"))

		assert.Equal(43, evalAlgebra("x + 10 = 53"))

		assert.Equal(3, evalAlgebra("-23 + x = -20"))

		assert.Equal(-5, evalAlgebra("10 + x = 5"))

		assert.Equal(131, evalAlgebra("-49 - x = -180"))

		assert.Equal(-34, evalAlgebra("x - 22 = -56"))

		assert.Equal(154, evalAlgebra("x - 54 = 100"))
	})
}
