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
- Go 1.21 or later
- Linear.app account with API access
- Linear Personal API Key

### Quick Install

Install directly from GitHub:
```bash
go install github.com/trip-zip/linear-tui/cmd/linear@latest
```

After installation, the `linear` command will be available in your terminal.

### Setup

1. Create a `.env` file in your working directory or home config:
   ```bash
   # Option 1: In your current project directory
   echo "LINEAR_API_KEY=your_personal_api_key_here" > .env
   
   # Option 2: In your home config (works globally)
   mkdir -p ~/.config/linear
   echo "LINEAR_API_KEY=your_personal_api_key_here" > ~/.config/linear/.env
   ```

2. Get your API key from [Linear Settings > API](https://linear.app/settings/api)

### Development Install

For development or local modifications:
```bash
git clone https://github.com/trip-zip/linear-tui.git
cd linear-tui
go install ./cmd/linear
```

## Usage

### Interactive TUI Mode
```bash
# Run the interactive terminal interface
linear
```

### CLI Commands

#### Get Your Assigned Issues
```bash
# List all issues assigned to you
linear me

# Filter your issues by status
linear me -s "In Progress"
linear me -s "Backlog"
linear me -s "Green Light"
```

#### Get All Workspace Issues
```bash
# List all issues in the workspace
linear list

# Filter all issues by status
linear list -s "Done"
linear list -s "Triage"
```

#### Help
```bash
# Show general help
linear --help

# Show command-specific help
linear me --help
linear list --help
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
linear me -s "In Progress"

# See what's ready to work on next
linear me -s "Green Light"

# Review your backlog
linear me -s "Backlog"

# Find all high-priority work in the team
linear list -s "In Progress"
```

### Claude Code Integration
This tool is specifically designed to work with Claude Code for AI-assisted development:

```bash
# Claude can fetch your current work
linear me -s "In Progress"

# Claude can find approved features to implement
linear list -s "Green Light" 

# Claude can help prioritize from your backlog
linear me -s "Backlog"
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
go build -o linear ./cmd/linear

# Run the built binary
./linear me -s "In Progress"
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