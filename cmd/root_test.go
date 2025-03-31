package cmd

import (
	"os"
	"testing"
)

func TestGetContent(t *testing.T) {
	t.Run("With arguments", func(t *testing.T) {
		args := []string{"hello", "world"}
		content, err := getContent(args)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		expected := "hello world"
		if content != expected {
			t.Errorf("Expected content to be %q, got %q", expected, content)
		}
	})

	t.Run("With empty arguments", func(t *testing.T) {
		args := []string{}
		content, err := getContent(args)

		if err == nil {
			t.Error("Expected error, got nil")
		}

		if content != "" {
			t.Errorf("Expected empty content, got %q", content)
		}

		if err.Error() != "no content provided" {
			t.Errorf("Expected error message to be 'no content provided', got %q", err.Error())
		}
	})

	t.Run("With piped input", func(t *testing.T) {
		// Save the original stdin
		oldStdin := os.Stdin
		defer func() { os.Stdin = oldStdin }()

		// Create a pipe
		r, w, _ := os.Pipe()
		os.Stdin = r

		// Write to the pipe
		expected := "piped content"
		go func() {
			w.Write([]byte(expected))
			w.Close()
		}()

		content, err := getContent([]string{})

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if content != expected {
			t.Errorf("Expected content to be %q, got %q", expected, content)
		}
	})
}
