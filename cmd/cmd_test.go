package cmd

import "testing"

func TestRun(t *testing.T) {
	logLn(Run("powershell ps -Name 'Sys*'"))
}
