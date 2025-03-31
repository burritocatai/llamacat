// services/config.go
package services

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/burritocatai/llamacat/providers"
	"github.com/burritocatai/llamacat/providers/fake"
	"github.com/burritocatai/llamacat/storage"
	"github.com/spf13/viper"
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

func GetPromptConfig(prompt string) (prompts.PromptTemplate, error) {
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

func GetOutputConfig(output string) (outputFunc func(content string, path string, target string), path string, target string, err error) {
	parts := strings.Split(output, ":")

	if len(parts) != 2 {
		return nil, "", "", fmt.Errorf("invalid output parameter, received %s", output)
	}

	alias := parts[0]
	target = parts[1]

	if !viper.InConfig("outputs") {
		return nil, "", "", fmt.Errorf("no outputs configured")
	}
	configuredOutputs := viper.Get("outputs").([]interface{})

	for _, config := range configuredOutputs {
		cfg := config.(map[string]interface{})
		if cfg["alias"] != alias {
			continue
		}
		switch cfg["destination"] {
		case "obsidian":
			return func(content, path, target string) {
				storage.WriteToObsidian(content, path, target)
			}, cfg["path"].(string), filepath.Join(target, cfg["file_name"].(string)), nil
		case "local":
			return func(content, path, target string) {
				storage.WriteToLocalStorage(content, path, target)
			}, cfg["path"].(string), filepath.Join(target, cfg["file_name"].(string)), nil
		default:
			return nil, "", "", nil
		}
	}

	return nil, "", "", fmt.Errorf("could not find output with alias %s", alias)

}
