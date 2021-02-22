package main

import (
	"os/exec"
	"strings"
)

type cmd struct {
}

func (c *cmd) RunCmd(command string) (res string, err error) {
	out, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		return "", err
	}
	return strings.Trim(string(out), "\n"), err
}
