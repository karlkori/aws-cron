package schedules

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aquasecurity/table"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/liamg/tml"
)

func SchedulesInfo(outputFormat string) {
	// Load the AWS SDK config from the environment or shared config
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load AWS SDK config, %v", err)
	}

	client := eventbridge.NewFromConfig(cfg)

	rules := listEventBridgeRules(client)

	switch {
	case outputFormat == "table":
		printAsTable(rules)
	case outputFormat == "text":
		printAsText(rules)
	}
}

func listEventBridgeRules(client *eventbridge.Client) []types.Rule {
	var limit int32 = 100
	var rules []types.Rule
	var nextToken *string = nil

	for {
		input := &eventbridge.ListRulesInput{NextToken: nextToken, Limit: &limit}

		result, err := client.ListRules(context.TODO(), input)
		if err != nil {
			log.Fatalf("failed to retrieve EventBridge rules, %v", err)
			break
		}

		nextToken = result.NextToken
		rules = append(rules, result.Rules...)

		if result.NextToken == nil {
			break
		}
	}

	return rules
}

func printAsTable(rules []types.Rule) {
	t := table.New(os.Stdout)
	t.SetHeaders("#", "Schedule", "Status", "Name")

	idx:=0
	for _, rule := range rules {
		if rule.ScheduleExpression != nil {
			idx = idx + 1
			scheduleExpression := *rule.ScheduleExpression

			statusStr := tml.Sprintf("<green>"+string(rule.State)+"</green>")
			if rule.State == "DISABLED" {
				statusStr = tml.Sprintf("<red>"+string(rule.State)+"</red>")
			}

			t.AddRow(strconv.FormatInt(int64(idx), 10), scheduleExpression, statusStr, *rule.Name+"\n"+tml.Sprintf("<dim>"+*rule.Description+"</dim>"))
		}
	}

	t.Render()
}

func printAsText(rules []types.Rule) {
	idx:=0
	for _, rule := range rules {
		if rule.ScheduleExpression != nil {
			idx = idx + 1
			scheduleExpression := *rule.ScheduleExpression

			fmt.Printf("%d, Schedule=%s, Status=%s, Name=%s\n", idx, scheduleExpression, rule.State, *rule.Name)
		}
	}
}
