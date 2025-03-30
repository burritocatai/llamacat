// services/llm.go
package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/tmc/langchaingo/prompts"
)

// For Model List
type ModelResponse struct {
	Data []struct {
		// Type        string `json:"type"`
		ID string `json:"id"`
		// DisplayName string `json:"display_name"`
		// CreatedAt   string `json:"created_at"`
	} `json:"data"`
	// HasMore bool   `json:"has_more"`
	// FirstID string `json:"first_id"`
	// LastID  string `json:"last_id"`
}

type AIProvider struct {
	APIKeyENV         string
	APIKeyPlaceholder string
	APIBaseURL        string
	APIModelEndpoint  string
	Name              string
	Id                string
	Website           string
	Call              func(model string, prompt prompts.PromptTemplate,
		content string, context context.Context) (string, error)
	GetModels func() ([]string, error)
}

var AIProviders []AIProvider

func NewAIProvider(apiKeyENV, apiKeyPlaceholder, apiBaseURL, apiModelEndpoint, name, id, website string) *AIProvider {
	return &AIProvider{
		APIKeyENV:         apiKeyENV,
		APIKeyPlaceholder: apiKeyPlaceholder,
		APIBaseURL:        apiBaseURL,
		APIModelEndpoint:  apiModelEndpoint,
		Name:              name,
		Id:                id,
		Website:           website,
	}
}

func RegisterAIProvider(newProvider AIProvider) {
	AIProviders = append(AIProviders, newProvider)
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
