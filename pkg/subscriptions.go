package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/mattytmn/azurenum/internal"
)

// func main() {
//	fmt.Println("fetching azure resources")
// cred, err := azidentity.NewDefaultAzureCredential(nil)
// if err != nil {
// 	// TODO: andl
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

func GetTenants() []*armsubscriptions.TenantIDDescription {
	cred, _ := internal.GetCredential()
	// var result armsubscription.TenantIDDescription
	// ctx := context.Background()
	clientFactory, err := armsubscriptions.NewClientFactory(cred, nil)
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
		for _, v := range page.Value {
			fmt.Printf("%v \n", *v.DefaultDomain)
			result = append(result, v)
		}
	}
	return result
}

func GetSubscriptions(AzSubscription string) []*armsubscriptions.Subscription {
	fmt.Println(AzSubscription)
	cred, _ := internal.GetCredential()
	ctx := context.TODO()
	var result []*armsubscriptions.Subscription
	clientFactory, err := armsubscriptions.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			fmt.Printf("%v \n", *v.DisplayName)
			result = append(result, v)
		}
	}

	return result
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
