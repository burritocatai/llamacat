package services

import (
	"testing"
)

func TestNewAIProvider(t *testing.T) {
	// Setup test data
	testCases := []struct {
		name              string
		apiKeyENV         string
		apiKeyPlaceholder string
		apiBaseURL        string
		apiModelEndpoint  string
		providerName      string
		id                string
		website           string
	}{
		{
			name:              "OpenAI Provider",
			apiKeyENV:         "OPENAI_API_KEY",
			apiKeyPlaceholder: "sk-xxxxxxxxxxxxxxxx",
			apiBaseURL:        "https://api.openai.com/v1",
			apiModelEndpoint:  "/chat/completions",
			providerName:      "OpenAI",
			id:                "openai",
			website:           "https://openai.com",
		},
		{
			name:              "Empty Provider",
			apiKeyENV:         "",
			apiKeyPlaceholder: "",
			apiBaseURL:        "",
			apiModelEndpoint:  "",
			providerName:      "",
			id:                "",
			website:           "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function
			provider := NewAIProvider(
				tc.apiKeyENV,
				tc.apiKeyPlaceholder,
				tc.apiBaseURL,
				tc.apiModelEndpoint,
				tc.providerName,
				tc.id,
				tc.website,
			)

			// Check if the provider is not nil
			if provider == nil {
				t.Fatal("Expected non-nil AIProvider")
			}

			// Check if all fields are set correctly
			if provider.APIKeyENV != tc.apiKeyENV {
				t.Errorf("Expected APIKeyENV to be %s, got %s", tc.apiKeyENV, provider.APIKeyENV)
			}
			if provider.APIKeyPlaceholder != tc.apiKeyPlaceholder {
				t.Errorf("Expected APIKeyPlaceholder to be %s, got %s", tc.apiKeyPlaceholder, provider.APIKeyPlaceholder)
			}
			if provider.APIBaseURL != tc.apiBaseURL {
				t.Errorf("Expected APIBaseURL to be %s, got %s", tc.apiBaseURL, provider.APIBaseURL)
			}
			if provider.APIModelEndpoint != tc.apiModelEndpoint {
				t.Errorf("Expected APIModelEndpoint to be %s, got %s", tc.apiModelEndpoint, provider.APIModelEndpoint)
			}
			if provider.Name != tc.providerName {
				t.Errorf("Expected Name to be %s, got %s", tc.providerName, provider.Name)
			}
			if provider.Id != tc.id {
				t.Errorf("Expected Id to be %s, got %s", tc.id, provider.Id)
			}
			if provider.Website != tc.website {
				t.Errorf("Expected Website to be %s, got %s", tc.website, provider.Website)
			}
		})
	}
}
