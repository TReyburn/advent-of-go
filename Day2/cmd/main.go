package main

import (
	"fmt"
	"github.com/TReyburn/advent-of-go/Day2/password"
	"github.com/TReyburn/advent-of-go/common/filehandler"
	"log"
)

func main() {
	pwm := password.NewPasswordManager()
	err := filehandler.LoadInputFile("assets/input.txt", pwm)
	if err != nil {
		log.Fatalln("Error loading passwords", err)
	}
	oRes, nRes := pwm.ValidatePasswords()
	fmt.Println("Valid passwords:", oRes)
	fmt.Println("Valid passwords:", nRes)
}
