package cmd

import (
	"testing"
)

func TestRun(t *testing.T) {
	logLn(Run("powershell ps -Name 'Sys*'"))
}

func TestFind(t *testing.T) {
	s := Find("agent-collector*")
	logLn(s)
}
