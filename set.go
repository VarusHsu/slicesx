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

// Intersection returns a new slice containing elements that are present
// in both a and b, preserving the order from a. Duplicate elements in a
// are included only once.
//
//	Intersection([]int{1, 2, 3, 4}, []int{3, 4, 5, 6})
//	// [3, 4]
func Intersection[S ~[]E, E comparable](a, b S) S {
	set := make(map[E]struct{}, len(b))
	for _, v := range b {
		set[v] = struct{}{}
	}
	seen := make(map[E]struct{})
	result := make(S, 0)
	for _, v := range a {
		if _, ok := set[v]; ok {
			if _, dup := seen[v]; !dup {
				seen[v] = struct{}{}
				result = append(result, v)
			}
		}
	}
	return result
}

// Difference returns a new slice containing elements that are in a
// but not in b, preserving the order from a. Duplicate elements in a
// are included only once.
//
//	Difference([]int{1, 2, 3, 4}, []int{3, 4, 5, 6})
//	// [1, 2]
func Difference[S ~[]E, E comparable](a, b S) S {
	set := make(map[E]struct{}, len(b))
	for _, v := range b {
		set[v] = struct{}{}
	}
	seen := make(map[E]struct{})
	result := make(S, 0)
	for _, v := range a {
		if _, ok := set[v]; !ok {
			if _, dup := seen[v]; !dup {
				seen[v] = struct{}{}
				result = append(result, v)
			}
		}
	}
	return result
}

// Union returns a new slice containing all unique elements from both
// a and b, preserving order (elements from a first, then new elements from b).
//
//	Union([]int{1, 2, 3}, []int{3, 4, 5})
//	// [1, 2, 3, 4, 5]
func Union[S ~[]E, E comparable](a, b S) S {
	seen := make(map[E]struct{}, len(a)+len(b))
	result := make(S, 0, len(a)+len(b))
	for _, v := range a {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	for _, v := range b {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}
