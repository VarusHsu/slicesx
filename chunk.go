// Copyright (c) 2025 Varus Hsu
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

// Chunk splits a slice into multiple sub-slices using one or more step sizes.
// If multiple step values are provided, they will be used in order.
// After the steps are exhausted, the last step size is reused.
// A step size of zero is allowed (produces empty chunk), but negative steps will panic.
// The final step must be positive.
//
//	var array = []int{
//		1, 2, 3, 4, 5, 6, 7, 8, 9, 10
//	}
//	chunks := Chunk(array, 3)
//	// chunks: [[1, 2, 3], [4, 5, 6], [7, 8, 9], [10]]
//
//	chunks = Chunk(array, 1, 2, 3) // reuse last step
//	// chunks: [[1], [2, 3], [4, 5, 6], [7, 8, 9], [10]]
//
//	chunks = Chunk(array, 1, 0, 3)
//	// chunks: [[1], [], [2, 3, 4], [5, 6, 7], [8, 9, 10]]
//
//	chunks = Chunk(array, 1, -1, 3)
//	// panic("step should not be a negative")
//
//	chunks = Chunk(array, 1, 3, 0)
//	// panic("last step must be a positive")
func Chunk[S ~[]E, E any, N number](array S, step ...N) []S {

	if step[len(step)-1] <= 0 {
		panic("last step must be a positive")
	}

	var stepIndex = 0

	nextStep := func() int {
		defer func() {
			stepIndex++
		}()

		if stepIndex >= len(step) {
			return int(step[len(step)-1])
		}
		return int(step[stepIndex])
	}

	if len(array) == 0 {
		return []S{}
	}

	start, end := 0, 0
	var ns int
	for ns == 0 {
		ns = nextStep()
	}
	// Preallocate capacity based on first step
	chunks := make([]S, 0, len(array)/ns+1)
	stepIndex = 0 // reset to re-use steps from the beginning

	for {
		ns = nextStep()
		if ns < 0 {
			panic("step should not be a negative")
		}
		end += ns

		if end > len(array) {
			end = len(array)
		}

		if end == start && end == len(array) {
			break
		}

		chunks = append(chunks, array[start:end])
		start = end
	}
	return chunks
}
