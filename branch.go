package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ParseBranches() []string {
	cmd := exec.Command("git", "branch")
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Print(string(out))
		os.Exit(1)
	}

	lines := strings.Split(string(out), "\n")
	branches := []string{}

	for _, line := range lines {
		// Exclude current working branch
		if strings.ContainsRune(line, '*') || len(line) == 0 {
			continue
		}

		line = strings.TrimSpace(line)
		branches = append(branches, line)
	}

	return branches
}

func DeleteBranches(branches []string) {
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
