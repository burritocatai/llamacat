// services/config_test.go
package services

import (
	"os"
	"testing"
)

func CreateTestProvider() *AIProvider {
	return NewAIProvider(
		"OPENAI_API_KEY",
		"sk-",
		"https://api.openai.com/v1",
		"models",
		"OpenAI",
		"openai",
		"https://platform.openai.com",
	)
}

func TestGetAPIKey(t *testing.T) {
	testAPIKey := "catsAreAPIs"
	os.Setenv("TEST_API_KEY", testAPIKey)

	testProvider := NewAIProvider("TEST_API_KEY", "", "", "", "", "", "")
	apikey, err := GetAPIKey(testProvider)
	expected := testAPIKey

	if apikey != expected || err != nil {
		t.Errorf("expected '%s' but got '%s'", expected, apikey)
	}
}
