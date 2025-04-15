// services/call_test.go
package services

import (
	"testing"

	lc_prompts "github.com/burritocatai/llamacat/prompts"
	"github.com/burritocatai/llamacat/providers"
	"github.com/burritocatai/llamacat/providers/fake"
	"github.com/tmc/langchaingo/prompts"
)

func TestProcessLLMRequest(t *testing.T) {
	content := "this is what the user has submitted"
	prompt := "default:extract_resume_points"
	model := "fakeai:fake-model-8b"

	_, _ = lc_prompts.DownloadDefaultPrompts()
	fakeAI := fake.CreateFakeAIProvider()
	providers.RegisterAIProvider(fakeAI)

	expectedResponse := "I'm a fake AI, but helpful"

	response, err := ProcessLLMRequest(content, model, prompt)
	if err != nil {
		t.Errorf("Expected to not receive an error. Received %v", err)
	}

	if response != expectedResponse {
		t.Errorf("expected to get %s, received %s", expectedResponse, response)
	}

}

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
