package main

import (
	"fmt"
	"github.com/TReyburn/advent-of-go/Day10/adapter"
	"github.com/TReyburn/advent-of-go/common/filehandler"
	"log"
)

func main() {
	a := adapter.NewAdapter()
	err := filehandler.LoadInputFile("assets/input.txt", a)
	if err != nil {
		log.Fatalln("Unexpected error while loading file:", err)
	}

	res, err := a.Summarize()
	if err != nil {
		log.Fatalln("Unexpected error while summarizing:", err)
	}
	fmt.Println(res)
	fmt.Println("Jolts multiplies:", res[1]*res[3])
}
