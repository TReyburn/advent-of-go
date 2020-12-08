package customs

import (
	"testing"
)

func TestGroupVisa_Load(t *testing.T) {
	gv := NewGroupVisa()
	str := "aabcad"

	gv.loadStr(str)

	if len(gv.Data) != 4 {
		t.Error("Expected len 4; got", len(gv.Data))
	}
}

func TestGroupVisa_SumCase1(t *testing.T) {
	gv := NewGroupVisa()
	str := "abc"

	gv.loadStr(str)

	if gv.Sum() != 3 {
		t.Error("Expected 3; got", gv.Sum())
	}
}

func TestGroupVisa_SumCase2(t *testing.T) {
	gv := NewGroupVisa()
	str := "aaa"

	gv.loadStr(str)

	if gv.Sum() != 1 {
		t.Error("Expected 1; got", gv.Sum())
	}
}

func TestVisaScanner_SumVisas(t *testing.T) {
	vs := NewVisaScanner()
	gv := NewGroupVisa()
	loadStrs := []string{"abc", "abc", "abac", "aaaa", "b"}
	for _, str := range loadStrs {
		gv.loadStr(str)
		vs.LoadVisa(*gv)
		gv = NewGroupVisa()
	}

	if vs.SumVisas() != 11 {
		t.Error("Expected 11; got", vs.SumVisas())
	}
}