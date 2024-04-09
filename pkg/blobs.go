package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

var (
	publicNetworkAccess = "nil"
	encrypted           = false
)

func AzBlobs(AzCred *azidentity.DefaultAzureCredential) error {
	// cred, _ := internal.GetCredential()

	// Iterate through all subscription to get all Storage Accounts
	subscriptions := GetSubscriptions(AzCred)
	for index, sub := range subscriptions {
		fmt.Printf("currently in subscription: %d of %d \n", index+1, len(subscriptions))
		subscriptionId := *sub.SubscriptionID
		clientFactory, err := armstorage.NewClientFactory(subscriptionId, AzCred, nil)
		if err != nil {
			log.Fatalf("failed to create client: %v \n", err)
		}
		pager := clientFactory.NewAccountsClient().NewListPager(nil)
		ctx := context.TODO()
		for pager.More() {
			page, err := pager.NextPage(ctx)
			if err != nil {
				log.Printf("Error occurred getting blobs for subscription... %v", err)

				// Go to the next subscription if an error or no storage accounts are returned
				break
			}
			// Return all storage accounts for given subscription
			for _, v := range page.Value {
				if v.Properties.PublicNetworkAccess != nil {
					fmt.Printf("%v | %v | %v \n", *v.Name, *v.Properties.NetworkRuleSet.DefaultAction, *v.Properties.PublicNetworkAccess)
				} else {
					fmt.Printf("%v | %v | %v \n", *v.Name, *v.Properties.NetworkRuleSet.DefaultAction, publicNetworkAccess)
				}
			}

		}

	}
	return nil
}

func getStorageAccounts() {
}
