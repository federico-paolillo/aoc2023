package utilities

import (
	"os"
	"path/filepath"
)

const AssetsBasePathEnvVar = "AOC2023_ASSETS_BASE"
const AssetsBasePathDefault = "../../assets"
const InputDataExtensions = ".txt"

func getAssetsBasePathOrDefault() string {

	maybeBasePathInEnvVar := os.Getenv(AssetsBasePathEnvVar)

	if maybeBasePathInEnvVar == "" {
		return AssetsBasePathDefault
	}

	return AssetsBasePathDefault

}

func getInputDataFilePath(inputDataName string) (string, error) {

	inputDataFileName := inputDataName + InputDataExtensions
	inputDataAssetsBasePath := getAssetsBasePathOrDefault()
	inputDataFilePath := filepath.Join(inputDataAssetsBasePath, inputDataFileName)

	inputDataAbsFilePath, absErr := filepath.Abs(inputDataFilePath)

	if absErr != nil {
		return "", absErr
	}

	return inputDataAbsFilePath, nil
}

func OpenInputData(inputDataName string) (*os.File, error) {

	inputDataFilePath, pathErr := getInputDataFilePath(inputDataName)

	if pathErr != nil {
		return nil, pathErr
	}

	_, statErr := os.Stat(inputDataFilePath)

	if statErr != nil {
		return nil, statErr
	}

	return os.Open(inputDataFilePath)
}
