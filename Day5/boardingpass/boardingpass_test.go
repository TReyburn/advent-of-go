package boardingpass

import "testing"

func TestParseRow(t *testing.T) {
	c := "FBFBBFF"
	res := parseRow(c)

	if res != 44 {
		t.Error("Expected row 44; got", res)
	}
}
