package models

import (
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type LoadingModel struct {
	progress    progress.Model
	percent     float64
	loadingStep int
	ready       bool
	mainModel   MainModel // Store the initialized main model
}

// Custom messages for loading steps
type loadMsg struct {
	step    int
	success bool
}

func InitialLoadingModel() LoadingModel {
	p := progress.New(
		progress.WithGradient("#2A2A2A", "#660066"),
		progress.WithWidth(40),
		progress.WithoutPercentage(),
	)

	// Pre-initialize the MainModel to avoid nil pointer
	mainModel := InitialModel()

	return LoadingModel{
		progress:    p,
		percent:     0.0,
		loadingStep: 0,
		ready:       false,
		mainModel:   mainModel, // Initialize this field
	}
}

func (m LoadingModel) Init() tea.Cmd {
	return m.initializeNextComponent()
}

func (m LoadingModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, nil

	case loadMsg:
		if msg.success {
			m.loadingStep++
			m.percent = float64(m.loadingStep) / 4 // Assuming 4 loading steps

			if m.loadingStep >= 4 {
				m.ready = true
				return m.mainModel, nil
			}

			return m, m.initializeNextComponent()
		}
		// Handle initialization failure
		return m, tea.Quit

	default:
		return m, nil
	}
}

func (m LoadingModel) View() string {
	var loadingText string
	switch m.loadingStep {
	case 0:
		loadingText = "Initializing hospital room..."
	case 1:
		loadingText = "Preparing diagnostic tools..."
	case 2:
		loadingText = "Initializing patient data..."
	case 3:
		loadingText = "Preparing medications..."
	}

	return lipgloss.NewStyle().
		Width(80).
		Height(7).
		Align(lipgloss.Center, lipgloss.Center).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				"Loading Illnesses...",
				loadingText,
				m.progress.ViewAs(m.percent),
			),
		)
}

func (m *LoadingModel) initializeNextComponent() tea.Cmd {
	return func() tea.Msg {
		// Simulate some work and actually initialize components
		switch m.loadingStep {
		case 0:
			// Initialize basic components
			time.Sleep(100 * time.Millisecond)
		case 1:
			// Load configuration
			time.Sleep(100 * time.Millisecond)
		case 2:
			// Set up UI components
			time.Sleep(100 * time.Millisecond)
		case 3:
			// Initialize the main model
			m.mainModel = InitialModel()
			time.Sleep(100 * time.Millisecond)
		}

		return loadMsg{step: m.loadingStep, success: true}
	}
}
