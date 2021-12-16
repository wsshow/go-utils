package cmd

import (
	"testing"
)

func TestRun(t *testing.T) {
	logLn(Run("powershell ps -Name 'Sys*'"))
}

func TestFind(t *testing.T) {
	ss := Find("agent-collector*")
	for i, s := range ss {
		logLn(i, s)
	}
}
