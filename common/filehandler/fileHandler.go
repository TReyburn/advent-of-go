package filehandler

import (
	"bytes"
	"github.com/TReyburn/advent-of-go/Day2/password"
	"io"
	"io/ioutil"
	"os"
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

func LoadInputFile(fp string, writer io.Writer) error {
	f, err := os.Open(fp)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, f)
	return err
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