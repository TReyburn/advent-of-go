package passport

import (
	"bytes"
	"fmt"
	"strings"
)

type Passport struct {
	Content map[string]string
}

func (pp *Passport) AddContent(k string, v string) {
	pp.Content[k] = v
}

func (pp Passport) Validate(req []string) bool {
	count := 0
	for _, required := range req {
		if _, ok := pp.Content[required]; ok {
			count++
		}
	}
	return count == len(req)
}

type PassportsScanner struct {
	Passports []Passport
}

func (ps *PassportsScanner) Write(p []byte) (n int, err error) {
	bss := bytes.Split(p, []byte{13})
	pp := Passport{Content: make(map[string]string, 0)}
	rb := len(p)

	for _, bString := range bss {
		rawString := string(bString)
		fmt.Println("RawString:", rawString)
		rawString = strings.Trim(rawString, "\n")
		if rawString != "" {
			splitS := strings.Split(rawString, " ")
			for _, pair := range splitS {
				kv := strings.Split(pair, ":")
				pp.AddContent(kv[0], kv[1])
			}
		} else {
			ps.Passports = append(ps.Passports, pp)
			pp = Passport{Content: make(map[string]string, 0)}
		}
	}
	return rb, nil
}

func NewPassportScanner() *PassportsScanner {
	ps := PassportsScanner{}
	return &ps
}