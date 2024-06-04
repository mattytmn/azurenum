package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

func GetResourceGroups(AzCred *azidentity.DefaultAzureCredential, AzSubscriptionId string) (resourceGroupsInSub []*armresources.ResourceGroup) {

	clientFactory, err := armresources.NewClientFactory(AzSubscriptionId, AzCred, nil)
	if err != nil {
		log.Fatalf("Failed to create resource group client: %v", err)
	}

	ctx := context.TODO()
	pager := clientFactory.NewResourceGroupsClient().NewListPager(nil)

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("Failed to advance page: %v", err)
		}
		resourceGroupsInSub = append(resourceGroupsInSub, page.Value...)
		for _, v := range page.Value {
			fmt.Printf("%v \n", *v.Name)

		}
	}
	return resourceGroupsInSub
}
