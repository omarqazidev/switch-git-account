package terminal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func AddSshKey(sshFileName string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return err
	}

	sshDir := filepath.Join(homeDir, ".ssh")
	sshKeyFullPath := filepath.Join(sshDir, sshFileName)

	cmd1 := exec.Command("ssh-add", sshKeyFullPath)

	err1 := cmd1.Run()
	if err1 != nil {
		return err1
	}

	fmt.Println("\nSSH key is added to your ssh-agent:", sshKeyFullPath)

	return nil
}
