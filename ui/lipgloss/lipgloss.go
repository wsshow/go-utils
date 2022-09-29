package lipgloss

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

func WorkFlow() {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingTop(2).
		PaddingLeft(4).
		Width(22)

	fmt.Println(style.Render("Hello, WorkFlow."))

	// 2 cells on all sides
	lipgloss.NewStyle().Padding(2)

	// 2 cells on the top and bottom, 4 cells on the left and right
	lipgloss.NewStyle().Margin(2, 4)

	// 1 cell on the top, 4 cells on the sides, 2 cells on the bottom
	lipgloss.NewStyle().Padding(1, 4, 2)

	// Clockwise, starting from the top: 2 cells on the top, 4 on the right, 3 on
	// the bottom, and 1 on the left
	lipgloss.NewStyle().Margin(2, 4, 3, 1)

	var style2 = lipgloss.NewStyle().
		Width(24).
		Align(lipgloss.Left).  // align it left
		Align(lipgloss.Right). // no wait, align it right
		Align(lipgloss.Center) // just kidding, align it in the center

	fmt.Println(style2.Render("test center"))

	// Add a purple, rectangular border
	var style3 = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63"))

	// Set a rounded, yellow-on-purple border to the top and left
	var anotherStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("228")).
		BorderBackground(lipgloss.Color("63")).
		BorderTop(true).
		BorderLeft(true)

	// Make your own border
	var myCuteBorder = lipgloss.Border{
		Top:         "._.:*:",
		Bottom:      "._.:*:",
		Left:        "|*",
		Right:       "|*",
		TopLeft:     "*",
		TopRight:    "*",
		BottomLeft:  "*",
		BottomRight: "*",
	}

	fmt.Println(style3.Render("style3"))
	fmt.Println(anotherStyle.Render("anotherStyle"))
	fmt.Println(lipgloss.NewStyle().Border(myCuteBorder).Render("myCuteBorder"))
}

const (
	// In real life situations we'd adjust the document to fit the width we've
	// detected. In the case of this example we're hardcoding the width, and
	// later using the detected width only to truncate in order to avoid jaggy
	// wrapping.
	width = 96

	columnWidth = 30
)

// Dialog.
var (
	subtle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}

	dialogBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(1, 0).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)

	buttonStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFF7DB")).
			Background(lipgloss.Color("#888B7E")).
			Padding(0, 3).
			MarginTop(1)

	activeButtonStyle = buttonStyle.Copy().
				Foreground(lipgloss.Color("#FFF7DB")).
				Background(lipgloss.Color("#F25D94")).
				MarginRight(2).
				Underline(true)
)

func Dialog() {
	doc := strings.Builder{}
	okButton := activeButtonStyle.Render("Yes")
	cancelButton := buttonStyle.Render("Maybe")

	question := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render("Are you sure you want to eat marmalade?")
	buttons := lipgloss.JoinHorizontal(lipgloss.Top, okButton, cancelButton)
	ui := lipgloss.JoinVertical(lipgloss.Center, question, buttons)

	dialog := lipgloss.Place(width, 9,
		lipgloss.Center, lipgloss.Center,
		dialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceChars("猫咪"),
		lipgloss.WithWhitespaceForeground(subtle),
	)

	doc.WriteString(dialog + "\n\n")
}
