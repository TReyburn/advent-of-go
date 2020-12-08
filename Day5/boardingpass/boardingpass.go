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
	Code   string
	Row    int
	Column int
	ID     int
}

func (bp *BoardingPass) Decode() {
	bp.Row = binarySearch(bp.Code[:7], 0, 127, "F", "B")
	bp.Column = binarySearch(bp.Code[7:], 0, 7, "L", "R")
	bp.ID = bp.Column + (bp.Row * 8)
}

func binarySearch(cs string, low int, high int, lowInd string, highInd string) int {
	css := strings.Split(cs, "")
	mid := 0

	for _, c := range css {
		mid = (low + high) / 2
		if low == high {
			return mid
		}
		switch {
		case c == lowInd:
			high = mid - 1
		case c == highInd:
			low = mid + 1
		default:
			return 0
		}
	}
	return mid
}