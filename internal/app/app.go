package app

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

// App represent all needed stuff for jntpdn
type App struct {
	Params *Params
}

// Params holds the customizable parameters
type Params struct {
	DirectoryToServe   string
	BindIP             int
	BindPort           int
	DocOutputDirectory string
}

// New create a new App instance
func New() *App {
	return &App{
		Params: &Params{},
	}
}

// ShellCompletion generate completion for supported Cobra's supported shell
func (a *App) ShellCompletion(cmd *cobra.Command, args []string) error {

	var err error

	if len(args) != 0 {
		switch args[0] {
		case "bash":
			err = cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			err = cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			err = cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			err = cmd.Root().GenPowerShellCompletion(os.Stdout)
		}
	} else {
		err = errors.New("Please select one shell to generate the completion file")
	}

	if err != nil {
		return err
	}
	return nil

}

// DocsGeneration generate documentation in any suported format
//func (a *App) DocsGeneration(cmd *cobra.Command, args []string) error {
func (a *App) DocsGeneration() {
}
