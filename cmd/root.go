package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	AzTenantId       string
	AzSubscriptionId string

	AzTenantsCmd = &cobra.Command{
		Use:     "tenants",
		Aliases: []string{},
		Long:    "Get all available Azure tenant IDs and display names",
		Short:   "Get all tenants",
		Run: func(cmd *cobra.Command, args []string) {
			err := GetTenants()
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
	rootCmd.AddCommand(
		AzTenantsCmd,
	)
}
