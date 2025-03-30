// internal/providers/anthropic.go
package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tmc/langchaingo/llms"
	anthropic_llm "github.com/tmc/langchaingo/llms/anthropic"
	"github.com/tmc/langchaingo/prompts"
)

type AnthropicModelResponse struct {
	Data []struct {
		Type        string `json:"type"`
		ID          string `json:"id"`
		DisplayName string `json:"display_name"`
		CreatedAt   string `json:"created_at"`
	} `json:"data"`
	HasMore bool   `json:"has_more"`
	FirstID string `json:"first_id"`
	LastID  string `json:"last_id"`
}

func GetAnthropicResponse(provider *AIProvider, model string, prompt prompts.PromptTemplate,
	content string, context context.Context) (string, error) {

	apiKey, err := GetAPIKey(provider)

	if err != nil {
		return content, err
	}

	llm, err := anthropic_llm.New(
		anthropic_llm.WithModel(model),
		anthropic_llm.WithBaseURL(provider.APIBaseURL),
		anthropic_llm.WithToken(apiKey),
	)
	if err != nil {
		return content, err
	}

	result, err := prompt.Format(map[string]any{"content": content})

	if err != nil {
		return content, err
	}

	completion, err := llm.Call(context, result, llms.WithTemperature(0.5))
	return completion, err

}

func GetAnthropicModels(provider *AIProvider) ([]string, error) {

	apiKey, err := GetAPIKey(provider)
	if err != nil {
		return nil, fmt.Errorf(err.Error() + provider.APIKey)
	}

	var modelRequestURL = provider.APIBaseURL + "/" + provider.APIModelEndpoint

	req, err := http.NewRequest("GET", modelRequestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %v\n")
	}

	req.Header.Add("x-api-key", apiKey)
	req.Header.Add("anthropic-version", "2023-06-01")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error making request: %v\n", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response: %v\n", err)
	}

	// Parse the JSON response
	var modelResp AnthropicModelResponse
	err = json.Unmarshal(body, &modelResp)
	if err != nil {
		return nil, fmt.Errorf("Error parsing JSON: %v\n", err)
	}

	// Print the parsed response
	var models []string
	for _, model := range modelResp.Data {
		models = append(models, model.ID)
	}

	return models, nil
}

func init() {
	anthropicProvider := NewAIProvider(
		"",
		"ANTHROPIC_API_KEY",
		"sk-",
		"https://api.anthropic.com/v1",
		"models",
		"Anthropic",
		"anthropic",
		"https://anthropic.com",
	)
	anthropicProvider.Call = func(model string, prompt prompts.PromptTemplate,
		content string, context context.Context) (string, error) {
		return GetAnthropicResponse(anthropicProvider, model, prompt, content, context)
	}

	anthropicProvider.GetModels = func() ([]string, error) {
		return GetAnthropicModels(anthropicProvider)
	}

	RegisterAIProvider(*anthropicProvider)
}
