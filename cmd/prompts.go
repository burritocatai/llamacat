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

	"github.com/spf13/cobra"
)

var promptAlias string
var promptRepo string

// promptsCmd represents the prompts command
var promptsCmd = &cobra.Command{
	Use:   "prompts",
	Short: "clone or update prompts for llamacat to use",
	Long: `Clone git repos containing prompts for use within llamacat:

Prompts just need a repo and an alias. Use clone to download, and update to 
pull updates for the prompts.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("prompts called")
	},
}

func init() {
	rootCmd.AddCommand(promptsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// promptsCmd.PersistentFlags().String("foo", "", "A help for foo")
	promptsCmd.PersistentFlags().StringVarP(&promptAlias, "alias", "a", "", "alias for the prompts")
	promptsCmd.PersistentFlags().StringVarP(&promptRepo, "repo", "r", "", "git repo for the prompts")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// promptsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
