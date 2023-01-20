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
		fmt.Print(errorFg(string(out)))
		os.Exit(0)
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

func GetBranches() []Branch {
	merged := parseGitBranchCommand(exec.Command("git", "branch", "--merged"), true)
	notMerged := parseGitBranchCommand(exec.Command("git", "branch", "--no-merged"), false)
	branches := append(merged, notMerged...)

	if len(branches) == 0 {
		fmt.Println(errorFg("No branches to delete"))
		os.Exit(0)
	}

	return branches
}

func DeleteBranches(branches []string) {
	for _, name := range branches {
		cmd := exec.Command("git", "branch", "-D", name)
		out, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Println(string(out))
			return
		}
	}
}
