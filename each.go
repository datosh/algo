package algo

// Each applies the provided function to each element of slice.
func Each[T any](function func(T)) func([]T) {
	return func(elements []T) {
		for i := range elements {
			function(elements[i])
		}
	}
}

// Each2 is the same as Each, but with single function call.
func Each2[T any](elements []T, function func(T)) {
	Each(function)(elements)
}
