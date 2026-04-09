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

func TestFlatten(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		got := Flatten([][]int{{1, 2}, {3, 4}, {5}})
		want := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})

	t.Run("SingleSlice", func(t *testing.T) {
		got := Flatten([][]int{{1, 2, 3}})
		want := []int{1, 2, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})

	t.Run("EmptyInner", func(t *testing.T) {
		got := Flatten([][]int{{1}, {}, {2}})
		want := []int{1, 2}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})

	t.Run("EmptyOuter", func(t *testing.T) {
		got := Flatten([][]int{})
		if len(got) != 0 {
			t.Errorf("Flatten() = %v, want empty slice", got)
		}
	})

	t.Run("NilOuter", func(t *testing.T) {
		got := Flatten([][]int(nil))
		if len(got) != 0 {
			t.Errorf("Flatten() = %v, want empty slice", got)
		}
	})

	t.Run("Strings", func(t *testing.T) {
		got := Flatten([][]string{{"a", "b"}, {"c"}})
		want := []string{"a", "b", "c"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})
}

func BenchmarkFlatten(b *testing.B) {
	data := make([][]int, 100)
	for i := range data {
		data[i] = genArray(i*100, (i+1)*100)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Flatten(data)
	}
}

func ExampleFlatten() {
	fmt.Println(Flatten([][]int{{1, 2}, {3, 4}, {5}}))
	// Output:
	// [1 2 3 4 5]
}
