package main

import (
	"fmt"
	"github.com/TReyburn/advent-of-go/common/fileHandler"
	"log"
)

func main() {
	data, err := fileHandler.LoadDay3File("assets/input.txt")
	if err != nil {
		log.Fatalln("Error opening file", err)
	}
	fmt.Println(data)
}
