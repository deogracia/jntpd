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
)

// docsCmd represents the docs command
var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Generate documentation",
	RunE: func(cmd *cobra.Command, args []string) error {
		return docsCommand(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)

	docsCmd.Flags().StringP("dir", "d", "", "Destination directory for docs")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// docsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// docsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func docsCommand(cmd *cobra.Command, args []string) error {
	dir, err := cmd.Flags().GetString("dir")
	if err != nil {
		return err
	}

	if dir == "" {
		if dir, err = ioutil.TempDir("", "jntpdn"); err != nil {
			return err
		}
	}
	return docsAction(os.Stdout, dir)
}

func docsAction(out io.Writer, dir string) error {

	/*  err := doc.GenerateMarkdownTree(rootCmd, dir)
	    if err != nil{
	      return err
	    }
	*/
	return nil
}
