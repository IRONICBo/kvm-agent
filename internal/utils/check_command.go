package utils

import "os/exec"

// IsCommandAvailable Check if command is available.
func IsCommandAvailable(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
