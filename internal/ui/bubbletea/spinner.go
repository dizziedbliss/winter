package bubbletea

import (
	"fmt"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type Model struct {
	spinner  spinner.Model
	quitting bool
	err      error
	status   string
	logs    []string
}

func initModel() Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return Model{spinner: s}
}

func (m Model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case StatusMsg:
		m.status = string(msg)
		return m, nil

	case ErrMsg:
		m.err = msg
		return m, nil

	case SuccessMsg:
		m.status = string(msg)
		return m, tea.Quit

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	}
}

func (m Model) View() tea.View {
	if m.err != nil {
		return tea.NewView(m.err.Error())
	}
	str := fmt.Sprintf("\n%s %s\n", m.spinner.View(), m.status)

	if m.quitting {
		return tea.NewView(str + "\n")
	}
	return tea.NewView(str)
}

