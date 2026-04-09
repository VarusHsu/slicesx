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

// Contains reports whether v is present in s.
//
//	Contains([]int{1, 2, 3}, 2)  // true
//	Contains([]int{1, 2, 3}, 4)  // false
func Contains[S ~[]E, E comparable](s S, v E) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the first occurrence of v in s,
// or -1 if v is not present.
//
//	IndexOf([]string{"a", "b", "c"}, "b")  // 1
//	IndexOf([]string{"a", "b", "c"}, "d")  // -1
func IndexOf[S ~[]E, E comparable](s S, v E) int {
	for i, e := range s {
		if e == v {
			return i
		}
	}
	return -1
}
