package cmd

import (
	"os"

	"github.com/karlkori/aws-cron/internal/schedules"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aws-cron",
	Short: "List scheduled jobs",
	Long:  `List scheduled jobs with their scheduled time, status, name and description`,
	Run: func(cmd *cobra.Command, args []string) {

		outputFormat, _ := cmd.PersistentFlags().GetString("output")

		schedules.SchedulesInfo(outputFormat)
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("output", "o", "table", "Output format")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
