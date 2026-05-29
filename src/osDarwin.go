package src

import (
	"os/exec"
	"strings"
)

func GetOSNameDarwin() string { // We'll run "sw_vers"

	cmd := exec.Command("sw_vers", "-productVersion") // fetching

	out, err := cmd.Output()
	if err != nil {
		return "unknow MacOS version"
	}

	version := strings.TrimSpace(string(out))

	return "MacOS" + version // final string
}

func GetUserNameDarwin() string {
	cmd := exec.Command("whoami")

	out, err := cmd.Output()
	if err != nil {
		return "Unknow username"
	}

	name := strings.TrimSpace(string(out))

	return name
}
