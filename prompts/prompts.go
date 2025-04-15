package prompts

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
)

type PromptStatus int

const (
	UpToDate = iota
	AlreadyExists
	Cloned
	Updated
	Error
)

const defaultPromptsRepo = "https://github.com/burritocatai/llamacat_prompts"

func DownloadDefaultPrompts() (PromptStatus, error) {
	return DownloadPromptRepo(defaultPromptsRepo, "default")
}

func UpdateDefaultPrompts() (PromptStatus, error) {
	return UpdatePromptRepo("default")
}

func DownloadPromptRepo(repoUrl, repoAlias string) (PromptStatus, error) {
	err := clonePromptsRepoToDir(repoUrl, repoAlias)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			var status PromptStatus = AlreadyExists
			return status, nil
		}
		var status PromptStatus = Error
		return status, err
	}
	var status PromptStatus = Cloned
	return status, err
}

func UpdatePromptRepo(repoAlias string) (PromptStatus, error) {
	err := pullExistingPromptsRepo(repoAlias)
	if err != nil {
		if strings.Contains(err.Error(), "up-to-date") {
			var status PromptStatus = UpToDate
			return status, nil
		}
		var status PromptStatus = Error
		return status, err
	}
	var status PromptStatus = Updated
	return status, err
}

func GetAvailablePrompts(repoAlias string) ([]string, error) {
	var foldersWithSystemMD []string

	if !promptsExist(repoAlias) {
		return make([]string, 0), fmt.Errorf("prompts with alias of %s do not exist", repoAlias)
	}

	// Get a list of directories in the specified path
	promptsDir, err := getOrCreatePromptsConfigDir()
	if err != nil {
		return nil, err
	}

	dirs, err := os.ReadDir(filepath.Join(promptsDir, repoAlias))
	if err != nil {
		return nil, err
	}

	// Iterate through each directory
	for _, dir := range dirs {
		if dir.IsDir() { // Check if it is a directory
			// Build the path to the SYSTEM.md file
			systemMDPath := filepath.Join(promptsDir, repoAlias, dir.Name(), "SYSTEM.md")

			// Check if the SYSTEM.md file exists (case insensitive)
			if fileExists(systemMDPath) {
				foldersWithSystemMD = append(foldersWithSystemMD, dir.Name())
			}
		}
	}

	return foldersWithSystemMD, nil
}

func GetPromptContent(repoAlias, promptName string) (string, error) {
	if !promptsExist(repoAlias) {
		return "", fmt.Errorf("prompts with alias of %s do not exist", repoAlias)
	}

	// Get a list of directories in the specified path
	promptsDir, err := getOrCreatePromptsConfigDir()
	if err != nil {
		return "", err
	}

	_, err = os.ReadDir(filepath.Join(promptsDir, repoAlias, promptName))
	if err != nil {
		return "", err
	}

	systemMDPath := filepath.Join(promptsDir, repoAlias, promptName, "SYSTEM.md")

	// Check if the SYSTEM.md file exists (case insensitive)
	if !fileExists(systemMDPath) {
		return "", fmt.Errorf("prompt not found")
	}

	data, err := os.ReadFile(systemMDPath)
	if err != nil {
		return "", err
	}

	// Convert the contents to a string
	fileContents := string(data)

	return fileContents, nil
}

// Helper function to check if a file exists (case insensitive)
func fileExists(filename string) bool {
	// Check if the file exists
	if _, err := os.Stat(filename); err == nil {
		return true
	}

	// Check if the file exists with a different case
	filenameWithoutExtension := strings.TrimSuffix(filename, filepath.Ext(filename))
	files, _ := os.ReadDir(filepath.Dir(filenameWithoutExtension))
	for _, file := range files {
		if strings.EqualFold(file.Name(), filepath.Base(filenameWithoutExtension)) {
			return true
		}
	}

	return false
}

func clonePromptsRepoToDir(repoUrl string, promptsAlias string) error {
	promptsDir, err := getOrCreatePromptsConfigDir()
	if err != nil {
		return err
	}

	// clone the repo
	_, err = git.PlainClone(filepath.Join(promptsDir, promptsAlias), false, &git.CloneOptions{
		URL:      repoUrl,
		Progress: os.Stdout,
	})
	if err != nil {
		return err
	}

	return nil
}

func promptsExist(repoAlias string) bool {
	path, err := getOrCreatePromptsConfigDir()
	if err != nil {
		return false
	}

	path = filepath.Join(path, repoAlias)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func pullExistingPromptsRepo(repoAlias string) error {
	path, err := getOrCreatePromptsConfigDir()
	if err != nil {
		return err
	}

	path = filepath.Join(path, repoAlias)
	r, err := git.PlainOpen(path)
	if err != nil {
		return err
	}
	w, err := r.Worktree()
	if err != nil {
		return err
	}
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		return err
	}
	ref, err := r.Head()
	if err != nil {
		return err
	}
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		return err
	}
	fmt.Println(commit)

	return nil
}

func getOrCreatePromptsConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	promptsDir := filepath.Join(home, ".config", "llamacat", "prompts")

	if _, err := os.Stat(promptsDir); os.IsNotExist(err) {
		// Create the directory
		err = os.MkdirAll(promptsDir, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	return promptsDir, nil
}
