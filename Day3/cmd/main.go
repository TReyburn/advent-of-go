package main

import (
	"fmt"
	"github.com/TReyburn/advent-of-go/Day3/traverse"
	"github.com/TReyburn/advent-of-go/common/filehandler"
	"log"
)

func main() {
	data, err := filehandler.LoadDay3File("assets/input.txt")
	if err != nil {
		log.Fatalln("Error opening file", err)
	}
	res := traverse.Traverse(data, "#", 1, 3)
	fmt.Println("We hit", res, "trees")

	angles := [][]int{
		[]int{1,1},
		[]int{1,3},
		[]int{1,5},
		[]int{1,7},
		[]int{2,1},
	}
	fres := traverse.MultiTraverse(data, "#", angles)
	fmt.Println("Multitraverse result of", fres)
}
