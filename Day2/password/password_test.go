package password

import "testing"

func TestPassword_ValidatePos(t *testing.T) {
	pwd := Password{
		Required: "a",
		Low:      1,
		High:     3,
		Password: "abcde",
	}

	b := pwd.Validate()

	if b != true {
		t.Error("Expected true; got", b)
	}
}

func TestPassword_ValidateNeg(t *testing.T) {
	pwd := Password{
		Required: "b",
		Low:      1,
		High:     3,
		Password: "cdefg",
	}

	b := pwd.Validate()

	if b != false {
		t.Error("Expected false; got", b)
	}
}

func TestPassword_ValidatePosLong(t *testing.T) {
	pwd := Password{
		Required: "c",
		Low:      2,
		High:     9,
		Password: "ccccccccc",
	}

	b := pwd.Validate()

	if b != true {
		t.Error("Expected true; got", b)
	}
}
