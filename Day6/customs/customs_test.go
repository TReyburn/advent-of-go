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

func TestGroupVisa_SumUniqueCase1(t *testing.T) {
	gv := NewGroupVisa()
	str := "abc"

	gv.loadStr(str)

	if gv.SumUnique() != 3 {
		t.Error("Expected 3; got", gv.SumUnique())
	}
}

func TestGroupVisa_SumUniqueCase2(t *testing.T) {
	gv := NewGroupVisa()
	str := "aaa"

	gv.loadStr(str)

	if gv.SumUnique() != 1 {
		t.Error("Expected 1; got", gv.SumUnique())
	}
}

func TestVisaScanner_SumVisasUniqueVisas(t *testing.T) {
	vs := NewVisaScanner()
	gv := NewGroupVisa()
	loadStrs := []string{"abc", "abc", "abac", "aaaa", "b"}
	for _, str := range loadStrs {
		gv.loadStr(str)
		vs.LoadVisa(*gv)
		gv = NewGroupVisa()
	}

	if vs.SumVisasUnique() != 11 {
		t.Error("Expected 11; got", vs.SumVisasUnique())
	}
}

func TestGroupVisa_SumCommonCase1(t *testing.T) {
	gv := NewGroupVisa()
	str := "abc"

	gv.loadStr(str)

	if gv.SumCommon() != 3 {
		t.Error("Expected 3; got", gv.SumUnique())
	}
}

func TestGroupVisa_SumCommonCase2(t *testing.T) {
	gv := NewGroupVisa()
	strS := []string{"a", "a", "a", "a"}

	for _, str := range strS {
		gv.loadStr(str)
	}

	if gv.SumUnique() != 1 {
		t.Error("Expected 1; got", gv.SumUnique())
	}
}

func TestVisaScanner_SumVisasCommon(t *testing.T) {
	vs := NewVisaScanner()
	gv := NewGroupVisa()
	loadStrs := [][]string{{"abc"}, {"a", "b", "c"}, {"ab", "ac"}, {"a", "a", "a", "a"}, {"b"}}
	for _, strS := range loadStrs {
		for _, str := range strS {
			gv.loadStr(str)
		}
		vs.LoadVisa(*gv)
		gv = NewGroupVisa()
	}

	if vs.SumVisasCommon() != 6 {
		t.Error("Expected 6; got", vs.SumVisasUnique())
	}
}