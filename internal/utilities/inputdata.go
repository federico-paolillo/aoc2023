package utilities

import (
	"os"
	"path/filepath"
)

const InputDataExtensions = ".txt"

func getInputDataFilePath(inputDataName string, basePath string) (string, error) {
	inputDataFileName := inputDataName + InputDataExtensions
	inputDataFilePath := filepath.Join(basePath, inputDataFileName)

	inputDataAbsFilePath, absErr := filepath.Abs(inputDataFilePath)

	if absErr != nil {
		return "", absErr
	}

	return inputDataAbsFilePath, nil
}

func OpenInputData(inputDataName string, inputDataBasePath string) (*os.File, error) {
	inputDataFilePath, pathErr := getInputDataFilePath(inputDataName, inputDataBasePath)

	if pathErr != nil {
		return nil, pathErr
	}

	_, statErr := os.Stat(inputDataFilePath)

	if statErr != nil {
		return nil, statErr
	}

	return os.Open(inputDataFilePath)
}
