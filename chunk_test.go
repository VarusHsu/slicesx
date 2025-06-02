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

// genArray generates a slice of ints from start (inclusive) to end (exclusive).
func genArray(start, end int) []int {
	array := make([]int, 0, end-start)
	for i := start; i < end; i++ {
		array = append(array, i)
	}
	return array
}
