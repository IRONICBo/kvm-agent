package utils

import (
	"errors"
	"os/exec"
)

// GetSNMPWalkList returns the output of the snmpwalk command.
func GetSNMPTranslateList(host, port, community, rootOID string) (string, error) {
	cmd := "snmptranslate"
	// args := []string{"-v", "2c", "-c", community, fmt.Sprintf("%s:%s", host, port), rootOID}
	args := []string{"-Tz"}

	if !IsCommandAvailable(cmd) {
		return "", errors.New("command '%s' is not available")
	}

	output, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
