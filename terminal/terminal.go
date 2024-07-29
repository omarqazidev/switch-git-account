package terminal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/omarqazidev/switch-git-account/file"
	gitaccount "github.com/omarqazidev/switch-git-account/git-account"
)

func Menu() {

	for {
		fmt.Print("\n\n____________ SWITCH GIT ACCOUNT ____________\n\n")

		fmt.Println("\nWhat do you want to do?")
		fmt.Println("\n\t1. Change Git Account")
		fmt.Println("\t2. List Git Accounts")
		fmt.Println("\t3. Add Git Account")
		fmt.Println("\t4. Remove Git Account")
		fmt.Print("\n\t0. Exit\n\n")

		var choice string
		getInput("Enter your choice: ", &choice)

		if choice == "1" {
			ChooseAccount()
			pause("\nPress Enter to continue...")
		}

		if choice == "2" {
			ListAccounts()
			pause("\nPress Enter to continue...")
		}

		if choice == "3" {
			AddAccount()
			pause("\nPress Enter to continue...")
		}

		if choice == "4" {
			RemoveAccount()
			pause("\nPress Enter to continue...")
		}

		if choice == "0" {
			break
		}
	}
}

func AddAccount() {
	fmt.Print("\n\n______________ ADD GIT ACCOUNT ______________\n\n")

	gitAccounts := []gitaccount.GitAccountDetails{}

	newGitAccount := gitaccount.GitAccountDetails{}

	fmt.Print("Enter your details below:\n\n")
	getInput("\tEnter your Username:\t\t", &newGitAccount.Username)
	getInput("\tEnter your Email:\t\t", &newGitAccount.Email)
	getInput("\tEnter your Git SSH file name:\t", &newGitAccount.SSHFileName)

	err := file.ReadJsonFile("git-accounts.json", &gitAccounts)

	if err != nil {
		gitAccounts = []gitaccount.GitAccountDetails{}
	}

	gitAccounts = append(gitAccounts, newGitAccount)

	err = file.WriteToJsonFile("git-accounts.json", gitAccounts)
	if err != nil {
		fmt.Println("Error writing to json file:", err)
	}

	fmt.Print("\n\n_____________________________________________\n\n")

	fmt.Println("\nAccount added successfully.")
}

func RemoveAccount() {
	fmt.Print("\n\n____________ REMOVE GIT ACCOUNT ____________\n\n")

	gitAccounts := ListAccounts()

	if len(gitAccounts) == 0 {
		fmt.Println("No accounts found.")
		return
	}

	for {
		var indexString string

		getInput("\nEnter the index of the account you want to remove: (-1 to exit): ", &indexString)

		index, err := strconv.Atoi(indexString)

		if err != nil {
			fmt.Println("Error parsing index:", err)
			continue
		}

		if index == -1 {
			break
		}

		if index > 0 && index <= len(gitAccounts) {
			gitAccounts = append(gitAccounts[:index-1], gitAccounts[index:]...)

			err := file.WriteToJsonFile("git-accounts.json", gitAccounts)
			if err != nil {
				fmt.Println("Error writing to json file:", err)
			}

			fmt.Println("\nRemoved account successfully.")

			break
		}

		fmt.Println("\nInvalid index. Please try again.")
	}

}

func ChooseAccount() error {
	fmt.Print("\n\n____________ SWITCH GIT ACCOUNT ____________\n\n")

	gitAccounts := ListAccounts()

	for {
		var indexString string

		getInput("\nEnter the index of the account you want to switch to: (-1 to exit): ", &indexString)

		index, err := strconv.Atoi(indexString)

		if err != nil {
			fmt.Println("Error parsing index:", err)
			continue
		}

		if index == -1 {
			break
		}

		if index > 0 && index <= len(gitAccounts) {

			account := gitAccounts[index-1]

			err := SetGitDefaults(account.Username, account.Email)

			if err != nil {
				fmt.Println("Error setting git defaults:", err)
				return err
			}

			err = AddSshKey(account.SSHFileName)

			if err != nil {
				fmt.Println("Error adding ssh key to ssh-agent:", err)
				return err
			}

			fmt.Println("")
			err = file.BackupExistingSshConfig()

			if err != nil {
				fmt.Println("Error backing up ssh config file:", err)
				return err
			}

			err = file.CreateSshConfig(account.SSHFileName)

			if err != nil {
				fmt.Println("Error creating ssh config file:", err)
				return err
			}

			fmt.Println("\nSwitched account successfully.")

			break
		}

		fmt.Println("Invalid index. Please try again.")
	}

	return nil
}

func ListAccounts() []gitaccount.GitAccountDetails {
	// fmt.Println("LIST GIT ACCOUNTS")

	gitAccounts := FetchAccounts()

	DisplayAccounts(gitAccounts)

	return gitAccounts
}

func FetchAccounts() []gitaccount.GitAccountDetails {
	gitAccounts := []gitaccount.GitAccountDetails{}

	err := file.ReadJsonFile("git-accounts.json", &gitAccounts)

	if err != nil {
		gitAccounts = []gitaccount.GitAccountDetails{}
	}

	return gitAccounts
}

func DisplayAccounts(gitAccounts []gitaccount.GitAccountDetails) {
	for i := 0; i < len(gitAccounts); i++ {
		fmt.Println("_____________________________________________________")
		fmt.Println("\n\t\tACCOUNT ", i+1)
		fmt.Println("\n\tUsername:\t", gitAccounts[i].Username)
		fmt.Println("\tEmail:\t\t", gitAccounts[i].Email)
		fmt.Println("\tSSHFilename:\t", gitAccounts[i].SSHFileName)
		fmt.Println("\n_____________________________________________________")
	}
}

func DisplayAccount(gitAccount gitaccount.GitAccountDetails) {
	fmt.Println("_____________________________________________________")
	fmt.Println("\n\t\tACCOUNT")
	fmt.Println("\n\tUsername:\t", gitAccount.Username)
	fmt.Println("\tEmail:\t\t", gitAccount.Email)
	fmt.Println("\tSSHFilename:\t", gitAccount.SSHFileName)
	fmt.Println("\n_____________________________________________________")
}

func getInput(prnt string, inputVarRef *string) error {
	fmt.Print(prnt)
	in := bufio.NewReader(os.Stdin)
	inputString, err := in.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return err
	}
	inputString = strings.Replace(inputString, "\n", "", -1)
	inputString = strings.Replace(inputString, "\r", "", -1)

	*inputVarRef = inputString

	return nil
}

func pause(prnt string) error {
	fmt.Print(prnt)
	in := bufio.NewReader(os.Stdin)
	_, err := in.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return err
	}

	return nil
}
