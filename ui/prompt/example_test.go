package prompt

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"testing"
)

func ExampleCliPrompt() {
	s := []prompt.Suggest{
		{Text: "users", Description: "Store the username and age"},
		{Text: "articles", Description: "Store the article text posted by user"},
		{Text: "comments", Description: "Store the text commented to articles"},
	}
	out := CliPrompt("> ", s)
	fmt.Println(out)

	// Input: users
	// Output: users
}

func TestExampleCliPrompt(t *testing.T) {
	ExampleCliPrompt()
}
