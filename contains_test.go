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

func TestContains(t *testing.T) {
	t.Run("Found", func(t *testing.T) {
		if !Contains([]int{1, 2, 3}, 2) {
			t.Error("Contains() = false, want true")
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		if Contains([]int{1, 2, 3}, 4) {
			t.Error("Contains() = true, want false")
		}
	})

	t.Run("Empty", func(t *testing.T) {
		if Contains([]int{}, 1) {
			t.Error("Contains() = true, want false")
		}
	})

	t.Run("Nil", func(t *testing.T) {
		if Contains([]int(nil), 1) {
			t.Error("Contains() = true, want false")
		}
	})

	t.Run("Strings", func(t *testing.T) {
		if !Contains([]string{"a", "b", "c"}, "b") {
			t.Error("Contains() = false, want true")
		}
	})

	t.Run("FirstElement", func(t *testing.T) {
		if !Contains([]int{1, 2, 3}, 1) {
			t.Error("Contains() = false, want true")
		}
	})

	t.Run("LastElement", func(t *testing.T) {
		if !Contains([]int{1, 2, 3}, 3) {
			t.Error("Contains() = false, want true")
		}
	})
}

func TestIndexOf(t *testing.T) {
	t.Run("Found", func(t *testing.T) {
		if got := IndexOf([]int{1, 2, 3}, 2); got != 1 {
			t.Errorf("IndexOf() = %v, want 1", got)
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		if got := IndexOf([]int{1, 2, 3}, 4); got != -1 {
			t.Errorf("IndexOf() = %v, want -1", got)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		if got := IndexOf([]int{}, 1); got != -1 {
			t.Errorf("IndexOf() = %v, want -1", got)
		}
	})

	t.Run("Nil", func(t *testing.T) {
		if got := IndexOf([]int(nil), 1); got != -1 {
			t.Errorf("IndexOf() = %v, want -1", got)
		}
	})

	t.Run("FirstOccurrence", func(t *testing.T) {
		if got := IndexOf([]int{1, 2, 2, 3}, 2); got != 1 {
			t.Errorf("IndexOf() = %v, want 1", got)
		}
	})

	t.Run("Strings", func(t *testing.T) {
		if got := IndexOf([]string{"a", "b", "c"}, "c"); got != 2 {
			t.Errorf("IndexOf() = %v, want 2", got)
		}
	})
}

func BenchmarkContains(b *testing.B) {
	array := genArray(0, 10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Contains(array, 9999)
	}
}

func BenchmarkIndexOf(b *testing.B) {
	array := genArray(0, 10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IndexOf(array, 9999)
	}
}

func ExampleContains() {
	fmt.Println(Contains([]int{1, 2, 3}, 2))
	fmt.Println(Contains([]int{1, 2, 3}, 4))
	// Output:
	// true
	// false
}

func ExampleIndexOf() {
	fmt.Println(IndexOf([]string{"a", "b", "c"}, "b"))
	fmt.Println(IndexOf([]string{"a", "b", "c"}, "d"))
	// Output:
	// 1
	// -1
}
