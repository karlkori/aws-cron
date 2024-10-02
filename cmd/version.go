package cmd

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"

	"github.com/karlkori/aws-cron/internal/version"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long:  `Show version number and other build related info.`,
	Example: `
	Display simple text version
	aws-cron version

	Display version in yaml format
	aws-cron version --output=yaml
	
	Display version in json format
	aws-cron version --output=json
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		outputFormat, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}

		if outputFormat != "" && outputFormat != "json" && outputFormat != "yaml" {
			return fmt.Errorf("must provide valid output")
		}

		versionInfo := version.Get()

		switch {
		case outputFormat == "json":
			marshaled, err := json.MarshalIndent(&versionInfo, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(marshaled))
		case outputFormat == "yaml":
			marshaled, err := yaml.Marshal(&versionInfo)
			if err != nil {
				return err
			}
			fmt.Println(string(marshaled))
		default:
			fmt.Println("Git Version:", versionInfo.GitVersion)
			fmt.Println("Git Commit:", versionInfo.GitCommit)
			fmt.Println("Build Date:", versionInfo.BuildDate)
			fmt.Println("Go Version:", versionInfo.GoVersion)
			fmt.Println("Compiler:", versionInfo.Compiler)
			fmt.Println("Platform:", versionInfo.Platform)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	versionCmd.Flags().StringP("output", "o", "", "json|yaml")
}
