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

func TestUnique(t *testing.T) {
	t.Run("WithDuplicates", func(t *testing.T) {
		got := Unique([]int{1, 2, 2, 3, 1, 4})
		want := []int{1, 2, 3, 4}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Unique() = %v, want %v", got, want)
		}
	})

	t.Run("NoDuplicates", func(t *testing.T) {
		got := Unique([]int{1, 2, 3})
		want := []int{1, 2, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Unique() = %v, want %v", got, want)
		}
	})

	t.Run("AllSame", func(t *testing.T) {
		got := Unique([]int{5, 5, 5, 5})
		want := []int{5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Unique() = %v, want %v", got, want)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		got := Unique([]int{})
		if len(got) != 0 {
			t.Errorf("Unique() = %v, want empty slice", got)
		}
	})

	t.Run("Nil", func(t *testing.T) {
		got := Unique([]int(nil))
		if len(got) != 0 {
			t.Errorf("Unique() = %v, want empty slice", got)
		}
	})

	t.Run("Strings", func(t *testing.T) {
		got := Unique([]string{"a", "b", "a", "c", "b"})
		want := []string{"a", "b", "c"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Unique() = %v, want %v", got, want)
		}
	})

	t.Run("SingleElement", func(t *testing.T) {
		got := Unique([]int{42})
		want := []int{42}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Unique() = %v, want %v", got, want)
		}
	})
}

func BenchmarkUnique(b *testing.B) {
	array := make([]int, 10000)
	for i := range array {
		array[i] = i % 100
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Unique(array)
	}
}

func ExampleUnique() {
	fmt.Println(Unique([]int{1, 2, 2, 3, 1, 4}))
	// Output:
	// [1 2 3 4]
}
