package cmd

import (
	"fmt"
	"log"

	"github.com/mattytmn/azurenum/internal"
	"github.com/mattytmn/azurenum/pkg"
	"github.com/spf13/cobra"
)

var (
	AzTenant          string
	AzSubscription    string
	TeamsNotification bool
	AzAuth, _         = internal.GetCredential()

	AzTenantsCmd = &cobra.Command{
		Use:     "tenants",
		Aliases: []string{},
		Long:    `Get all available Azure tenant IDs and display names`,
		Short:   "Get all tenants",
		Run: func(cmd *cobra.Command, args []string) {
			err := pkg.ListTenants(AzAuth)
			if err != nil {
				log.Fatalf("%T", err)
			}
		},
	}

	AzSubscriptionsCmd = &cobra.Command{
		Use:     "subscriptions [-s | --subscrition subscriptionID]",
		Aliases: []string{"subs"},
		Long:    `Get all available Azure subscriptions and display names`,
		Short:   "Get all subscriptions",
		Run: func(cmd *cobra.Command, args []string) {
			err := pkg.OutputSubscriptions(AzAuth)
			if err != nil {
				log.Fatalf("An error occurred: %v", err)
			}
		},
	}

	AzStorageAccountCmd = &cobra.Command{
		Use:     "storage-account",
		Aliases: []string{"sa"},
		Long: `Get all blobs that the given account has access to.
        Specifying a subscription ID will get the storage accounts in only those subscription`,
		Short: "Get all Storage Accounts",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Enumerating blobs...")
			err := pkg.AzStorageAccount(AzAuth, AzSubscription)
			if err != nil {
				log.Fatal(err)
			}
			// err := pkg.GetBl
		},
	}
	KeyVaultListSecrets      bool
	KeyvaultListCertificates bool
	ExpiryDays               int
	AzKeyvaultCmd            = &cobra.Command{
		Use:     "keyvault",
		Aliases: []string{"kv"},
		Long:    `Get all keyvaults that the given account has access to. Specifying a subscription will only get the keyvaults in that subscription`,
		Short:   "Get all keyvaults",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Enumerating keyvaults...")
			err := pkg.AzKeyVaults(AzAuth, AzTenant, AzSubscription, KeyVaultListSecrets, KeyvaultListCertificates, TeamsNotification, ExpiryDays)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	// External attack surface management monitor
	AzEasmCmd = &cobra.Command{
		Use:     "easm",
		Aliases: []string{"m"},
		Long:    `Get all externally exposed resources and associated URIs`,
		Short:   "EASM module",
	}

	AzResourceGroupCmd = &cobra.Command{
		Use:     "resource-groups",
		Aliases: []string{"rg"},
		Long:    `Get all resource groups`,
		Short:   "Get all resource groups",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Getting resource groups...")
			err := pkg.GetResourceGroups(AzAuth, AzSubscription)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	AzKVSecretsCmd = &cobra.Command{
		Use:     "secrets",
		Aliases: []string{"sec"},
		Long:    `Get all secrets in Azure Key Vault`,
		Short:   "Get all secrets",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Getting resource groups...")
			err := pkg.GetKeyVaultSecretsForSubscription(AzAuth, AzSubscription, TeamsNotification)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	AzContainerAppsCmd = &cobra.Command{
		Use:     "container-apps",
		Aliases: []string{"ca"},
		Short:   "Get all container apps",
		Long:    `Get all container apps and associated URLs`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Getting container app...")
			err := pkg.AzContainerApps(AzAuth, AzTenant, AzSubscription)
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
	rootCmd.PersistentFlags().BoolVarP(&TeamsNotification, "notify", "n", false, "Send result notification to teams")

	AzKeyvaultCmd.Flags().BoolVarP(&KeyVaultListSecrets, "secrets", "x", false, "Only list secrets for key vault")

	AzKeyvaultCmd.Flags().BoolVarP(&KeyvaultListCertificates, "certificates", "c", false, "Only list certificates for key vault")
	AzKeyvaultCmd.Flags().IntVarP(&ExpiryDays, "days", "d", 30, "Number of days to secret expiry")
	rootCmd.AddCommand(
		AzTenantsCmd,
		AzSubscriptionsCmd,
		AzStorageAccountCmd,
		AzKeyvaultCmd,
		AzResourceGroupCmd,
		AzContainerAppsCmd,
		AzKVSecretsCmd,
	)
}
