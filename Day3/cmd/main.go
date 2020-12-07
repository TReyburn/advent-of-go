package main

import (
	"fmt"
	"github.com/TReyburn/advent-of-go/Day3/traverse"
	"github.com/TReyburn/advent-of-go/common/fileHandler"
	"log"
)

func main() {
	data, err := fileHandler.LoadDay3File("assets/input.txt")
	if err != nil {
		log.Fatalln("Error opening file", err)
	}
	res := traverse.Traverse(data, "#", 1, 3)
	fmt.Println("We hit", res, "trees")
}
