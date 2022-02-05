package main

import (
	"fmt"

	"github.com/datosh/algo"
)

func main() {
	sum := algo.Fold(func(a, b int) int { return a + b })
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sum(numbers)) // 45
}
