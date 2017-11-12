package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var WhoAmICmd = &cobra.Command{
	Use:     "whoami",
	Aliases: []string{"w", "who"},
	Short:   "Lists the username, GUID, and email of the logged in user.",
	Long:    "Lists the username, GUID, and email of the logged in user. It will also confirm if you are not currently logged in.",
	Run: func(cmd *cobra.Command, args []string) {
		isLoggedIn, err := CLIConnection.IsLoggedIn()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if !isLoggedIn {
			fmt.Println("Not logged in. Use 'cf login' to log in.")
			os.Exit(0)
		}
		username, err := CLIConnection.Username()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		userGUID, err := CLIConnection.UserGuid()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		userEmail, err := CLIConnection.UserEmail()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf("Username  : %s\nUser GUID : %s\nUser Email: %s\n", username, userGUID, userEmail)
	},
}

var WhoAmIHelpCmd = &cobra.Command{
	Use: "help",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(`
Usage:
   cf user whoami

Aliases:
   w, who

Description:
   Lists the username, GUID, and email of the logged in user. It will also confirm if you are not currently logged in.
   It takes no arguments and has no flags.
`)
	},
}
