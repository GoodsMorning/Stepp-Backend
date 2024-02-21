package initializer

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadCredential() (*Credential, error) {
	file, err := os.Open("./credential/postgres.json")
	if err != nil {
		fmt.Println("Error opening config file: ./credential/postgres.json", err)
		return nil, err
	}
	defer file.Close()

	var credential Credential
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&credential); err != nil {
		fmt.Println("Error decoding config file:", err)
		return nil, err
	}

	return &credential, nil
}

type Credential struct {
	DbName      string `json:"db_name"`
	DbIpAddress string `json:"db_ip_address"`
	DbPort      string `json:"db_port"`
	DbUsername  string `json:"db_username"`
	DbPassword  string `json:"db_password"`
}
