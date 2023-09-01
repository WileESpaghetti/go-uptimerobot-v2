package main

import "github.com/WileESpaghetti/go-uptimerobot-v2/cmd/uptimerobot/cmd"

// TODO handle signals
// TODO add Context
// TODO don't use hard-coded root command name (example: see bash script best practices)
// TODO consider matching how gcloud and kubectl commands work. ($cmd [verb] [resource]) vs how we do it ($cmd [resource] [verb]). would also want to flip the command/file name if we do that
func main() {
	cmd.Execute()
}
