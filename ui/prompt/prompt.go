package prompt

import (
	"github.com/c-bata/go-prompt"
)

func CliPrompt(prefix string, sugs []prompt.Suggest) string {
	completer := func(d prompt.Document) []prompt.Suggest {
		return prompt.FilterHasPrefix(sugs, d.GetWordBeforeCursor(), true)
	}
	t := prompt.Input(prefix, completer)
	return t
}
