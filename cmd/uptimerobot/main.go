package main

import (
	"context"
	"flag"
	"github.com/WileESpaghetti/go-uptimerobot-v2/cmd/uptimerobot/commands"
	"github.com/google/subcommands"
	"os"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&commands.Account{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
