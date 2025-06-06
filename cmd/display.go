package cmd

import "fmt"

func displayIssues(issues []Issue, description string) {
	if len(issues) == 0 {
		fmt.Printf("No %s found", description)
		if statusFilter != "" {
			fmt.Printf(" with status '%s'", statusFilter)
		}
		fmt.Println(".")
		return
	}

	fmt.Printf("Found %d %s", len(issues), description)
	if statusFilter != "" {
		fmt.Printf(" with status '%s'", statusFilter)
	}
	fmt.Println(":")
	fmt.Println()

	for i, issue := range issues {
		fmt.Printf("%d. %s\n", i+1, issue.Title)
		fmt.Printf("   Team: %s | State: %s", issue.Team.Name, issue.State.Name)
		if issue.Assignee.Name != "" {
			fmt.Printf(" | Assignee: %s", issue.Assignee.Name)
		}
		if issue.Priority > 0 {
			fmt.Printf(" | Priority: %d", issue.Priority)
		}
		fmt.Printf("\n   ID: %s", issue.ID)
		
		if showDescription && issue.Description != "" {
			fmt.Printf("\n   Description: %s", issue.Description)
		}
		fmt.Printf("\n\n")
	}
}
