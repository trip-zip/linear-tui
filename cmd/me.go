package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "List issues assigned to you",
	Long: `Fetch and display all issues assigned to the current user (based on the API key).
This uses the 'viewer' GraphQL query to identify the current user and filter issues accordingly.`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getLinearClient()
		if err != nil {
			fmt.Printf("Error creating Linear client: %v\n", err)
			return
		}

		user, err := client.GetCurrentUser()
		if err != nil {
			fmt.Printf("Error getting current user: %v\n", err)
			return
		}

		fmt.Printf("Current user: %s (%s)\n\n", user.Name, user.Email)

		fmt.Println("Fetching your assigned Linear issues...")
		var issues []Issue

		if statusFilter != "" {
			issues, err = client.GetMyIssuesByStatus(statusFilter)
		} else {
			issues, err = client.GetMyIssues()
		}

		if err != nil {
			fmt.Printf("Error fetching your issues: %v\n", err)
			return
		}

		displayIssues(issues, "issues assigned to you")
	},
}

func init() {
	meCmd.Flags().StringVarP(&statusFilter, "status", "s", "", "Filter your issues by status (e.g., 'In Progress', 'Backlog', 'Done')")
}
