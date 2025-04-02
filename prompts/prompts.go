package prompts

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

const defaultPromptsRepo = "https://github.com/burritocatai/llamacat_prompts"

func DownloadDefaultPrompts() error {
	return DownloadPromptRepo(defaultPromptsRepo, "default")
}

func UpdateDefaultPrompts() error {
	return UpdatePromptRepo("default")
}

func DownloadPromptRepo(repoUrl, repoAlias string) error {
	err := clonePromptsRepoToDir(repoUrl, repoAlias)
	return err
}

func UpdatePromptRepo(repoAlias string) error {
	err := pullExistingPromptsRepo(repoAlias)
	return err
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
