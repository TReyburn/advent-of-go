package passport

import (
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
	pp1 := NewPassport()
	pp1.loadStringValues("hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm")
	res := pp1.Validate(required)

	if res != true {
		t.Error("Case 3 Validation failed but should have passed - allowed to be missing CID as long as nothing else is missing")
	}
}

func TestPassport_ValidateCase4(t *testing.T) {
	pp1 := NewPassport()
	pp1.loadStringValues("hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in")
	res := pp1.Validate(required)

	if res != false {
		t.Error("Case 4 Validation passed but should have failed")
	}
}