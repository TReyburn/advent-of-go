package customs

import (
	"bytes"
	"strings"
)

type GroupVisa struct {
	Data map[string]bool
}

func (v *GroupVisa) loadStr(s string) {
	for _, char := range strings.Split(s, "") {
		v.Data[char] = true
	}
}

func (v GroupVisa) SumUnique() int {
	return len(v.Data)
}

func NewGroupVisa() *GroupVisa {
	v := GroupVisa{Data: make(map[string]bool)}
	return &v
}

type VisaScanner struct {
	Visas []GroupVisa
}

func (vs *VisaScanner) LoadVisa(gv GroupVisa) {
	vs.Visas = append(vs.Visas, gv)
}

func (vs VisaScanner) SumVisas() int {
	sum := 0
	for _, gv := range vs.Visas {
		sum += gv.SumUnique()
	}
	return sum
}

func (vs *VisaScanner) Write(p []byte) (int, error) {
	rb := len(p)
	gv := NewGroupVisa()
	bss := bytes.Split(p, []byte{13})
	for _, bString := range bss {
		rawString := string(bString)
		rawString = strings.Trim(rawString, "\n")
		if rawString != "" {
			gv.loadStr(rawString)
		} else {
			vs.LoadVisa(*gv)
			gv = NewGroupVisa()
		}
	}
	vs.LoadVisa(*gv)
	return rb, nil
}

func NewVisaScanner() *VisaScanner {
	vs := VisaScanner{Visas: make([]GroupVisa, 0)}
	return &vs
}