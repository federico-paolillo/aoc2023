package day1_test

import (
	"testing"

	"github.com/federico-paolillo/aoc2023/internal/days/day1"
	"github.com/federico-paolillo/aoc2023/internal/utilities"
)

func TestGetsCalibrationValues(t *testing.T) {
	testInputs := []struct {
		line     string
		expected int
	}{
		{line: "1abc2", expected: 12},
		{line: "pqr3stu8vwx", expected: 38},
		{line: "a1b2c3d4e5f", expected: 15},
		{line: "treb7uchet", expected: 77},
	}

	for _, testInput := range testInputs {
		value, err := day1.GetCalibrationValue(false, testInput.line)

		if err != nil {
			t.Fatal("An unexpected error was returned when finding calibration values")
		}

		if testInput.expected != value {
			t.Errorf("Expected %d, got %d", testInput.expected, value)
		}
	}
}

func TestGetsComplexCalibrationValues(t *testing.T) {
	testInputs := []struct {
		line     string
		expected int
	}{
		{line: "two1nine", expected: 29},
		{line: "eightwothree", expected: 83},
		{line: "abcone2threexyz", expected: 13},
		{line: "xtwone3four", expected: 24},
		{line: "4nineeightseven2", expected: 42},
		{line: "zoneight234", expected: 14},
		{line: "7pqrstsixteen", expected: 76},
	}

	for _, testInput := range testInputs {
		value, err := day1.GetCalibrationValue(true, testInput.line)

		if err != nil {
			t.Fatal("An unexpected error was returned when finding calibration values")
		}

		if testInput.expected != value {
			t.Errorf("Expected %d, got %d", testInput.expected, value)
		}
	}
}

func TestDayOnePartOneWorksWithExampleValues(t *testing.T) {
	expectedResult := 142

	day1ExampleDataScanner, openErr := utilities.NewFileInputDataScanner("day1_example", "../../../assets")

	if openErr != nil {
		t.Fatalf("Could not open input data. Error is %s", openErr)
	}

	result := day1.PartOne(day1ExampleDataScanner)

	if result != expectedResult {
		t.Fatalf("Expected result to be %d but got %d", expectedResult, result)
	}
}

func TestDayOnePartTwoWorksWithExampleValues(t *testing.T) {
	expectedResult := 281

	day1ExampleDataScanner, openErr := utilities.NewFileInputDataScanner("day1_example2", "../../../assets")

	if openErr != nil {
		t.Fatalf("Could not open input data. Error is %s", openErr)
	}

	result := day1.PartTwo(day1ExampleDataScanner)

	if result != expectedResult {
		t.Fatalf("Expected result to be %d but got %d", expectedResult, result)
	}
}

func TestDayOnePartOneWorksWithInputValues(t *testing.T) {
	expectedResult := 55621

	day1ExampleDataScanner, openErr := utilities.NewFileInputDataScanner("day1", "../../../assets")

	if openErr != nil {
		t.Fatalf("Could not open input data. Error is %s", openErr)
	}

	result := day1.PartOne(day1ExampleDataScanner)

	if result != expectedResult {
		t.Fatalf("Expected result to be %d but got %d", expectedResult, result)
	}
}

func TestDayOnePartTwoWorksWithInputValues(t *testing.T) {
	expectedResult := 53592

	day1ExampleDataScanner, openErr := utilities.NewFileInputDataScanner("day1", "../../../assets")

	if openErr != nil {
		t.Fatalf("Could not open input data. Error is %s", openErr)
	}

	result := day1.PartTwo(day1ExampleDataScanner)

	if result != expectedResult {
		t.Fatalf("Expected result to be %d but got %d", expectedResult, result)
	}
}
