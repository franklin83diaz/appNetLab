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

func CreateVethPair() error {
	nIfaces := NextIfaces()
	cmd := exec.Command("sudo", "ip", "link", "add", nIfaces[0], "type", "veth", "peer", "name", nIfaces[1])
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create veth pair: %w", err)
	}
	return nil
}

func DeleteVethPair(veth string) error {
	cmd := exec.Command("sudo", "ip", "link", "delete", veth)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to delete veth pair: %w", err)
	}
	return nil
}
