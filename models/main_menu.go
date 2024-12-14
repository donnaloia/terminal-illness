package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"terminal-illness/api_requests"
	"terminal-illness/utils"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Main Menu - This is the entrance point to the CLI.
// This screen provides a menu to select the different types of API tests.

// Main application model
type MainModel struct {
	first           SidebarModel
	second          Model
	third           SendRequestModel
	focus           string
	showOverlay     bool
	viewport        viewport.Model
	requestDuration time.Duration
	currentURL      string
	status          string
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c": // Check for the 'q' key to quit
			utils.ClearTerminal()
			return m, tea.Quit // Return the quit command
		case "tab", "right", "enter": // Use the Tab key to switch focus
			if m.focus == "sidebar" {
				switch m.first.selected {
				case 4:
					utils.ClearTerminal()
					return m, tea.Quit // Return the quit command
				}
				m.focus = "setup"
				if setupModel, ok := m.second.(SetupModel); ok {
					setupModel.Focus()
					m.second = setupModel
				}
				return m, nil
			} else if m.focus == "setup" {
				if msg.String() == "enter" {
					if setupModel, ok := m.second.(SetupModel); ok && setupModel.selected == 2 {
						m.focus = "third"
						m.third.Focus()
						setupModel.Blur()
						m.second = setupModel
						return m, nil
					}
				}
				updatedModel, cmd := m.second.Update(msg)
				if updatedSetupModel, ok := updatedModel.(SetupModel); ok {
					m.second = updatedSetupModel
				}
				return m, cmd
			} else if m.focus == "third" {
				if msg.String() == "enter" {
					var url string
					switch model := m.second.(type) {
					case SetupModel:
						url = model.inputs[0].Value()
					default:
						url = "https://api.thecatapi.com/v1/images/search?limit=10"
					}
					m.currentURL = url

					startTime := time.Now()
					resp, err := api_requests.MakeRequest(
						url,
						api_requests.GET,
						"", // Add bearer token if needed
					)
					m.requestDuration = time.Since(startTime)

					if err != nil {
						m.status = err.Error()
					} else {
						status, _ := api_requests.ReadStatus(resp)
						body, err := api_requests.ReadResponse(resp)

						m.status = status
						if err != nil {
							m.status = err.Error()
						} else {
							// Pretty print the JSON
							var prettyJSON bytes.Buffer
							if err := json.Indent(&prettyJSON, []byte(body), "", "  "); err == nil {
								// Add color to JSON keys and values
								content := prettyJSON.String()
								lines := strings.Split(content, "\n")
								for i, line := range lines {
									// Find JSON keys (text before :)
									if idx := strings.Index(line, ":"); idx != -1 {
										key := line[:idx]
										rest := line[idx:]
										// Style the key in dark purple (color "90") and the value in dark gray (color "240")
										lines[i] = lipgloss.NewStyle().Foreground(lipgloss.Color("90")).Render(key) +
											":" +
											lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render(rest[1:])
									}
								}
								m.viewport.SetContent(strings.Join(lines, "\n"))
							} else {
								m.viewport.SetContent(body) // Fallback to raw if JSON parsing fails
							}
							m.showOverlay = true
						}
					}
					return m, nil
				}
			}
		case "left":
			if m.focus == "setup" {
				if setupModel, ok := m.second.(SetupModel); ok {
					// Get the current input value
					currentValue := setupModel.inputs[setupModel.selected].Value()

					// If we're at the start of the input (empty) and it's the first field
					if len(currentValue) == 0 && setupModel.selected == 0 {
						setupModel.Blur()
						m.focus = "sidebar"
						m.second = setupModel
						return m, nil
					}

					// Otherwise, let the input handle the left arrow key
					updatedModel, cmd := setupModel.Update(msg)
					m.second = updatedModel.(SetupModel)
					return m, cmd
				}
			} else if m.focus == "third" {
				m.focus = "setup"
				m.third.Blur()
				if setupModel, ok := m.second.(SetupModel); ok {
					setupModel.Focus()
					m.second = setupModel

				}
			}
		case "up", "k":
			if m.showOverlay {
				m.viewport.LineUp(1)
			} else if m.focus == "sidebar" {
				m.first.selected = (m.first.selected - 1 + len(m.first.options)) % len(m.first.options)
			} else if m.focus == "setup" {
				updatedModel, cmd := m.second.Update(msg)
				m.second = updatedModel.(SetupModel)
				return m, cmd
			}
		case "pgup":
			if m.showOverlay {
				m.viewport.LineUp(1)
			}
		case "pgdown", "down", "j":
			if m.showOverlay {
				m.viewport.LineDown(1)
			} else if m.focus == "sidebar" {
				m.first.selected = (m.first.selected + 1) % len(m.first.options)
			} else if m.focus == "setup" {
				updatedModel, cmd := m.second.Update(msg)
				m.second = updatedModel.(SetupModel)
				return m, cmd
			}
		case "esc":
			m.showOverlay = false
			return m, nil
		case "home":
			if m.showOverlay {
				m.viewport.GotoTop()
			}
		case "end", "G":
			if m.showOverlay {
				m.viewport.GotoBottom()
			}
		default:
			updatedModel, cmd := m.second.Update(msg) // Call Update on the SetupModel
			m.second = updatedModel.(SetupModel)      // Update the second model
			return m, cmd

		}

	}
	switch m.focus {
	case "sidebar":
		// var cmd tea.Cmd
		updatedModel, _ := m.first.Update(msg) // Call Update on the sidebar model
		m.first = updatedModel.(SidebarModel)

		switch m.first.selected {
		// case api test selected
		case 0:
			m.second = InitialSetupModel(0)
			switch model := updatedModel.(type) {
			case SetupModel:
				m.second = model // Safe assignment
			}
		// case load test
		case 1:
			m.second = InitialAuditSetupModel(0)
			switch model := updatedModel.(type) {
			case AuditSetupModel:
				// Handle DemoModel case if needed
				m.second = model // Safe assignment
			}
			// Type assertion to SidebarModel
			return m, nil
		case 2:
			m.second = InitialLoadSetupModel(0)
			switch model := updatedModel.(type) {
			case LoadSetupModel:
				m.second = model
			}
			// Type assertion to SidebarModel
			return m, nil
		case 3:
			// Settings: todo

		}

	}
	return m, nil
}

