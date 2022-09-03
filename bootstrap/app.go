package bootstrap

import (
	"github.com/idrpambudi/fita-appointment/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              "fita-appointment",
	Short:            "Fita appointment API for take-home assignment ",
	TraverseChildren: true,
}

// App root of application
type App struct {
	*cobra.Command
}

func NewApp() App {
	cmd := App{
		Command: rootCmd,
	}
	cmd.AddCommand(commands.GetSubCommands(CommonModules)...)
	return cmd
}

var RootApp = NewApp()
