package cmd

import (
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot"
	"os"

	"github.com/spf13/cobra"
)

var (
	apiKey    string
	apiClient *uptime_robot.Client
)

// flags
const (
	FlagApiKey = "api-key"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "uptimerobot",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&apiKey, FlagApiKey, "", "config file (default is $HOME/.uptimerobot.yaml)")
	cobra.OnInitialize(func() {
		apiClient = uptime_robot.NewClient(apiKey)
	})
}
