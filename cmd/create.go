package cmd

import (
	"appnetlab/pkg"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "name of the lab")
	createCmd.Flags().StringP("ip", "i", "", "ip of the lab")
	createCmd.Flags().Bool("int", false, "interface of the lab")
	createCmd.Flags().StringP("sh", "s", "", "shell use in the lab, example: -sh=bash")
	createCmd.Flags().IntP("upload", "u", 0, "bandwidth upload")
	createCmd.Flags().IntP("download", "d", 0, "bandwidth download")

	err := createCmd.MarkFlagRequired("name")
	if err != nil {
		fmt.Println(err)
	}

	err = createCmd.MarkFlagRequired("ip")
	if err != nil {
		fmt.Println(err)
	}
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create network lab",
	Long:  `create network lab`,
	Run: func(cmd *cobra.Command, args []string) {
		gwIpFull := "192.168.137.254/24"
		gwIp := "192.168.137.254"
		ip := cmd.Flag("ip").Value.String()
		ipFull := ip + "/24"
		name := cmd.Flag("name").Value.String()
		namespace := "net-lab-" + name
		dns := "8.8.8.8"
		enableInternet, err := cmd.Flags().GetBool("int")
		if err != nil {
			fmt.Println(err)
		}
		bandwidthUpload, err := cmd.Flags().GetInt("upload")
		if err != nil {
			fmt.Println(err)
		}
		bandwidthDownload, err := cmd.Flags().GetInt("download")
		if err != nil {
			fmt.Println(err)
		}

		sh, err := cmd.Flags().GetString("sh")
		if err != nil {
			fmt.Println(err)
		}
		if err != nil {
			fmt.Println(err)
		}

		//validate ip
		err = pkg.ValidateIp(ipFull)
		if err != nil {
			fmt.Println(err)
			return
		}
		//validate ip inside the range
		if gwIpFull[0:10] != ipFull[0:10] {
			fmt.Println("Invalid ip, must be in the range 192.168.137.x")
			return
		}

		//Create Bridge
		err = pkg.CreateBridge(gwIpFull)
		if err != nil {
			fmt.Println(err)
			return
		}

		//create namespace
		err = pkg.CreateNamespace(namespace)
		if err != nil {
			fmt.Println(err)
			return
		}

		//create Veth pair
		nIfaces, err := pkg.CreateVethPair()
		if err != nil {
			//delete namespace
			defer pkg.DeleteNamespace(namespace)
			fmt.Println(err)
			return
		}

		//add interface to namespace
		err = pkg.AddIfaceToNamespace(namespace, nIfaces[0])
		if err != nil {
			//delete namespace and interfaces
			defer pkg.DeleteNamespace(namespace)
			defer pkg.DeleteVethPair(nIfaces[1])
			fmt.Println(err)
			return
		}

		//set ip to interface in namespace
		err = pkg.SetIpInNamespace(ipFull, nIfaces[0], namespace)
		if err != nil {
			fmt.Println(err)
			//delete namespace and interface
			defer pkg.DeleteNamespace(namespace)
			defer pkg.DeleteVethPair(nIfaces[1])
			return
		}

		//set interface up in namespace
		err = pkg.UpIfaceInNamespace(namespace, nIfaces[0])
		if err != nil {
			fmt.Println(err)
		}

		//set gateway
		err = pkg.SetDefaultGatewayInNamespace(gwIp, nIfaces[0], namespace)
		if err != nil {
			fmt.Println(err)
		}

		//ser dns server
		err = pkg.SetDefaultDNSInNamespace(dns, namespace)
		if err != nil {
			fmt.Println(err)
		}

		//enable internet
		if enableInternet {
			//set interface up
			err = pkg.EnableNat(ip)
			if err != nil {
				fmt.Println(err)
			}
		}

		//set bandwidth Upload
		if bandwidthUpload > 0 {
			err = pkg.SetBandwidth(nIfaces[1], bandwidthUpload)
			if err != nil {
				fmt.Println(err)
			}
		}

		//set bandwidth Download
		if bandwidthDownload > 0 {
			err = pkg.SetBandwidthInNamespace(namespace, nIfaces[0], bandwidthDownload)
			if err != nil {
				fmt.Println(err)
			}
		}

		//run command in namespace
		if sh != "" {
			err = pkg.RunCmdInNamespace(namespace, sh)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err = pkg.RunCmdInNamespace(namespace, "bash")
			if err != nil {
				fmt.Println(err)
			}
		}

		//End of the lab
		//////////////////////

		//Delete nat
		if enableInternet {
			//set interface up
			err = pkg.DisableNat(ip)
			if err != nil {
				fmt.Println(err)
			}
		}

		//delete interface
		defer pkg.DeleteVethPair(nIfaces[1])

		//delete namespace
		defer pkg.DeleteNamespace(namespace)

	},
}