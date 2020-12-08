package boardingpass

import "testing"

func TestBinarySearchRows(t *testing.T) {
	c := "FBFBBFFRLR"
	res := binarySearch(c[:7], 0, 127, "F", "B")

	if res != 44 {
		t.Error("Expected row 44; got", res)
	}
}

func TestBinarySearchCols(t *testing.T) {
	c := "FBFBBFFRLR"
	res := binarySearch(c[7:], 0, 7, "L", "R")

	if res != 5 {
		t.Error("Expected col 5; got", res)
	}
}

func TestBoardingPass_DecodeCase1(t *testing.T) {
	bp := BoardingPass{
		Code: "BFFFBBFRRR",
	}
	bp.Decode()

	if bp.Row != 70 {
		t.Error("Expected row 70; got", bp.Row)
	}

	if bp.Column != 7 {
		t.Error("Expected col 7; got", bp.Column)
	}

	if bp.ID != 567 {
		t.Error("Expected id 567; got", bp.ID)
	}
}

func TestBoardingPass_DecodeCase2(t *testing.T) {
	bp := BoardingPass{
		Code: "FFFBBBFRRR",
	}
	bp.Decode()

	if bp.Row != 14 {
		t.Error("Expected row 14; got", bp.Row)
	}

	if bp.Column != 7 {
		t.Error("Expected col 7; got", bp.Column)
	}

	if bp.ID != 119 {
		t.Error("Expected id 119; got", bp.ID)
	}
}

func TestBoardingPass_DecodeCase3(t *testing.T) {
	bp := BoardingPass{
		Code: "BBFFBBFRLL",
	}
	bp.Decode()

	if bp.Row != 102 {
		t.Error("Expected row 102; got", bp.Row)
	}

	if bp.Column != 4 {
		t.Error("Expected col 4; got", bp.Column)
	}

	if bp.ID != 820 {
		t.Error("Expected id 820; got", bp.ID)
	}
}