package validator

import (
	"regexp"
	"strconv"
	"strings"
)

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

// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
func eyrValidate(eyr string) (bool, error) {
	if len(eyr) != 4 {
		return false, nil
	}
	nEyr, err := strconv.Atoi(eyr)
	if err != nil {
		return false, err
	}
	return nEyr >= 2020 && nEyr <= 2030, nil
}

// hgt (Height) - a number followed by either cm or in:
//    If cm, the number must be at least 150 and at most 193.
//    If in, the number must be at least 59 and at most 76.
func hgtValidate(hgt string) (bool, error) {
	cmb := strings.HasSuffix(hgt, "cm")
	inb := strings.HasSuffix(hgt, "in")
	if cmb {
		nHgt, err := strconv.Atoi(strings.TrimSuffix(hgt, "cm"))
		if err != nil {
			return false, err
		}
		return nHgt >= 150 && nHgt <= 193, nil
	}
	if inb {
		nHgt, err := strconv.Atoi(strings.TrimSuffix(hgt, "in"))
		if err != nil {
			return false, err
		}
		return nHgt >= 59 && nHgt <= 76, nil
	}
	return false, nil
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func hclValidate(hcl string) (bool, error) {
	res, err := regexp.MatchString("#[0-9a-f]{6}", hcl)
	if err != nil {
		return false, err
	}
	return res, nil
}

// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func eclValidate(hcl string) (bool, error) {
	vc := 0
	allowed := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, c := range allowed {
		if hcl == c {
			vc++
		}
	}
	return vc == 1, nil
}