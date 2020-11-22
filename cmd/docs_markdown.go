/*Package cmd heart of cobra

Copyright © 2020 Lionel Félicité <deogracia@free.fr>
This file is licensed under the BSD 3-Clause Clear License.
The full text can also be found:
  * in the LICENSE file at the root directory of this project
  * at https://spdx.org/licenses/BSD-3-Clause-Clear.html

*/
package cmd

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// markdownCmd represents the markdown command
var markdownCmd = &cobra.Command{
	Use:   "markdown",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return markdownDocsCommand(cmd, args)
	},
}

func init() {
	docsCmd.AddCommand(markdownCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markdownCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markdownCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func markdownDocsCommand(cmd *cobra.Command, args []string) error {
	dir, err := cmd.Flags().GetString("dir")
	if err != nil {
		return err
	}

	if dir == "" {
		if dir, err = ioutil.TempDir("", "jntpdn"); err != nil {
			return err
		}
	} else {
		if os.MkdirAll(dir, 0750) != nil {
			return err
		}
	}
	return markdownDocsAction(os.Stdout, dir)
}

func markdownDocsAction(out io.Writer, dir string) error {
	err := doc.GenMarkdownTree(rootCmd, dir)
	if err != nil {
		return err
	}

	return nil
}
