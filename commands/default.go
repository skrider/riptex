/*
This is the root level command. Its function is to bring up a screen.
*/
package commands

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	input textinput.Model
}

func initialModel() model {
	input_model := textinput.New()
	input_model.Focus()
	return model{
		input: input_model,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEscape:
			return m, tea.Quit
		default:
			m.input, _ = m.input.Update(msg)
		}
	}
	return m, nil
}

func (m model) View() string {
	return m.input.View()
}

func Default() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
