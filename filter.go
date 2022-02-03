package algo

func Filter[Type any](predicate func(t Type) bool) func([]Type) []Type {
	return func(elements []Type) []Type {
		i := 0
		for _, element := range elements {
			if predicate(element) {
				elements[i] = element
				i++
			}
		}
		for j := i; j < len(elements); j++ {
			elements[j] = *new(Type)
		}
		return elements[:i]
	}
}

func Filter2[Type any](elements []Type, predicate func(Type) bool) []Type {
	return Filter(predicate)(elements)
}
