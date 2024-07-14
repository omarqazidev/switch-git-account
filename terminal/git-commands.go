package terminal

import (
	"fmt"
	"os/exec"
)

func SetGitDefaults(username, email string) error {
	// Set Username
	cmd1 := exec.Command("git", "config", "--global", "user.name", username)

	err1 := cmd1.Run()
	if err1 != nil {
		fmt.Println("Error executing command:", err1)
		return err1
	}

	// Set Email
	cmd2 := exec.Command("git", "config", "--global", "user.email", email)

	err2 := cmd2.Run()
	if err2 != nil {
		fmt.Println("Error executing command:", err2)
		return err2
	}

	fmt.Println("\ngit user.name is set to :", username)
	fmt.Println("git user.email is set to :", email)

	return nil
}
