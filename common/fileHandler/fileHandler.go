package fileHandler

import (
	"bytes"
	"io/ioutil"
	"strconv"
	"strings"
)

func LoadFile(fp string) ([]int, error) {
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