package main

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

type UserConfigurationInfo struct {
	ConfigurationType string  `json:"type"`
	Version           string  `json:"version"`
	Content           *string `json:"content"`
}
