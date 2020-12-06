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
	counter := 0

	for _, pwd := range pws {
		b := pwd.Validate()
		if b == true {
			counter++
		}
	}
	fmt.Println("Valid passwords:", counter)
}
