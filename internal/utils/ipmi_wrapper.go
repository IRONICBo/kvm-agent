package utils

import (
	"errors"
	"os/exec"
)

// GetIPMISensorList returns the output of the ipmitool sensor list command.
func GetIPMISensorList(host, port, username, password string) (string, error) {
	cmd := "ipmitool"
	args := []string{"-I", "lanplus", "-H", host, "-p", port, "-U", username, "-P", password, "sensor", "list"}

	if !isCommandAvailable(cmd) {
		return "", errors.New("command '%s' is not available")
	}

	output, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

// GetIPMISelList returns the output of the ipmitool sel list command.
func GetIPMISelList(host, port, username, password string) (string, error) {
	cmd := "ipmitool"
	args := []string{"-I", "lanplus", "-H", host, "-p", port, "-U", username, "-P", password, "sel", "list"}

	if !isCommandAvailable(cmd) {
		return "", errors.New("command '%s' is not available")
	}

	output, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
