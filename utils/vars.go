package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	GraffTitle = `

  __                            .__                 .__      .__ .__   .__                                   
_/  |_   ____  _______   _____  |__|  ____  _____   |  |     |__||  |  |  |    ____    ____    ______  ______
\   __\_/ __ \ \_  __ \ /     \ |  | /    \ \__  \  |  |     |  ||  |  |  |   /    \ _/ __ \  /  ___/ /  ___/
 |  |  \  ___/  |  | \/|  Y Y  \|  ||   |  \ / __ \_|  |__   |  ||  |__|  |__|   |  \\  ___/  \___ \  \___ \ 
 |__|   \___  > |__|   |__|_|  /|__||___|  /(____  /|____/   |__||____/|____/|___|  / \___  >/____  >/____  >
			\/               \/          \/      \/                               \/      \/      \/      \/ 
`

	SendButton = `\\\\\\\\\\\\\\\\\\
\\		      \\
\\     Send     \\
\\    Request   \\
\\			  \\
\\\\\\\\\\\\\\\\\\`

	P = `        ___
	  /	 \
    / (-) (-) \
   |    ( )    |
    __ (---) __
	    \-/	  `
	Pp = `         _
	    / \
	/(O>) (<O)\
   |   (o o)   |
	__ (---) __
		\-/	  `

	AuthorStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("235"))
	FocusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	BlurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	CursorStyle         = FocusedStyle
	NoStyle             = lipgloss.NewStyle()
	HelpStyle           = BlurredStyle
	CursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	FocusedButton = FocusedStyle.Render("[ Run Test ]")
	BlurredButton = fmt.Sprintf("[ %s ]", BlurredStyle.Render("Run Test"))
)

// Function to clear the terminal screen
func ClearTerminal() {
	cmd := exec.Command("clear") // For Unix/Linux/Mac
	if strings.Contains(strings.ToLower(os.Getenv("OS")), "windows") {
		cmd = exec.Command("cmd", "/c", "cls") // For Windows
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// "
// 	        ______
//         .-"      "-.
//        /            \
//       |              |
//       |,  .-.  .-.  ,|
//       | )(__/  \__)( |
//       |/     /\     \|
//       (_     ^^     _)
//        \__|IIIIII|__/
//         | \\IIIIII/ |
//         \\          /
//          `--------`
//   "
