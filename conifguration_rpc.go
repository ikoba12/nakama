package main

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"github.com/heroiclabs/nakama-common/runtime"
)

type ConfigurationInfoResponse struct {
	ConfigurationType string  `json:"type"`
	Version           string  `json:"version"`
	Hash              *string `json:"hash"`
	Content           *string `json:"content"`
}

type ConfigurationInfoRequest struct {
	ConfigurationType string  `json:"type"`
	Version           string  `json:"version"`
	Hash              *string `json:"hash"`
}

func configurationInfoRpc(ctx context.Context, logger runtime.Logger, db *sql.DB, module runtime.NakamaModule, payload string) (string, error) {
	logger.Debug("Configuration content RPC called")
	var request = ConfigurationInfoRequest{
		ConfigurationType: "core",
		Version:           "1.0.0",
	}
	err := json.Unmarshal([]byte(payload), &request)
	if err != nil {
		logger.Error("Error unmarshalling request payload: %v", err)
		return "", runtime.NewError("Invalid request", 01)
	}

	fileContent, fileFetchError := getFileContent(request.ConfigurationType, request.Version, logger, err)
	if fileFetchError != nil {
		logger.Error("Error fetching file: %v", err)
		return "", runtime.NewError("Unable to find the data for given request", 02)
	}
	contentString := string(fileContent)
	response := &ConfigurationInfoResponse{
		ConfigurationType: request.ConfigurationType,
		Version:           request.Version,
		Hash:              request.Hash,
		Content:           &contentString,
	}
	// Calculate file hash
	hash := sha256.Sum256(fileContent)
	calculatedHash := hex.EncodeToString(hash[:])

	if request.Hash != nil && calculatedHash != *request.Hash {
		logger.Debug("Calculated hash not equal to request hash. Calculated hash : %v ", calculatedHash)
		response.Content = nil
	}

	out, err := json.Marshal(response)
	if err != nil {
		logger.Error("Error marshalling response type to JSON: %v", err)
		return "", runtime.NewError("Error while fetching content", 13)
	}

	return string(out), nil
}
