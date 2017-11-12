package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var UserCmd = &cobra.Command{
	Use: "user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("User called with args:", args)
	},
}
