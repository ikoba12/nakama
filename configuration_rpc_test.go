package main

import (
	"context"
	"errors"
	"fmt"
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
	if configType == "core" && version == "1.0.0" {
		return []byte(`{"key": "value"}`), nil
	}
	return nil, runtime.NewError("File not found", 14)
}

func mockSaveDataToDb(ctx context.Context, logger runtime.Logger, data []byte, module runtime.NakamaModule, userID string) error {
	return nil
}

func TestConfigurationInfoRpc(t *testing.T) {
	subtests := []struct {
		name                      string
		payload, expectedResponse string
	}{
		{
			name:             "when all the fields are passed in should return content",
			payload:          `{"type": "core", "version": "1.0.0", "hash": "9724c1e20e6e3e4d7f57ed25f9d4efb006e508590d528c90da597f6a775c13e5"}`,
			expectedResponse: `{"type":"core","version":"1.0.0","hash":"9724c1e20e6e3e4d7f57ed25f9d4efb006e508590d528c90da597f6a775c13e5","content":"{\"key\": \"value\"}"}`,
		},
		{
			name:             "when type and version are omitted should use default values and return content",
			payload:          `{"hash": "9724c1e20e6e3e4d7f57ed25f9d4efb006e508590d528c90da597f6a775c13e5"}`,
			expectedResponse: `{"type":"core","version":"1.0.0","hash":"9724c1e20e6e3e4d7f57ed25f9d4efb006e508590d528c90da597f6a775c13e5","content":"{\"key\": \"value\"}"}`,
		},
		{
			name:             "when hash not passed in should not return content",
			payload:          `{}`,
			expectedResponse: `{"type":"core","version":"1.0.0","hash":null,"content":null}`,
		},
	}
	DbUpdate = mockSaveDataToDb
	ReadFile = mockGetFileContent
	ctx := context.WithValue(context.Background(), "user_id", "test")
	logger := &MockLogger{}
	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {

			response, err := configurationInfoRpc(ctx, logger, nil, nil, subtest.payload)

			if err != nil {
				t.Log(err)
				t.Fatalf("Unexpected error: %v", err)
			}

			if response != subtest.expectedResponse {
				t.Errorf("Error with response! expected: %v, actual: %v", subtest.expectedResponse, response)
			}
		})
	}
}

func TestConfigurationInfoRpcErrors(t *testing.T) {
	subtests := []struct {
		name                           string
		payload, expectedError, userId string
		fileContentReader              func(configType string, version string, logger runtime.Logger) ([]byte, *runtime.Error)
		dbReader                       func(ctx context.Context, logger runtime.Logger, data []byte, module runtime.NakamaModule, userID string) error
	}{
		{
			name:              "when malformed json passed in should return invalid request error",
			payload:           `{"type": "core", "version": "1.0.0", "hash": "9724c1e20e6e3e4d7f57ed25f9d4efb006e508590d528c90da597f6a775c13e5"`,
			expectedError:     `Invalid request payload`,
			userId:            "test",
			fileContentReader: mockGetFileContent,
			dbReader:          mockSaveDataToDb,
		},
		{
			name:          "when file read throws error should return error for not finding data",
			payload:       `{"type": "core", "version": "1.0.0", "hash": "9724c1e20e6e3e4d7f57ed25f9d4efb006e508590d528c90da597f6a775c13e5"}`,
			expectedError: `Unable to find the data for given request`,
			userId:        "test",
			fileContentReader: func(configType string, version string, logger runtime.Logger) ([]byte, *runtime.Error) {
				return nil, runtime.NewError("test error", 1)
			},
			dbReader: mockSaveDataToDb,
		},
		{
			name:              "when Db save fails should return general error",
			payload:           `{"type": "core", "version": "1.0.0", "hash": "9724c1e20e6e3e4d7f57ed25f9d4efb006e508590d528c90da597f6a775c13e5"}`,
			expectedError:     `Error while fetching content`,
			userId:            "test",
			fileContentReader: mockGetFileContent,
			dbReader: func(ctx context.Context, logger runtime.Logger, data []byte, module runtime.NakamaModule, userID string) error {
				return errors.New("testError")
			},
		},
		{
			name:              "when no user id should throw user not found error",
			payload:           `{"type": "core", "version": "1.0.0", "hash": "9724c1e20e6e3e4d7f57ed25f9d4efb006e508590d528c90da597f6a775c13e5"}`,
			expectedError:     `No user ID in context`,
			userId:            "",
			fileContentReader: mockGetFileContent,
			dbReader: func(ctx context.Context, logger runtime.Logger, data []byte, module runtime.NakamaModule, userID string) error {
				return errors.New("testError")
			},
		},
	}
	logger := &MockLogger{}
	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), "user_id", subtest.userId)

			ReadFile = subtest.fileContentReader
			DbUpdate = subtest.dbReader
			_, err := configurationInfoRpc(ctx, logger, nil, nil, subtest.payload)
			errorMessage := fmt.Sprint(err)
			if errorMessage != subtest.expectedError {
				t.Errorf("incorrect error message! expected: %v, actual: %v", subtest.expectedError, errorMessage)
			}
		})
	}
}
