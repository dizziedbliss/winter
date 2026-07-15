package ui

import (
	"fmt"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type errMsg error

type model struct {
	spinner  spinner.Model
	quitting bool
	err      error
}

func initModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return model{spinner: s}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case errMsg:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	}
}

func (m model) View() tea.View {
	if m.err != nil {
		return tea.NewView(m.err.Error())
	}
	str := fmt.Sprintf("\n\n  %s Deploying...\n\n", m.spinner.View())

	if m.quitting {
		return tea.NewView(str + "\n")
	}
	return tea.NewView(str)
}

func runSpinner() error {
	p := tea.NewProgram(initModel())
	if _, err := p.Run(); err != nil {
		return fmt.Errorf("failed to run spinner: %v", err)
	}
	return nil
}
