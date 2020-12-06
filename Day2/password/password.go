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
			return count >= p.Low && count <= p.High
		}
	}
	return count >= p.Low && count <= p.High
}
