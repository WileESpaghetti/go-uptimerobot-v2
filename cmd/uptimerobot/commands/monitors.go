package commands

import (
	"context"
	"flag"
	"fmt"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/api"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/models"
	"github.com/google/subcommands"
	"os"
	"text/tabwriter"
)

type Monitors struct {
	Client   *uptime_robot.Client
	monitors models.Monitors
}

func (*Monitors) Name() string     { return "monitors" }
func (*Monitors) Synopsis() string { return "View monitor information." }
func (*Monitors) Usage() string {
	return `monitors :
  Print monitor information to stdout.
`
}

func (p *Monitors) SetFlags(f *flag.FlagSet) {
	f.Var(&p.monitors, "monitor", "Monitor ID")
}

func (p *Monitors) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if p.Client == nil {
		_, _ = fmt.Fprintf(os.Stderr, "No API client found")
		return subcommands.ExitFailure
	}

	var query *api.GetMonitorsRequest
	if len(p.monitors) > 0 {
		query = &api.GetMonitorsRequest{Monitors: p.monitors}
	}
	//fmt.Println("query: ", query)

	monitors, err := p.Client.GetMonitors(query)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Could not get monitor details: %s\n", err)
		return subcommands.ExitFailure
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
	_, _ = fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", "ID", "STATUS", "FRIENDLY NAME", "URL", "TYPE", "SUB TYPE", "KEYWORD TYPE", "KEYWORD", "USERNAME", "PASSWORD", "PORT", "INTERVAL", "CREATED")
	for _, monitor := range *monitors {
		_, _ = fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%d\t%s\n", monitor.Id, monitor.Status, monitor.FriendlyName, monitor.Url, monitor.Type, monitor.SubType, monitor.KeywordType, monitor.KeywordValue, monitor.HttpUsername, monitor.HttpPassword, monitor.Port, monitor.Interval, monitor.CreateDatetime)
	}
	_ = w.Flush()

	return subcommands.ExitSuccess
}
