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

	branchesDeleted := []string{}
	staleBranches := listStaleBranches()
	if len(staleBranches) == 0 {
		fmt.Println("No stale branches found.")
	} else {
		fmt.Println("Stale branches:")
		for _, branch := range staleBranches {
			fmt.Println(branch)
			if removeStaleBranch(branch) {
				branchesDeleted = append(branchesDeleted, branch)
			}
		}
	}

	localBranches := listLocaleBranchesWithoutRemoteTracking()
	fmt.Println("\nLocal untracked branches:")
	if len(localBranches) > 0 {
		fmt.Println("We found some local branches.")
		for _, branch := range localBranches {
			fmt.Println(branch)
		}
		fmt.Println("Do you want to delete these local branches? (y/n)")
		var response string
		fmt.Scanln(&response)
		if response == "y" || response == "Y" {
			for _, branch := range localBranches {
				if removeLocalBranch(branch) {
					branchesDeleted = append(branchesDeleted, branch)
				}
			}
		}
	}

	const colorGreen = "\033[32m"
	const colorReset = "\033[0m"

	if len(branchesDeleted) == 0 {
		fmt.Println("\nNo branches were deleted.")
		return
	}

	fmt.Printf("\n%sDeleted %d branches%s\n", colorGreen, len(branchesDeleted), colorReset)
	fmt.Printf("\n%sDeleted branches:%s\n", colorGreen, colorReset)
	for _, branch := range branchesDeleted {
		fmt.Printf("%s%s%s\n", colorGreen, branch, colorReset)
	}
}
