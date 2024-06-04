package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
)

// func main() {
//	fmt.Println("fetching azure resources")
// cred, err := azidentity.NewDefaultAzureCredential(nil)
// if err != nil {
// 	// TODO:
// 	fmt.Println("Auth error")
// } else {
// 	fmt.Println(&cred)
// }

// client, _ := GetCredential()
// fmt.Println(client)
//	tenants := GetTenants()
//	fmt.Println(tenants)
//	for _, v := range tenants {
//		fmt.Println(*v.TenantID, *v.DisplayName)
//	}
//	GetSubscriptions()
//}

func ListTenants(AzCred *azidentity.DefaultAzureCredential) (err error) {
	tenants := GetTenants(AzCred)

	for _, v := range tenants {
		fmt.Printf("%v %v \n", *v.TenantID, *v.DisplayName)
	}
	return nil
}

func GetTenants(AzCred *azidentity.DefaultAzureCredential) []*armsubscriptions.TenantIDDescription {
	// cred, _ := internal.GetCredential()
	// var result armsubscription.TenantIDDescription
	// ctx := context.Background()
	clientFactory, err := armsubscriptions.NewClientFactory(AzCred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.TODO()
	pager := clientFactory.NewTenantsClient().NewListPager(nil)
	var result []*armsubscriptions.TenantIDDescription
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		result = append(result, page.Value...)
		// for _, v := range page.Value {
		//	fmt.Printf("%v \n", *v.DefaultDomain)
		//	result = append(result, v)
		//	}
	}
	return result
}

func GetSubscriptions(AzCred *azidentity.DefaultAzureCredential) []*armsubscriptions.Subscription {
	// cred, _ := internal.GetCredential()
	ctx := context.TODO()
	var result []*armsubscriptions.Subscription
	clientFactory, err := armsubscriptions.NewClientFactory(AzCred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		// for _, v := range page.Value {
		// 	fmt.Printf("%v %v\n", *v.SubscriptionID, *v.DisplayName)
		// 	result = append(result, v)
		// }
		result = append(result, page.Value...)
	}

	return result
}

// Given a Subscription ID or name, verify that it exists and return
func getSubscriptionFromId(AzCred *azidentity.DefaultAzureCredential, AzSubscriptionId string) (isFound bool, subId *armsubscriptions.Subscription) {
	isFound = false
	subscriptions := GetSubscriptions(AzCred)

	for _, sub := range subscriptions {
		if *sub.DisplayName == AzSubscriptionId {
			isFound = true
			return isFound, sub

		} else if *sub.SubscriptionID == AzSubscriptionId {
			isFound = true
			return isFound, sub
		}

	}
	return
}

// Get all subs with matching tenantID, OR find a function that takes Tenant as an argument
func getSubscriptionsForTenant(AzCred *azidentity.DefaultAzureCredential, AzTenantId string) (tenantValid bool, subs []*armsubscriptions.Subscription) {
	tenants := GetTenants(AzCred)

	tenantValid = false

	for _, t := range tenants {
		if *t.DisplayName == AzTenantId {
			tenantValid = true
			break
		} else if *t.TenantID == AzTenantId {
			tenantValid = true
			break
		}
	}
	if tenantValid {
		subs := GetSubscriptions(AzCred)

		for _, s := range subs {
			if *s.TenantID == AzTenantId {
				subs = append(subs, s)
			}
			//fmt.Println(*s.TenantID)
		}
		return tenantValid, subs

	}
	return tenantValid, nil
}

// client, err := armsubscription.NewSubscriptionsClient(cred, nil)
// if err != nil {
// 	// TODO: Handle Error
// 	fmt.Println("New Subscription client error")
// }

// _, err = client.Get(context.TODO(), subscriptionID, nil)
// if err != nil {
// 	// TODO: handle error
// 	fmt.Println("Subscription error")
// }
