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

func NewInputDataScanner(r io.Reader) *InputDataScanner {
	innerScanner := bufio.NewScanner(r)

	return &InputDataScanner{innerScanner}
}

func NewFileInputDataScanner(inputDataName string, inputDataBasePath string) (*InputDataScanner, error) {
	day1ExampleDataFile, openErr := OpenInputData(inputDataName, inputDataBasePath)

	if openErr != nil {
		return nil, openErr
	}

	return NewInputDataScanner(day1ExampleDataFile), nil
}
