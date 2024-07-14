package file

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func createFile(filePath, filename, contents string) error {
	fileNameWithPath := filepath.Join(filePath, filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err = os.Mkdir(filePath, 0700)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return err
		}
	}

	file, err := os.Create(fileNameWithPath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString(contents)
	if err != nil {
		fmt.Println("Error writing to config file:", err)
		return err
	}

	fmt.Printf("%s file created successfully at %s\n", filename, fileNameWithPath)

	return nil
}

func backupExistingFile(filePath, filename string) error {
	fileNameWithPath := filepath.Join(filePath, filename)

	if _, err := os.Stat(fileNameWithPath); os.IsNotExist(err) {
		fmt.Println("No existing file to back up.")
		return err
	}

	timestamp := time.Now().Format("20060102_150405")
	backupFile := filepath.Join(filePath, fmt.Sprintf("%s_backup_%s", filename, timestamp))

	srcFile, err := os.Open(fileNameWithPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(backupFile)
	if err != nil {
		fmt.Println("Error creating backup file:", err)
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		fmt.Println("Error copying to backup file:", err)
		return err
	}

	fmt.Printf("%s backup created successfully at %s\n", filename, backupFile)

	return nil
}
