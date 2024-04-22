package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

func AzKeyVaults(AzCred *azidentity.DefaultAzureCredential) error {
	subscriptions := GetSubscriptions(AzCred)

	for _, sub := range subscriptions {
		fmt.Printf("getting subscriptions for keyvault enum... \n")
		subscriptionId := *sub.SubscriptionID
		clientFactory, err := armkeyvault.NewClientFactory(subscriptionId, AzCred, nil)
		if err != nil {
			log.Fatalf("failed to create client: %v \n", err)
		}
		pager := clientFactory.NewVaultsClient().NewListPager(nil)
		ctx := context.TODO()

		for pager.More() {
			page, err := pager.NextPage(ctx)
			if err != nil {
				log.Fatalf("error occurred getting keyvaults... %v \n", err)
			}

			for _, v := range page.Value {
				fmt.Printf("%v \n", *v.Name)
			}
		}

	}
	return nil
}
