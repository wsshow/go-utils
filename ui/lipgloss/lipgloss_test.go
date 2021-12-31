package lipgloss

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"strings"
	"testing"
)

func TestWorkFlow(t *testing.T) {
	WorkFlow()
}

func TestSetColorProfile(t *testing.T) {
	t.Parallel()

	tt := []struct {
		profile  termenv.Profile
		input    string
		style    lipgloss.Style
		expected string
	}{
		{
			termenv.Ascii,
			"hello",
			lipgloss.NewStyle().Foreground(lipgloss.Color("#5A56E0")),
			"hello",
		},
		{
			termenv.ANSI,
			"hello",
			lipgloss.NewStyle().Foreground(lipgloss.Color("#5A56E0")),
			"\x1b[94mhello\x1b[0m",
		},
		{
			termenv.ANSI256,
			"hello",
			lipgloss.NewStyle().Foreground(lipgloss.Color("#5A56E0")),
			"\x1b[38;5;62mhello\x1b[0m",
		},
		{
			termenv.TrueColor,
			"hello",
			lipgloss.NewStyle().Foreground(lipgloss.Color("#5A56E0")),
			"\x1b[38;2;89;86;224mhello\x1b[0m",
		},
	}

	for i, tc := range tt {
		lipgloss.SetColorProfile(tc.profile)
		res := tc.style.Render(tc.input)
		if res != tc.expected {
			t.Errorf("Test %d, expected:\n\n`%s`\n`%s`\n\nActual output:\n\n`%s`\n`%s`\n\n",
				i, tc.expected, formatEscapes(tc.expected),
				res, formatEscapes(res))
		}
	}
}

func formatEscapes(str string) string {
	return strings.ReplaceAll(str, "\x1b", "\\x1b")
}
