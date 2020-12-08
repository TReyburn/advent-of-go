package boardingpass

import (
	"bytes"
	"strings"
)

type BPManager struct {
	Passes []*BoardingPass
}

func (bpm BPManager) GetMaxID() int {
	max := 0
	for _, bp := range bpm.Passes {
		if bp.ID > max {
			max = bp.ID
		}
	}
	return max
}

func (bpm *BPManager) DecodeAll() {
	for _, bp := range bpm.Passes {
		bp.Decode()
	}
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
		bpm.Passes = append(bpm.Passes, &bp)
	}
	return rb, nil
}

func NewBoardingPassManager() *BPManager {
	bpm := BPManager{Passes: make([]*BoardingPass, 0)}
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

	for _, c := range css {
		switch {
		case c == lowInd:
			// Go division is floor division - so no need to subtract by 1
			high = (low + high) / 2
		case c == highInd:
			low = (low + high) / 2 + 1
		default:
			return 0
		}
	}
	// low == high at this point
	return low
}