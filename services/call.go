// services/call.go
package services

import (
	"context"

	"github.com/burritocatai/llamacat/providers"
	"github.com/tmc/langchaingo/prompts"
)

func CallLLM(provider providers.AIProvider, model string, content string, prompt prompts.PromptTemplate) (string, error) {
	ctx := context.Background()
	response, err := provider.Call(model, prompt, content, ctx)

	if err != nil {
		return "", err
	}

	return response, nil
}
