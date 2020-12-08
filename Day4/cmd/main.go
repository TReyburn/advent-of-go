package main

import (
	"fmt"
	"github.com/TReyburn/advent-of-go/Day4/passport"
	"github.com/TReyburn/advent-of-go/Day4/validator"
	"github.com/TReyburn/advent-of-go/common/filehandler"
	"log"
)

func main() {
	var required = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
	ps := passport.NewPassportScanner()
	err := filehandler.LoadInputFile("assets/input.txt", ps)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	res := ps.ValidatePassports(required)
	fmt.Println("Valid Passports with just field validation", res)

	v := validator.NewValidator()
	dres := ps.ValidatePassportsData(*v)
	fmt.Println("Valid Passports with data validation", dres)
}
