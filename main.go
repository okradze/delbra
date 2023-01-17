package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	Info("Git Branch:")
	cmd := exec.Command("git", "branch")
	out, err := cmd.CombinedOutput()

	if err != nil {
		panic(err)
	}

	branches := parseBranches(string(out))
	fmt.Println(branches)

	deleteBranches(branches)
}

func parseBranches(s string) []string {
	lines := strings.Split(s, "\n")
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

func deleteBranches(branches []string) {
	Info("Delete Branches:")

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

func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}
