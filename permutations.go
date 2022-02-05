package algo

// Permutations uses non-recursive Heap's algorithm to calculate all
// permutations of the provided input slice **eagerly**.
// The argument will not be manipulated.
// The result will be of size: factorial of number of input elements.
// See also: https://en.wikipedia.org/wiki/Heap%27s_algorithm
func Permutations[T any](elements []T) [][]T {
	localElements := makeCopy(elements)

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
				Swap(&localElements[0], &localElements[i])
			} else {
				Swap(&localElements[counters[i]], &localElements[i])
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

// Should Swap be provided? Swapping is a simple as
// a, b = b, a
// but having a speaking function name might be easier to read?
func Swap[T any](a, b *T) {
	*b, *a = *a, *b
}

// factorial calculates n!
// See also: https://en.wikipedia.org/wiki/Factorial
func factorial(n int) int {
	// TODO: Test for n < 0
	// TODO: What is the upper limit for int? 20?
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// makeCopy extends go's build-in copy by creating a slice (of correct size)
// as copy target and returns it.
func makeCopy[T any](src []T) []T {
	dest := make([]T, len(src))
	copy(dest, src)
	return dest
}
