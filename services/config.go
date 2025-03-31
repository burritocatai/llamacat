// services/config.go
package services

import (
	"fmt"
	"os"

	"github.com/burritocatai/llamacat/providers"
)

// ai configs
func GetAPIKey(aiProvider *providers.AIProvider) (string, error) {

	apiKey := os.Getenv(aiProvider.APIKeyENV)

	if apiKey != "" {
		return apiKey, nil
	}

	err := fmt.Errorf("error finding api key from: %s", aiProvider.APIKeyENV)

	return "", err
}
