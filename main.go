package main

import (
	"fmt"
	"os"

	ui "main/ui-components"

	tea "github.com/charmbracelet/bubbletea"
)

// model is the top-level model for the application.
type model struct {
	tabs ui.TabsModel
}

func New() model {
	return model{tabs: *ui.NewTabsModel(
		[]ui.TabModel{
			ui.NewTab1(),
			ui.NewTab2(),
			ui.NewTab3()},
		true)}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "u":
			return m, tea.ClearScreen
		}
	}

	// Update components
	_, cmd = m.tabs.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return m.tabs.View() +
		"\n\nPress left/right to switch tabs" +
		"\nPress q to quit, i to add '<' to the active view."
}

func main() {
	p := tea.NewProgram(
		New(),
		tea.WithAltScreen(), // Use alternate screen buffer
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
