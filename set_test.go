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

func TestIntersection(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		got := Intersection([]int{1, 2, 3, 4}, []int{3, 4, 5, 6})
		want := []int{3, 4}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Intersection() = %v, want %v", got, want)
		}
	})

	t.Run("NoOverlap", func(t *testing.T) {
		got := Intersection([]int{1, 2}, []int{3, 4})
		if len(got) != 0 {
			t.Errorf("Intersection() = %v, want empty slice", got)
		}
	})

	t.Run("FullOverlap", func(t *testing.T) {
		got := Intersection([]int{1, 2, 3}, []int{1, 2, 3})
		want := []int{1, 2, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Intersection() = %v, want %v", got, want)
		}
	})

	t.Run("EmptyFirst", func(t *testing.T) {
		got := Intersection([]int{}, []int{1, 2, 3})
		if len(got) != 0 {
			t.Errorf("Intersection() = %v, want empty slice", got)
		}
	})

	t.Run("EmptySecond", func(t *testing.T) {
		got := Intersection([]int{1, 2, 3}, []int{})
		if len(got) != 0 {
			t.Errorf("Intersection() = %v, want empty slice", got)
		}
	})

	t.Run("DuplicatesInFirst", func(t *testing.T) {
		got := Intersection([]int{1, 1, 2, 2, 3}, []int{1, 2})
		want := []int{1, 2}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Intersection() = %v, want %v", got, want)
		}
	})

	t.Run("Strings", func(t *testing.T) {
		got := Intersection([]string{"a", "b", "c"}, []string{"b", "c", "d"})
		want := []string{"b", "c"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Intersection() = %v, want %v", got, want)
		}
	})
}

func TestDifference(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		got := Difference([]int{1, 2, 3, 4}, []int{3, 4, 5, 6})
		want := []int{1, 2}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Difference() = %v, want %v", got, want)
		}
	})

	t.Run("NoOverlap", func(t *testing.T) {
		got := Difference([]int{1, 2}, []int{3, 4})
		want := []int{1, 2}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Difference() = %v, want %v", got, want)
		}
	})

	t.Run("FullOverlap", func(t *testing.T) {
		got := Difference([]int{1, 2, 3}, []int{1, 2, 3})
		if len(got) != 0 {
			t.Errorf("Difference() = %v, want empty slice", got)
		}
	})

	t.Run("EmptyFirst", func(t *testing.T) {
		got := Difference([]int{}, []int{1, 2, 3})
		if len(got) != 0 {
			t.Errorf("Difference() = %v, want empty slice", got)
		}
	})

	t.Run("EmptySecond", func(t *testing.T) {
		got := Difference([]int{1, 2, 3}, []int{})
		want := []int{1, 2, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Difference() = %v, want %v", got, want)
		}
	})

	t.Run("DuplicatesInFirst", func(t *testing.T) {
		got := Difference([]int{1, 1, 2, 2, 3}, []int{2})
		want := []int{1, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Difference() = %v, want %v", got, want)
		}
	})
}

func TestUnion(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		got := Union([]int{1, 2, 3}, []int{3, 4, 5})
		want := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Union() = %v, want %v", got, want)
		}
	})

	t.Run("NoOverlap", func(t *testing.T) {
		got := Union([]int{1, 2}, []int{3, 4})
		want := []int{1, 2, 3, 4}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Union() = %v, want %v", got, want)
		}
	})

	t.Run("FullOverlap", func(t *testing.T) {
		got := Union([]int{1, 2, 3}, []int{1, 2, 3})
		want := []int{1, 2, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Union() = %v, want %v", got, want)
		}
	})

	t.Run("EmptyFirst", func(t *testing.T) {
		got := Union([]int{}, []int{1, 2, 3})
		want := []int{1, 2, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Union() = %v, want %v", got, want)
		}
	})

	t.Run("EmptyBoth", func(t *testing.T) {
		got := Union([]int{}, []int{})
		if len(got) != 0 {
			t.Errorf("Union() = %v, want empty slice", got)
		}
	})

	t.Run("DuplicatesInBoth", func(t *testing.T) {
		got := Union([]int{1, 1, 2}, []int{2, 3, 3})
		want := []int{1, 2, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Union() = %v, want %v", got, want)
		}
	})

	t.Run("Strings", func(t *testing.T) {
		got := Union([]string{"a", "b"}, []string{"b", "c"})
		want := []string{"a", "b", "c"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Union() = %v, want %v", got, want)
		}
	})
}

func BenchmarkIntersection(b *testing.B) {
	a := genArray(0, 5000)
	bSlice := genArray(2500, 7500)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Intersection(a, bSlice)
	}
}

func BenchmarkDifference(b *testing.B) {
	a := genArray(0, 5000)
	bSlice := genArray(2500, 7500)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Difference(a, bSlice)
	}
}

func BenchmarkUnion(b *testing.B) {
	a := genArray(0, 5000)
	bSlice := genArray(2500, 7500)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Union(a, bSlice)
	}
}

func ExampleIntersection() {
	fmt.Println(Intersection([]int{1, 2, 3, 4}, []int{3, 4, 5, 6}))
	// Output:
	// [3 4]
}

func ExampleDifference() {
	fmt.Println(Difference([]int{1, 2, 3, 4}, []int{3, 4, 5, 6}))
	// Output:
	// [1 2]
}

func ExampleUnion() {
	fmt.Println(Union([]int{1, 2, 3}, []int{3, 4, 5}))
	// Output:
	// [1 2 3 4 5]
}
