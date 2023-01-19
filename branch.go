package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func parseGitBranchCommand(cmd *exec.Cmd, merged bool) []Branch {
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Print(string(out))
		os.Exit(1)
	}

	lines := strings.Split(string(out), "\n")
	branches := []Branch{}

	for _, line := range lines {
		// Exclude current working branch
		if strings.ContainsRune(line, '*') || len(line) == 0 {
			continue
		}

		line = strings.TrimSpace(line)
		branches = append(branches, Branch{name: line, merged: merged})
	}

	return branches
}

func ParseBranches() []Branch {
	merged := parseGitBranchCommand(exec.Command("git", "branch", "--merged"), true)
	notMerged := parseGitBranchCommand(exec.Command("git", "branch", "--no-merged"), false)
	return append(merged, notMerged...)
}

func DeleteBranches(branches []string) {
	for _, name := range branches {
		cmd := exec.Command("git", "branch", "-D", name)
		out, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Println(string(out))
			return
		}

		fmt.Print(string(out))
	}
}
