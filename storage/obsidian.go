package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type ObsidianConfig struct {
	Vaults map[string]struct {
		Path string `json:"path"`
	} `json:"vaults"`
}

type VaultInfo struct {
	Name string
	Path string
}

func WriteToObsidian(content, vaultAlias, targetPath string) {
	// TODO:
	fmt.Println("this is where I would write to obsidian...")
	// TODO: GetVaultPathFromAlias
	// WriteToLocalStorage(content, vaultpath, targetpath)
}

func GetObsidianVaults() []VaultInfo {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return []VaultInfo{{Name: "Default", Path: ""}}
	}

	configPath := filepath.Join(homeDir, ".config", "obsidian", "obsidian.json")
	_, err = os.Stat(configPath)
	if err != nil {
		// no config in .config, check application support on windows or macos
		appSupport, _ := getUserApplicationSupportFolder()
		configPath = filepath.Join(appSupport, "obsidian", "obsidian.json")
	}
	data, err := os.ReadFile(configPath)
	if err != nil {
		return []VaultInfo{{Name: "Default", Path: ""}}
	}

	var config ObsidianConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return []VaultInfo{{Name: "Default", Path: ""}}
	}

	vaults := make([]VaultInfo, 0, len(config.Vaults))
	for name, info := range config.Vaults {
		vaults = append(vaults, VaultInfo{
			Name: name,
			Path: info.Path,
		})
	}

	return vaults
}

func getUserApplicationSupportFolder() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = os.Getenv("HOME")
		if homeDir == "" {
			homeDir = os.Getenv("USERPROFILE")
		}
		if homeDir == "" {
			return "", err
		}
	}

	// On Windows, use the AppData folder instead of Application Support
	if os.Getenv("OS") == "Windows_NT" {
		applicationSupportFolder := filepath.Join(homeDir, "AppData", "Roaming")
		return applicationSupportFolder, nil
	}

	applicationSupportFolder := filepath.Join(homeDir, "Library", "Application Support")
	return applicationSupportFolder, nil
}
