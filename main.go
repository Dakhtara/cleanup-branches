package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
)

var errorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Bold(true).MarginTop(1)
var successStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#05DF72")).MarginTop(1).Bold(true)
var infoStyle = lipgloss.NewStyle().Underline(true).MarginTop(1)
var warningStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFA500")).Bold(true).MarginTop(1)

func main() {
	var headStyle = lipgloss.NewStyle().
		Padding(2, 4).
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#5D5FEF")).
		Align(lipgloss.Center).
		Bold(true)

	fmt.Println(headStyle.Render(" Git Branch Cleaner \n This tool only deletes branches that have been merged to main or master,\n or local branches without remote tracking."))

	argsWithProg := os.Args
	if len(argsWithProg) < 2 || argsWithProg[1] != "-f" {
		fmt.Println(warningStyle.Render("Before starting, you should ensure that you have committed or stashed any changes,\n and that you have pushed any branches you want to keep to the remote repository.\n Are you sure you want to continue? (y/n)"))

		var proceed string
		fmt.Scanln(&proceed)
		if proceed != "y" && proceed != "Y" {
			fmt.Println(infoStyle.Render("Exiting Git Branch Cleaner. No changes were made."))
			return
		} else {
			fmt.Println(infoStyle.Render("Proceeding with Git Branch Cleaner...\n Next time you can run with -f flag to skip this prompt."))
		}
	}

	if !checkGitCommand() {
		fmt.Println(errorStyle.Render("Error: Git is not installed or not found in PATH."))
		return
	}
	fmt.Println(successStyle.Render("✅ Git found"))

	if !areWeInGitRepo() {
		fmt.Println(errorStyle.Render("Error: Not inside a Git repository."))
		return
	}
	fmt.Println(successStyle.Render("✅ Inside a Git repository"))

	branchesDeleted := []string{}
	// TODO: find a better way to understand if a branch is stale or not, now it uses --merged but it doesn't mean the PR was merged.
	// staleBranches := listStaleBranches()
	// if len(staleBranches) == 0 {
	// 	fmt.Println(infoStyle.Render("No stale branches found."))
	// } else {
	// 	fmt.Println(infoStyle.Render("Stale branches:"))
	// 	for _, branch := range staleBranches {
	// 		fmt.Println(branch)
	// 		if removeStaleBranch(branch) {
	// 			branchesDeleted = append(branchesDeleted, branch)
	// 		}
	// 	}
	// }

	localBranches := listLocaleBranchesWithoutRemoteTracking()
	if len(localBranches) > 0 {
		fmt.Println(infoStyle.Render("We found some local branches."))

		listBranch := list.New(localBranches)

		fmt.Println(listBranch)

		var style = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFA500")).Bold(true)

		fmt.Println(style.Render("Do you want to delete these local branches? (y/n)"))
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

	if len(branchesDeleted) == 0 {
		fmt.Println(successStyle.Render("No branches were deleted."))
		return
	}

	fmt.Println(successStyle.Render(fmt.Sprintf("Deleted %d branches", len(branchesDeleted))))
	listDeletedBranches := list.New(branchesDeleted)

	fmt.Println(listDeletedBranches)
}
