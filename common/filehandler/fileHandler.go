package filehandler

import (
	"bytes"
	"github.com/TReyburn/advent-of-go/Day2/password"
	"github.com/TReyburn/advent-of-go/Day4/passport"
	"io/ioutil"
	"strconv"
	"strings"
)

// I should refactor LoadFile to take in a Reader/Writer interface - then I can refactor all of these LoadDayXFile funcs
//	into a single function

func LoadDay1File(fp string) ([]int, error) {
	fb, err := loadFileBytes(fp)
	if err != nil {
		return []int{}, err
	}
	res, err := convertByteSToIntS(fb)
	if err != nil {
		return []int{}, err
	}
	return res, nil
}

func LoadDay2File(fp string) ([]password.Password, error) {
	fb, err := loadFileBytes(fp)
	if err != nil {
		return []password.Password{}, err
	}
	pws, err := convertByteStoPasswordS(fb)
	return pws, err
}

func LoadDay3File(fp string) ([]string, error) {
	fb, err := loadFileBytes(fp)
	if err != nil {
		return []string{}, err
	}
	res := convertByteSToStringS(fb)
	return res, nil
}

func LoadDay4File(fp string) ([]passport.Passport, error) {
	_, err := loadFileBytes(fp)
	if err != nil {
		return []passport.Passport{}, err
	}
	return []passport.Passport{}, nil
}

func loadFileBytes(fp string) ([][]byte, error) {
	// Produces a raw []bytes of our file
	fc, err := ioutil.ReadFile(fp)
	if err != nil {
		return [][]byte{}, err
	}
	// Produces a [][]byte
	res := bytes.Split(fc, []byte{13})
	return res, nil
}

func convertByteSToIntS(bss [][]byte) ([]int, error) {
	ns := make([]int, 0)
	// Iterating over [][]byte, converting each []byte to str, converting str to int, and then appending int to []int
	for _, bString := range bss {
		sInt := string(bString)
		// Removing newline chars
		sInt = strings.Trim(sInt, "\n")
		n, err := strconv.Atoi(sInt)
		if err != nil {
			return []int{}, err
		}
		ns = append(ns, n)
	}
	return ns, nil
}

func convertByteStoPasswordS(bss [][]byte) ([]password.Password, error) {
	pws := make([]password.Password, 0)
	for _, bstring := range bss {
		rawString := string(bstring)
		// Removing newline chars
		rawString = strings.Trim(rawString, "\n")
		pwStringS := strings.Split(rawString, " ")
		lh := strings.Split(pwStringS[0], "-")
		low, err := strconv.Atoi(lh[0])
		if err != nil {
			return []password.Password{}, err
		}
		high, err := strconv.Atoi(lh[1])
		if err != nil {
			return []password.Password{}, err
		}
		ltr := strings.Trim(pwStringS[1], ":")
		pwd := password.Password{
			Required: ltr,
			Low:      low,
			High:     high,
			Password: pwStringS[2],
		}
		pws = append(pws, pwd)
	}
	return pws, nil
}

func convertByteSToStringS (bss [][]byte) []string {
	res := make([]string, 0)
	for _, bString := range bss {
		rawString := string(bString)
		rawString = strings.Trim(rawString, "\n")
		res = append(res, rawString)
	}
	return res
}