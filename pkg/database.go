package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

func AzDatabase(AzCred *azidentity.DefaultAzureCredential, AzTenantId, AzSubscriptionId string, notify bool) error {

	if AzTenantId == "" && AzSubscriptionId != "" {

	}
	if AzTenantId == "" && AzSubscriptionId == "" {
		subscriptions := GetSubscriptions(AzCred)
		getAllDatabaseTables(AzCred, subscriptions)
	}
	return nil
}

// Get all servers for each resource group, then can get database
func getAllDatabaseTables(cred *azidentity.DefaultAzureCredential, subs []*armsubscriptions.Subscription) {

	for _, sub := range subs {
		resourceGroups := GetResourceGroups(cred, *sub.SubscriptionID)

		clientFactory, err := armsql.NewClientFactory(*sub.SubscriptionID, cred, nil)
		if err != nil {
			log.Fatalf("failed to create client: %v \n", err)
		}

		for _, rg := range resourceGroups {
			pager := clientFactory.NewServersClient().NewListByResourceGroupPager(*rg.Name, nil)
			ctx := context.TODO()

			for pager.More() {
				page, err := pager.NextPage(ctx)
				if err != nil {
					log.Fatalf("error occurred gettings SQL Servers... %v \n", err)
				}
				for _, v := range page.Value {
					fmt.Printf("SQL Server: %s \n", *v.Name)
					getDatabaseTablesInRg(*rg.Name, *v.Name, clientFactory)
				}
			}
		}
	}

}

// Return all tables for a given subscription
func getDatabaseTablesInRg(resourGroupName, serverName string, clientFactory *armsql.ClientFactory) {
	pager := clientFactory.NewDatabasesClient().NewListByServerPager(resourGroupName, serverName, nil)

	ctx := context.TODO()

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to get database tables... %v \n", err)
		}
		for _, v := range page.Value {
			fmt.Printf("%s  \n", *v.Name)
		}
	}
}
