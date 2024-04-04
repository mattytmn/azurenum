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

	AzBlobsCmd = &cobra.Command{
		Use:     "blobs [-s | --subscrition-id subscriptionID]",
		Aliases: []string{},
		Long: `Get all blobs that the given account has access to.
        Specifying a subscription ID will get the blobs only those tenants`,
		Short: "Get all blobs",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("getting blobs...")
			err := pkg.AzBlobs(AzAuth)
			if err != nil {
				log.Fatal(err)
			}
			// err := pkg.GetBl
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
		AzBlobsCmd,
	)
}
