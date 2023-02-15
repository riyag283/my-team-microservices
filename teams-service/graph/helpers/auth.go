package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	util "teams/utils"

	"go.uber.org/zap"
)

func Authenticate(logger *zap.Logger) (string, error) {
	config, err := util.LoadConfig(".")
	if err != nil {
		logger.Error("failed to load config", zap.Error(err))
		return "", fmt.Errorf("failed to load config: %w", err)
	}

	// Prepare request body
	requestBody, err := json.Marshal(map[string]string{
		"username": config.AuthUsername,
		"password": config.AuthPassword,
	})
	if err != nil {
		logger.Error("failed to marshal request body", zap.Error(err))
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Build HTTP request
	url := config.AuthService + config.Login
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		logger.Error("failed to create http request", zap.Error(err))
		return "", fmt.Errorf("failed to create http request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send HTTP request and handle response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("failed to make http request", zap.Error(err))
		return "", fmt.Errorf("failed to make http request: %w", err)
	}
	defer resp.Body.Close()

	// Decode response body
	var responseBody struct {
		Error bool   `json:"error"`
		Token string `json:"token"`
	}
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		logger.Error("failed to decode response body", zap.Error(err))
		return "", fmt.Errorf("failed to decode response body: %w", err)
	}

	// Return token if authentication was successful
	if !responseBody.Error {
		logger.Info("Authentication Successful", zap.Any("Token", responseBody.Token))
		return responseBody.Token, nil
	}

	return "", fmt.Errorf("authentication failed")
}
