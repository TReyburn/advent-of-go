package validator

import "testing"

func TestByrValidatePos(t *testing.T) {
	v := "2002"

	res, err := byrValidate(v)
	if err != nil {
		t.Error("Unexpected error", err)
	}

	if res != true {
		t.Error("Failed to validate byr:2002")
	}
}

func TestByrValidateNeg(t *testing.T) {
	iv := "2003"

	res, err := byrValidate(iv)
	if err != nil {
		t.Error("Unexpected error", err)
	}

	if res != false {
		t.Error("Failed to invalidate byr:2003")
	}
}

func TestIyrValidatePos(t *testing.T) {
	v := "2010"

	res, err := iyrValidate(v)
	if err != nil {
		t.Error("Unexpected error", err)
	}

	if res != true {
		t.Error("Failed to validate iyr:2010")
	}
}

func TestIyrValidateMeh(t *testing.T) {
	iv := "2009"

	res, err := iyrValidate(iv)
	if err != nil {
		t.Error("Unexpected error", err)
	}

	if res != false {
		t.Error("Failed to invalidate iyr:2009")
	}
}

func TestEyrValidatePos(t *testing.T) {
	v := "2020"

	res, err := eyrValidate(v)
	if err != nil {
		t.Error("Unexpected error", err)
	}

	if res != true {
		t.Error("Failed to validate eyr:2020")
	}
}

func TestEyrValidateMeh(t *testing.T) {
	iv := "2031"

	res, err := eyrValidate(iv)
	if err != nil {
		t.Error("Unexpected error", err)
	}

	if res != false {
		t.Error("Failed to invalidate eyr:2031")
	}
}

func TestHgtValidatePosIN(t *testing.T) {
	v := "60in"

	res, err := hgtValidate(v)
	if err != nil {
		t.Error("Unexpected error", err)
	}

	if res != true {
		t.Error("Failed to validate hgt:60in")
	}
}

func TestHgtValidateNegIN(t *testing.T) {
	iv := "190in"

	res, err := hgtValidate(iv)
	if err != nil {
		t.Error("Unexpected error", err)
	}

	if res != false {
		t.Error("Failed to invalidate hgt:190in")
	}
}

func TestHgtValidatePosCM(t *testing.T) {
	v := "190cm"

	res, err := hgtValidate(v)
	if err != nil {
		t.Error("Unexpected error", err)
	}

	if res != true {
		t.Error("Failed to validate hgt:190cm")
	}
}

func TestHgtValidateNeg(t *testing.T) {
	iv := "190"

	res, err := hgtValidate(iv)
	if err != nil {
		t.Error("Unexpected error", err)
	}

	if res != false {
		t.Error("Failed to validate hgt:190")
	}
}

func TestHclValidatePos(t *testing.T) {
	v := "#123abc"

	res, err := hclValidate(v)
	if err != nil {
		t.Error("Unexpected error", err)
	}

	if res != true {
		t.Error("Failed to validate hcl:#123abc")
	}
}

func TestHclValidateNeg1(t *testing.T) {
	iv := "#123abz"

	res, err := hclValidate(iv)
	if err != nil {
		t.Error("Unexpected error", err)
	}

	if res != false {
		t.Error("Failed to invalidate hcl:#123abz")
	}
}

func TestHclValidateNeg2(t *testing.T) {
	iv := "123abc"

	res, err := hclValidate(iv)
	if err != nil {
		t.Error("Unexpected error", err)
	}

	if res != false {
		t.Error("Failed to invalidate hcl:123abc")
	}
}

func TestEclValidatePos(t *testing.T) {
	v := "brn"

	res, err := eclValidate(v)
	if err != nil {
		t.Error("Unexpected error", err)
	}

	if res != true {
		t.Error("Failed to validate ecl:brn")
	}
}

func TestEclValidateNeg(t *testing.T) {
	iv := "wat"

	res, err := eclValidate(iv)
	if err != nil {
		t.Error("Unexpected error", err)
	}

	if res != false {
		t.Error("Failed to invalidate ecl:wat")
	}
}

func TestPidValidatePos(t *testing.T) {
	v := "000000001"

	res, err := pidValidate(v)
	if err != nil {
		t.Error("Unexpected error", err)
	}

	if res != true {
		t.Error("Failed to validate pid:000000001")
	}
}

func TestPidValidateNeg1(t *testing.T) {
	iv := "0123456789"

	res, err := pidValidate(iv)
	if err != nil {
		t.Error("Unexpected error", err)
	}

	if res != false {
		t.Error("Failed to invalidate pid:0123456789")
	}
}

func TestPidValidateNeg2(t *testing.T) {
	iv := "0123a56789"

	res, err := pidValidate(iv)
	if err != nil {
		t.Error("Unexpected error", err)
	}

	if res != false {
		t.Error("Failed to invalidate pid:0123a56789")
	}
}