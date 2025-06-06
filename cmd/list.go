package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all issues in the workspace",
	Long:  `Fetch and display all issues from your Linear workspace with optional filtering.`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getLinearClient()
		if err != nil {
			fmt.Printf("Error creating Linear client: %v\n", err)
			return
		}

		fmt.Println("Fetching Linear issues...")
		var issues []Issue

		if statusFilter != "" {
			issues, err = client.GetIssuesByStatus(statusFilter)
		} else {
			issues, err = client.GetIssues()
		}

		if err != nil {
			fmt.Printf("Error fetching issues: %v\n", err)
			return
		}

		displayIssues(issues, "all issues")
	},
}

func init() {
	listCmd.Flags().StringVarP(&statusFilter, "status", "s", "", "Filter issues by status (e.g., 'In Progress', 'Backlog', 'Done')")
}
