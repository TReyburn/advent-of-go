package main

import (
	"fmt"
	"github.com/TReyburn/advent-of-go/Day9/xmas"
	"github.com/TReyburn/advent-of-go/common/filehandler"
	"log"
)

func main() {
	d := xmas.NewDecoder(25)
	err := filehandler.LoadInputFile("assets/input.txt", d)
	if err != nil {
		log.Fatalln("Unexpected error loading file:", err)
	}

	res, err := d.Attack()
	if err != nil {
		log.Fatalln("Unexpected error while attacking cipher:", err)
	}

	fmt.Println("Found non-compliant int:", res)
}
