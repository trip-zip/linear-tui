package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	statusFilter string
	showAll      bool
)

var rootCmd = &cobra.Command{
	Use:   "linear-cli",
	Short: "A CLI for Linear.app integration",
	Long: `Linear CLI is a command-line interface for interacting with Linear.app. 
It allows you to fetch issues, filter by status, and view detailed information 
about your Linear workspace.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Show help if no subcommand is provided
		cmd.Help()
	},
}

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

		// Get current user info
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
	fmt.Println(":\n")

	for i, issue := range issues {
		fmt.Printf("%d. %s\n", i+1, issue.Title)
		fmt.Printf("   Team: %s | State: %s", issue.Team.Name, issue.State.Name)
		if issue.Assignee.Name != "" {
			fmt.Printf(" | Assignee: %s", issue.Assignee.Name)
		}
		if issue.Priority > 0 {
			fmt.Printf(" | Priority: %d", issue.Priority)
		}
		fmt.Printf("\n   ID: %s\n\n", issue.ID)
	}
}

func initCLI() {
	// Add status filter flag to both list and me commands
	listCmd.Flags().StringVarP(&statusFilter, "status", "s", "", "Filter issues by status (e.g., 'In Progress', 'Backlog', 'Done')")
	meCmd.Flags().StringVarP(&statusFilter, "status", "s", "", "Filter your issues by status (e.g., 'In Progress', 'Backlog', 'Done')")

	// Add commands to root
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(meCmd)

	// Add help information
	rootCmd.SetHelpTemplate(`{{.Long}}

Usage:
  {{.UseLine}}

Available Commands:
  list        List all issues in the workspace
  me          List issues assigned to you
  help        Help about any command

Flags:
  -h, --help   Help for {{.Name}}

Use "{{.CommandPath}} [command] --help" for more information about a command.

Examples:
  linear-cli list               # List all issues
  linear-cli list -s "Backlog"  # List all backlog issues
  linear-cli me                 # List your assigned issues
  linear-cli me -s "In Progress" # List your in-progress issues
`)
}

func executeCLI() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}