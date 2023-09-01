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
	client := uptime_robot.New()

	account, err := client.GetAccountDetails()
	if (err != nil) {
		fmt.Fprintf(os.Stderr, "Could not get account details")
		return subcommands.ExitFailure
	}

	fmt.Println(account)
	return subcommands.ExitSuccess
}
