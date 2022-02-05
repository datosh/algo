package main

import (
	"fmt"
	"strings"

	"github.com/datosh/algo"
)

func main() {
	// Preconfigure functions
	onlyShortWords := algo.Filter(func(s string) bool { return len(s) <= 5 })
	shout := algo.Map(func(s string) string { return strings.ToUpper(s) })
	concat := algo.Fold(func(t1, t2 string) string { return t1 + " " + t2 })

	words := []string{"Hello", "gophers", "World"}

	sentence := concat(shout(onlyShortWords(words)))
	fmt.Println(sentence)
}
