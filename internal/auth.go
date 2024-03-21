package internal

import (
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

func GetCredential() (*azidentity.DefaultAzureCredential, error) {
	client, err := azidentity.NewDefaultAzureCredential(nil)
	fmt.Println("getting token from Azure...")
	if err != nil {
		log.Fatalf("authentication to Azure failed with: %v", err)
	}
	return client, nil
}
