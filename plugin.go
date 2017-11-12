package main

import (
	"github.com/spf13/cobra"

	"code.cloudfoundry.org/cli/plugin"
)

type UserPlugin struct {
	cmd *cobra.Command
}

func (up *UserPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	up.cmd.SetArgs(args)
	up.cmd.Execute()
}

func (up *UserPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "user",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 1,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 7,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "user",
				HelpText: "Manage CloudFoundry users",
			},
		},
	}
}
