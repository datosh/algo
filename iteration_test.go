package algo_test

import (
	"testing"

	"github.com/datosh/algo"
	"github.com/stretchr/testify/assert"
)

func Test_Fold(t *testing.T) {
	t.Run("summing slice of int", func(t *testing.T) {
		numbers := []int{2, 3, 4, 5}
		sumOp := func(a, b int) int { return a + b }

		sum := algo.Fold(sumOp)(numbers)

		assert.Equal(t, 14, sum)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		words := []string{"this", "is", "freakin", "awesome"}
		cat := func(s1, s2 string) string { return s1 + " " + s2 }

		sentence := algo.Fold(cat)(words)

		assert.Equal(t, "this is freakin awesome", sentence)
	})
}
