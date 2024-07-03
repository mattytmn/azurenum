package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/mattytmn/azurenum/internal"
)

// Gets all keyvaults in Azure
// TODO check for no keyvaults in subscription
// Entry method to enumerate keyvaults
func AzKeyVaults(AzCred *azidentity.DefaultAzureCredential, AzTenantID, AzSubscriptionID string, AzKvSecrets, AzKvCerts bool) error {

	// var subscriptions []*armsubscriptions.
	// Subscription given

	if AzTenantID == "" && AzSubscriptionID != "" {
		// verify ID exists
		isValid, subId := getSubscriptionFromId(AzCred, AzSubscriptionID)

		if isValid {
			fmt.Println("ID is valid, fetching keyvaults.. ")
			fmt.Println(*subId.DisplayName)
			sub := []*armsubscriptions.Subscription{subId}

			// If secret flag specified, get secrets for sub
			// else only get keyvaults
			if AzKvSecrets {
				fmt.Println("Getting secrets in subscription...")
				GetKeyVaultSecretsForSubscription(AzCred, *subId.SubscriptionID)

			} else {
				getKeyVaultsForSubscriptionSlice(AzCred, sub)
			}

		} else {
			log.Fatalf("No subscription found for ID: %s", AzSubscriptionID)
		}
	} else if AzTenantID != "" && AzSubscriptionID == "" {
		tenantExists, subs := getSubscriptionsForTenant(AzCred, AzTenantID)

		if tenantExists {
			for _, s := range subs {
				fmt.Printf("%v \n", *s.DisplayName)
			}
		}

	} else if AzTenantID == "" && AzSubscriptionID == "" {
		subscriptions := GetSubscriptions(AzCred)
		//getKeyVaultsForSubscriptionSlice(AzCred, subscriptions)

		// TODO
		// For each subscription get keyvault secrets
		if AzKvSecrets {
			for _, s := range subscriptions {
				GetKeyVaultSecretsForSubscription(AzCred, *s.SubscriptionID)
			}
		} else {
			getKeyVaultsForSubscriptionSlice(AzCred, subscriptions)
		}
		// for _, sub := range subscriptions {
		// 	fmt.Printf("getting subscriptions for keyvault enum... \n")
		// 	subscriptionId := *sub.SubscriptionID
		// 	clientFactory, err := armkeyvault.NewClientFactory(subscriptionId, AzCred, nil)
		// 	if err != nil {
		// 		log.Fatalf("failed to create client: %v \n", err)
		// 	}
		// 	pager := clientFactory.NewVaultsClient().NewListPager(nil)
		// 	ctx := context.TODO()

		// 	for pager.More() {
		// 		page, err := pager.NextPage(ctx)
		// 		if err != nil {
		// 			log.Fatalf("error occurred getting keyvaults... %v \n", err)
		// 		}

		// 		for _, v := range page.Value {
		// 			fmt.Printf("%T \n", v)
		// 			fmt.Printf("%T \n", *v)
		// 		}
		// 	}

		// }
	}

	return nil
}

// Currently not used
func AzKeyVaultSecrets(AzCred *azidentity.DefaultAzureCredential, AzTenantID, AzSubscriptionID string) error {
	if AzTenantID == "" && AzSubscriptionID != "" {
		// verify ID exists
		isValid, subId := getSubscriptionFromId(AzCred, AzSubscriptionID)

		if isValid {
			fmt.Println("ID is valid, fetching secrets.. ")
			fmt.Println(*subId.DisplayName)
			// sub := []*armsubscriptions.Subscription{subId}
			// getKeyVaultSecretsForSubscription(AzCred, sub)

		} else {
			log.Fatalf("No subscription found for ID: %s", AzSubscriptionID)
		}
	} else if AzTenantID != "" && AzSubscriptionID == "" {
		tenantExists, subs := getSubscriptionsForTenant(AzCred, AzTenantID)

		if tenantExists {
			for _, s := range subs {
				fmt.Printf("%v \n", *s.DisplayName)
			}
		}

	} else if AzTenantID == "" && AzSubscriptionID == "" {
		// subscriptions := GetSubscriptions(AzCred)
		// getKeyVaultSecretsForSubscription(AzCred, subscriptions)
	}

	return nil
}

