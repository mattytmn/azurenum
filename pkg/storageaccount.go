package pkg

import (
	"context"
	"log"
	"strconv"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/mattytmn/azurenum/internal"
	"github.com/schollz/progressbar/v3"
)

var (
	publicNetworkAccess = "nil"
	encrypted           = false
)

func AzStorageAccount(AzCred *azidentity.DefaultAzureCredential, AzSubscriptionID string) error {
	// cred, _ := internal.GetCredential()

	// Iterate through all subscription to get all Storage Accounts
	outTable := internal.TableClient{}
	outTable.Header = []string{"Subscription", "Name", "Network Rule", "Public Access", "Blob Access"}
	subscriptions := GetSubscriptions(AzCred)
	bar := progressbar.Default(int64(len(subscriptions)), "Subscriptions")
	for _, sub := range subscriptions {

		//fmt.Printf("currently in subscription: %d of %d \n", index+1, len(subscriptions))
		bar.Add(1)
		subscriptionId := *sub.SubscriptionID
		clientFactory, err := armstorage.NewClientFactory(subscriptionId, AzCred, nil)
		if err != nil {
			log.Fatalf("failed to create client: %v \n", err)
		}

		pager := clientFactory.NewAccountsClient().NewListPager(nil)
		ctx := context.TODO()

		for pager.More() {
			page, err := pager.NextPage(ctx)
			if err != nil {
				defer log.Printf("Error occurred getting Storage accounts for subscription %v... %v", *sub.DisplayName, err)

				// Go to the next subscription if an error or no storage accounts are returned
				break
			}
			// Return all storage accounts for given subscription
			for _, v := range page.Value {
				if v.Properties.PublicNetworkAccess != nil {
					outTable.Body = append(outTable.Body, []string{
						*sub.DisplayName,
						*v.Name,
						string(*v.Properties.NetworkRuleSet.DefaultAction),
						string(*v.Properties.PublicNetworkAccess), strconv.FormatBool(*v.Properties.AllowBlobPublicAccess)})
					//fmt.Printf("%v | %v | %v \n", *v.Name, *v.Properties.NetworkRuleSet.DefaultAction, *v.Properties.PublicNetworkAccess)
				} else {
					outTable.Body = append(outTable.Body, []string{
						*sub.DisplayName,
						*v.Name,
						string(*v.Properties.NetworkRuleSet.DefaultAction),
						"N/A"})
					//fmt.Printf("%v | %v | %v \n", *v.Name, *v.Properties.NetworkRuleSet.DefaultAction, publicNetworkAccess)
				}
			}
		}
	}

	outTable.PrintResultAsTable(outTable)
	return nil
}

func getStorageAccounts() {
}
