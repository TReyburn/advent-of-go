package main

import (
	"fmt"
	"github.com/TReyburn/advent-of-go/Day5/boardingpass"
	"github.com/TReyburn/advent-of-go/common/filehandler"
	"log"
)

func main() {
	bpm := boardingpass.NewBoardingPassManager()
	err := filehandler.LoadInputFile("assets/input.txt", bpm)
	if err != nil {
		log.Fatalln("Unexpected error loading file", err)
	}
	fmt.Println(bpm.Passes)
}
