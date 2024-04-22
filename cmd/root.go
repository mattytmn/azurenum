package cmd

import (
	"fmt"
	"log"

	"github.com/mattytmn/azurenum/internal"
	"github.com/mattytmn/azurenum/pkg"
	"github.com/spf13/cobra"
)

var (
	AzTenant       string
	AzSubscription string
	AzAuth, _      = internal.GetCredential()

	AzTenantsCmd = &cobra.Command{
		Use:     "tenants",
		Aliases: []string{},
		Long:    `Get all available Azure tenant IDs and display names`,
		Short:   "Get all tenants",
		Run: func(cmd *cobra.Command, args []string) {
			err := pkg.GetTenants(AzAuth)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	AzSubscriptionsCmd = &cobra.Command{
		Use:     "subscriptions [-s | --subscrition-id subscriptionID]",
		Aliases: []string{"subs"},
		Long:    `Get all available Azure subscriptions and display names`,
		Short:   "Get all subscriptions",
		Run: func(cmd *cobra.Command, args []string) {
			err := pkg.GetSubscriptions(AzAuth)
			if err != nil {
				log.Fatalf("An error occurred: %v", err)
			}
		},
	}

	AzStorageAccountCmd = &cobra.Command{
		Use:     "storage-account [-s | --subscrition-id subscriptionID]",
		Aliases: []string{"sa"},
		Long: `Get all blobs that the given account has access to.
        Specifying a subscription ID will get the storage accounts in only those subscription`,
		Short: "Get all Storage Accounts",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("getting blobs...")
<<<<<<< HEAD
			err := pkg.AzBlobs()
=======
			err := pkg.AzBlobs(AzAuth)
>>>>>>> feature/keyvault
			if err != nil {
				log.Fatal(err)
			}
			// err := pkg.GetBl
		},
	}

	AzKeyvaultCmd = &cobra.Command{
		Use:     "keyvault [-s | subscription-id --subscriptionID] [-k | --keyvault-id keyvaultID]",
		Aliases: []string{"kv"},
		Long:    `Get all keyvaults that the given account has access to. Specifying a subscription will only get the keyvaults in that subscription`,
		Short:   "Get all keyvaults",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("getting keyvaults...")
			err := pkg.AzKeyVaults(AzAuth)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

var rootCmd = &cobra.Command{
	Use:   "azurenum",
	Short: "Enumerate all Azure resources",
	Long:  "A fast Azure enumeration tool to identify and monitor Azure resources",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Flags
	rootCmd.PersistentFlags().StringVarP(&AzSubscription, "subscription", "s", "", "Subscription ID")
	rootCmd.PersistentFlags().StringVarP(&AzTenant, "tenant", "t", "", "Tenant name or ID")
	rootCmd.AddCommand(
		AzTenantsCmd,
		AzSubscriptionsCmd,
		AzStorageAccountCmd,
		AzKeyvaultCmd,
	)
}
