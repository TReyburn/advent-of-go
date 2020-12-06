package main

import (
	"fmt"
	"github.com/TReyburn/advent-of-go/Day1/fileHandler"
	"github.com/TReyburn/advent-of-go/Day1/intSearch"
	"log"
	"sort"
)

// Basic algo steps:
// 1. Load data as []byte
// 2. Convert to []int
// 3. Sort []int
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
	sort.Ints(intSlice)
	sol, err := intSearch.IntSearch(intSlice, 2020)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Solution: Values:", sol.N1, sol.N2, "Final Multiplied value:", sol.MultVal)
}