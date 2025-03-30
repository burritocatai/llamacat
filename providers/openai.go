// internal/providers/openai.go
package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tmc/langchaingo/llms"
	openai_llm "github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/prompts"
)

type OpenAIModelResponse struct {
	Data []struct {
		Object    string `json:"object"`
		ID        string `json:"id"`
		OwnedBy   string `json:"owned_by"`
		CreatedAt int32  `json:"created"`
	} `json:"data"`
}

func GetOpenAIResponse(provider *AIProvider, model string, prompt prompts.PromptTemplate,
	content string, context context.Context) (string, error) {

	apiKey, err := GetAPIKey(provider)

	if err != nil {
		return content, err
	}

	llm, err := openai_llm.New(
		openai_llm.WithModel(model),
		openai_llm.WithBaseURL(provider.APIBaseURL),
		openai_llm.WithToken(apiKey),
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

func GetOpenAIModels(provider *AIProvider) ([]string, error) {

	apiKey, err := GetAPIKey(provider)
	if err != nil {
		return nil, fmt.Errorf("error getting api key: %v\n")
	}

	var modelRequestURL = provider.APIBaseURL + "/" + provider.APIModelEndpoint

	req, err := http.NewRequest("GET", modelRequestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %v\n")
	}

	req.Header.Add("Authorization", "Bearer "+apiKey)

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
	var modelResp OpenAIModelResponse
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
	openAIProvider := NewAIProvider(
		"",
		"OPENAI_API_KEY",
		"sk-",
		"https://api.openai.com/v1",
		"models",
		"OpenAI",
		"openai",
		"https://platform.openai.com",
	)
	openAIProvider.Call = func(model string, prompt prompts.PromptTemplate,
		content string, context context.Context) (string, error) {
		return GetOpenAIResponse(openAIProvider, model, prompt, content, context)
	}

	openAIProvider.GetModels = func() ([]string, error) {
		return GetOpenAIModels(openAIProvider)
	}

	RegisterAIProvider(*openAIProvider)
}
