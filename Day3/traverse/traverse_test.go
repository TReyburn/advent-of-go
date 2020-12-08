package traverse

import "testing"

var testData = []string{
	"..##.......",
	"#...#...#..",
	".#....#..#.",
	"..#.#...#.#",
	".#...##..#.",
	"..#.##.....",
	".#.#.#....#",
	".#........#",
	"#.##...#...",
	"#...##....#",
	".#..#...#.#",
}

func TestMultiTraverse(t *testing.T) {
	angles := [][]int{
		[]int{1,1},
		[]int{1,3},
		[]int{1,5},
		[]int{1,7},
		[]int{2,1},
	}
	res := MultiTraverse(testData, "#", angles)

	if res != 336 {
		t.Error("Expected result of 336; got", res)
	}
}

func TestTraverseCase1(t *testing.T) {
	count := Traverse(testData, "#", 1, 3)

	if count != 7 {
		t.Error("Expected 7; got", count)
	}
}

func TestTraverseCase2(t *testing.T) {
	count := Traverse(testData, "#", 1, 1)

	if count != 2 {
		t.Error("Expected 2; got", count)
	}
}

func TestTraverseCase3(t *testing.T) {
	count := Traverse(testData, "#", 1, 5)

	if count != 3 {
		t.Error("Expected 3; got", count)
	}
}

func TestTraverseCase4(t *testing.T) {
	count := Traverse(testData, "#", 1, 7)

	if count != 4 {
		t.Error("Expected 4; got", count)
	}
}

func TestTraverseCase5(t *testing.T) {
	count := Traverse(testData, "#", 2, 1)

	if count != 2 {
		t.Error("Expected 2; got", count)
	}
}