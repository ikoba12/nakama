package main

import (
	"fmt"
	"github.com/heroiclabs/nakama-common/runtime"
	"os"
)

// Define as a variable for testing purposes
var osReadFile = os.ReadFile

func getFileContent(configType string, version string, logger runtime.Logger) ([]byte, *runtime.Error) {
	// Construct file path
	filePath := fmt.Sprintf("/nakama/data/configs/%s/%s.json", configType, version)
	// Read file from disk
	fileContent, err := osReadFile(filePath)

	if err != nil {
		if os.IsNotExist(err) {
			return nil, runtime.NewError("File not found", 14)
		}
		logger.Error("Error reading file: %v", err)
		return nil, runtime.NewError("Cannot read file", 13)
	}
	logger.Debug("File content : ", string(fileContent))
	return fileContent, nil
}
