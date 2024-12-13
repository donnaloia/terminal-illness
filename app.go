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
	menu := models.InitialModel()
	if _, err := tea.NewProgram(menu).Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
