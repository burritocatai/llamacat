// internal/providers/openai/openai.go
package openai

import (
	"context"
	"fmt"

	"github.com/burritocatai/llamacat/services"
	"github.com/tmc/langchaingo/llms"
	openai_llm "github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/prompts"
)

type OpenAIModel struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type OpenAIModels struct {
	Models []OpenAIModel `json:"models"`
}

var OpenAIModelList = OpenAIModels{
	Models: []OpenAIModel{
		{
			ID:          "o3-mini",
			Description: "fast, reasoning",
		},
		{
			ID:          "o1-mini",
			Description: "smaller, fast reasoning",
		},
		{
			ID:          "o1",
			Description: "high intelligence reasoning model",
		},
		{
			ID:          "o1-pro",
			Description: "o1 with more compute",
		},
		{
			ID:          "gpt-4.5-preview",
			Description: "largest and most capable gpt model",
		},
		{
			ID:          "gpt-4o",
			Description: "fast intelligence flexible gpt model",
		},
		{
			ID:          "chatgpt-4o",
			Description: "model used in chat gpt",
		},
		{
			ID:          "gpt-4o-mini",
			Description: "fast affordable small model",
		},
	},
}

func CreateOpenAIProvider() services.AIProvider {
	openAIProvider := services.NewAIProvider(
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

	return *openAIProvider
}

func GetOpenAIResponse(provider *services.AIProvider, model string, prompt prompts.PromptTemplate,
	content string, context context.Context) (string, error) {

	apiKey, err := services.GetAPIKey(provider)

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

func GetOpenAIModels(provider *services.AIProvider) ([]string, error) {
	modelList := make([]string, 0)
	for _, model := range OpenAIModelList.Models {
		modelList = append(modelList, model.ID)
	}
	if len(modelList) == 0 {
		return modelList, fmt.Errorf("no models returned for %s", provider.Id)
	}

	return modelList, nil
}

func init() {
	services.RegisterAIProvider(CreateOpenAIProvider())
}
