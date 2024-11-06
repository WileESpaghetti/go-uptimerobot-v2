/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/public_status_page"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// pspListCmd represents the pspList command
var pspListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return PSPListAction(cmd.OutOrStdout(), apiClient, args)
	},
}

func init() {
	pspCmd.AddCommand(pspListCmd)
}

const errBadPSPListFormat = "could not get monitor list: %w"

func PSPListAction(out io.Writer, apiClient *uptime_robot.Client, args []string) error {
	psps := make(public_status_page.PublicStatusPages, 0, len(args))
	if len(args) > 0 {
		contactIDs := strings.Join(args, "-")
		if err := psps.Set(contactIDs); err != nil {
			return fmt.Errorf(errBadPSPListFormat, err)
		}
	}

	getPSPResults, err := apiClient.GetPublicStatusPages(public_status_page.WithPublicStatusPages(psps))
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(out, 0, 0, 2, ' ', tabwriter.Debug)
	_, _ = fmt.Fprintln(w, strings.Join([]string{"ID", "STATUS", "NAME", "STANDARD URL", "CUSTOM URL"}, "\t"))
	for _, psp := range getPSPResults {
		_, _ = fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n", psp.ID, psp.Status, psp.FriendlyName, psp.StandardURL, psp.CustomURL)
	}

	err = w.Flush()
	if err != nil {
		return err
	}

	return nil
}
