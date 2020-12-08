package main

import (
	"fmt"
	"github.com/TReyburn/advent-of-go/Day6/customs"
	"github.com/TReyburn/advent-of-go/common/filehandler"
	"log"
)

func main() {
	vs := customs.NewVisaScanner()
	err := filehandler.LoadInputFile("assets/input.txt", vs)
	if err != nil {
		log.Fatalln("Unexpected error loading file:", err)
	}
	fmt.Println("Sum of Visa counts (old):", vs.SumVisas())
}
