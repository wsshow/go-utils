package lipgloss

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
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
