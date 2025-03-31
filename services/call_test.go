// services/call_test.go
package services

import (
	"testing"

	"github.com/burritocatai/llamacat/providers/fake"
	"github.com/tmc/langchaingo/prompts"
)

func TestCallLLM(t *testing.T) {
	content := "this is what the user has submitted"
	sysPrompt := "You are a friendly expert, {content}"

	fakeAI := fake.CreateFakeAIProvider()
	prompt := prompts.NewPromptTemplate(sysPrompt+"\n\nCONTENT: {{.content}}", []string{"content"})

	response, err := CallLLM(fakeAI, "fakemodel", content, prompt)

	if err != nil {
		t.Errorf("expected no error, but received %v", err)
	}

	expectedResponse := "I'm a fake AI, but helpful"
	if response != expectedResponse {
		t.Errorf("expected %s, but received %s", expectedResponse, response)
	}
}
