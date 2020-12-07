package main

import (
	"fmt"
	"github.com/TReyburn/advent-of-go/common/filehandler"
	"log"
)

func main() {
	pws, err := filehandler.LoadDay2File("assets/input.txt")
	if err != nil {
		log.Fatalln("Error loading passwords", err)
	}
	oldCounter := 0

	for _, pwd := range pws {
		b := pwd.Validate()
		if b == true {
			oldCounter++
		}
	}
	fmt.Println("Valid passwords:", oldCounter)

	newCounter := 0

	for _, pwd := range pws {
		b := pwd.NewValidate()
		if b == true {
			newCounter++
		}
	}
	fmt.Println("Valid passwords:", newCounter)
}
