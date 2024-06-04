package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
)

// Gets all keyvaults in Azure
// TODO change this to private method, create a list keyvaults method instead
// TODO check for no keyvaults in subscription
func AzKeyVaults(AzCred *azidentity.DefaultAzureCredential, AzTenantID, AzSubscriptionID string, AzKvSecrets, AzKvCerts bool) error {

	// var subscriptions []*armsubscriptions.
	if AzTenantID == "" && AzSubscriptionID != "" {
		// verify ID exists
		isValid, subId := getSubscriptionFromId(AzCred, AzSubscriptionID)

		if isValid {
			fmt.Println("ID is valid, fetching keyvaults.. ")
			fmt.Println(*subId.DisplayName)
			sub := []*armsubscriptions.Subscription{subId}
			if AzKvSecrets {
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
		getKeyVaultsForSubscriptionSlice(AzCred, subscriptions)

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

func getKeyVaultsForSubscriptionSlice(AzCred *azidentity.DefaultAzureCredential, subscription []*armsubscriptions.Subscription) (keyvaults []*armkeyvault.Resource) {

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
	return
}

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

		for _, v := range page.Value {
			keyvaults = append(keyvaults, page.Value...)
			fmt.Printf("Keyvault %v \n", *v.Name)
		}
	}

	return
}

// Given a list of keyvaults, return secrets
func GetKeyVaultSecretsForSubscription(AzCred *azidentity.DefaultAzureCredential, subscription string) (secrets []*armkeyvault.Secret) {
	// for _, sub := range subscription {
	// 	fmt.Printf("Getting secrets in Keyvault: %v | %v\n", *sub.DisplayName, *sub.SubscriptionID)
	// 	subscriptionId := *sub.SubscriptionID

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
					if s.Properties.Attributes.Expires != nil {
						fmt.Printf("Secret name: %v | Secret expiry: %v \n", *s.Name, *s.Properties.Attributes.Expires)
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