// Given a slice of subscriptions, returns a slice of those keyvaults
// Should be used as an input to get secrets, certs etc
func getKeyVaultsForSubscriptionSlice(AzCred *azidentity.DefaultAzureCredential, subscription []*armsubscriptions.Subscription) (keyvaults []*armkeyvault.Resource) {
	outTable := internal.TableClient{
		Header: []string{"Name"},
	}
	for _, sub := range subscription {
		//fmt.Printf("Getting keyvaults in subcription: %v | %v\n", *sub.DisplayName, *sub.SubscriptionID)
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
				keyvaults = append(keyvaults, page.Value...)
				fmt.Printf("%v \n", *v.Name)
			}
		}
	}
	outTable.PrintResultAsTable(outTable)
	return
}

// Get keyvaults for single subscription
func getKeyVaultsForSubscription(AzCred *azidentity.DefaultAzureCredential, subscription string, ResourceGroupId string) (keyvaults []*armkeyvault.Vault) {

	//fmt.Printf("Getting keyvaults in subcription: %v \n", subscription)

	clientFactory, err := armkeyvault.NewClientFactory(subscription, AzCred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v \n", err)
	}
	pager := clientFactory.NewVaultsClient().NewListByResourceGroupPager(ResourceGroupId, nil)
	ctx := context.TODO()

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("error occurred getting keyvaults... %v \n", err)
		}

		keyvaults = append(keyvaults, page.Value...)
		// for _, v := range page.Value {
		// 	keyvaults = append(keyvaults, page.Value...)
		// 	// fmt.Printf("Keyvault %v \n", *v.Name)
		// }
	}

	return
}

// Given a list of keyvaults, return secrets
// TODO
// Make this run concurrently to speed up getting multiple secrets for multiple keyvaults
func GetKeyVaultSecretsForSubscription(AzCred *azidentity.DefaultAzureCredential, subscription string) (secrets []*armkeyvault.Secret) {
	// for _, sub := range subscription {
	// 	fmt.Printf("Getting secrets in Keyvault: %v | %v\n", *sub.DisplayName, *sub.SubscriptionID)
	// 	subscriptionId := *sub.SubscriptionID
	secretsCounter := 0
	resourceGroups := GetResourceGroups(AzCred, subscription)
	//_, subs := getSubscriptionFromId(AzCred, subscription)

	// get all resource groups and iterate through
	for _, rg := range resourceGroups {

		// fmt.Printf("Getting resource group key vaults: %v \n", *rg.Name)
		keyvaults := getKeyVaultsForSubscription(AzCred, subscription, *rg.Name)

		ctx := context.TODO()
		clientFactory, err := armkeyvault.NewClientFactory(subscription, AzCred, nil)

		if err != nil {
			log.Fatalf("failed to obtain credential: %v", err)
		}

		for _, kv := range keyvaults {
			pager := clientFactory.NewSecretsClient().NewListPager(*rg.Name, *kv.Name, nil)
			fmt.Printf("%v \n", *rg.Name)
			for pager.More() {
				page, err := pager.NextPage(ctx)
				if err != nil {
					log.Fatalf("Failed to advance page: %v\n", err)
				}
				for _, s := range page.Value {
					secretsCounter++
					if s.Properties.Attributes.Expires != nil {
						fmt.Printf("Secret name: %v | Secret expiry: %v | Secret URI: %v \n", *s.Name, *s.Properties.Attributes.Expires, *s.ID)
					} else {
						fmt.Printf("Secret name: %v | Secret enabled: %v \n", *s.Name, *s.Properties.Attributes.Enabled)

					}
				}
				secrets = append(secrets, page.Value...)
			}
		}

	}
	return nil
}

func AzKeyVaultCertificates() {}
