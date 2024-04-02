package utilities

import (
	"bufio"
	"io"
)

type InputDataScanner struct {
	innerScanner *bufio.Scanner
}

func (inputDataScanner *InputDataScanner) Scan() (string, error) {
	bufioScanner := inputDataScanner.innerScanner

	if bufioScanner.Scan() {

		lineScanned := bufioScanner.Text()

		return lineScanned, nil

	}

	bufioErr := bufioScanner.Err()

	if bufioErr == nil {

		return "", io.EOF

	}

	return "", bufioErr
}

func NewFileInputDataScanner(r io.Reader) *InputDataScanner {

	innerScanner := bufio.NewScanner(r)

	return &InputDataScanner{innerScanner}

}
