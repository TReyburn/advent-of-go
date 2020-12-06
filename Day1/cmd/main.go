package main

import (
	"fmt"
	"github.com/TReyburn/advent-of-go/Day1/fileHandler"
	"log"
)

// Basic algo steps:
// 1. Load data as []byte (Can I leave the data as bytes and do byte math?)
// 2. Convert to []int (I can skip this if I do byte math)
// 3. Sort []int (what sorting algo?)
// 4. var MaxV = 2020 - (first value in []int)
// 5. for value in []int;
//		if value <= MaxV;
//		searchV := 2020 - value;
//		res := search([]int, searchV;
//		if res != nil found answer

func main() {
	intSlice, err := fileHandler.LoadFile("assets/input.txt")
	if err != nil {
		log.Fatalln("Error reading file", err)
	}
	fmt.Println(intSlice)

}