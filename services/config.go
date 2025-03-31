// services/config.go
package services

import (
	"fmt"
	"os"
	"strings"

	"github.com/burritocatai/llamacat/providers"
	"github.com/burritocatai/llamacat/providers/fake"
	"github.com/tmc/langchaingo/prompts"
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

func GetProviderAndModel(model string) (providers.AIProvider, string, error) {
	parts := strings.Split(model, ":")

	if len(parts) != 2 {
		return fake.CreateFakeAIProvider(), "", fmt.Errorf("invalid input provider and model selection, received %s", model)
	}

	providerId := parts[0]
	modelName := parts[1]
	var selectedProvider providers.AIProvider

	for _, provider := range providers.AIProviders {
		if provider.Id == providerId {
			selectedProvider = provider
			break
		}
	}
	if selectedProvider.Id != providerId {
		return fake.CreateFakeAIProvider(), "", fmt.Errorf("unable to find provider with Id %s", providerId)
	}

	return selectedProvider, modelName, nil

}

func GetPrompt(prompt string) (prompts.PromptTemplate, error) {
	parts := strings.Split(prompt, ":")

	if len(parts) != 2 {
		return prompts.NewPromptTemplate("error prompt, {{.content}}", []string{"content"}),
			fmt.Errorf("invalid input prompt source and prompt, received %s", prompt)
	}

	// TODO: implement actual prompt service
	return prompts.NewPromptTemplate(
			"You are a helpful assistant. Help the user with their content.\n\nCONTENT: {{.content}}", []string{"content"}),
		nil

}
