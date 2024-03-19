package main

import (
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"

	//    "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	// "context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
)

const subscriptionID = "0be01d84-8432-4558-9aba-ecd204a3ee61"

func main() {
	fmt.Println("Fetching Azure Resources")
	// cred, err := azidentity.NewDefaultAzureCredential(nil)
	// if err != nil {
	// 	// TODO: handle
	// 	fmt.Println("Auth error")
	// } else {
	// 	fmt.Println(&cred)
	// }

	client, _ := GetAuthz()
	fmt.Println(client)

}

func GetAuthz() (*azidentity.DefaultAzureCredential, error) {

	client, err := azidentity.NewDefaultAzureCredential(nil)

	if err != nil {
		return nil, fmt.Errorf("Authentication to Azure failed with: %s", err)
	}
	return client, nil
}

func GetTenantClient() armsubscription.TenantsClient {
	cred, _ := GetAuthz()
	// var result armsubscription.TenantIDDescription
	client, err := armsubscription.NewTenantsClient(cred, nil)
	if err != nil {
		log.Fatalf("Failed to fetch tenant information: %s", err)
	}
	return *client
}

func GetTenants() armsubscription.TenantListResult {
	var result armsubscription.TenantListResult

	return result
}

func GetTenantSubscriptions() {

}

// client, err := armsubscription.NewSubscriptionsClient(cred, nil)
// if err != nil {
// 	// TODO: Handle Error
// 	fmt.Println("New Subscription client error")
// }

// _, err = client.Get(context.TODO(), subscriptionID, nil)
// if err != nil {
// 	// TODO: handle error
// 	fmt.Println("Subscription error")
// }
