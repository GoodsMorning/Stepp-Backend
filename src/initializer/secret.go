package initializer

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

func GetSecret(ctx context.Context) (*Credential, error) {

	fmt.Println("GET CREDENTIAL")

	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create Secret Manager client: %v", err)
	}
	defer client.Close()

	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: os.Getenv("PROJECT_ID"),
	}

	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to access secret version: %v", err)
	}

	var credential Credential
	if err := json.Unmarshal(result.Payload.Data, &credential); err != nil {
		return nil, fmt.Errorf("failed to decode secret value: %v", err)
	}

	return &credential, nil
}
