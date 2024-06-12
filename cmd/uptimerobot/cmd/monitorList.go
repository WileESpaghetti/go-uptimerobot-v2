/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/monitors"
	"github.com/spf13/cobra"
	"io"
)

import (
	"fmt"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/api"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/models"
	"os"
	"strings"
	"text/tabwriter"
)

// monitorListCmd represents the monitorList command
var monitorListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return MonitorListAction(os.Stdout, apiClient, monitorTypes, monitorStatuses, args)
	},
}

var monitorTypes monitors.Types
var monitorStatuses monitors.Statuses

func init() {
	monitorCmd.AddCommand(monitorListCmd)

	monitorListCmd.Flags().VarP(newMonitorTypesFlag(monitors.Types{}, &monitorTypes), "type", "t", "Monitor types: 1 - HTTP(s), 2 - Keyword, 3 - Ping, 4 - Port, 5 - Heartbeat")
	monitorListCmd.Flags().VarP(newMonitorStatusesFlag(monitors.Statuses{}, &monitorStatuses), "status", "s", "Monitor Statuses: 0 - Paused, 1 - Not Checked, 2 - Up, 8 - Seems Down, 9 - Down")
}

const errBadMonitorListFormat = "could not get monitor list: %w"

func MonitorListAction(out io.Writer, apiClient *uptime_robot.Client, types monitors.Types, statuses monitors.Statuses, args []string) error {
	ms := &models.Monitors{}
	if len(args) > 0 {
		monitorStr := strings.Join(args, "-")
		if err := ms.Set(monitorStr); err != nil {
			return fmt.Errorf(errBadMonitorListFormat, err)
		}
	}

	options := make([]api.MonitorOptions, 0, 2) // number of flags in the command
	options = append(options,
		api.WithMonitors(*ms),
		api.WithTypes(types),
		api.WithStatuses(statuses))

	// TODO does not handle pagination
	getMonitorResults, err := apiClient.GetMonitors(options...)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(out, 0, 0, 2, ' ', tabwriter.Debug)
	_, _ = fmt.Fprintln(w, strings.Join([]string{"ID", "STATUS", "FRIENDLY NAME", "URL", "TYPE", "SUB TYPE", "KEYWORD TYPE", "KEYWORD", "USERNAME", "PASSWORD", "PORT", "INTERVAL", "CREATED"}, "\t"))
	for _, monitor := range getMonitorResults {
		_, _ = fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%d\t%s\n", monitor.ID, monitor.Status, monitor.FriendlyName, monitor.Url, monitor.Type, monitor.SubType, monitor.KeywordType, monitor.KeywordValue, monitor.HttpUsername, monitor.HttpPassword, monitor.Port, monitor.Interval, monitor.CreateDatetime)
	}
	_ = w.Flush()
	return nil
}
