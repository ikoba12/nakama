package main

import (
	"fmt"
	"github.com/heroiclabs/nakama-common/runtime"
	"os"
)

func getFileContent(configType string, version string, logger runtime.Logger, err error) ([]byte, error) {
	// Construct file path
	filePath := fmt.Sprintf("/nakama/data/configs/%s/%s.json", configType, version)
	logger.Debug("file path", filePath)
	logger.Debug("curDir", os.TempDir())
	// Read file from disk
	fileContent, err := os.ReadFile(filePath)

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
