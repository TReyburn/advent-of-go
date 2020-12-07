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

func TestDumbTraverse(t *testing.T) {
	count := DumbTraverse(testData, "#", 1, 3)

	if count != 7 {
		t.Error("Expected 7; got", count)
	}
}
