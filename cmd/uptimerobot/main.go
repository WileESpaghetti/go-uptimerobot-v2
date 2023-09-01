package main

import (
	"context"
	"flag"
	"github.com/WileESpaghetti/go-uptimerobot-v2/cmd/uptimerobot/commands"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot"
	"github.com/google/subcommands"
	"os"
)

func main() {
	var apiKeyFlag string

	flag.StringVar(&apiKeyFlag, "api-key", "", "Account or Monitor Specific API key")

	client := uptime_robot.New("")

	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&commands.Account{Client: client}, "")
	subcommands.Register(&commands.Monitors{Client: client}, "")

	flag.Parse()

	client.ApiKey = apiKeyFlag

	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
