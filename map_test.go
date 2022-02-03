package algo_test

import (
	"strings"
	"testing"

	"github.com/datosh/algo"
	"github.com/stretchr/testify/assert"
)

var result int

func Test_Mapper(t *testing.T) {
	t.Run("double slice of int", func(t *testing.T) {
		numbers := []int{0, 1, 2, 3, 4, 5, 7, 8, 9}
		double := func(i int) int { return i * 2 }

		doubles := algo.Map(double)(numbers)

		assert.Equal(t, 0, doubles[0])
		assert.Equal(t, 2, doubles[1])
		assert.Equal(t, 4, doubles[2])
		assert.Equal(t, 6, doubles[3])
	})

	t.Run("convert slice of int to float64", func(t *testing.T) {
		integers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		toFloat := func(i int) float64 {
			return float64(i)
		}

		floats := algo.Map(toFloat)(integers)

		assert.Equal(t, 2.0, floats[2])
	})

	t.Run("strings to uppercase", func(t *testing.T) {
		words := []string{"apples", "bananas", "oranges"}

		upperWords := algo.Map(strings.ToUpper)(words)

		assert.Equal(t, "APPLES", upperWords[0])
		assert.Equal(t, "BANANAS", upperWords[1])
		assert.Equal(t, "ORANGES", upperWords[2])
	})
}

func Benchmark_Map(b *testing.B) {
	// Setup
	numbers := []int{0, 1, 2, 3, 4, 5, 7, 8, 9}
	add := func(i int) int { return i + 1 }
	AddMapper := algo.Map(add)

	for n := 0; n < b.N; n++ {
		numbers = AddMapper(numbers)
	}

	result = numbers[0]
}

func Benchmark_Map2(b *testing.B) {
	// Setup
	numbers := []int{0, 1, 2, 3, 4, 5, 7, 8, 9}
	add := func(i int) int { return i + 1 }

	for n := 0; n < b.N; n++ {
		numbers = algo.Map2(numbers, add)
	}

	result = numbers[0]
}

func Benchmark_MapBaseline(b *testing.B) {
	// Setup
	numbers := []int{0, 1, 2, 3, 4, 5, 7, 8, 9}

	for n := 0; n < b.N; n++ {
		for i := range numbers {
			numbers[i] += 1
		}
	}

	result = numbers[0]
}
