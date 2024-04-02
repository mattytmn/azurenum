package pkg

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/mattytmn/azurenum/internal"
)

func AzBlobs() {
	cred, _ := internal.GetCredential()
	subscriptions := GetSubscriptions("")
	clientFactory, err := armstorage.NewClientFactory("1234", cred, nil)
}

func getStorageAccounts() {
}
