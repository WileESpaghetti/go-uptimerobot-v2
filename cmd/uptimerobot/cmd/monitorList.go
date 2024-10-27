/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/monitor"
	"github.com/spf13/cobra"
	"io"
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
		return MonitorListAction(os.Stdout, apiClient, &monitorListRequest, args)
	},
}

var monitorListRequest = monitor.GetMonitorsRequest{}

func init() {
	monitorCmd.AddCommand(monitorListCmd)

	monitorListCmd.Flags().VarP(newMonitorTypesFlag(monitor.Types{}, &monitorListRequest.Types), "type", "t", "Monitor types: 1 - HTTP(s), 2 - Keyword, 3 - Ping, 4 - Port, 5 - Heartbeat")
	monitorListCmd.Flags().VarP(newMonitorStatusesFlag(monitor.Statuses{}, &monitorListRequest.Statuses), "status", "s", "Monitor Statuses:\n\t0 - Paused,\n\t1 - Not Checked,\n\t2 - Up,\n\t8 - Seems Down,\n\t9 - Down")
	monitorListCmd.Flags().Int64SliceVar(&monitorListRequest.UptimeRatios, "uptime-ratios", []int64{}, "Number of days to calculate the uptime ratio(s)") // FIXME might need a custom flag so that we can include hyphen separated
	monitorListCmd.Flags().BoolVar(&monitorListRequest.HasAllTimeUptimeRatio, "all-time-uptime-ratio", false, "Includes the all time uptime ratio")
	monitorListCmd.Flags().BoolVar(&monitorListRequest.HasLogs, "logs", false, "Include event logs with monitors")
	monitorListCmd.Flags().Var(newMonitorLogTypesFlag(monitor.LogTypes{}, &monitorListRequest.LogTypes), "log-type", "Log Types:\n\t1 - Down\n\t2 - Up\n\t98 - Started\n\t99 - Paused")
	monitorListCmd.Flags().Int64Var(&monitorListRequest.LogsLimit, "logs-limit", 0, "The number of logs to be returned in descending order")
	monitorListCmd.Flags().BoolVar(&monitorListRequest.HasResponseTimes, "response-times", false, "Include response times with monitors")
	monitorListCmd.Flags().Int64Var(&monitorListRequest.ResponseTimesLimit, "response-times-limit", 0, "The number of response times to be returned")
	monitorListCmd.Flags().Int64Var(&monitorListRequest.ResponseTimesAverage, "response-times-average", 0, "The minute interval to use to calculate average response time")
	monitorListCmd.Flags().BoolVar(&monitorListRequest.HasAlertContacts, "alert-contacts", false, "Include alert contacts with monitors")
	monitorListCmd.Flags().BoolVar(&monitorListRequest.HasAuthType, "auth-type", false, "Include authentication type with monitors")
	monitorListCmd.Flags().BoolVar(&monitorListRequest.HasCustomHTTPStatuses, "custom-http-statuses", false, "Include list of HTTP statuses that are considered Up/Down")
}

const errBadMonitorListFormat = "could not get monitor list: %w"

func MonitorListAction(out io.Writer, apiClient *uptime_robot.Client, options *monitor.GetMonitorsRequest, args []string) error {
	ms := &monitor.Monitors{}
	if len(args) > 0 {
		monitorStr := strings.Join(args, "-")
		if err := ms.Set(monitorStr); err != nil {
			return fmt.Errorf(errBadMonitorListFormat, err)
		}
	}

	// TODO does not handle pagination
	getMonitorResults, err := apiClient.GetMonitors(monitor.WithOptions(options))
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
