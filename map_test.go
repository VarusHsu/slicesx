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
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	t.Run("Double", func(t *testing.T) {
		got := Map([]int{1, 2, 3}, func(v int) int { return v * 2 })
		want := []int{2, 4, 6}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Map() = %v, want %v", got, want)
		}
	})

	t.Run("IntToString", func(t *testing.T) {
		got := Map([]int{1, 2, 3}, func(v int) string { return strconv.Itoa(v) })
		want := []string{"1", "2", "3"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Map() = %v, want %v", got, want)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		got := Map([]int{}, func(v int) int { return v })
		if len(got) != 0 {
			t.Errorf("Map() = %v, want empty slice", got)
		}
	})

	t.Run("Nil", func(t *testing.T) {
		got := Map([]int(nil), func(v int) int { return v })
		if len(got) != 0 {
			t.Errorf("Map() = %v, want empty slice", got)
		}
	})

	t.Run("StringLength", func(t *testing.T) {
		got := Map([]string{"a", "bb", "ccc"}, func(v string) int { return len(v) })
		want := []int{1, 2, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Map() = %v, want %v", got, want)
		}
	})
}

func BenchmarkMap(b *testing.B) {
	array := genArray(0, 10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Map(array, func(v int) int { return v * 2 })
	}
}

func ExampleMap() {
	doubled := Map([]int{1, 2, 3}, func(v int) int {
		return v * 2
	})
	fmt.Println(doubled)
	// Output:
	// [2 4 6]
}
