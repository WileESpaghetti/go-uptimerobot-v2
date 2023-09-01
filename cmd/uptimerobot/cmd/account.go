package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Display one or many resourcesPrint account information",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("account called")
		if apiClient == nil { // FIXME this validation needs to be moved up a level so that it happens for all commands
			_, _ = fmt.Fprintf(os.Stderr, "No API client found")
			return
		}

		account, err := apiClient.GetAccountDetails()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Could not get account details: %s\n", err)
			return
		}

		fmt.Println(account)
		return
	},
}

func init() {
	rootCmd.AddCommand(accountCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// accountCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// accountCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
