# Linear TUI

A dual CLI and Terminal User Interface (TUI) application for Linear.app integration. Built with Go and designed to enable programmatic access to Linear issues for development workflows.

## Features

- **Interactive TUI**: Browse Linear issues with a beautiful terminal interface using bubbletea
- **CLI Commands**: Programmatic access to Linear data for automation and scripting
- **Status Filtering**: Filter issues by any Linear state (In Progress, Backlog, Done, etc.)
- **User-specific Queries**: Get issues assigned to you using the Linear API's viewer functionality
- **Claude Code Integration**: Designed to work seamlessly with Claude Code for AI-assisted development

## Installation

### Prerequisites
- Go 1.24.3 or later
- Linear.app account with API access
- Linear Personal API Key

### Setup

1. Clone or download this repository
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Create a `.env` file with your Linear API key:
   ```
   LINEAR_API_KEY=your_personal_api_key_here
   ```

## Usage

### Interactive TUI Mode
```bash
# Run the interactive terminal interface
go run .
```

### CLI Commands

#### Get Your Assigned Issues
```bash
# List all issues assigned to you
go run . me

# Filter your issues by status
go run . me -s "In Progress"
go run . me -s "Backlog"
go run . me -s "Green Light"
```

#### Get All Workspace Issues
```bash
# List all issues in the workspace
go run . list

# Filter all issues by status
go run . list -s "Done"
go run . list -s "Triage"
```

#### Help
```bash
# Show general help
go run . --help

# Show command-specific help
go run . me --help
go run . list --help
```

### Common Status Values
- `In Progress` - Currently being worked on
- `Backlog` - Planned work
- `Green Light` - Approved and ready to implement
- `Done` - Completed work
- `Triage` - Needs review/prioritization
- `Proposals` - Ideas under consideration

## Examples

### Developer Workflow
```bash
# Check what you're currently working on
go run . me -s "In Progress"

# See what's ready to work on next
go run . me -s "Green Light"

# Review your backlog
go run . me -s "Backlog"

# Find all high-priority work in the team
go run . list -s "In Progress"
```

### Claude Code Integration
This tool is specifically designed to work with Claude Code for AI-assisted development:

```bash
# Claude can fetch your current work
go run . me -s "In Progress"

# Claude can find approved features to implement
go run . list -s "Green Light" 

# Claude can help prioritize from your backlog
go run . me -s "Backlog"
```

## Output Format

Issues are displayed with:
- **Title**: The issue title
- **Team**: Which Linear team owns the issue
- **State**: Current status
- **Assignee**: Who it's assigned to (if anyone)
- **Priority**: Priority level (if set)
- **ID**: Linear issue ID for API reference

## Building

```bash
# Build the application
go build -o linear-tui

# Run the built binary
./linear-tui me -s "In Progress"
```

## Development

The application is structured with:
- `main.go` - Entry point and environment setup
- `cmd.go` - CLI command definitions using Cobra
- `linear.go` - Linear GraphQL API client
- `tui.go` - Bubbletea terminal interface
- `CLAUDE.md` - Instructions for Claude Code integration

## License

This project is for internal use and Linear.app integration.