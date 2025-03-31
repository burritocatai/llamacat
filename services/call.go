// services/call.go
package services

import (
	"context"

	"github.com/burritocatai/llamacat/providers"
	"github.com/tmc/langchaingo/prompts"
)

func ProcessLLMRequest(content, model, prompt string) (string, error) {
	selectedProvider, selectedModel, err := GetProviderAndModel(model)
	if err != nil {
		return "", err
	}

	selectedPrompt, err := GetPrompt(prompt)
	if err != nil {
		return "", err
	}

	response, err := CallLLM(selectedProvider, selectedModel, content, selectedPrompt)

	return response, err
}

func CallLLM(provider providers.AIProvider, model string, content string, prompt prompts.PromptTemplate) (string, error) {
	ctx := context.Background()
	response, err := provider.Call(model, prompt, content, ctx)

	if err != nil {
		return "", err
	}

	return response, nil
}
