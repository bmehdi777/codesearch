package utils

import "os/exec"

func ExecCommand(command string, options []string) (string, error) {
	cmd := exec.Command(command, options...)

	stdout, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(stdout), nil
}
