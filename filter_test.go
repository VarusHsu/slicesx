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
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	t.Run("FilterEvens", func(t *testing.T) {
		got := Filter([]int{1, 2, 3, 4, 5, 6}, func(v int) bool { return v%2 == 0 })
		want := []int{2, 4, 6}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Filter() = %v, want %v", got, want)
		}
	})

	t.Run("FilterNone", func(t *testing.T) {
		got := Filter([]int{1, 2, 3}, func(v int) bool { return v > 10 })
		if len(got) != 0 {
			t.Errorf("Filter() = %v, want empty slice", got)
		}
	})

	t.Run("FilterAll", func(t *testing.T) {
		got := Filter([]int{1, 2, 3}, func(v int) bool { return true })
		want := []int{1, 2, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Filter() = %v, want %v", got, want)
		}
	})

	t.Run("FilterEmpty", func(t *testing.T) {
		got := Filter([]int{}, func(v int) bool { return true })
		if len(got) != 0 {
			t.Errorf("Filter() = %v, want empty slice", got)
		}
	})

	t.Run("FilterNil", func(t *testing.T) {
		got := Filter([]int(nil), func(v int) bool { return true })
		if len(got) != 0 {
			t.Errorf("Filter() = %v, want empty slice", got)
		}
	})

	t.Run("FilterStrings", func(t *testing.T) {
		got := Filter([]string{"apple", "banana", "avocado", "blueberry"}, func(v string) bool {
			return v[0] == 'a'
		})
		want := []string{"apple", "avocado"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Filter() = %v, want %v", got, want)
		}
	})
}

func BenchmarkFilter(b *testing.B) {
	array := genArray(0, 10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Filter(array, func(v int) bool { return v%2 == 0 })
	}
}

func ExampleFilter() {
	evens := Filter([]int{1, 2, 3, 4, 5}, func(v int) bool {
		return v%2 == 0
	})
	fmt.Println(evens)
	// Output:
	// [2 4]
}
