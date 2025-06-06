package cmd

import (
	"fmt"

	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 1)

	itemStyle = lipgloss.NewStyle().
			PaddingLeft(4)

	selectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Foreground(lipgloss.Color("170"))

	paginationStyle = lipgloss.NewStyle().
			PaddingLeft(4)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
)

type model struct {
	issues   []Issue
	cursor   int
	selected map[int]struct{}
	loading  bool
	err      error
}

func initialModel() model {
	return model{
		selected: make(map[int]struct{}),
		loading:  true,
	}
}

func (m model) Init() tea.Cmd {
	return loadIssues
}

func loadIssues() tea.Msg {
	client, err := getLinearClient()
	if err != nil {
		return errMsg{err}
	}

	issues, err := client.GetIssues()
	if err != nil {
		return errMsg{err}
	}

	return issuesLoadedMsg{issues}
}

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

type issuesLoadedMsg struct{ issues []Issue }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.issues)-1 {
				m.cursor++
			}

		case "enter", " ":
			if len(m.issues) > 0 {
				_, ok := m.selected[m.cursor]
				if ok {
					delete(m.selected, m.cursor)
				} else {
					m.selected[m.cursor] = struct{}{}
				}
			}
		}

	case issuesLoadedMsg:
		m.issues = msg.issues
		m.loading = false

	case errMsg:
		m.err = msg.err
		m.loading = false
	}

	return m, nil
}

func (m model) View() string {
	if m.loading {
		return "\n  Loading Linear issues...\n"
	}

	if m.err != nil {
		return fmt.Sprintf("\n  Error: %v\n", m.err)
	}

	s := titleStyle.Render("Linear Issues")
	s += "\n\n"

	if len(m.issues) == 0 {
		s += "  No issues found.\n"
	} else {
		for i, issue := range m.issues {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			checked := " "
			if _, ok := m.selected[i]; ok {
				checked = "x"
			}

			line := fmt.Sprintf("%s [%s] %s - %s (%s)", cursor, checked, issue.Title, issue.State.Name, issue.Team.Name)

			if m.cursor == i {
				s += selectedItemStyle.Render(line)
			} else {
				s += itemStyle.Render(line)
			}
			s += "\n"
		}
	}

	s += "\n" + helpStyle.Render("Press q to quit, ↑/↓ to navigate, space to select")
	return s
}

func runTUI() error {
	p := tea.NewProgram(initialModel())
	_, err := p.Run()
	return err
}
