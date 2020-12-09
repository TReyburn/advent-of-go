package main

import (
	"fmt"
	"github.com/TReyburn/advent-of-go/Day7/graph"
	"github.com/TReyburn/advent-of-go/common/filehandler"
	"log"
)

// Algo steps
// 1, load file as []string
// 2. iterate over each string
// 3. parse each string and load into graph
// 3. iterate over each graph node (skip shiny)
// 4. do breadth first search of shiny golden
func main() {
	g := graph.NewGraph()
	err := filehandler.LoadInputFile("assets/input.txt", g)
	if err != nil {
		log.Fatalln("Unexpected err:", err)
	}
	res := g.CountPossiblePaths("shiny gold bag")
	fmt.Println("Count of bags which can eventual contain a shiny gold bag:", res)
}
