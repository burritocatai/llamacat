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
package cmd

import (
	"fmt"
	"os"

	"github.com/burritocatai/llamacat/prompts"
	"github.com/spf13/cobra"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "clone a prompt repo for use",
	Long: `Clones a prompt repo with the selected alias and link into the prompts folder:

Repos are kept in ~/.config/llamacat/prompts and are updated with the update command`,
	Run: func(cmd *cobra.Command, args []string) {
		if promptAlias != "" && promptRepo != "" {
			promptStatus, err := prompts.DownloadPromptRepo(promptRepo, promptAlias)
			if err != nil {
				fmt.Printf("received error trying to download %s from %s. error: %v", promptAlias, promptRepo, err)
				os.Exit(1)
			}
			if promptStatus != prompts.Cloned {
				fmt.Printf("could not clone prompt repo")
				os.Exit(1)
			}
			fmt.Printf("cloned repo %s from %s. now available as %s:prompt_name", promptAlias, promptRepo, promptAlias)
			os.Exit(0)
		} else if promptAlias == "default" {
			fmt.Printf("cloned %s prompts. now available as %s:prompt_name", promptAlias, promptAlias)
			prompts.DownloadDefaultPrompts()
		}
	},
}

func init() {
	promptsCmd.AddCommand(cloneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cloneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cloneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
