package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	command := exec.Command(cmd[0], cmd[1:]...)

	for k, v := range env {
		command.Env = append(command.Env, k+"="+v.Value)
	}

	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()

	if err != nil {
		if exitCode, ok := err.(*exec.ExitError); ok {
			returnCode = exitCode.ExitCode()
		}
	}

	return
}
