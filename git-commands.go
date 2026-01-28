package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func checkGitCommand() bool {
	_, err := exec.Command("git", "version").Output()

	if err != nil {
		fmt.Println("Git is not installed or not found in PATH.")
		return false
	}

	return true
}

func areWeInGitRepo() bool {
	_, err := exec.Command("git", "rev-parse", "--is-inside-work-tree").Output()

	if err != nil {
		fmt.Println("Not inside a Git repository.")
		return false
	}

	return true
}

func listStaleBranches() []string {
	_, err := exec.Command("git", "fetch", "--prune").Output()
	if err != nil {
		fmt.Println("Failed to fetch and prune branches.")
		return nil
	}

	output, err := exec.Command("git", "branch", "-r", "--merged").Output()
	if err != nil {
		fmt.Println("Failed to list merged branches.")
		return nil
	}

	var staleBranches []string
	lines := string(output)
	for _, line := range strings.Split(lines, "\n") {
		branch := strings.TrimSpace(line)
		if branch != "" && !strings.Contains(branch, "main") && !strings.Contains(branch, "master") {
			staleBranches = append(staleBranches, branch)
		}
	}

	return staleBranches
}

func listLocalBranches() []string {
	output, err := exec.Command("git", "branch").Output()
	if err != nil {
		fmt.Println("Failed to list local branches.")
		return nil
	}

	var localBranches []string
	lines := string(output)
	for _, line := range strings.Split(lines, "\n") {
		branch := strings.TrimSpace(line)
		branch = strings.TrimPrefix(branch, "* ")
		if branch != "" {
			localBranches = append(localBranches, branch)
		}
	}

	return localBranches
}
