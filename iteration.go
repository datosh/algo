package algo

func Each[T any](function func(T)) func([]T) {
	return func(elements []T) {
		for i := range elements {
			function(elements[i])
		}
	}
}

func Each2[T any](elements []T, function func(T)) {
	Each(function)(elements)
}

// There is no need to implement Swap in go, a simple
// a, b = b, a
// can be written in line
// func Swap[T any](a, b T) (T, T) {
// 	return b, a
// }

func factorial(n int) int {
	// TODO: Test for n < 0
	// TODO: What is the upper limit for int? 20?
	// https://en.wikipedia.org/wiki/Factorial
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func Permutations[T any](elements []T) [][]T {
	localElements := make([]T, len(elements))
	copy(localElements, elements)

	permutations := make([][]T, factorial(len(localElements)))
	for i := 0; i < len(permutations); i++ {
		permutations[i] = make([]T, len(localElements))
	}
	permutationIdx := 0
	counters := make([]int, len(localElements))

	copy(permutations[permutationIdx], localElements)
	permutationIdx++

	// i acts similarly to a stack pointer
	i := 0
	for i < len(localElements) {
		if counters[i] < i {
			if i%2 == 0 {
				localElements[0], localElements[i] = localElements[i], localElements[0]
			} else {
				localElements[counters[i]], localElements[i] = localElements[i], localElements[counters[i]]
			}

			copy(permutations[permutationIdx], localElements)
			permutationIdx++
			counters[i] += 1

			i = 0
		} else {
			counters[i] = 0
			i++
		}
	}
	return permutations
}

// Fold (also called Reduce) applies a combining operator to each
// element of the slice.
// Also see: https://en.wikipedia.org/wiki/Fold_(higher-order_function)
func Fold[T any](combiner func(T, T) T) func([]T) T {
	return func(elements []T) T {
		result := elements[0]
		for i := 1; i < len(elements); i++ {
			result = combiner(result, elements[i])
		}
		return result
	}
}

func Fold2[T any](elements []T, combiner func(T, T) T) T {
	result := elements[0]
	for i := 1; i < len(elements); i++ {
		result = combiner(result, elements[i])
	}
	return result
}
