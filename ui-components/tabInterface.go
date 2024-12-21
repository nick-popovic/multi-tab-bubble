package ui_components

import (
	tea "github.com/charmbracelet/bubbletea"
)

// TabModel represents the interface for a tab component in the UI.
// It defines the necessary methods that a tab should implement to
// initialize, update, render its view, and provide its title.
type TabModel interface {
	// Init initializes the tab and returns an initial command.
	Init() tea.Cmd

	// Update handles incoming messages and updates the tab's state.
	// It returns the updated model and an optional command.
	Update(tea.Msg) (tea.Model, tea.Cmd)

	// View renders the tab's view as a string.
	View() string

	// Title returns the title of the tab.
	Title() string

	// SetActive handles the activation state of the tab.
	SetActive(bool)
}
