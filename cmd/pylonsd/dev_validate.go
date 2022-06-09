package cli

import (
	"fmt"

	"github.com/Pylons-tech/pylons/x/pylons/types"
	"github.com/spf13/cobra"
)

func CmdDevValidate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate [path]",
		Short: "Validates all Pylons recipe or cookbook files in the provided path",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			ForFiles(path, perCookbook, perRecipe)
		},
	}
	return cmd
}

func perCookbook(path string, _ types.Cookbook) {
	fmt.Println(path, "is a valid cookbook")
}

func perRecipe(path string, _ types.Recipe) {
	fmt.Println(path, "is a valid recipe")
}
