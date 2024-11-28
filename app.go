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

// var (
// 	graffTitle = `

// 	__                            .__                 .__     .__ .__   .__
// 	_/  |_   ____  _______   _____  |__|  ____  _____   |  |    |__||  |  |  |    ____    ____    ______  ______
// 	\   __\_/ __ \ \_  __ \ /     \ |  | /    \ \__  \  |  |    |  ||  |  |  |   /    \ _/ __ \  /  ___/ /  ___/
// 	 |  |  \  ___/  |  | \/|  Y Y  \|  ||   |  \ / __ \_|  |__  |  ||  |__|  |__|   |  \\  ___/  \___ \  \___ \
// 	 |__|   \___  > |__|   |__|_|  /|__||___|  /(____  /|____/  |__||____/|____/|___|  / \___  >/____  >/____  >
// 				\/               \/          \/      \/                             \/      \/      \/      \/

// `
// 	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
// 	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
// 	cursorStyle         = focusedStyle
// 	noStyle             = lipgloss.NewStyle()
// 	helpStyle           = blurredStyle
// 	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

// 	focusedButton = focusedStyle.Render("[ Submit ]")
// 	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
// )

// type model struct {
// 	focusIndex int
// 	inputs     []textinput.Model
// 	cursorMode cursor.Mode
// }

// func initialModel() model {
// 	m := model{
// 		inputs: make([]textinput.Model, 3),
// 	}

// 	var t textinput.Model
// 	for i := range m.inputs {
// 		t = textinput.New()
// 		t.Cursor.Style = cursorStyle
// 		t.CharLimit = 32

// 		switch i {
// 		case 0:
// 			t.Placeholder = "Nickname"
// 			t.Focus()
// 			t.PromptStyle = focusedStyle
// 			t.TextStyle = focusedStyle
// 		case 1:
// 			t.Placeholder = "Email"
// 			t.CharLimit = 64
// 		case 2:
// 			t.Placeholder = "Password"
// 			t.EchoMode = textinput.EchoPassword
// 			t.EchoCharacter = '•'
// 		}

// 		m.inputs[i] = t
// 	}

// 	return m
// }

// func (m model) Init() tea.Cmd {
// 	return textinput.Blink
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "ctrl+c", "esc":
// 			return m, tea.Quit

// 		// Change cursor mode
// 		case "ctrl+r":
// 			m.cursorMode++
// 			if m.cursorMode > cursor.CursorHide {
// 				m.cursorMode = cursor.CursorBlink
// 			}
// 			cmds := make([]tea.Cmd, len(m.inputs))
// 			for i := range m.inputs {
// 				cmds[i] = m.inputs[i].Cursor.SetMode(m.cursorMode)
// 			}
// 			return m, tea.Batch(cmds...)

// 		// Set focus to next input
// 		case "tab", "shift+tab", "enter", "up", "down":
// 			s := msg.String()

// 			// Did the user press enter while the submit button was focused?
// 			// If so, exit.
// 			if s == "enter" && m.focusIndex == len(m.inputs) {
// 				return m, tea.Quit
// 			}

// 			// Cycle indexes
// 			if s == "up" || s == "shift+tab" {
// 				m.focusIndex--
// 			} else {
// 				m.focusIndex++
// 			}

// 			if m.focusIndex > len(m.inputs) {
// 				m.focusIndex = 0
// 			} else if m.focusIndex < 0 {
// 				m.focusIndex = len(m.inputs)
// 			}

// 			cmds := make([]tea.Cmd, len(m.inputs))
// 			for i := 0; i <= len(m.inputs)-1; i++ {
// 				if i == m.focusIndex {
// 					// Set focused state
// 					cmds[i] = m.inputs[i].Focus()
// 					m.inputs[i].PromptStyle = focusedStyle
// 					m.inputs[i].TextStyle = focusedStyle
// 					continue
// 				}
// 				// Remove focused state
// 				m.inputs[i].Blur()
// 				m.inputs[i].PromptStyle = noStyle
// 				m.inputs[i].TextStyle = noStyle
// 			}

// 			return m, tea.Batch(cmds...)
// 		}
// 	}

// 	// Handle character input and blinking
// 	cmd := m.updateInputs(msg)

// 	return m, cmd
// }

// func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
// 	cmds := make([]tea.Cmd, len(m.inputs))

// 	// Only text inputs with Focus() set will respond, so it's safe to simply
// 	// update all of them here without any further logic.
// 	for i := range m.inputs {
// 		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
// 	}

// 	return tea.Batch(cmds...)
// }

// func (m model) View() string {
// 	var b strings.Builder

// 	for i := range m.inputs {
// 		b.WriteString(m.inputs[i].View())
// 		if i < len(m.inputs)-1 {
// 			b.WriteRune('\n')
// 		}
// 	}

// 	button := &blurredButton
// 	if m.focusIndex == len(m.inputs) {
// 		button = &focusedButton
// 	}
// 	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

// 	b.WriteString(helpStyle.Render("cursor mode is "))
// 	b.WriteString(cursorModeHelpStyle.Render(m.cursorMode.String()))
// 	b.WriteString(helpStyle.Render(" (ctrl+r to change style)"))

// 	return b.String()
// }

// type menuModel struct {
// 	options  []string
// 	selected int
// }

