package customs

import (
	"bytes"
	"strings"
)

type GroupVisa struct {
	Data map[string]int
	Count int
}

func (v *GroupVisa) loadStr(s string) {
	for _, char := range strings.Split(s, "") {
		if val, ok := v.Data[char]; ok {
			v.Data[char] = val + 1
		} else {
			v.Data[char] = 1
		}
	}
	v.Count++
}

func (v GroupVisa) SumUnique() int {
	return len(v.Data)
}

func (v GroupVisa) SumCommon() int {
	common := 0
	for _, val := range v.Data {
		if val == v.Count {
			common++
		}
	}
	return common
}

func NewGroupVisa() *GroupVisa {
	v := GroupVisa{Data: make(map[string]int), Count: 0}
	return &v
}

type VisaScanner struct {
	Visas []GroupVisa
}

func (vs *VisaScanner) LoadVisa(gv GroupVisa) {
	vs.Visas = append(vs.Visas, gv)
}

func (vs VisaScanner) SumVisasUnique() int {
	sum := 0
	for _, gv := range vs.Visas {
		sum += gv.SumUnique()
	}
	return sum
}

func (vs VisaScanner) SumVisasCommon() int {
	sum := 0
	for _, gv := range vs.Visas {
		sum += gv.SumCommon()
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