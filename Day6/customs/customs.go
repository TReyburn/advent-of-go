package customs

import (
	"bytes"
	"strings"
)

type GroupVisa struct {
	Data map[string]bool
}

func (v *GroupVisa) Load(k string) {
	v.Data[k] = true
}

func (v GroupVisa) Sum() int {
	return len(v.Data)
}

func NewGroupVisa() *GroupVisa {
	v := GroupVisa{Data: make(map[string]bool)}
	return &v
}

type VisaScanner struct {
	Visas []GroupVisa
}

func (vs *VisaScanner) Write(p []byte) (int, error) {
	rb := len(p)
	gv := NewGroupVisa()
	bss := bytes.Split(p, []byte{13})
	for _, bString := range bss {
		rawString := string(bString)
		rawString = strings.Trim(rawString, "\n")
		if rawString != "" {
			for _, char := range strings.Split(rawString, "") {
				gv.Load(char)
			}
		} else {
			vs.Visas = append(vs.Visas, *gv)
			gv = NewGroupVisa()
		}
	}
	return rb, nil
}

func NewVisaScanner() *VisaScanner {
	vs := VisaScanner{Visas: make([]GroupVisa, 0)}
	return &vs
}