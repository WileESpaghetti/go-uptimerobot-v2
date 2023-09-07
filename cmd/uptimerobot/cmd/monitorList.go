/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/api"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/models"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
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
	Run: func(cmd *cobra.Command, args []string) {
		// TODO does not handle pagination
		fmt.Println("monitor list called")

		monitorStr := strings.Join(args, "-")

		ms := &models.Monitors{}
		if len(args) > 0 {
			if err := ms.Set(monitorStr); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "could not get monitor list\n")
				return
			}
		}

		query := api.GetMonitorsRequest{}
		if len(*ms) > 0 {
			query.Monitors = *ms
		}

		monitors, err := apiClient.GetMonitors(&query)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Could not get monitor details: %s\n", err)
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
		_, _ = fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", "ID", "STATUS", "FRIENDLY NAME", "URL", "TYPE", "SUB TYPE", "KEYWORD TYPE", "KEYWORD", "USERNAME", "PASSWORD", "PORT", "INTERVAL", "CREATED")
		for _, monitor := range monitors {
			_, _ = fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%d\t%s\n", monitor.ID, monitor.Status, monitor.FriendlyName, monitor.Url, monitor.Type, monitor.SubType, monitor.KeywordType, monitor.KeywordValue, monitor.HttpUsername, monitor.HttpPassword, monitor.Port, monitor.Interval, monitor.CreateDatetime)
		}
		_ = w.Flush()
	},
}

func init() {
	monitorCmd.AddCommand(monitorListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// monitorListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// monitorListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
