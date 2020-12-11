package main

import (
	"fmt"
	"github.com/TReyburn/advent-of-go/Day8/console"
	"github.com/TReyburn/advent-of-go/common/filehandler"
	"log"
)

func main() {
	c := console.NewConsole()
	err := filehandler.LoadInputFile("assets/input.txt", c)
	if err != nil {
		log.Fatalln("Unexpected error reading file:", err)
	}
	res := c.Run()
	fmt.Println("Accumulator amount before repeat:", res)

	c1 := console.NewConsole()
	err = filehandler.LoadInputFile("assets/input.txt", c1)
	if err != nil {
		log.Fatalln("Unexpected error reading file:", err)
	}
	res2 := c1.DFSDebug()
	fmt.Println("Accumulator after fixed bug:", res2)
}
