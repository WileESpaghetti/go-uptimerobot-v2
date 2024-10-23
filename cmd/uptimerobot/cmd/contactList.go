/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// contactListCmd represents the contactList command
var contactListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return ContactListAction(cmd.OutOrStdout(), apiClient, args)
	},
}

func init() {
	contactCmd.AddCommand(contactListCmd)
}

func ContactListAction(out io.Writer, apiClient *uptime_robot.Client, args []string) error {
	getAlertContactResults, err := apiClient.GetAlertContacts()
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(out, 0, 0, 2, ' ', tabwriter.Debug)
	_, _ = fmt.Fprintln(w, strings.Join([]string{"ID", "STATUS", "NAME", "TYPE", "VALUE"}, "\t"))
	for _, contact := range getAlertContactResults {
		_, _ = fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", contact.ID, contact.Status, contact.FriendlyName, contact.Type, contact.Value)
	}

	err = w.Flush()
	if err != nil {
		return err
	}

	return nil
}
