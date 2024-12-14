package main

// A simple example demonstrating the use of multiple text input components
// from the Bubbles component library.

import (
	"fmt"
	"os"
	"terminal-illness/models"
	"terminal-illness/utils"

	tea "github.com/charmbracelet/bubbletea"
)

// Main function to run the menu
func main() {
	utils.ClearTerminal() // Clear the terminal before starting the application
	menu := models.InitialLoadingModel()
	p := tea.NewProgram(menu,
		tea.WithAltScreen(),       // Use alternate screen buffer
		tea.WithMouseCellMotion(), // Enable mouse support
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
