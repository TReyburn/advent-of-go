package filehandler

import (
	"io"
	"os"
)

func LoadInputFile(fp string, writer io.Writer) error {
	f, err := os.Open(fp)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, f)
	return err
}