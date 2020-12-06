package password

import (
	"sort"
	"strings"
)

type Password struct {
	Required string
	Low int
	High int
	Password string
}

func (p Password) Validate() bool {
	str := strings.Split(p.Password, "")
	sort.Strings(str)
	count := 0
	for _, s := range str {
		if s == p.Required {
			count++
		}
		if count > 0 && s != p.Required {
			break
		}
	}
	return count >= p.Low && count <= p.High
}

func (p Password) NewValidate() bool {
	counter := 0
	fi := p.Low - 1
	si := p.High - 1
	pws := strings.Split(p.Password, "")
	if pws[fi] == p.Required {
		counter ++
	}
	if pws[si] == p.Required {
		counter++
	}
	return counter == 1
}
