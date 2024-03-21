package pkg

import (
	"fmt"
	"net"
	"os/exec"
	"strconv"
)

func GetIfaces() (list []string) {

	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, iface := range interfaces {
		//start with vth
		if len(iface.Name) > 8 && iface.Name[:8] == "lab-veth" {
			list = append(list, iface.Name)
		}
	}

	return
}

func NextIfaces() []string {

	s := []int{}
	returnIfaces := []string{}

	ifaces := GetIfaces()

	//get the number of the interfaces
	for _, iface := range ifaces {
		n, _ := strconv.Atoi(iface[8:])
		s = append(s, n)
	}

	//get the max number
	max := 0
	for _, n := range s {
		if n > max {
			max = n
		}
	}
	returnIfaces = append(returnIfaces, "lab-veth"+strconv.Itoa(max+1))
	returnIfaces = append(returnIfaces, "lab-veth"+strconv.Itoa(max+2))
	return returnIfaces

}

func CreateVethPair() (string, error) {
	nIfaces := NextIfaces()
	//create veth pair
	cmd := exec.Command("sudo", "ip", "link", "add", nIfaces[0], "type", "veth", "peer", "name", nIfaces[1])
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to create veth pair: %w", err)
	}
	//create bridge
	cmd = exec.Command("sudo", "ip", "link", "set", nIfaces[0], "master", "lab-bridge")
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to set veth pair to bridge: %w", err)
	}
	return nIfaces[1], nil
}

func DeleteVethPair(veth string) error {
	cmd := exec.Command("sudo", "ip", "link", "delete", veth)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to delete veth pair: %w", err)
	}
	return nil
}

func CreateBridge(ip string) error {
	cmd := exec.Command("sudo", "ip", "link", "add", "name", "lab-bridge", "type", "bridge")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to delete veth pair: %w", err)
	}
	//set ip to bridge
	cmd = exec.Command("sudo", "ip", "addr", "add", ip, "dev", "lab-bridge")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to add ip to bridge: %w", err)
	}
	//set bridge up
	cmd = exec.Command("sudo", "ip", "link", "set", "lab-bridge", "up")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set bridge up: %w", err)
	}
	return nil
}
