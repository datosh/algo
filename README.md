# algo

Algorithm library of golang1.18+ using
[type parameters.](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)

## Installation

```bash
go1.18beta2 get ...
```

## Usage

```go

# show off the amazing things

```

## Goal

Explore how a go library for algorithms 6 data structures COULD look like.

To get an idea what is usually implemented in such a library I looked at
* [D standard library](https://dlang.org/phobos/std_algorithm_iteration.html)

## Limitations

* Currently only support slices

## Roadmap / Ideas

* Iterator based approach to support any data structure

## Performance

benchmarks comparing custom, type parameter based & interface{} based
implementations. 

## Contributing

## License
[MIT](https://choosealicense.com/licenses/mit/)