package algo

// Map takes a mapping function and applies it to each element of the slice,
// returning a new slice with the mapped elements.
func Map[FromType, ToType any](mapper func(t FromType) ToType) func([]FromType) []ToType {
	return func(elements []FromType) []ToType {
		result := make([]ToType, len(elements))
		for i := range elements {
			result[i] = mapper(elements[i])
		}
		return result
	}
}

// Map2 is the same as Map, but with single function call.
func Map2[FromType, ToType any](elements []FromType, mapper func(t FromType) ToType) []ToType {
	return Map(mapper)(elements)
}
