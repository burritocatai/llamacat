/*
Copyright © 2025 Dan Wiseman

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// cmd/config.go
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/burritocatai/llamacat/providers"
	"github.com/burritocatai/llamacat/providers/groq"
	"github.com/burritocatai/llamacat/providers/openai"
	"github.com/burritocatai/llamacat/services"
	"github.com/burritocatai/llamacat/storage"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure llamacat via interactive TUI",
	Long:  `Launch the interactive terminal UI to create and configure your llamacat.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Launch the config TUI
		showConfig()
	},
}

type Config struct {
	Destinations   []string          `yaml:"destinations"`
	Obsidian       ObsidianConfig    `yaml:"obsidian"`
	Github         GitHubConfig      `yaml:"github"`
	Ai             AIProvidersConfig `yaml:"ai"`
	DefaultCommand string            `yaml:"default_command"`
}

type ObsidianConfig struct {
	DefaultVault string                `yaml:"default_vault"`
	Vaults       []ObsidianVaultConfig `yaml:"vaults"`
}

type ObsidianVaultConfig struct {
	VaultName  string `yaml:"vault_name"`
	VaultPath  string `yaml:"vault_path"`
	VaultAlias string `yaml:"vault_alias"`
}

type GitHubConfig struct {
	Description string `yaml:"description"`
	Visible     bool   `yaml:"visible"`
	Open        bool   `yaml:"open"`
}

type AIProvidersConfig struct {
	Providers       []string           `yaml:"providers"`
	Configs         []AIProviderConfig `yaml:"configs"`
	DefaultProvider string             `yaml:"default_provider"`
}

type AIProviderConfig struct {
	Id     string `yaml:"id"`
	APIKey string `yaml:"api_key"`
	URL    string `yaml:"url"`
	Model  string `yaml:"model"`
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func showConfig() {
	theme := huh.ThemeCatppuccin()

	var newConfig Config

	destForm := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("llamacat Config").
				Description("Configure llamacat by selecting the destinations you want to use"),
			huh.NewMultiSelect[string]().
				Title("Destination(s)").
				Key("destinations").
				Options(
					huh.NewOption("Obsidian", "obsidian"),
					// huh.NewOption("GitHub Gist", "github"),
				).
				Validate(func(t []string) error {
					if len(t) <= 0 {
						return fmt.Errorf("at least one destination is required")
					}
					return nil
				}).
				Value(&newConfig.Destinations),
			huh.NewSelect[string]().
				Title("Default Destination").
				Description("Choose a default from the one(s) you selected. This will be chosen when you run just `llamacat`").
				OptionsFunc(func() []huh.Option[string] {
					options := []huh.Option[string]{}
					for _, dest := range newConfig.Destinations {
						options = append(options, huh.NewOption(dest, dest))
					}
					return options
				}, &newConfig.Destinations).
				Height(3).
				Value(&newConfig.DefaultCommand),
		),
	).WithTheme(theme)

	if err := destForm.Run(); err != nil {
		fail(err)
	}

	if services.Contains(newConfig.Destinations, "obsidian") {
		// Get Obsidian Vaults
		vaultOptions := make([]huh.Option[string], 0)
		for _, vault := range storage.GetObsidianVaults() {
			// Display name and path, but store just the name as the value
			displayName := fmt.Sprintf("%s (%s)", vault.Name, vault.Path)
			vaultOptions = append(vaultOptions, huh.NewOption(displayName, vault.Path))
		}

		obsidianForm := huh.NewForm(
			huh.NewGroup(
				huh.NewNote().
					Title("Obsidian Options"),
				huh.NewMultiSelect[string]().
					Title("Vaults to Connect with").
					Key("obsidianVaults").
					Description("Vaults that llamacat can interact with. Choose as many as you want. A default will be chosen later.").
					Validate(func(t []string) error {
						if len(t) <= 0 {
							return fmt.Errorf("at least one vault is required")
						}
						return nil
					}).
					Options(vaultOptions...),
			),
		).WithTheme(theme)

		if err := obsidianForm.Run(); err != nil {
			fail(err)
		}

		// loop through vaults here to assign aliases.
		chosenVaults := obsidianForm.Get("obsidianVaults").([]string)
		vaultAliases := make([]huh.Option[string], 0)

		for _, chosenVault := range chosenVaults {
			vaultForm := huh.NewForm(
				huh.NewGroup(
					huh.NewNote().
						Title("Vault Options").
						Description("Options for vault: "+chosenVault),
					huh.NewInput().
						Title("Vault Alias").
						Key("vaultAlias").
						Description("Alias to use for this vault, used with calling `llamacat obsidian vault <alias>`"),
					huh.NewInput().
						Title("Obsidian Folder").
						Key("vaultPath").
						Description("Path inside Obsidian to write the new llamacat Notes. If blank, defaults to llamacats/yyyy/mm/llamacats-yyyy-mm-dd.md"),
				),
			)

			if err := vaultForm.Run(); err != nil {
				fail(err)
			}

			// save vault info here
			vaultConfig := ObsidianVaultConfig{
				chosenVault, vaultForm.GetString("vaultPath"), vaultForm.GetString("vaultAlias"),
			}

			newConfig.Obsidian.Vaults = append(newConfig.Obsidian.Vaults, vaultConfig)
			vaultAliases = append(vaultAliases, huh.NewOption(vaultForm.GetString("vaultAlias"), vaultForm.GetString("vaultAlias")))
		}

		defaultVaultForm := huh.NewForm(
			huh.NewGroup(
				huh.NewNote().
					Title("Default Vault").
					Description("Choose a default vault to use when not specified."),
				huh.NewSelect[string]().
					Title("Default").
					Key("defaultVault").
					Options(vaultAliases...).
					Value(&newConfig.Obsidian.DefaultVault),
			),
		)

		if err := defaultVaultForm.Run(); err != nil {
			fail(err)
		}

	}

	if services.Contains(newConfig.Destinations, "github") {

		githubForm := huh.NewForm(
			huh.NewGroup(
				huh.NewNote().
					Title("GitHub Options"),
				huh.NewInput().
					Title("Description").
					Key("githubDefaultDesc").
					Description("Default description for gists, Override with --desc").
					Value(&newConfig.Github.Description),
				huh.NewConfirm().
					Title("Default Visibility of Gist").
					Key("githubDefaultVis").
					Description("Create by default public or private gists? Override with --public or --private").
					Affirmative("Public").Negative("Private").
					Value(&newConfig.Github.Visible),
				huh.NewConfirm().
					Title("Open Gist After Creation?").
					Key("githubDefaultOpen").
					Description("Open github gist after running? override with --open or --notopen").
					Affirmative("Yes").Negative("No").
					Value(&newConfig.Github.Open),
			),
		).WithTheme(theme)

		if err := githubForm.Run(); err != nil {
			fail(err)
		}

	}

	openAIProvider := openai.CreateOpenAIProvider()
	grokAIProvider := groq.CreateGroqProvider()
	providers.RegisterAIProvider(openAIProvider)
	providers.RegisterAIProvider(grokAIProvider)

	availableProviders := make([]huh.Option[string], 0)
	for _, provider := range providers.AIProviders {
		// Display name and path, but store just the name as the value
		availableProviders = append(availableProviders, huh.NewOption(provider.Name, provider.Id))
	}

	providersSelectForm := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("Available AI Providers").
				Description("Select from the below providers that llamacat currently supports."),
			huh.NewMultiSelect[string]().
				Title("Providers").
				Options(availableProviders...).
				Validate(func(t []string) error {
					if len(t) <= 0 {
						return fmt.Errorf("at least one provider is required")
					}
					return nil
				}).
				Value(&newConfig.Ai.Providers),
			huh.NewSelect[string]().
				Title("Default Provider").
				Description("Choose a default from the ones you selected").
				OptionsFunc(func() []huh.Option[string] {
					options := []huh.Option[string]{}
					for _, provider := range newConfig.Ai.Providers {
						options = append(options, huh.NewOption(provider, provider))
					}
					return options
				}, &newConfig.Ai.Providers).
				Value(&newConfig.Ai.DefaultProvider),
		),
	).WithTheme(theme)

	if err := providersSelectForm.Run(); err != nil {
		fail(err)
	}

	selectedProviders := newConfig.Ai.Providers
	for _, provider := range providers.AIProviders {
		if services.Contains(selectedProviders, provider.Id) {
			apiKey := ""
			providerForm := buildAIProviderForm(provider, apiKey)
			if err := providerForm.Run(); err != nil {
				fail(err)
			}
			newApiKey := providerForm.GetString("APIKey")
			if newApiKey != "" && apiKey != newApiKey {
				apiKey = newApiKey
			}
			newProviderConfig := AIProviderConfig{
				provider.Id, apiKey, providerForm.GetString("APIURL"), "",
			}

			// set env to grab models
			if apiKey != "" {
				err := os.Setenv(provider.APIKeyENV, apiKey)
				if err != nil {
					fail(err)
				}
			}

			providerModelForm := buildModelSelectForm(provider, newProviderConfig.APIKey)
			if err := providerModelForm.Run(); err != nil {
				fail(err)
			}
			newProviderConfig.Model = providerModelForm.GetString("Model")
			newConfig.Ai.Configs = append(newConfig.Ai.Configs, newProviderConfig)
		}

	}

	log.Debug("configs are: %s", newConfig)

	err := spinner.New().
		Title("Saving your configs...").
		Action(func() {
			saveConfigs(&newConfig)
			time.Sleep(time.Second * 2)
		}).
		Run()

	if err != nil {
		fail(err)
	}
	fmt.Println("Configs Saved.")

}

func fail(err error) {
	log.Error("failed with", "err", err)
	os.Exit(1) // Exit with a non-zero code
}

func saveConfigs(config *Config) error {

	// log.Debug("configs are: %s", config)

	// viper.Set("destinations", config.Destinations)
	// viper.Set("obsidian", config.Obsidian)
	// viper.Set("github", config.Github)
	// viper.Set("ai", config.Ai)
	// viper.Set("default_command", config.DefaultCommand)

	// configDir, err := services.GetConfigDir()
	// if err != nil {
	// 	return err
	// }

	// if err := os.MkdirAll(configDir, 0755); err != nil {
	// 	return err
	// }

	// configFile := filepath.Join(configDir, "config.yaml")
	// viper.SetConfigFile(configFile)

	// // Write the configuration to disk
	// if err := viper.WriteConfig(); err != nil {
	// 	// If the config file doesn't exist, create it
	// 	if os.IsNotExist(err) {
	// 		if err := viper.SafeWriteConfig(); err != nil {
	// 			return fmt.Errorf("failed to write config file: %w", err)
	// 		}
	// 	} else {
	// 		return fmt.Errorf("failed to write config file: %w", err)
	// 	}
	// }

	return nil
}

func buildAIProviderForm(provider providers.AIProvider, apiKey string) *huh.Form {
	envKey := os.Getenv(provider.APIKeyENV)
	if envKey == "" {
		return huh.NewForm(
			huh.NewGroup(
				huh.NewNote().
					Title(provider.Name+" Options"),
				huh.NewInput().
					Title("API Key").
					Key("APIKey").
					Description("Get your API key from "+provider.Website+". Can also set env "+provider.APIKeyENV).
					PlaceholderFunc(func() string {
						// Check if environment variable exists
						if envKey != "" {
							// {}
							return fmt.Sprintf("API key **%s...%s** found. Leave Blank", envKey[:4], envKey[len(envKey)-4:])
						}
						// Default placeholder when env var is not set
						return provider.APIKeyPlaceholder
					}, &apiKey).
					Value(&apiKey),
				huh.NewInput().
					Title(provider.Name+" API URL").
					Key("APIURL").
					Placeholder(provider.APIBaseURL),
			),
		).WithTheme(huh.ThemeCatppuccin())
	} else {

		return huh.NewForm(
			huh.NewGroup(
				huh.NewNote().
					Title(provider.Name+" Options"),
				huh.NewNote().
					Title("API Key").
					DescriptionFunc(func() string {
						// Check if environment variable exists
						if envKey := os.Getenv(provider.APIKeyENV); envKey != "" {
							// {}
							return fmt.Sprintf("Environment API key: %s...%s found", envKey[:4], envKey[len(envKey)-4:])
						}
						// Default placeholder when env var is not set
						return "error"
					}, &provider.APIKeyENV),
				huh.NewInput().
					Title(provider.Name+" API URL").
					Key("APIURL").
					Placeholder(provider.APIBaseURL),
			),
		)
	}
}

func buildModelSelectForm(provider providers.AIProvider, apiKey string) *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title(provider.Name+" Models"),
			huh.NewSelect[string]().
				Title("Default Model").
				Description("Choose the default model to use for this provider.").
				Key("Model").
				OptionsFunc(func() []huh.Option[string] {
					models, err := provider.GetModels()
					options := []huh.Option[string]{}
					if err != nil {
						options = append(options, huh.NewOption(err.Error(), "error"))
						return options
					}
					for _, model := range models {
						options = append(options, huh.NewOption(model, model))
					}
					return options
				}, &apiKey),
		),
	).WithTheme(huh.ThemeCatppuccin())
}
