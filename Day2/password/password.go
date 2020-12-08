package password

import (
	"bytes"
	"sort"
	"strconv"
	"strings"
)

type PwdManager struct {
	Passwords []Password
}

func (pm PwdManager) ValidatePasswords() (int, int) {
	ocnt := 0
	ncnt := 0

	for _, pwd := range pm.Passwords {
		ob := pwd.Validate()
		nb := pwd.NewValidate()
		if ob {
			ocnt++
		}
		if nb {
			ncnt++
		}
	}
	return ocnt, ncnt
}

func (pm *PwdManager) Write(p []byte) (int, error) {
	bss := bytes.Split(p, []byte{13})
	rb := len(p)
	for _, bString := range bss {
		rawString := string(bString)
		// Removing newline chars
		rawString = strings.Trim(rawString, "\n")
		pwStringS := strings.Split(rawString, " ")
		lh := strings.Split(pwStringS[0], "-")
		low, err := strconv.Atoi(lh[0])
		if err != nil {
			return 0, err
		}
		high, err := strconv.Atoi(lh[1])
		if err != nil {
			return 0, err
		}
		ltr := strings.Trim(pwStringS[1], ":")
		pwd := Password{
			Required: ltr,
			Low:      low,
			High:     high,
			Password: pwStringS[2],
		}
		pm.Passwords = append(pm.Passwords, pwd)
	}
	return rb, nil
}

func NewPasswordManager() *PwdManager {
	pm := PwdManager{Passwords: make([]Password, 0)}
	return &pm
}

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
