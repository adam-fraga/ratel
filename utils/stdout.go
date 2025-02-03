package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func RunCommand(command string, showOutput bool, args ...string) error {
	// Split the first argument into individual arguments
	splitArgs := strings.Fields(args[0])
	// Combine the split arguments with the rest of the arguments
	allArgs := append(splitArgs, args[1:]...)

	fmt.Println(allArgs)
	cmd := exec.Command(command, allArgs...)
	// Get the output
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Failed getting the output of the command: %s", err.Error())
	}
	// Print the output
	if showOutput {
		PrintInfoMsg(string(output))
	}

	return nil
}
