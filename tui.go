package todolist

import (
	tea "github.com/charmbracelet/bubbletea"
)

type ListView struct {
	list List
}

// Init implements tea.Model.
func (l ListView) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (l ListView) Update(tea.Msg) (tea.Model, tea.Cmd) {
	panic("unimplemented")
}

// View implements tea.Model.
func (l ListView) View() string {
	return ""
}

