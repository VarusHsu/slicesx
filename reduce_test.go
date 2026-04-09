// Copyright (c) 2026 Varus Hsu
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package slicesx

import (
	"fmt"
	"testing"
)

func TestReduce(t *testing.T) {
	t.Run("Sum", func(t *testing.T) {
		got := Reduce([]int{1, 2, 3, 4}, 0, func(acc, v int) int { return acc + v })
		if got != 10 {
			t.Errorf("Reduce() = %v, want 10", got)
		}
	})

	t.Run("Product", func(t *testing.T) {
		got := Reduce([]int{1, 2, 3, 4}, 1, func(acc, v int) int { return acc * v })
		if got != 24 {
			t.Errorf("Reduce() = %v, want 24", got)
		}
	})

	t.Run("Concat", func(t *testing.T) {
		got := Reduce([]string{"a", "b", "c"}, "", func(acc string, v string) string { return acc + v })
		if got != "abc" {
			t.Errorf("Reduce() = %v, want abc", got)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		got := Reduce([]int{}, 42, func(acc, v int) int { return acc + v })
		if got != 42 {
			t.Errorf("Reduce() = %v, want 42", got)
		}
	})

	t.Run("Nil", func(t *testing.T) {
		got := Reduce([]int(nil), 0, func(acc, v int) int { return acc + v })
		if got != 0 {
			t.Errorf("Reduce() = %v, want 0", got)
		}
	})

	t.Run("SingleElement", func(t *testing.T) {
		got := Reduce([]int{5}, 10, func(acc, v int) int { return acc + v })
		if got != 15 {
			t.Errorf("Reduce() = %v, want 15", got)
		}
	})
}

func BenchmarkReduce(b *testing.B) {
	array := genArray(0, 10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce(array, 0, func(acc, v int) int { return acc + v })
	}
}

func ExampleReduce() {
	sum := Reduce([]int{1, 2, 3, 4}, 0, func(acc, v int) int {
		return acc + v
	})
	fmt.Println(sum)
	// Output:
	// 10
}
