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