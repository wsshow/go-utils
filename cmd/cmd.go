package cmd

import (
	"log"
	"os/exec"
	"runtime"
	"strings"
)

const logTitle = "[cmd]"

func logLn(v ...interface{}) {
	log.Println(logTitle, v)
}

func Run(cmd string) string {
	var result []byte
	var err error
	curOS := runtime.GOOS
	if curOS == "linux" {
		result, err = exec.Command("/bin/sh", "-c", cmd).Output()
	} else if curOS == "windows" {
		result, err = exec.Command("cmd", "/c", cmd).Output()
	}
	if err != nil {
		logLn("cmd:", cmd, "error:", err)
	}
	return strings.TrimSpace(string(result))
}
