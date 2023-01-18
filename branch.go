package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func ParseBranches() []string {
	// Info("Git Branch:")
	cmd := exec.Command("git", "branch")
	out, err := cmd.CombinedOutput()

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(out), "\n")
	branches := []string{}

	for _, line := range lines {
		line = strings.TrimPrefix(line, "*")
		line = strings.TrimSpace(line)

		if len(line) != 0 {
			branches = append(branches, line)
		}
	}

	return branches
}

func DeleteBranches(branches []string) {
	// Info("Delete Branches:")

	for _, name := range branches {
		cmd := exec.Command("git", "branch", "-d", name)
		out, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Print(string(out))
			return
		}

		fmt.Print(string(out))
	}
}
