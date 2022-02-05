package algo

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

// Fold2 is the same as Fold, but with single function call.
func Fold2[T any](elements []T, combiner func(T, T) T) T {
	return Fold(combiner)(elements)
}
