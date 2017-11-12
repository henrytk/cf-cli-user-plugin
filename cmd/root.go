package cmd

import (
	"code.cloudfoundry.org/cli/plugin"
	"github.com/spf13/cobra"
)

var CLIConnection plugin.CliConnection

// NewRootCmd returns the root of the command tree. It has no `Run` function
// defined as we never expect it to be executed. The CF CLI will call
// this plugin with 'user ...' as the args, which will first match the UserCmd.
func NewRootCmd() *cobra.Command {
	RootCmd.AddCommand(UserCmd)
	UserCmd.AddCommand(WhoAmICmd)
	WhoAmICmd.AddCommand(WhoAmIHelpCmd)
	return RootCmd
}

var RootCmd = &cobra.Command{
	Use: "cf",
}
