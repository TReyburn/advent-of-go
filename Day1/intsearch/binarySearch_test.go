package intsearch

import (
	"sort"
	"testing"
)

func TestIntSearchPos(t *testing.T) {
	ns := []int{1721, 979, 366, 299, 675, 1456}
	sort.Ints(ns)
	sv := 2020
	res, err := IntSearch(ns, sv, 2)

	if err != nil {
		t.Error("Shouldn't have gotten an error", err)
	}

	if res.MultVal != int64(514579) {
		t.Error("Expected MultVal of 514579; got", res.MultVal)
	}
}

func TestIntSearchNeg(t *testing.T) {
	ns := []int{1721, 979, 366, 299, 675, 1456}
	sort.Ints(ns)
	sv := 2021
	res, err := IntSearch(ns, sv, 2)

	if err == nil {
		t.Error("Shouldn have gotten an error; got solution of:", res)
	}
}

func TestIntSearchPosRecurse(t *testing.T) {
	ns := []int{1721, 979, 366, 299, 675, 1456}
	sort.Ints(ns)
	sv := 2020
	res, err := IntSearch(ns, sv, 3)

	if err != nil {
		t.Error("Shouldn't have gotten an error", err)
	}

	if res.MultVal != int64(241861950) {
		t.Error("Expected MultVal of 514579; got", res.MultVal)
	}
}

func TestIntSearchNegRecurse(t *testing.T) {
	ns := []int{1721, 979, 366, 299, 675, 1456}
	sort.Ints(ns)
	sv := 2021
	res, err := IntSearch(ns, sv, 3)

	if err == nil {
		t.Error("Shouldn have gotten an error; got solution of:", res)
	}
}

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