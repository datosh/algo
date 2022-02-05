package algo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_factorial(t *testing.T) {
	t.Run("0!", func(t *testing.T) {
		assert.Equal(t, 1, factorial(0))
	})

	t.Run("1!", func(t *testing.T) {
		assert.Equal(t, 1, factorial(1))
	})

	t.Run("2!", func(t *testing.T) {
		assert.Equal(t, 2, factorial(2))
	})

	t.Run("3!", func(t *testing.T) {
		assert.Equal(t, 6, factorial(3))
	})

	t.Run("4!", func(t *testing.T) {
		assert.Equal(t, 24, factorial(4))
	})

	t.Run("5!", func(t *testing.T) {
		assert.Equal(t, 120, factorial(5))
	})

	t.Run("6!", func(t *testing.T) {
		assert.Equal(t, 720, factorial(6))
	})
}

func Test_Permutations(t *testing.T) {
	t.Run("permutate slice of 3 ints", func(t *testing.T) {
		numbers := []int{2, 3, 4}

		permutations := Permutations(numbers)

		assert.Len(t, permutations, 6)
		assert.Equal(t, []int{2, 3, 4}, permutations[0])
		assert.Equal(t, []int{3, 2, 4}, permutations[1])
		assert.Equal(t, []int{4, 2, 3}, permutations[2])
		assert.Equal(t, []int{2, 4, 3}, permutations[3])
		assert.Equal(t, []int{3, 4, 2}, permutations[4])
		assert.Equal(t, []int{4, 3, 2}, permutations[5])

		assert.Equal(t, []int{2, 3, 4}, numbers)
	})

	t.Run("permutate slice runes", func(t *testing.T) {
		s := "hello"

		permutations := Permutations(strings.Split(s, ""))

		assert.Len(t, permutations, 120)
		assert.Equal(t, "hello", s)
	})
}

func Test_Fold(t *testing.T) {
	t.Run("summing slice of int", func(t *testing.T) {
		numbers := []int{2, 3, 4, 5}
		sumOp := func(a, b int) int { return a + b }

		sum := Fold(sumOp)(numbers)

		assert.Equal(t, 14, sum)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		words := []string{"this", "is", "freakin", "awesome"}
		cat := func(s1, s2 string) string { return s1 + " " + s2 }

		sentence := Fold(cat)(words)

		assert.Equal(t, "this is freakin awesome", sentence)
	})
}
