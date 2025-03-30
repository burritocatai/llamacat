// services/config.go
package services

import (
	"fmt"
	"os"
)

// ai configs
func GetAPIKey(aiProvider *AIProvider) (string, error) {

	apiKey := os.Getenv(aiProvider.APIKeyENV)

	if apiKey != "" {
		return apiKey, nil
	}

	err := fmt.Errorf("error finding api key from: %s", aiProvider.APIKeyENV)

	return "", err
}
