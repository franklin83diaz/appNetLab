package cmd

import (
	"appnetlab/pkg"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "netlab",
	Short: "netlab creator of network labs",
	Long: `
	Netlab is net lab tool for creating network labs,
	Confugure isolated network labs with the desired bandwidth for testing.
	and net monitoring tools.

	netlab whiout arguments will show current configuration running.
	
	Use the netlab command to configure the network lab.

	netlab -int -ip 192.168.137.2 -u 1024 -d 1024 ./bash

	this command will create a shell with the ip and the bandwidth specified 1Mbps.
	and internet access with the default gateway 192.168.137.254 and the dns 8.8.8.8
	the ips will are inside the range 192.168.137.1-253
	
	`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Current Configuration:")
		fmt.Println("Interfaces Bridge: lab-bridge")
		fmt.Printf("Interfaces: lab-veth\033[33mX\033[0m\n")
		fmt.Println("IPs: 192.168.137.1-253")
		fmt.Println("Default Gateway: 192.168.137.1.254")
		fmt.Println("DNS: 8.8.8.8")
		fmt.Println()
		fmt.Println("-------------------------")
		fmt.Println("Namespaces Runnig:")
		pkg.ShowNamespaces()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
