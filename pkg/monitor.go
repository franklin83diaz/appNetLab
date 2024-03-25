package pkg

import (
	"fmt"
	"os"
	"os/exec"
)

func Monitor(iface string, filename string) {

	if filename != "" {
		//sudo tcpdump -i lab-veth4 -XX -vvv -w file.pcap
		cmd := exec.Command("sudo", "tcpdump", "-i", iface, "-XX", "-vvv", "-w", filename+".pcap")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("Error running tcpdump")
		}
	} else {
		cmd := exec.Command("sudo", "tcpdump", "-i", iface, "-XX", "-vvv")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("Error running tcpdump")
		}
	}

}