// // Initialize the menu model
// func initialMenuModel() menuModel {
// 	return menuModel{
// 		options: []string{"Fetal Alcohol Syndrome (API Test)", "PTSD (API Load Test)", "Schizophrenia (API Probe)", "Settings", "Quit"},
// 	}
// }

// // Add a new model for the input screen
// type inputModel struct {
// 	model // Embed the existing model for inputs
// }

// // Update method for menu navigation
// func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "down":
// 			m.selected = (m.selected + 1) % len(m.options)
// 		case "up":
// 			m.selected = (m.selected - 1 + len(m.options)) % len(m.options)
// 		case "enter":
// 			// Handle menu selection
// 			switch m.selected {
// 			case 0:
// 				// Start environment selection
// 				return initialEnvModel(), nil // Transition to the environment model
// 			case 1:
// 				// Start API Blaster
// 			case 2:
// 				// Start API Probe
// 			case 3:
// 				// Open Settings
// 			}
// 		}
// 	}
// 	return m, nil
// }

// // Create a light gray style for the author description
// var authorStyle lipgloss.Style

// func init() {
// 	// Create a darker gray style for the author description with smaller font
// 	authorStyle = lipgloss.NewStyle().
// 		Foreground(lipgloss.Color("235"))
// }

// // View method for the menu
// func (m menuModel) View() string {
// 	var b strings.Builder
// 	// Create a bordered title
// 	title := lipgloss.NewStyle().
// 		Bold(true).
// 		Foreground(lipgloss.Color("205")).
// 		Border(lipgloss.NormalBorder(), true).
// 		Padding(1, 2).
// 		Render(graffTitle + "\n" + authorStyle.Render("author: ben \"legdonor\" donnaloia"))

// 	b.WriteString(title + "\n\n") // Add the bordered title

// 	for i, option := range m.options {
// 		if i == m.selected {
// 			b.WriteString(fmt.Sprintf("> %s\n", lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Render(option))) // Highlight selected option in purple
// 		} else {
// 			b.WriteString(fmt.Sprintf("  %s\n", lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Render(option))) // Regular option in purple
// 		}
// 	}
// 	return b.String()
// }

// // Add the Init method to menuModel
// func (m menuModel) Init() tea.Cmd {
// 	return nil // No initial command needed
// }

// // Update the envModel to include authentication options
// type envModel struct {
// 	options      []string
// 	authOptions  []string
// 	selected     int
// 	authSelected int // Track the selected auth option
// }

// // Initialize the environment model with auth options
// func initialEnvModel() envModel {
// 	return envModel{
// 		options:     []string{"dev", "prod"},
// 		authOptions: []string{"auth", "no auth"},
// 	}
// }

// // Update method for environment selection
// func (m envModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "down":
// 			m.selected = (m.selected + 1) % len(m.options)
// 		case "up":
// 			m.selected = (m.selected - 1 + len(m.options)) % len(m.options)
// 		case "enter":
// 			// Handle environment selection
// 			return initialModel(), nil // Transition to the input model
// 		case "right":
// 			m.authSelected = (m.authSelected + 1) % len(m.authOptions) // Move right in auth options
// 		case "left":
// 			m.authSelected = (m.authSelected - 1 + len(m.authOptions)) % len(m.authOptions) // Move left in auth options
// 		}
// 	}
// 	return m, nil
// }

// // View method for the environment selection
// func (m envModel) View() string {
// 	var b strings.Builder

// 	// Create a bordered title for the environment selection
// 	title := lipgloss.NewStyle().
// 		Bold(true).
// 		Foreground(lipgloss.Color("205")).
// 		Border(lipgloss.NormalBorder(), true).
// 		Padding(1, 2).
// 		Render("Select Environment")

// 	b.WriteString(title + "\n\n") // Add the bordered title

// 	// Environment options
// 	for i, option := range m.options {
// 		if i == m.selected {
// 			b.WriteString(fmt.Sprintf("> %s  ", option)) // Highlight selected option
// 		} else {
// 			b.WriteString(fmt.Sprintf("  %s  ", option))
// 		}
// 	}

// 	// Add space between the two menus
// 	b.WriteString("\n\n")

// 	// Authentication options
// 	for i, option := range m.authOptions {
// 		if i == m.authSelected {
// 			b.WriteString(fmt.Sprintf("> %s  ", option)) // Highlight selected auth option
// 		} else {
// 			b.WriteString(fmt.Sprintf("  %s  ", option))
// 		}
// 	}

// 	b.WriteString("\n") // Add a newline at the end for better formatting

// 	return b.String()
// }

// // Add the Init method to envModel
// func (m envModel) Init() tea.Cmd {
// 	return nil // No initial command needed
// }

// // Function to clear the terminal screen
// func clearTerminal() {
// 	cmd := exec.Command("clear") // For Unix/Linux/Mac
// 	if strings.Contains(strings.ToLower(os.Getenv("OS")), "windows") {
// 		cmd = exec.Command("cmd", "/c", "cls") // For Windows
// 	}
// 	cmd.Stdout = os.Stdout
// 	cmd.Run()
// }

// Main function to run the menu
func main() {
	utils.ClearTerminal() // Clear the terminal before starting the application
	menu := models.InitialMenuModel()
	if _, err := tea.NewProgram(menu).Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
