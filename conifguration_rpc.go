package main

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"github.com/heroiclabs/nakama-common/runtime"
)

var (
	ReadFile = getFileContent
	DbUpdate = saveDataToDb
)

func configurationInfoRpc(ctx context.Context, logger runtime.Logger, db *sql.DB, module runtime.NakamaModule, payload string) (string, error) {
	logger.Debug("Configuration content RPC called")

	// parse request
	request, err := parseRequest(payload, logger)
	if err != nil {
		return "", err
	}

	// fetch file content
	fileContent, fileFetchError := ReadFile(request.ConfigurationType, request.Version, logger)
	if fileFetchError != nil {
		logger.Error("Error fetching file: %v", err)
		// return invalid argument code
		return "", runtime.NewError("Unable to find the data for given request", 3)
	}

	// build response
	contentString := string(fileContent)
	response := &ConfigurationInfoResponse{
		ConfigurationType: request.ConfigurationType,
		Version:           request.Version,
		Hash:              request.Hash,
		Content:           &contentString,
	}

	// check hashes
	calculatedHash := calculateHash(fileContent)
	if request.Hash == nil || calculatedHash != *request.Hash {
		logger.Debug("Calculated hash not equal to request hash. Calculated hash : %v ", calculatedHash)
		response.Content = nil
	}

	// marshal the response
	out, err := json.Marshal(response)
	if err != nil {
		logger.Error("Error marshalling response type to JSON: %v", err)
		// return internal error
		return "", runtime.NewError("Error while fetching content", 13)
	}

	// save data to DB
	dbSaveError := DbUpdate(ctx, logger, out, err, module)
	if dbSaveError != nil {
		return "", dbSaveError
	}

	return string(out), nil
}

func parseRequest(payload string, logger runtime.Logger) (ConfigurationInfoRequest, error) {
	var request = ConfigurationInfoRequest{
		ConfigurationType: "core",
		Version:           "1.0.0",
	}
	err := json.Unmarshal([]byte(payload), &request)
	if err != nil {
		logger.Error("Error unmarshalling request payload: %v", err)
		// return invalid argument code
		return ConfigurationInfoRequest{}, runtime.NewError("Invalid request", 3)
	}
	return request, nil
}

func calculateHash(input []byte) string {
	hash := sha256.Sum256(input)
	calculatedHash := hex.EncodeToString(hash[:])
	return calculatedHash
}
