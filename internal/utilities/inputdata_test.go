package utilities_test

import (
	"testing"

	"github.com/federico-paolillo/aoc2023/internal/utilities"
)

func TestOpenInputDataFile(t *testing.T) {
	osFile, err := utilities.OpenInputData("day1", "../../assets")

	if err != nil {
		t.Fatalf("Could not locate input data 'day1'. Error was %s", err)
	}

	_ = osFile.Close()
}
