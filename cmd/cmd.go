package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

const logTitle = "[cmd]"

func logLn(v ...interface{}) {
	log.Println(logTitle, v)
}

var curOS = runtime.GOOS

func Run(cmd string) string {
	var result []byte
	var err error
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

func Find(s string) string {
	var rt string
	if curOS == "linux" {
		rt = Run(fmt.Sprintf("find / -name \"%s\" 2>/dev/null", s))
	} else if curOS == "windows" {
		rt = Run(fmt.Sprintf("powershell ls / -Recurse -Include \"%s\"", s))
	}
	return rt
}
