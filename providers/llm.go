// internal/providers/llm.go
package providers

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/viper"
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
	APIKey            string
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

func NewAIProvider(apiKey, apiKeyENV, apiKeyPlaceholder, apiBaseURL, apiModelEndpoint, name, id, website string) *AIProvider {
	return &AIProvider{
		APIKey:            apiKey,
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

func ProcessPurrWithLLM(content string, provider AIProvider, model string) (string, error) {
	ctx := context.Background()
	prompt := prompts.NewPromptTemplate(GetSystemPrompt()+"\n\nCONTENT: {{.content}}", []string{"content"})

	return provider.Call(model, prompt, content, ctx)
}

func GetSystemPrompt() string {
	// prompt := viper.GetString("default_prompt")

	// if prompt == "" {
	var prompt = `You are an expert programmer, business person, and shell scripter. You receive notes from
					the user that may be code, snippets of code, shell commands, or regular text notes. You
					take the INPUT and output markdown formatted text. You will analyze the input, if it is
					code or a shell command, you provide an explaination of the code in less than 3 sentences. 
					you will then include the code or command EXACTLY as the user gave it to you, in a Markdown 
					code block with the language or shell properly marked. If it is not code or a shell command,
					you will just format the note in markdown and output it without changing the content. Include
					a short subject of no more than three with a header of ## SUBJECT`
	// }

	return prompt

}

func GetAPIUrl(provider *AIProvider) (string, error) {

	if viper.InConfig("ai.configs") {
		configs := viper.Get("ai.configs").([]interface{})

		// Loop through the configs
		for _, config := range configs {
			cfg := config.(map[string]interface{})
			if cfg["id"] == provider.Id {
				url, ok := cfg["url"].(string)
				if !ok {
					return "", fmt.Errorf("error getting url for %s", provider.Name)
				}
				if url != "" {
					return url, nil
				}
			}
		}
	}

	url := provider.APIBaseURL
	if url == "" {
		return "", fmt.Errorf("no url found for provider %s", provider.Name)
	}

	return url, nil
}

func GetAPIKey(provider *AIProvider) (string, error) {

	if viper.InConfig("ai.configs") {
		configs := viper.Get("ai.configs").([]interface{})

		// Loop through the configs
		for _, config := range configs {
			cfg := config.(map[string]interface{})
			if cfg["id"] == provider.Id {
				apiKey, ok := cfg["api_key"].(string)
				if !ok {
					return "", fmt.Errorf("error getting key for %s", provider.Name)
				}
				if apiKey != "" {
					return apiKey, nil
				}
			}
		}
	}

	envApiKey := os.Getenv(provider.APIKeyENV)

	if envApiKey == "" {
		return "", fmt.Errorf("no key found. %s environment variable also not set", provider.APIKeyENV)
	}

	return envApiKey, nil
}

func GetDefaultModel(provider *AIProvider) (string, error) {

	configs := viper.Get("ai.configs").([]interface{})

	// Loop through the configs
	for _, config := range configs {
		cfg := config.(map[string]interface{})
		if cfg["id"] == provider.Id {
			model, ok := cfg["model"].(string)
			if !ok {
				return "", fmt.Errorf("error getting model for %s", provider.Name)
			}
			if model != "" {
				return model, nil
			}
		}
	}

	return "", fmt.Errorf("no model found for provider %s", provider.Name)
}
