package main

import (
	"fmt"
	"github.com/TReyburn/advent-of-go/common/fileHandler"
	"log"
)

func main() {
	pws, err := fileHandler.LoadDay2File("assets/input.txt")
	if err != nil {
		log.Fatalln("Error loading passwords", err)
	}
	fmt.Println(pws)
}
