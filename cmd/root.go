package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aquasecurity/table"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/liamg/tml"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aws-cron",
	Short: "List scheduled jobs",
	Long: `List scheduled jobs with their scheduled time, status, name and description`,
	Run: func(cmd *cobra.Command, args []string) {

		// Load the AWS SDK config from the environment or shared config
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatalf("unable to load SDK config, %v", err)
		}

		client := eventbridge.NewFromConfig(cfg)

		rules := listEventBridgeRules(client)

		outputFormat, _ := cmd.PersistentFlags().GetString("output")

		switch {
		case outputFormat == "table":
			printAsTable(rules)
		case outputFormat == "text":
			printAsText(rules)
		}
	},
}

func listEventBridgeRules(client *eventbridge.Client) []types.Rule {
	input := &eventbridge.ListRulesInput{}

	result, err := client.ListRules(context.TODO(), input)
	if err != nil {
		log.Fatalf("failed to retrieve EventBridge rules, %v", err)
	}

	if result.NextToken != nil {
		fmt.Println("...")
	}

	return result.Rules
}

func printAsTable(rules []types.Rule) {
	t := table.New(os.Stdout)
	t.SetHeaders("Schedule", "Status", "Name")

	for _, rule := range rules {
		if rule.ScheduleExpression != nil {
			scheduleExpression := *rule.ScheduleExpression

			if rule.State == "ENABLED" {
				t.AddRow(scheduleExpression, tml.Sprintf("<green>"+string(rule.State)+"</green>"), *rule.Name + "\n" + tml.Sprintf("<dim>"+ *rule.Description + "</dim>"))
			} else {
				t.AddRow(scheduleExpression, tml.Sprintf("<red>"+string(rule.State)+"</red>"), *rule.Name + "\n" + tml.Sprintf("<dim>"+ *rule.Description + "</dim>"))
			}
		}
	}

	t.Render()
}

func printAsText(rules []types.Rule) {
	for _, rule := range rules {
		// jsonData, _ := json.Marshal(&rule)
		// fmt.Println(string(jsonData))

		if rule.ScheduleExpression != nil {
			scheduleExpression := *rule.ScheduleExpression

			fmt.Printf("ScheduleExpression=%s, Status=%s, Name=%s\n", scheduleExpression, rule.State, *rule.Name)
		}
	}
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aws-cron.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringP("output", "o", "table", "Output format")
}
