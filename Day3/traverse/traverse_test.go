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

func TestDumbTraverseCase1(t *testing.T) {
	count := Traverse(testData, "#", 1, 3)

	if count != 7 {
		t.Error("Expected 7; got", count)
	}
}

func TestDumbTraverseCase2(t *testing.T) {
	count := Traverse(testData, "#", 1, 1)

	if count != 2 {
		t.Error("Expected 2; got", count)
	}
}

func TestDumbTraverseCase3(t *testing.T) {
	count := Traverse(testData, "#", 1, 5)

	if count != 3 {
		t.Error("Expected 3; got", count)
	}
}

func TestDumbTraverseCase4(t *testing.T) {
	count := Traverse(testData, "#", 1, 7)

	if count != 4 {
		t.Error("Expected 4; got", count)
	}
}

func TestDumbTraverseCase5(t *testing.T) {
	count := Traverse(testData, "#", 2, 1)

	if count != 2 {
		t.Error("Expected 2; got", count)
	}
}