func (m MainModel) View() string {
	baseView := lipgloss.JoinVertical(
		lipgloss.Left,
		Header(utils.GraffTitle),
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.first.View(),
			m.second.View(),
			m.third.View(),
		),
	)

	if m.showOverlay {
		headerStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Bold(true).
			Padding(0, 1)

		dividerStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Render(strings.Repeat("─", 78))

		// Status styling
		statusStyle := lipgloss.NewStyle()
		if strings.HasPrefix(m.status, "2") {
			statusStyle = statusStyle.Foreground(lipgloss.Color("10"))
		} else if strings.HasPrefix(m.status, "4") || strings.HasPrefix(m.status, "5") {
			statusStyle = statusStyle.Foreground(lipgloss.Color("9"))
		}

		// Duration styling
		durationStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("12"))

		content := lipgloss.JoinVertical(lipgloss.Left,
			headerStyle.Render("Test Results"),
			dividerStyle,
			fmt.Sprintf("Endpoint: %s", lipgloss.NewStyle().Foreground(lipgloss.Color("223")).Render(m.currentURL)),
			fmt.Sprintf("Status: %s", statusStyle.Render(m.status)),
			fmt.Sprintf("Duration: %s", durationStyle.Render(fmt.Sprintf("%dms", m.requestDuration.Milliseconds()))),
			"Body:",
			lipgloss.NewStyle().
				Foreground(lipgloss.Color("240")).
				Render(m.viewport.View()),
		)

		styledContent := lipgloss.NewStyle().
			Width(80).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("205")).
			Padding(1).
			Render(content)

		contentHeight := lipgloss.Height(styledContent)

		return lipgloss.Place(
			116, contentHeight,
			lipgloss.Center, lipgloss.Center,
			styledContent,
		)
	}

	return baseView
}

func Header(title string) string {
	return lipgloss.NewStyle().
		Width(116).
		Bold(true).
		Foreground(lipgloss.Color("205")).
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2).
		Render(title + "\n" + utils.AuthorStyle.Render("author: legdonor"))
}

type Model interface {
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
	View() string
}

func InitialModel() MainModel {
	m := MainModel{
		first:  SidebarModel{options: []string{"Paranoia (API Client Test)", "Schizophrenia (API Security Audit)", "PTSD (API Load Test)", "Settings", "Quit"}},
		second: InitialSetupModel(0),
		third:  SendRequestModel{},
		focus:  "sidebar",
	}
	m.viewport = viewport.New(78, 13)
	m.viewport.YPosition = 0
	m.viewport.MouseWheelEnabled = true
	return m
}

type SidebarModel struct {
	options  []string
	selected int
}

// View method for the menu
func (m SidebarModel) View() string {
	var b strings.Builder

	for i, option := range m.options {
		if i == m.selected {
			b.WriteString(fmt.Sprintf("> %s\n", lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Render(option))) // Highlight selected option in purple
		} else {
			b.WriteString(fmt.Sprintf("  %s\n", lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Render(option))) // Regular option in purple
		}
	}
	var menu = b.String()

	return lipgloss.NewStyle().
		Width(40).
		Border(lipgloss.RoundedBorder()).
		Padding(1).
		Render(menu)
}

// Update method for menu navigation
func (m SidebarModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m SidebarModel) Init() tea.Cmd {
	return nil // Return nil or any initial command if needed
}
