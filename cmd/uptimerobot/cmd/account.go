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

	a, err := urClient.GetAccountDetails()
	if err != nil {
		return fmt.Errorf(errGetAccountFailure, err)
	}

	accountFormat := "Account Details:" +
		"\n Email:            %s" +
		"\n Monitor Limit:    %d" +
		"\n Monitor Interval: %d minute(s)" +
		"\n\nMonitor Details:" +
		"\n Up:     %d" +
		"\n Down:   %d" +
		"\n Paused: %d\n"
	_, err = fmt.Fprintf(out, accountFormat, a.Email,
		a.MonitorLimit,
		a.MonitorInterval,
		a.UpMonitors,
		a.DownMonitors,
		a.PausedMonitors)
	if err != nil {
		return err
	}

	return nil
}
