// services/llm.go
package services

import (
	"context"

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
