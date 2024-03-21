package main

import (
	"context"
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

	// client, _ := GetAuthz()
	// fmt.Println(client)
	fmt.Println(GetTenantClients())
}

func GetAuthz() (*azidentity.DefaultAzureCredential, error) {
	client, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("authentication to Azure failed with: %v", err)
	}
	return client, nil
}

func GetTenantClients() []armsubscription.TenantIDDescription {
	cred, _ := GetAuthz()
	// var result armsubscription.TenantIDDescription
	ctx := context.Background()
	clientFactory, err := armsubscription.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	pager := clientFactory.NewTenantsClient().NewListPager(nil)
	var result []armsubscription.TenantIDDescription
	for pager.More() {
		page, err := pager.NextPage(ctx)
		fmt.Printf("%T\n", page)
		fmt.Printf("%T\n", page.TenantListResult)

		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			fmt.Printf("%T \n", *v.TenantID)
			result = append(result, *v)
		}
	}
	return result
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
