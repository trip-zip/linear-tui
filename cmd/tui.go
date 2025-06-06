package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Run the interactive TUI",
	Long:  `Launch the interactive terminal user interface for browsing Linear issues.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runTUI(); err != nil {
			fmt.Printf("Error running TUI: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
