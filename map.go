package algo

func Map[FromType, ToType any](mapper func(t FromType) ToType) func([]FromType) []ToType {
	return func(elements []FromType) []ToType {
		result := make([]ToType, len(elements))
		for i := range elements {
			result[i] = mapper(elements[i])
		}
		return result
	}
}

func Map2[FromType, ToType any](elements []FromType, mapper func(t FromType) ToType) []ToType {
	return Map(mapper)(elements)
}
