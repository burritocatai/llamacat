// internal/providers/ollama.go
package providers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetOllamaModels(apiKey string, baseURL string) ([]string, error) {

	// Get API key from environment variable
	if apiKey == "" {
		return nil, fmt.Errorf("ANTHROPIC_API_KEY environment variable not set")
	}

	// Create a new request
	req, err := http.NewRequest("GET", "https://api.anthropic.com/v1/models", nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %v\n")
	}

	// Add headers
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
	var modelResp ModelResponse
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
