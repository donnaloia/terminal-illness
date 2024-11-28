package models

import (
	"fmt"
	"strings"
	"terminal-illness/utils"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Main Menu - This is the entrance point to the CLI.
// This screen provides a menu to select the different types of API tests.

type menuModel struct {
	options  []string
	selected int
}

// Initialize the menu model
func InitialMenuModel() menuModel {
	return menuModel{
		options: []string{"Fetal Alcohol Syndrome (API Test)", "PTSD (API Load Test)", "Schizophrenia (API Probe)", "Settings", "Quit"},
	}
}

// View method for the menu
func (m menuModel) View() string {
	var b strings.Builder
	// Create a bordered title
	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		Border(lipgloss.NormalBorder(), true).
		Padding(1, 2).
		Render(utils.GraffTitle + "\n" + utils.AuthorStyle.Render("author: ben \"legdonor\" donnaloia"))

	b.WriteString(title + "\n\n") // Add the bordered title

	for i, option := range m.options {
		if i == m.selected {
			b.WriteString(fmt.Sprintf("> %s\n", lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Render(option))) // Highlight selected option in purple
		} else {
			b.WriteString(fmt.Sprintf("  %s\n", lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Render(option))) // Regular option in purple
		}
	}
	return b.String()
}

// Update method for menu navigation
func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "down":
			m.selected = (m.selected + 1) % len(m.options)
		case "up":
			m.selected = (m.selected - 1 + len(m.options)) % len(m.options)
		case "enter":
			// Handle menu selection
			switch m.selected {
			case 0:
				// Start environment selection
				// this is the model that gets called that renders a new screen from the models view function
				// rename initialEnvModel to APITestSetupModel
				return InitialTestSetupModel(), nil // Transition to the environment model
			case 1:
				// Start API Blaster
			case 2:
				// Start API Probe
			case 3:
				// Open Settings
			}
		}
	}
	return m, nil
}

func (m menuModel) Init() tea.Cmd {
	return nil // Return nil or any initial command if needed
}
