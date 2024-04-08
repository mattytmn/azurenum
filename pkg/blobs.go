package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/mattytmn/azurenum/internal"
)

func AzBlobs() error {
	cred, _ := internal.GetCredential()

	subscriptions := GetSubscriptions("")
	for _, sub := range subscriptions {
		subscriptionId := *sub.SubscriptionID
		clientFactory, err := armstorage.NewClientFactory(subscriptionId, cred, nil)
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
			for _, v := range page.Value {
				fmt.Printf("%v \n", *v.Name)
			}

		}

	}
	return nil
}

func getStorageAccounts() {
}
