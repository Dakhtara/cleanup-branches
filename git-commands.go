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

func listLocaleBranchesWithoutRemoteTracking() []string {
	output, err := exec.Command("git", "branch", "-vv").Output()
	if err != nil {
		fmt.Println("Failed to list local branches with verbose info.")
		return nil
	}

	var branchesWithoutRemoteTracking []string
	lines := string(output)
	for _, line := range strings.Split(lines, "\n") {
		branch := strings.TrimSpace(line)
		if branch != "" && ((strings.Contains(branch, "gone]") || (strings.Contains(branch, "disparue]"))) || !strings.Contains(branch, "[origin/")) && !strings.Contains(branch, "main") && !strings.Contains(branch, "master") {
			branchName := strings.Split(branch, " ")[0]
			branchName = strings.TrimPrefix(branchName, "* ")
			branchesWithoutRemoteTracking = append(branchesWithoutRemoteTracking, branchName)
		}
	}

	return branchesWithoutRemoteTracking
}

// Stales branches contains the origin/ prefix, we need to remove it before deletion
func removeStaleBranch(branch string) bool {
	branch = strings.TrimPrefix(branch, "origin/")
	if !checkIfBranchExists(branch) {
		// fmt.Printf("Branch %s does not exist locally, skipping deletion.\n", branch)
		return false
	}
	_, err := exec.Command("git", "branch", "-D", branch).Output()
	if err != nil {
		fmt.Printf("Failed to delete branch %s: %v\n", branch, err)
		return false
	} else {
		fmt.Printf("Deleted branch %s\n", branch)
		return true
	}
}

func checkIfBranchExists(branch string) bool {
	output, err := exec.Command("git", "branch", "--list", branch).Output()
	if err != nil {
		fmt.Printf("Failed to check if branch %s exists: %v\n", branch, err)
		return false
	}
	return strings.TrimSpace(string(output)) != ""
}

func removeLocalBranch(branch string) bool {
	if !checkIfBranchExists(branch) {
		fmt.Printf("Branch %s does not exist locally, skipping deletion.\n", branch)
		return false
	}
	_, err := exec.Command("git", "branch", "-D", branch).Output()
	if err != nil {
		fmt.Printf("Failed to delete local branch %s: %v\n", branch, err)
		return false
	} else {
		fmt.Printf("Deleted local branch %s\n", branch)
		return true
	}
}
