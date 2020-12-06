package binarySearch

import "testing"

func TestBinarySearchEvenArrayLow(t *testing.T) {
	ns := []int{1, 5, 7, 11, 13, 14}
	s := 5

	b, l := binarySearch(ns, s)

	if b != true {
		t.Error("Expected true; got", b)
	}

	if l != 1 {
		t.Error("Expected 1; got", l)
	}
}

func TestBinarySearchEvenArrayHigh(t *testing.T) {
	ns := []int{1, 5, 7, 11, 13, 14}
	s := 13

	b, l := binarySearch(ns, s)

	if b != true {
		t.Error("Expected true; got", b)
	}

	if l != 4 {
		t.Error("Expected 1; got", l)
	}
}

func TestBinarySearchEvenArrayInBoundsNoValueHigh(t *testing.T) {
	ns := []int{1, 5, 7, 11, 13, 20}
	s := 15

	b, l := binarySearch(ns, s)

	if b != false {
		t.Error("Expected false; got", b)
	}

	if l != 0 {
		t.Error("Expected 0; got", l)
	}
}

func TestBinarySearchEvenArrayInBoundsNoValueLow(t *testing.T) {
	ns := []int{1, 5, 7, 11, 13, 20}
	s := 2

	b, l := binarySearch(ns, s)

	if b != false {
		t.Error("Expected false; got", b)
	}

	if l != 0 {
		t.Error("Expected 0; got", l)
	}
}

func TestBinarySearchEvenArrayOutBoundsNoValueHigh(t *testing.T) {
	ns := []int{1, 5, 7, 11, 13, 20}
	s := 25

	b, l := binarySearch(ns, s)

	if b != false {
		t.Error("Expected false; got", b)
	}

	if l != 0 {
		t.Error("Expected 0; got", l)
	}
}

func TestBinarySearchEvenArrayOutBoundsNoValueLow(t *testing.T) {
	ns := []int{1, 5, 7, 11, 13, 20}
	s := -1

	b, l := binarySearch(ns, s)

	if b != false {
		t.Error("Expected false; got", b)
	}

	if l != 0 {
		t.Error("Expected 0; got", l)
	}
}

func TestBinarySearchOddArrayLow(t *testing.T) {
	ns := []int{1, 5, 7, 11, 13, 14, 20}
	s := 5

	b, l := binarySearch(ns, s)

	if b != true {
		t.Error("Expected true; got", b)
	}

	if l != 1 {
		t.Error("Expected 1; got", l)
	}
}

func TestBinarySearchOddArrayHigh(t *testing.T) {
	ns := []int{1, 5, 7, 11, 13, 14, 20}
	s := 14

	b, l := binarySearch(ns, s)

	if b != true {
		t.Error("Expected true; got", b)
	}

	if l != 5 {
		t.Error("Expected 5; got", l)
	}
}

func TestBinarySearchOddArrayInBoundsNoValueHigh(t *testing.T) {
	ns := []int{1, 5, 7, 11, 13, 14, 20}
	s := 15

	b, l := binarySearch(ns, s)

	if b != false {
		t.Error("Expected false; got", b)
	}

	if l != 0 {
		t.Error("Expected 0; got", l)
	}
}

func TestBinarySearchOddArrayInBoundsNoValueLow(t *testing.T) {
	ns := []int{1, 5, 7, 11, 13, 14, 20}
	s := 2

	b, l := binarySearch(ns, s)

	if b != false {
		t.Error("Expected false; got", b)
	}

	if l != 0 {
		t.Error("Expected 0; got", l)
	}
}

func TestBinarySearchOddArrayOutBoundsNoValueHigh(t *testing.T) {
	ns := []int{1, 5, 7, 11, 13, 14, 20}
	s := 25

	b, l := binarySearch(ns, s)

	if b != false {
		t.Error("Expected false; got", b)
	}

	if l != 0 {
		t.Error("Expected 0; got", l)
	}
}

func TestBinarySearchOddArrayOutBoundsNoValueLow(t *testing.T) {
	ns := []int{1, 5, 7, 11, 13, 14, 20}
	s := -1

	b, l := binarySearch(ns, s)

	if b != false {
		t.Error("Expected false; got", b)
	}

	if l != 0 {
		t.Error("Expected 0; got", l)
	}
}