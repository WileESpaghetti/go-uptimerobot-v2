package commands

import (
	"context"
	"flag"
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
	return subcommands.ExitSuccess
}
