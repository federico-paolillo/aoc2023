package utilities_test

import (
	"io"
	"strings"
	"testing"

	"github.com/federico-paolillo/aoc2023/internal/utilities"
)

func TestScannerScanLinesUntilEOF(t *testing.T) {

	stringReader := strings.NewReader("line1\nline2\nline3\n")

	scanner := utilities.NewFileInputDataScanner(stringReader)

	expectations := []struct {
		line string
		eof  bool
	}{
		{line: "line1", eof: false},
		{line: "line2", eof: false},
		{line: "line3", eof: false},
		{line: "", eof: true},
	}

	for _, expectation := range expectations {

		lineScanned, scanErr := scanner.Scan()

		if scanErr != nil {
			if scanErr == io.EOF {
				if expectation.eof == false {
					t.Fatalf("Unexpected EOF encountered")
				}
			} else {
				t.Fatalf("Scan line returned an error: %s", scanErr)
			}
		}

		if lineScanned != expectation.line {
			t.Fatalf("Expected to scan line %s instead got %s", expectation.line, lineScanned)
		}
	}
}
