/*
This is the root level command. Its function is to bring up a screen.
*/
package commands

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/skrider/riptex/buffer"
)

type model struct {
	input  buffer.Buffer
	cursor int
}

func initialModel() model {
	return model{
		input:  buffer.NewBuffer(),
		cursor: 0,
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
		case tea.KeyRunes:
			m.input.Insert(msg.Runes)
		}
	}
	return m, nil
}

// var cursorStyle = lipgloss.NewStyle().Background(lipgloss.Color("#fa695f"))

func (m model) View() string {
	var s strings.Builder
	fmt.Fprint(&s, "Write something: ")
	fmt.Fprint(&s, m.input.String())
	return s.String()
}

func Default() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
