package cmd

import (
	"appnetlab/pkg"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(goinCmd)
	goinCmd.Flags().StringP("name", "n", "", "name of the namespace")
	goinCmd.Flags().StringP("sh", "s", "bash", "command to run in the namespace, example: -sh=bash")

	err := goinCmd.MarkFlagRequired("name")
	if err != nil {
		fmt.Println(err)
	}

}

var goinCmd = &cobra.Command{
	Use:   "go-in",
	Short: "go-in namespace",
	Long:  `go-in namespace`,
	Run: func(cmd *cobra.Command, args []string) {
		subName := "net-lab-"
		name := cmd.Flag("name").Value.String()
		sh := cmd.Flag("sh").Value.String()

		if len(name) >= 8 {
			if name[:8] != subName {
				name = subName + name
			}
		} else {
			name = subName + name
		}

		if sh != "" {
			err := pkg.RunCmdInNamespace(name, sh)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err := pkg.RunCmdInNamespace(name, "bash")
			if err != nil {
				fmt.Println(err)
			}
		}

	},
}
