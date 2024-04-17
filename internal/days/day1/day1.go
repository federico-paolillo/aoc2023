package day1

import (
	"io"
	"strconv"
	"unicode"

	"github.com/federico-paolillo/aoc2023/internal/utilities"
)

type DigitWord struct {
	runes  []rune
	symbol string
}

type directionalIndexCalculator = func(scanIndex int) int

var DigitWords = []DigitWord{
	{runes: []rune{'z', 'e', 'r', 'o'}, symbol: "0"},
	{runes: []rune{'o', 'n', 'e'}, symbol: "1"},
	{runes: []rune{'t', 'w', 'o'}, symbol: "2"},
	{runes: []rune{'t', 'h', 'r', 'e', 'e'}, symbol: "3"},
	{runes: []rune{'f', 'o', 'u', 'r'}, symbol: "4"},
	{runes: []rune{'f', 'i', 'v', 'e'}, symbol: "5"},
	{runes: []rune{'s', 'i', 'x'}, symbol: "6"},
	{runes: []rune{'s', 'e', 'v', 'e', 'n'}, symbol: "7"},
	{runes: []rune{'e', 'i', 'g', 'h', 't'}, symbol: "8"},
	{runes: []rune{'n', 'i', 'n', 'e'}, symbol: "9"},
}

func tryMatchDigitWord(someRunesToMatch []rune) (string, bool) {
	for _, digitWord := range DigitWords {
		for runeToMatchIndex, runeToMatch := range someRunesToMatch {
			// No digit word is empty. We don't have to fear that we may go out of range
			// The guards after this check will ensure that we stop once we fully consume the digit word

			runeDiffers := digitWord.runes[runeToMatchIndex] != runeToMatch

			if runeDiffers {
				break
			}

			// If we get here we did not have any mismatch up until now
			// We are looking for the longest prefix, consuming the whole digit word is necessary to have a match

			matchesAllDigitWordRunes := runeToMatchIndex == len(digitWord.runes)-1

			if matchesAllDigitWordRunes {
				return digitWord.symbol, true
			}
		}
	}

	return "", false
}

func getDigit(complexMatching bool, calculateDirectionalIndex directionalIndexCalculator, someRunes []rune) string {

	for scanIndex := range someRunes {
		runeIndex := calculateDirectionalIndex(scanIndex)

		currentRune := someRunes[runeIndex]

		if unicode.IsDigit(currentRune) {
			return string(currentRune)
		}

		if complexMatching {
			runesSubstring := someRunes[runeIndex:]

			if digit, ok := tryMatchDigitWord(runesSubstring); ok {
				return digit
			}
		}
	}

	return ""
}

func getDigitFromLeft(complexMatching bool, someRunes []rune) string {
	return getDigit(complexMatching, func(scanIndex int) int { return scanIndex }, someRunes)
}

func getDigitFromRight(complexMatching bool, someRunes []rune) string {
	return getDigit(complexMatching, func(scanIndex int) int { return (len(someRunes) - 1) - scanIndex }, someRunes)
}

func GetCalibrationValue(complexMatching bool, line string) (int, error) {
	lineRunes := []rune(line)

	leftDigit := getDigitFromLeft(complexMatching, lineRunes)
	rightDigit := getDigitFromRight(complexMatching, lineRunes)

	completeDigit := leftDigit + rightDigit

	return strconv.Atoi(completeDigit)
}

func calculateCalibrationTotal(complexMatching bool, scanner *utilities.InputDataScanner) int {
	total := 0

	for {
		line, scanErr := scanner.Scan()

		if scanErr == io.EOF {
			break
		}

		if scanErr != nil {
			panic(scanErr)
		}

		value, getErr := GetCalibrationValue(complexMatching, line)

		if getErr != nil {
			panic(getErr)
		}

		total += value
	}

	return total
}

func PartOne(scanner *utilities.InputDataScanner) int {
	return calculateCalibrationTotal(false, scanner)
}

func PartTwo(scanner *utilities.InputDataScanner) int {
	return calculateCalibrationTotal(true, scanner)
}
