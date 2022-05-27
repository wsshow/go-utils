package log

import "testing"

func TestWorkflow(t *testing.T) {
	Error("error...")
	Info("info...")
	Debug("debug...")
	Warn("warn...")

	Errorf("error...%s %d", "string", 123)
	Infof("info...")
	Debugf("debug...")
}
