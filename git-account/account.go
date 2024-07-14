package gitaccount

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type GitAccountDetails struct {
	Username    string
	Email       string
	SSHFileName string
}

func CreateAccount(username, email, sshFileName string) error {
	fmt.Println("CREATE GIT ACCOUNT")

	personalGitAccount := GitAccountDetails{
		Username:    username,
		Email:       email,
		SSHFileName: sshFileName,
	}

	// personalGitAccount := GitAccountDetails{
	// 	Username:    "Omar Qazi",
	// 	Email:       "omarqazidev@gmail.com",
	// 	sshFileName: "omarqazidev-git-ed25519",
	// }

	// jsonFile, err := os.Create(".git-account.json")
	// if err != nil {
	// 	fmt.Println("Error creating json file:", err)
	// 	return err
	// }
	// defer jsonFile.Close()

	// // Write configuration content to the file
	// configContent := fmt.Sprintf(`{
	// 	"username": "%s",
	// 	"email": "%s",
	// 	"sshFileName": "%s"
	// }`, personalGitAccount.Username, personalGitAccount.Email, personalGitAccount.sshFileName)

	// _, err = jsonFile.WriteString(configContent)
	// if err != nil {
	// 	fmt.Println("Error writing to json file:", err)
	// 	return err
	// }

	// return nil

	acc, _ := json.Marshal(personalGitAccount)

	fmt.Println(acc)
	fmt.Println(string(acc))

	return nil

}

func createSshConfig() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	sshDir := filepath.Join(homeDir, ".ssh")
	configFile := filepath.Join(sshDir, "configZZZZ")

	// Create .ssh directory if it doesn't exist
	if _, err := os.Stat(sshDir); os.IsNotExist(err) {
		err = os.Mkdir(sshDir, 0700)
		if err != nil {
			fmt.Println("Error creating .ssh directory:", err)
			return
		}
	}

	// Create config file
	file, err := os.Create(configFile)
	if err != nil {
		fmt.Println("Error creating config file:", err)
		return
	}
	defer file.Close()

	// Write configuration content to the file
	configContent := `# SSH Config File
Host example
    HostName example.com
    User your-username
    IdentityFile ~/.ssh/id_rsa
`
	_, err = file.WriteString(configContent)
	if err != nil {
		fmt.Println("Error writing to config file:", err)
		return
	}

	fmt.Println("SSH config file created successfully at", configFile)
}

func backupExistingSshConfig() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return err
	}

	sshDir := filepath.Join(homeDir, ".ssh")
	configFile := filepath.Join(sshDir, "config")

	// Check if the config file exists
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		fmt.Println("No existing config file to back up.")
		return err
	}

	// Generate a timestamp string
	timestamp := time.Now().Format("20060102_150405")
	backupFile := filepath.Join(sshDir, fmt.Sprintf("config_backup_%s", timestamp))

	// Create the backup file
	srcFile, err := os.Open(configFile)
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(backupFile)
	if err != nil {
		fmt.Println("Error creating backup file:", err)
		return err
	}
	defer destFile.Close()

	// Copy the contents of the config file to the backup file
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		fmt.Println("Error copying to backup file:", err)
		return err
	}

	fmt.Println("Backup created successfully at", backupFile)

	return nil
}
