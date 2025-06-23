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

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	t.Run("Step=100", func(t *testing.T) {
		original := genArray(0, 1000)
		expect := [][]int{
			genArray(0, 100),
			genArray(100, 200),
			genArray(200, 300),
			genArray(300, 400),
			genArray(400, 500),
			genArray(500, 600),
			genArray(600, 700),
			genArray(700, 800),
			genArray(800, 900),
			genArray(900, 1000),
		}
		if !reflect.DeepEqual(Chunk(original, 100), expect) {
			t.Errorf("Step=100 failed: got %v, want %v", Chunk(original, 100), expect)
		}
	})

	t.Run("NotDivisible", func(t *testing.T) {
		original := genArray(0, 1020)
		expect := [][]int{
			genArray(0, 100),
			genArray(100, 200),
			genArray(200, 300),
			genArray(300, 400),
			genArray(400, 500),
			genArray(500, 600),
			genArray(600, 700),
			genArray(700, 800),
			genArray(800, 900),
			genArray(900, 1000),
			genArray(1000, 1020),
		}
		if !reflect.DeepEqual(Chunk(original, 100), expect) {
			t.Errorf("NotDivisible failed: got %v, want %v", Chunk(original, 100), expect)
		}
	})

	t.Run("MultipleSteps", func(t *testing.T) {
		original := genArray(0, 1020)
		expect := [][]int{
			{0},
			{1, 2},
			{3, 4, 5},
			{6, 7, 8, 9},
			{10, 11, 12, 13, 14},
			genArray(15, 115),
			genArray(115, 215),
			genArray(215, 315),
			genArray(315, 415),
			genArray(415, 515),
			genArray(515, 615),
			genArray(615, 715),
			genArray(715, 815),
			genArray(815, 915),
			genArray(915, 1015),
			genArray(1015, 1020),
		}
		if !reflect.DeepEqual(Chunk(original, 1, 2, 3, 4, 5, 100), expect) {
			t.Errorf("MultipleSteps failed: got %v, want %v", Chunk(original, 1, 2, 3, 4, 5, 100), expect)
		}
	})

	t.Run("ZeroStep", func(t *testing.T) {
		if !reflect.DeepEqual(Chunk([]int{1}, 0, 100), [][]int{{}, {1}}) {
			t.Errorf("ZeroStep case 1 failed")
		}
		if !reflect.DeepEqual(Chunk([]int{1, 2, 3, 4, 5}, 1, 0, 1, 0, 1), [][]int{{1}, {}, {2}, {}, {3}, {4}, {5}}) {
			t.Errorf("ZeroStep case 2 failed")
		}
	})

	t.Run("MixedSteps", func(t *testing.T) {
		expected := [][]int{{1, 2}, {}, {}, {3, 4}, {5, 6}, {7, 8}, {9}}
		if !reflect.DeepEqual(Chunk([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2, 0, 0, 2), expected) {
			t.Errorf("MixedSteps failed")
		}
	})

	t.Run("NormalCase", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		expected := [][]int{{1}, {}, {2, 3, 4}, {5, 6, 7}, {8, 9, 10}}
		if !reflect.DeepEqual(Chunk(array, 1, 0, 3), expected) {
			t.Errorf("NormalCase failed")
		}
	})
}

func TestChunkPanicCases(t *testing.T) {
	t.Run("LastStepZero", func(t *testing.T) {
		defer func() {
			if v := recover(); v != "last step must be a positive" {
				t.Errorf("Expected panic 'last step must be a positive', got %v", v)
			}
		}()
		Chunk([]int{1}, 0)
	})

	t.Run("LastStepNegative", func(t *testing.T) {
		defer func() {
			if v := recover(); v != "last step must be a positive" {
				t.Errorf("Expected panic 'last step must be a positive', got %v", v)
			}
		}()
		Chunk([]int{1}, -1)
	})

	t.Run("NegativeStepInMiddle", func(t *testing.T) {
		defer func() {
			if v := recover(); v != "step should not be a negative" {
				t.Errorf("Expected panic 'step should not be a negative', got %v", v)
			}
		}()
		Chunk([]int{1, 2, 3, 4, 5}, 1, 0, -1, 1)
	})
}

func TestChunkZeroLen(t *testing.T) {
	t.Run("ZeroLen1", func(t *testing.T) {
		var array []int
		excepted := make([][]int, 0)
		if !reflect.DeepEqual(Chunk(array, 1), excepted) {
			t.Errorf("ZeroLen1 failed")
		}
	})

	t.Run("ZeroLen2", func(t *testing.T) {
		var array = make([]int, 0)
		excepted := make([][]int, 0)
		if !reflect.DeepEqual(Chunk(array, 1), excepted) {
			t.Errorf("ZeroLen2 failed")
		}
	})
}

func TestChunkCapacityAndLength(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	tests := []struct {
		name       string
		steps      []int
		wantChunks [][]int
		wantCapGE  bool
	}{
		{
			name:       "Single step 3",
			steps:      []int{3},
			wantChunks: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10}},
			wantCapGE:  true,
		},
		{
			name:       "Multiple steps 1,2,3",
			steps:      []int{1, 2, 3},
			wantChunks: [][]int{{1}, {2, 3}, {4, 5, 6}, {7, 8, 9}, {10}},
			wantCapGE:  true,
		},
		{
			name:       "Step with zero",
			steps:      []int{1, 0, 3},
			wantChunks: [][]int{{1}, {}, {2, 3, 4}, {5, 6, 7}, {8, 9, 10}},
			wantCapGE:  true,
		},
		{
			name:       "Single element steps",
			steps:      []int{1},
			wantChunks: [][]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}, {10}},
			wantCapGE:  true,
		},
		{
			name:       "Step larger than array",
			steps:      []int{20},
			wantChunks: [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			wantCapGE:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Chunk(array, tt.steps...)

			if len(got) != len(tt.wantChunks) {
				t.Errorf("Chunk() got length %d, want %d", len(got), len(tt.wantChunks))
			}

			for i := range got {
				if !reflect.DeepEqual(got[i], tt.wantChunks[i]) {
					t.Errorf("Chunk() chunk[%d] = %v, want %v", i, got[i], tt.wantChunks[i])
				}
			}

			if tt.wantCapGE && cap(got) != len(got) {
				t.Errorf("Chunk() result capacity %d < length %d", cap(got), len(got))
			}
		})
	}
}

// genArray generates a slice of ints from start (inclusive) to end (exclusive).
func genArray(start, end int) []int {
	array := make([]int, 0, end-start)
	for i := start; i < end; i++ {
		array = append(array, i)
	}
	return array
}
