package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

func AzBlobs(AzCred *azidentity.DefaultAzureCredential) error {
	// cred, _ := internal.GetCredential()

	subscriptions := GetSubscriptions(AzCred)
	for _, sub := range subscriptions {
		subscriptionId := *sub.SubscriptionID
		clientFactory, err := armstorage.NewClientFactory(subscriptionId, AzCred, nil)
		if err != nil {
			log.Fatalf("Error occurred in Blobs... %v", err)
		}
		pager := clientFactory.NewAccountsClient().NewListPager(nil)
		ctx := context.TODO()
		for pager.More() {
			page, err := pager.NextPage(ctx)
			if err != nil {
				return err
			}
			// Return all storage accounts for given subscription
			for _, v := range page.Value {
				fmt.Printf("%v \n", *v.Name)
				fmt.Printf("%v \n", *v.Properties.NetworkRuleSet.DefaultAction)
			}

		}

	}
	return nil
}

func getStorageAccounts() {
}
