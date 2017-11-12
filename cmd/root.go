package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewRootCmd returns the root of the command tree. It has no `Run` function
// defined as we never expect it to be executed. The CF CLI will call
// this plugin with 'user ...' as the args, which will first match the UserCmd.
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{Use: "cf"}
	rootCmd.AddCommand(UserCmd)
	return rootCmd
}

var UserCmd = &cobra.Command{
	Use: "user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("User called with args:", args)
	},
}
