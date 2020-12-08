package boardingpass

import (
	"bytes"
	"strings"
)

type BPManager struct {
	Passes []BoardingPass
}

func (bpm *BPManager) Write(p []byte) (int, error) {
	rb := len(p)
	bss := bytes.Split(p, []byte{13})
	for _, bString := range bss {
		rawString := string(bString)
		rawString = strings.Trim(rawString, "\n")
		bp := BoardingPass{
			Code: rawString,
		}
		bpm.Passes = append(bpm.Passes, bp)
	}
	return rb, nil
}

func NewBoardingPassManager() *BPManager {
	bpm := BPManager{Passes: make([]BoardingPass, 0)}
	return &bpm
}

type BoardingPass struct {
	Code string
	Row int
	Column int
	ID int
}

func (bp *BoardingPass) Decode() {

}

func parseRow(cs string) int {
	low := 0
	high := 127
	css := strings.Split(cs, "")
	mid := 0

	for _, c := range css {
		mid = (low + high) / 2
		if low == high {
			return mid
		}
		if c == "F" {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return mid
}