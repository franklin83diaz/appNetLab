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
	return nIfaces[0], nil
}

func SetIp(ip string, iface string) error {
	cmd := exec.Command("sudo", "ip", "addr", "add", ip, "dev", iface)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set ip to interface: %w", err)
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

// set mode router
func SetModeRoute() error {
	//echo 1 > /proc/sys/net/ipv4/ip_forward
	cmd := exec.Command("sudo", "echo", "1", ">", "/proc/sys/net/ipv4/ip_forward")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set mode router: %w", err)
	}
	return nil
}

// create namespace
func CreateNamespace(name string) error {
	cmd := exec.Command("sudo", "ip", "netns", "add", name)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create namespace: %w", err)
	}
	return nil
}

// add interface to namespace
func AddIfaceToNamespace(name string, iface string) error {
	cmd := exec.Command("sudo", "ip", "link", "set", iface, "netns", name)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to add interface to namespace: %w", err)
	}
	return nil
}

// set ip to interface in namespace
func SetIpInNamespace(ip string, iface string, name string) error {
	//ip netns exec ns-lab ip addr add
	command := "sudo ip netns exec " + name + " ip addr add " + ip + " dev " + iface
	cmd := exec.Command("sh", "-c", command)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set ip to interface in namespace: %w", err)
	}

	//set default route
	command = "sudo ip netns exec " + name + " ip route add default via " + ip
	cmd = exec.Command("sh", "-c", command)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set ip to interface in namespace: %w", err)
	}

	return nil
}

// up interfaces in namespace
func UpIfaceInNamespace(name string, iface string) error {
	//up interface
	command := "sudo ip netns exec " + name + " ip link set dev " + iface + " up"
	cmd := exec.Command("sh", "-c", command)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set ip to interface in namespace: %w", err)
	}

	//up interface loopback
	command = "sudo ip netns exec " + name + " ip link set dev lo up"
	cmd = exec.Command("sh", "-c", command)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set ip to interface in namespace: %w", err)
	}

	return nil
}

// run command in namespace
func RunCmdInNamespace(name string, command string) error {
	commandString := "sudo ip netns exec " + name + " " + command
	cmd := exec.Command("sh", "-c", commandString)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run command in namespace: %w", err)
	}
	return nil
}

// delete namespace
func DeleteNamespace(name string) error {
	cmd := exec.Command("sudo", "ip", "netns", "delete", name)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to delete namespace: %w", err)
	}
	return nil
}
