package passport

import (
	"github.com/TReyburn/advent-of-go/Day4/validator"
	"github.com/TReyburn/advent-of-go/common/filehandler"
	"testing"
)

var required = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}

func TestPassport_AddContent(t *testing.T) {
	pp := NewPassport()
	pp.addContent("Test", "Value")
	pp.addContent("Test2", "Value2")
	v, ok := pp.Content["Test"]

	if !ok {
		t.Error("Could not get value from 'Test' key")
	}

	if v != "Value" {
		t.Error("Value set incorrectly. Expected 'Value'; got", v)
	}
}

func TestPassport_LoadStringValues(t *testing.T) {
	pp := NewPassport()
	pp.loadStringValues("ecl:gry pid:860033327")
	v1, ok1 := pp.Content["ecl"]
	v2, ok2 := pp.Content["pid"]

	if !ok1 {
		t.Error("Error retrieving value")
	}

	if v1 != "gry" {
		t.Error("Expected ecl:gry; got", v1)
	}

	if !ok2 {
		t.Error("Error retrieving value")
	}

	if v2 != "860033327" {
		t.Error("Expected pid:860033327; got", v1)
	}
}

func TestPassport_ValidateCase1(t *testing.T) {
	pp := NewPassport()
	pp.loadStringValues("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm")
	res := pp.Validate(required)

	if res != true {
		t.Error("Case 1 Validation failed")
	}
}

func TestPassport_ValidateCase2(t *testing.T) {
	pp := NewPassport()
	pp.loadStringValues("iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929")
	res := pp.Validate(required)

	if res != false {
		t.Error("Case 2 Validation passed but should have failed")
	}
}

func TestPassport_ValidateCase3(t *testing.T) {
	pp := NewPassport()
	pp.loadStringValues("hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm")
	res := pp.Validate(required)

	if res != true {
		t.Error("Case 3 Validation failed but should have passed - allowed to be missing CID as long as nothing else is missing")
	}
}

func TestPassport_ValidateCase4(t *testing.T) {
	pp := NewPassport()
	pp.loadStringValues("hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in")
	res := pp.Validate(required)

	if res != false {
		t.Error("Case 4 Validation passed but should have failed")
	}
}

func TestPassportsScanner_ValidatePassports(t *testing.T) {
	ps := NewPassportScanner()
	err := filehandler.LoadInputFile("testdata/input.txt", ps)
	if err != nil {
		t.Error("Unexpected error reading file:", err)
	}

	res := ps.ValidatePassports(required)
	if res != 2 {
		t.Error("Expected 2 valid; got", res)
	}
}

func TestPassportsScanner_Write(t *testing.T) {
	ps := NewPassportScanner()
	err := filehandler.LoadInputFile("testdata/input.txt", ps)
	if err != nil {
		t.Error("Unexpected error reading file:", err)
	}

	l := len(ps.Passports)
	if l != 4 {
		t.Error("Expected 4 passports; got", l)
	}
}

func TestPassport_ValidateDataPos1(t *testing.T) {
	v := validator.NewValidator()
	pp := NewPassport()
	pp.loadStringValues("pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980 hcl:#623a2f")
	res := pp.ValidateData(*v)

	if res != true {
		t.Error("Case 1 - expected true; got", res)
	}
}

func TestPassport_ValidateDataPos2(t *testing.T) {
	v := validator.NewValidator()
	pp := NewPassport()
	pp.loadStringValues("eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm")
	res := pp.ValidateData(*v)

	if res != true {
		t.Error("Case 2 - expected true; got", res)
	}
}

func TestPassport_ValidateDataPos3(t *testing.T) {
	v := validator.NewValidator()
	pp := NewPassport()
	pp.loadStringValues("hcl:#888785 hgt:164cm byr:2001 iyr:2015 cid:88 pid:545766238 ecl:hzl eyr:2022")
	res := pp.ValidateData(*v)

	if res != true {
		t.Error("Case 3 - expected true; got", res)
	}
}

func TestPassport_ValidateDataPos4(t *testing.T) {
	v := validator.NewValidator()
	pp := NewPassport()
	pp.loadStringValues("iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719")
	res := pp.ValidateData(*v)

	if res != true {
		t.Error("Case 4 - expected true; got", res)
	}
}

//func TestPassport_ValidateDataNeg1(t *testing.T) {
//	v := validator.NewValidator()
//	pp := NewPassport()
//	pp.loadStringValues("eyr:1972 cid:100 hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926")
//	res := pp.ValidateData(*v)
//
//	if res != false {
//		t.Error("Case 1 - expected false; got", res)
//	}
//}
//
//func TestPassport_ValidateDataNeg2(t *testing.T) {
//	v := validator.NewValidator()
//	pp := NewPassport()
//	pp.loadStringValues("iyr:2019 hcl:#602927 eyr:1967 hgt:170cm ecl:grn pid:012533040 byr:1946")
//	res := pp.ValidateData(*v)
//
//	if res != false {
//		t.Error("Case 2 - expected false; got", res)
//	}
//}
//
//func TestPassport_ValidateDataNeg3(t *testing.T) {
//	v := validator.NewValidator()
//	pp := NewPassport()
//	pp.loadStringValues("hcl:dab227 iyr:2012 ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277")
//	res := pp.ValidateData(*v)
//
//	if res != false {
//		t.Error("Case 3 - expected false; got", res)
//	}
//}
//
//func TestPassport_ValidateDataNeg4(t *testing.T) {
//	v := validator.NewValidator()
//	pp := NewPassport()
//	pp.loadStringValues("hgt:59cm ecl:zzz eyr:2038 hcl:74454a iyr:2023 pid:3556412378 byr:2007")
//	res := pp.ValidateData(*v)
//
//	if res != false {
//		t.Error("Case 4 - expected false; got", res)
//	}
//}