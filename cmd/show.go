package cmd

import (
	"appnetlab/pkg"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(showCmd)
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show names",
	Long:  `show namespaces running of App Net Lab`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg.ShowNamespaces()
	},
}
