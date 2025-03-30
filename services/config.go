// services/config.go
package services

import (
	"fmt"
	"os"
	"strings"
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

func GetProviderAndModel(providerAndModel string) (AIProvider, string, error) {

	var modelName string
	var providerId string
	var aiProvider AIProvider

	parts := strings.Split(providerAndModel, ":")
	if len(parts) >= 2 {
		providerId = parts[0]
		modelName = parts[1]
		// Now you can use provider_id and model_name
	}

	for _, provider := range AIProviders {
		if provider.Id == providerId {
			aiProvider = provider
			break
		}
	}
	if aiProvider.Id != providerId {
		return aiProvider, modelName, fmt.Errorf("unable to find a provider of Id %s", providerId)
	}

	return aiProvider, modelName, nil

}
