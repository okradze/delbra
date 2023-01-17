package main

import (
	"fmt"
	"os/exec"
)

func main() {
	Info("Git Branch:")
	err := exec.Command("cd", "~/Desktop/go-learning/cyoa").Run()

	if err != nil {
		panic(err)
	}

	cmd := exec.Command("git", "branch")
	out, err := cmd.CombinedOutput()

	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}

func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}
