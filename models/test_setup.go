package models

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Test Setup - This is the first screen out of two for setting up our api test.
// This screen just asks a few basic questions about auth/no-auth and localhost/prod

// Update the TestSetupModel to include authentication options
type TestSetupModel struct {
	options      []string
	authOptions  []string
	selected     int
	authSelected int // Track the selected auth option
}

// Initialize the environment model with auth options
func InitialTestSetupModel() TestSetupModel {
	return TestSetupModel{
		options:     []string{"dev", "prod"},
		authOptions: []string{"auth", "no auth"},
	}
}

// Update method for environment selection
func (m TestSetupModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "down":
			m.selected = (m.selected + 1) % len(m.options)
		case "up":
			m.selected = (m.selected - 1 + len(m.options)) % len(m.options)
		case "enter":
			// Handle environment selection
			return InitialTestDetailModel(), nil // Transition to the input model
		case "right":
			m.authSelected = (m.authSelected + 1) % len(m.authOptions) // Move right in auth options
		case "left":
			m.authSelected = (m.authSelected - 1 + len(m.authOptions)) % len(m.authOptions) // Move left in auth options
		}
	}
	return m, nil
}

// View method for the environment selection
func (m TestSetupModel) View() string {
	var b strings.Builder

	// Create a bordered title for the environment selection
	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		Border(lipgloss.NormalBorder(), true).
		Padding(1, 2).
		Render("Select Environment")

	b.WriteString(title + "\n\n") // Add the bordered title

	// Environment options
	for i, option := range m.options {
		if i == m.selected {
			b.WriteString(fmt.Sprintf("> %s  ", option)) // Highlight selected option
		} else {
			b.WriteString(fmt.Sprintf("  %s  ", option))
		}
	}

	// Add space between the two menus
	b.WriteString("\n\n")

	// Authentication options
	for i, option := range m.authOptions {
		if i == m.authSelected {
			b.WriteString(fmt.Sprintf("> %s  ", option)) // Highlight selected auth option
		} else {
			b.WriteString(fmt.Sprintf("  %s  ", option))
		}
	}

	b.WriteString("\n") // Add a newline at the end for better formatting

	return b.String()
}

// Add the Init method to TestSetupModel
func (m TestSetupModel) Init() tea.Cmd {
	return nil // No initial command needed
}
