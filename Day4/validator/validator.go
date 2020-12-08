package validator

import "strconv"

type Validator struct {
	Validators map[string]func(string)(bool, error)
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
func byrValidate(byr string) (bool, error) {
	if len(byr) != 4 {
		return false, nil
	}
	nByr, err := strconv.Atoi(byr)
	if err != nil {
		return false, err
	}
	return nByr >= 1920 && nByr <= 2002, nil
}

// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
func iyrValidate(iyr string) (bool, error) {
	if len(iyr) != 4 {
		return false, nil
	}
	nIyr, err := strconv.Atoi(iyr)
	if err != nil {
		return false, err
	}
	return nIyr >= 2010 && nIyr <= 2020, nil
}
