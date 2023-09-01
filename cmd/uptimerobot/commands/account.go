package commands

import (
	"context"
	"flag"
	"fmt"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot"
	"github.com/google/subcommands"
	"os"
)

type Account struct {
	Client *uptime_robot.Client
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
	if p.Client == nil {
		fmt.Fprintf(os.Stderr, "No API client found")
		return subcommands.ExitFailure
	}

	account, err := p.Client.GetAccountDetails()
	if (err != nil) {
		fmt.Fprintf(os.Stderr, "Could not get account details: %s\n", err)
		return subcommands.ExitFailure
	}

	fmt.Println(account)
	return subcommands.ExitSuccess
}
