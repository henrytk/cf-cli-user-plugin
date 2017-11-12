package main

import "code.cloudfoundry.org/cli/plugin"

func main() {
	plugin.Start(new(UserPlugin))
}
