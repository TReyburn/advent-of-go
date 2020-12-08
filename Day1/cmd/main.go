package main

import (
	"fmt"
	"github.com/TReyburn/advent-of-go/Day1/intsearch"
	"github.com/TReyburn/advent-of-go/common/filehandler"
	"log"
	"sort"
	"strconv"
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
	var filePath string
	var sumValue string
	var numCount string

	fmt.Println("Input filepath")
	_, err := fmt.Scanln(&filePath)
	if err != nil {
		log.Fatalln("User input error", err)
	}

	fmt.Println("What sum value?")
	_, err = fmt.Scanln(&sumValue)
	if err != nil {
		log.Fatalln("User input error", err)
	}
	intSumValue, err := strconv.Atoi(sumValue)
	if err != nil {
		log.Fatalln("User input error", err)
	}

	fmt.Println("How many values to sum up?")
	_, err = fmt.Scanln(&numCount)
	if err != nil {
		log.Fatalln("User input error", err)
	}
	intNumCount, err := strconv.Atoi(numCount)
	if err != nil {
		log.Fatalln("User input error", err)
	}

	dm := intsearch.NewDataManager()
	err = filehandler.LoadInputFile(filePath, dm)
	if err != nil {
		log.Fatalln("Error reading file", err)
	}

	sort.Ints(dm.Data)
	sol, err := intsearch.IntSearch(dm.Data, intSumValue, intNumCount)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Solution: Values:", sol.Values, "Final Multiplied value:", sol.MultVal)
}