package cmd

import (
	"appnetlab/pkg"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(monitorCmd)
	monitorCmd.Flags().StringP("file", "f", "", "file to save the traffic")
	monitorCmd.Flags().StringP("name", "l", "", "name of the lab")

	err := monitorCmd.MarkFlagRequired("name")
	if err != nil {
		panic(err)
	}
}

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "monitor network traffic",
	Long:  `monitor network traffic and save to file if needed`,
	Run: func(cmd *cobra.Command, args []string) {
		fileName := cmd.Flag("file").Value.String()
		name := "net-lab-" + cmd.Flag("name").Value.String()
		fmt.Println("Monitoring traffic in lab", name)
		if fileName != "" {
			fmt.Println("Saving to file", fileName+".pcap")
		}
		fmt.Println()
		iface, err := pkg.GetIfacesLinkNs(name)

		if err != nil {
			panic(err)
		}
		pkg.Monitor(iface, fileName)
		if fileName != "" {
			fmt.Println("\nopen with: \n wireshark", fileName+".pcap")
		}
	},
}
