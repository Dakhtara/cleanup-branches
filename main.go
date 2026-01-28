package main

import (
	"fmt"
)

func main() {

	if !checkGitCommand() {
		return
	}
	fmt.Println("Git found")

	if !areWeInGitRepo() {
		return
	}
	fmt.Println("Inside a Git repository")

	staleBranches := listStaleBranches()
	if len(staleBranches) == 0 {
		fmt.Println("No stale branches found.")
	} else {
		fmt.Println("Stale branches:")
		for _, branch := range staleBranches {
			fmt.Println(branch)
		}
	}

	localBranches := listLocalBranches()
	fmt.Println("\nLocal branches:")
	for _, branch := range localBranches {
		fmt.Println(branch)
	}
}
