package main

import (
	"code.cloudfoundry.org/cli/plugin"
	"github.com/henrytk/cf-cli-user-plugin/cmd"
)

func main() {
	plugin.Start(&UserPlugin{cmd.NewRootCmd()})
}
