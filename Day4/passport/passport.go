package passport

import (
	"bytes"
	"github.com/TReyburn/advent-of-go/Day4/validator"
	"strings"
)

type Passport struct {
	Content map[string]string
}

func (pp *Passport) addContent(k string, v string) {
	pp.Content[k] = v
}

func (pp *Passport) loadStringValues(v string) {
	split := strings.Split(v, " ")
	for _, pair := range split {
		kv := strings.Split(pair, ":")
		pp.addContent(kv[0], kv[1])
	}
}

func (pp Passport) Validate(req []string) bool {
	count := 0
	missed := 0
	for _, required := range req {
		if _, ok := pp.Content[required]; ok {
			count++
		} else if required != "cid" {
			missed++
		}
	}
	return missed == 0
}

func (pp Passport) ValidateData(v validator.Validator) bool {
	vm := v.Validators
	invalid := 0
	missed := 0
	for key, fnc := range vm {
		if val, ok := pp.Content[key]; ok {
			res, err := fnc(val)
			if err != nil {
				invalid++
			}
			if !res {
				invalid++
			}
		} else if key != "cid" {
			missed++
		}
	}
	return missed == 0 && invalid == 0
}

type PassportsScanner struct {
	Passports []Passport
}

func (ps PassportsScanner) ValidatePassports(req []string) int {
	vc := 0

	for _, pp := range ps.Passports {
		res := pp.Validate(req)
		if res {
			vc++
		}
	}
	return vc
}

func (ps PassportsScanner) ValidatePassportsData(v validator.Validator) int {
	vc := 0

	for _, pp := range ps.Passports {
		res := pp.ValidateData(v)
		if res {
			vc++
		}
	}
	return vc
}

func (ps *PassportsScanner) Write(p []byte) (n int, err error) {
	bss := bytes.Split(p, []byte{13})
	pp := NewPassport()
	rb := len(p)

	for _, bString := range bss {
		rawString := string(bString)
		rawString = strings.Trim(rawString, "\n")
		if rawString != "" {
			pp.loadStringValues(rawString)
		} else {
			ps.Passports = append(ps.Passports, *pp)
			pp = NewPassport()
		}
	}
	ps.Passports = append(ps.Passports, *pp)
	return rb, nil
}

func NewPassportScanner() *PassportsScanner {
	ps := PassportsScanner{}
	return &ps
}

func NewPassport() *Passport {
	pp := Passport{Content: make(map[string]string)}
	return &pp
}