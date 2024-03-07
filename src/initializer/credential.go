package initializer

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func LoadCredential(ctx context.Context) (*Credential, error) {

	fmt.Println("try use env RELEASE = ", os.Getenv("RELEASE"))

	env := os.Environ()
	for i, v := range env {
		fmt.Println("envi : " + strconv.Itoa(i) + " " + v)
	}

	if os.Getenv("RELEASE") == "true" {
		credential, err := GetSecret(ctx)
		if err != nil {
			return nil, err
		}
		return credential, nil
	}

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
