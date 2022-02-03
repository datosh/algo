package algo_test

import (
	"reflect"
	"testing"

	"github.com/datosh/algo"
	"github.com/stretchr/testify/assert"
)

func filterInterface(list interface{}, pred func(interface{}) bool) interface{} {
	contentType := reflect.TypeOf(list)
	contentValue := reflect.ValueOf(list)

	newContent := reflect.MakeSlice(contentType, 0, 0)
	for i := 0; i < contentValue.Len(); i++ {
		if content := contentValue.Index(i); pred(content.Interface()) {
			newContent = reflect.Append(newContent, content)
		}
	}
	return newContent.Interface()
}

func onlyEvenInterface(i interface{}) bool {
	number := i.(int)
	return (number % 2) == 0
}

func onlyEven(i int) bool {
	return (i % 2) == 0
}

// generateNumbers in range (including) from (excluding) to.
func generateNumbers(from, to int) []int {
	var numbers []int
	for i := from; i < to; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}

func Test_Filter(t *testing.T) {
	t.Run("keep even numbers", func(t *testing.T) {
		numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

		filteredNumbers := algo.Filter(onlyEven)(numbers)

		assert.Equal(t, 0, filteredNumbers[0])
		assert.Equal(t, 2, filteredNumbers[1])
		assert.Equal(t, 4, filteredNumbers[2])
		assert.Equal(t, 6, filteredNumbers[3])
		assert.Equal(t, 8, filteredNumbers[4])
		assert.Len(t, filteredNumbers, 5)
	})
}

func Benchmark_Filter(b *testing.B) {
	// Setup
	var listsOfNumbers [][]int
	for i := 0; i < b.N; i++ {
		listsOfNumbers = append(listsOfNumbers, generateNumbers(0, 100000))
	}

	for n := 0; n < b.N; n++ {
		filteredList := algo.Filter(onlyEven)(listsOfNumbers[n])
		result = filteredList[0]
	}
}

func Benchmark_Filter2(b *testing.B) {
	// Setup
	var listsOfNumbers [][]int
	for i := 0; i < b.N; i++ {
		listsOfNumbers = append(listsOfNumbers, generateNumbers(0, 100000))
	}

	for n := 0; n < b.N; n++ {
		filteredList := algo.Filter2(listsOfNumbers[n], onlyEven)
		result = filteredList[0]
	}
}

func Benchmark_FilterInterfaced(b *testing.B) {
	// Setup
	var listsOfNumbers [][]int
	for i := 0; i < b.N; i++ {
		listsOfNumbers = append(listsOfNumbers, generateNumbers(0, 100000))
	}

	for n := 0; n < b.N; n++ {
		filteredList := filterInterface(listsOfNumbers[n], onlyEvenInterface)
		result = (filteredList).([]int)[0]
	}
}

func Benchmark_FilterBaseline(b *testing.B) {
	// Setup
	var listsOfNumbers [][]int
	for i := 0; i < b.N; i++ {
		listsOfNumbers = append(listsOfNumbers, generateNumbers(0, 100000))
	}

	for n := 0; n < b.N; n++ {
		i := 0
		for _, element := range listsOfNumbers[n] {
			if (element % 2) == 0 {
				listsOfNumbers[n][i] = element
				i++
			}
		}
		for j := i; j < len(listsOfNumbers[n]); j++ {
			listsOfNumbers[n][j] = 0
		}
		listsOfNumbers[n] = listsOfNumbers[n][i:]
		result = listsOfNumbers[n][0]
	}
}
