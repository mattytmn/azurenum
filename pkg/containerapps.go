package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
)

func AzContainerApps(AzCred *azidentity.DefaultAzureCredential, AzTenantId, AzSubId string) error {

	if AzTenantId == "" && AzSubId != "" {
		isValid, subId := getSubscriptionFromId(AzCred, AzSubId)

		if isValid {
			fmt.Println("ID is valid, enumerating container apps...")
			sub := []*armsubscriptions.Subscription{subId}
			getContainerAppsForSub(AzCred, sub)
		}
	}
	if AzTenantId == "" && AzSubId == "" {
		subs := GetSubscriptions(AzCred)
		getContainerAppsForSub(AzCred, subs)
	}
	return nil
}

func getContainerAppsForSub(azCred *azidentity.DefaultAzureCredential, subs []*armsubscriptions.Subscription) {
	for _, sub := range subs {
		subId := *sub.SubscriptionID
		clientFactory, err := armappcontainers.NewClientFactory(subId, azCred, nil)
		if err != nil {
			log.Fatalf("failed to create client: %v \n", err)
		}
		pager := clientFactory.NewContainerAppsClient().NewListBySubscriptionPager(nil)
		ctx := context.TODO()

		for pager.More() {
			page, err := pager.NextPage(ctx)
			if err != nil {
				defer log.Printf("error occurred getting container apps... %v \n", err)
				break
			}
			for _, v := range page.Value {
				fmt.Printf("%v \n", *v.ID)
				// Check that Ingress settings exist
				if v.Properties.Configuration.Ingress != nil {
					fmt.Printf("%v \n", *v.Properties.Configuration.Ingress.Fqdn)
				}
			}
		}
	}
}
