package main

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"

	//    "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
)

const subscriptionID = "0be01d84-8432-4558-9aba-ecd204a3ee61"

func main() {
	fmt.Println("Fetching subscription", subscriptionID)
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		// TODO: handle
		fmt.Println("Auth error")
	} else {
		fmt.Println(&cred)
	}

	client, err := armsubscription.NewSubscriptionsClient(cred, nil)
	if err != nil {
		// TODO: Handle Error
		fmt.Println("New Subscription client error")
	}

	_, err = client.Get(context.TODO(), subscriptionID, nil)
	if err != nil {
		// TODO: handle error
		fmt.Println("Subscription error")
	}
}
