# slicesx

[![Go Reference](https://pkg.go.dev/badge/github.com/varushsu/slicesx.svg)](https://pkg.go.dev/github.com/varushsu/slicesx)
[![Go Report Card](https://goreportcard.com/badge/github.com/varushsu/slicesx)](https://goreportcard.com/report/github.com/varushsu/slicesx)
[![Code Coverage](https://img.shields.io/codecov/c/github/varushsu/slicesx.svg)](https://codecov.io/gh/VarusHsu/slicesx)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A Go library providing generic slice utility functions that extend the standard library.

## Installation

```bash
go get github.com/varushsu/slicesx
```

Requires Go 1.21 or later.

## Functions

### Chunk

`Chunk` splits a slice into multiple sub-slices using one or more step sizes.

```go
import "github.com/varushsu/slicesx"

array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// Single step size
chunks := slicesx.Chunk(array, 3)
// [[1 2 3] [4 5 6] [7 8 9] [10]]

// Multiple step sizes (last step is reused)
chunks = slicesx.Chunk(array, 1, 2, 3)
// [[1] [2 3] [4 5 6] [7 8 9] [10]]

// Step with zero (produces empty chunk)
chunks = slicesx.Chunk(array, 1, 0, 3)
// [[1] [] [2 3 4] [5 6 7] [8 9 10]]
```

### Filter

`Filter` returns a new slice containing only the elements that satisfy the predicate function.

```go
evens := slicesx.Filter([]int{1, 2, 3, 4, 5}, func(v int) bool {
    return v%2 == 0
})
// [2 4]
```

### Map

`Map` transforms each element using the provided function and returns a new slice.

```go
doubled := slicesx.Map([]int{1, 2, 3}, func(v int) int {
    return v * 2
})
// [2 4 6]
```

### Reduce

`Reduce` reduces a slice to a single value using an accumulator function.

```go
sum := slicesx.Reduce([]int{1, 2, 3, 4}, 0, func(acc, v int) int {
    return acc + v
})
// 10
```

### Contains

`Contains` reports whether a value is present in the slice.

```go
slicesx.Contains([]int{1, 2, 3}, 2)  // true
slicesx.Contains([]int{1, 2, 3}, 4)  // false
```

### IndexOf

`IndexOf` returns the index of the first occurrence of a value, or -1 if not found.

```go
slicesx.IndexOf([]string{"a", "b", "c"}, "b")  // 1
slicesx.IndexOf([]string{"a", "b", "c"}, "d")  // -1
```

### Unique

`Unique` returns a new slice with duplicate elements removed, preserving order.

```go
slicesx.Unique([]int{1, 2, 2, 3, 1, 4})
// [1 2 3 4]
```

### Flatten

`Flatten` merges a slice of slices into a single flat slice.

```go
slicesx.Flatten([][]int{{1, 2}, {3, 4}, {5}})
// [1 2 3 4 5]
```

### Intersection

`Intersection` returns a new slice containing elements present in both input slices.

```go
slicesx.Intersection([]int{1, 2, 3, 4}, []int{3, 4, 5, 6})
// [3 4]
```

### Difference

`Difference` returns elements that are in the first slice but not in the second.

```go
slicesx.Difference([]int{1, 2, 3, 4}, []int{3, 4, 5, 6})
// [1 2]
```

### Union

`Union` returns a new slice containing all unique elements from both input slices.

```go
slicesx.Union([]int{1, 2, 3}, []int{3, 4, 5})
// [1 2 3 4 5]
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
