package algo

func Each[T any](elements []T, function func(T)) {
	for i := range elements {
		function(elements[i])
	}
}

func Permutations[T any](elements []T) [][]T {
	return *new([][]T)
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
