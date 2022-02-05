# algo

Algorithm library for golang1.18+ using
[type parameters.](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)

## Installation

```bash
# Or any other way to install go1.18 beta.
# https://go.dev/blog/go1.18beta2
go install golang.org/dl/go1.18beta2@latest
go1.18beta2 download
# Install library
go1.18beta2 get github.com/datosh/algo
```

## Usage

```go
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

```

```go
package main

import (
	"fmt"
	"strings"

	"github.com/datosh/algo"
)

func main() {
	onlyShortWords := algo.Filter(func(s string) bool { return len(s) <= 5 })
	shout := algo.Map(func(s string) string { return strings.ToUpper(s) })
	concat := algo.Fold(func(t1, t2 string) string { return t1 + " " + t2 })

	words := []string{"Hello", "gophers", "World"}

	sentence := concat(shout(onlyShortWords(words)))
	fmt.Println(sentence) // HELLO WORLD
}
```

## Goal

Explore how a go library for algorithms COULD look like.

To get an idea what is usually implemented in such a library I looked at
* [D Algorithms library](https://dlang.org/phobos/std_algorithm_iteration.html)
* [C++ Algorithms library](https://en.cppreference.com/w/cpp/algorithm)

## Limitations

* Currently only support slices

## Roadmap / Ideas

* "Type cast" a slice of build-in types
* "Type cast" a slice of complex types
* Iterator based approach to support any data structure
    * [On Iteration - Andrei Alexandrescu](https://www.informit.com/articles/printerfriendly/1407357)

## Performance

> :warning: **No performance optimizations were implemented, yet!**

These benchmarks were created to get a feeling for the rough performance
characteristics of golangs new type parameters.

**Baseline** is a custom implementation for a specific type.  
**Filter/Map** is the generic implementation returning a filter/map function.  
**Filter2/Map2** is the generic implementation doing the operation in one call.  
**Interfaced** is a `interface{}` based implementations.  


**cpu: 11th Gen Intel(R) Core(TM) i7-11800H @ 2.30GHz**
| Test | Iterations | Speed |
| :- | :-: | -: |
| Filter Baseline | 506 | 2989602 ns/op |
| Filter | 172 | 6455209 ns/op |
| Filter2 | 451 | 8465454 ns/op |
| Filter Interfaced | 48 | 25998184 ns/op |
| Map Baseline | 97946026 | 12.43 ns/op |
| Map | 6051519 | 182.3 ns/op |
| Map2 | 4241458 | 292.0 ns/op |

My take-away is that a performance hit is to be expected when using type
paremters compared to a custom implementation. On the other hand, if you
relied on `interface{}` & `reflection` based implementations in the past,
type parameters will provide you both with a performance as well as
safety (compiler checks) improvement.

## Contributing

## License
[MIT](https://choosealicense.com/licenses/mit/)