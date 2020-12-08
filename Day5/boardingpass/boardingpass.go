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