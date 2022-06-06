package term

import (
	"fmt"
	"golang.org/x/term"
	"os"
)

type MockTerminal struct {
}

func (c *MockTerminal) Read(data []byte) (n int, err error) {
	return os.Stdin.Read(data)
}

func (c *MockTerminal) Write(data []byte) (n int, err error) {
	os.Stdout.Write(data)
	return len(data), nil
}

func WorkFlow() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)
	c := &MockTerminal{}
	t := term.NewTerminal(c, "> ")
	for {
		line, err := t.ReadLine()
		if err != nil {
			return
		}
		if line == "exit" {
			break
		}
		fmt.Println(line)
	}
	line, err := t.ReadPassword("Enter your password: ")
	if err != nil {
		return
	}
	fmt.Println(line)
}
