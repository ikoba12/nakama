package main

import (
	"context"
	"github.com/heroiclabs/nakama-common/runtime"
	"testing"
)

func TestHashFunction(t *testing.T) {
	// Test case 1: Empty input
	input := []byte("Gandalf")
	expectedHash := "955d2922801f650dace9ceb671d058fd0e350d1f593b1a4536b107de7d8482ae"
	actualHash := calculateHash(input)
	if actualHash != expectedHash {
		t.Errorf("Error with hash! expected: \"%v\", actual: \"%v\"", expectedHash, actualHash)
	}
}

func mockGetFileContent(configType string, version string, logger runtime.Logger) ([]byte, *runtime.Error) {
	if configType == "test" && version == "1.0.0" {
		return []byte(`{"key": "value"}`), nil
	}
	return nil, runtime.NewError("File not found", 14)
}

func mockSaveDataToDb(ctx context.Context, logger runtime.Logger, data []byte, err error, module runtime.NakamaModule) error {
	return nil
}

func TestConfigurationInfoRpc(t *testing.T) {

	logger := &MockLogger{}
	DbUpdate = mockSaveDataToDb
	ReadFile = mockGetFileContent
	payload := `{"type": "test", "version": "1.0.0", "hash": "9724c1e20e6e3e4d7f57ed25f9d4efb006e508590d528c90da597f6a775c13e5"}`
	response, err := configurationInfoRpc(nil, logger, nil, nil, payload)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedResponse := `{"type":"test","version":"1.0.0","hash":"9724c1e20e6e3e4d7f57ed25f9d4efb006e508590d528c90da597f6a775c13e5","content":"{\"key\": \"value\"}"}`
	if response != expectedResponse {
		t.Errorf("Error with response! expected: %v, actual: %v", expectedResponse, response)
	}
}
