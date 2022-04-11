/*
This is the root level command. Its function is to bring up a screen.
*/
package commands

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	input  []rune
	cursor int
}

func initialModel() model {
	return model{
		input:  []rune("abcdefg"),
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
		case tea.KeyEscape, 'q':
			return m, tea.Quit
		case tea.KeyLeft, 'h':
			if m.cursor > 0 {
				m.cursor--
			}
		case tea.KeyRight, 'l':
			if m.cursor < len(m.input)-1 {
				m.cursor++
			}
		case tea.KeyBackspace:
			if m.cursor > 0 {
				m.input = append(m.input[:(m.cursor-1)], m.input[(m.cursor):]...)
				m.cursor--
			}
		case tea.KeyRunes:
			input := append(m.input[0:(m.cursor)], msg.Runes...)
			if m.cursor < len(m.input)-1 {
				input = append(input, m.input[m.cursor:len(m.input)]...)
			}
			m.input = input
			m.cursor++
		}
	}
	return m, nil
}

var cursorStyle = lipgloss.NewStyle().Background(lipgloss.Color("#fa695f"))

func (m model) View() string {
	s := "Type a string: "
	if m.cursor > 0 {
		s += string(m.input[:m.cursor])
	}
	s += cursorStyle.Render(string(m.input[m.cursor]))
	s += string(m.input[(m.cursor + 1):])
	return s
}

func Default() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
