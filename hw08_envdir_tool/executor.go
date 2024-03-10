package main

import (
	"errors"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	if len(cmd) == 0 {
		returnCode = 1
		return
	}

	inputCommand := cmd[0]
	command := exec.Command(inputCommand, cmd[1:]...)

	for k, v := range env {
		command.Env = append(command.Env, k+"="+v.Value)
	}

	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		var exitErr *exec.ExitError

		if errors.As(err, &exitErr) {
			returnCode = exitErr.ExitCode()
		}
	}

	return
}
