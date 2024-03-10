package initializer

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
)

func GetSecret(ctx context.Context) (*Credential, error) {

	fmt.Println("GET CREDENTIAL")

	//client, err := secretmanager.NewClient(ctx)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to create Secret Manager client: %v", err)
	//}
	//defer client.Close()
	//
	//req := &secretmanagerpb.AccessSecretVersionRequest{
	//	Name: os.Getenv("PROJECT_ID"),
	//}
	//
	//result, err := client.AccessSecretVersion(ctx, req)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to access secret version: %v", err)
	//}
	//
	//var credential Credential
	//if err := json.Unmarshal(result.Payload.Data, &credential); err != nil {
	//	return nil, fmt.Errorf("failed to decode secret value: %v", err)
	//}

	file, err := os.Open("/etc/secrets/postgres.json")
	if err != nil {
		fmt.Println("Error opening config file: /etc/secrets/postgres.json", err)
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
