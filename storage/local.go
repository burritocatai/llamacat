package storage

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func WriteToLocalStorage(content, rootPath, targetPath string) {
	targetPath = ReplaceStringWithDate(targetPath)
	path := path.Join(rootPath, targetPath)
	appendToFile(path, content)
}

func appendToFile(filePath, content string) error {
	// Create all directories in the path if they don't exist
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	// Check if the file exists
	_, err := os.Stat(filePath)
	if err == nil {
		// If the file exists, add three dashes and two linebreaks to the content
		content = "\n\n---\n" + content
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}
