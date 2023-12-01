package utils

import "os/exec"

// isCommandAvailable Check if command is available.
func isCommandAvailable(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
