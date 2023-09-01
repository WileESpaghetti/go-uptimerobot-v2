package commands

import (
	"context"
	"flag"
	"fmt"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/models"
	"github.com/google/subcommands"
)

type Account struct {
}

func (*Account) Name() string     { return "account" }
func (*Account) Synopsis() string { return "View account information." }
func (*Account) Usage() string {
	return `account :
  Print account information to stdout.
`
}

func (p *Account) SetFlags(f *flag.FlagSet) {
}

func (p *Account) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Println(models.Account{Email: "example@example.com",
		DownMonitors: 2,
		MonitorLimit: 50,
		MonitorInterval: 5,
		PausedMonitors: 1,
		UpMonitors: 5})
	return subcommands.ExitSuccess
}
