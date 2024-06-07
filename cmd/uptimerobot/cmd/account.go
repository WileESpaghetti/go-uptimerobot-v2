package cmd

import (
	"errors"
	"fmt"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Print information about your account",
	Long: `The account command displays your user information, account limits, and number
of up, down, and paused monitors.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return AccountAction(os.Stdout, apiClient)
	},
}

func init() {
	rootCmd.AddCommand(accountCmd)
}

var ErrNoApiClient = errors.New("no usable API client")

const errGetAccountFailure = "could not get account details: %w"

func AccountAction(out io.Writer, urClient *uptime_robot.Client) error {
	if urClient == nil {
		return ErrNoApiClient
	}

	account, err := urClient.GetAccountDetails()
	if err != nil {
		return fmt.Errorf(errGetAccountFailure, err)
	}

	_, err = fmt.Fprintln(out, account)
	if err != nil {
		return err
	}

	return nil
}
