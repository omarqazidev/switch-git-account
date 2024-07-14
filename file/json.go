package file

import (
	"encoding/json"
	"fmt"
	"os"

	gitaccount "github.com/omarqazidev/switch-git-account/git-account"
)

func WriteToJsonFile(fileName string, data []gitaccount.GitAccountDetails) error {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating json file:", err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("Error encoding json data:", err)
		return err
	}

	return nil
}

func ReadJsonFile(fileName string, data *[]gitaccount.GitAccountDetails) error {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening json file:", err)
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(data)
	if err != nil {
		fmt.Println("Error decoding json data:", err)
		return err
	}

	return nil
}
