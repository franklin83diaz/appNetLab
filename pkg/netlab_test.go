package pkg_test

//test for getIfaces
import (
	"appnetlab/pkg"
	"os/exec"
	"testing"
)

func TestGetIfaces(t *testing.T) {
	exec.Command("sudo", "ip", "link", "add", "lab-veth0", "type", "veth", "peer", "name", "veth1").Run()
	ifaces := pkg.GetIfaces()
	if len(ifaces) == 0 {
		t.Errorf("No interfaces found")
	}
	exec.Command("sudo", "ip", "link", "delete", "lab-veth0").Run()
}

// test for nextIfaces
func TestNextIfaces(t *testing.T) {
	exec.Command("sudo", "ip", "link", "add", "lab-veth0", "type", "veth", "peer", "name", "lab-veth1").Run()
	iface := pkg.NextIfaces()
	if iface[0] != "lab-veth2" {
		t.Errorf("Expected lab-veth2, got %s", iface)
	}
	if iface[1] != "lab-veth3" {
		t.Errorf("Expected lab-veth3, got %s", iface)
	}
	exec.Command("sudo", "ip", "link", "delete", "lab-veth0").Run()
}

// test for createVethPair
func TestCreateVethPair(t *testing.T) {
	err := pkg.CreateBridge("192.168.137.1/24")
	if err != nil {
		t.Errorf("Error creating bridge: %s", err)
	}

	iface, err := pkg.CreateVethPair()
	if err != nil {
		t.Errorf("Error creating veth pair: %s", err)
	}
	exec.Command("sudo", "ip", "link", "delete", iface).Run()
	exec.Command("sudo", "ip", "link", "delete", "lab-bridge").Run()
}

// // Create 2 interfaces and a bridge and make ping between them
func TestPing(t *testing.T) {

	namespace := "ns-lab-01"

	//set mode route
	err := pkg.SetModeRoute()
	if err != nil {
		t.Errorf("Error setting mode route: %s", err)
	}

	//create bridge
	err = pkg.CreateBridge("192.168.137.254/24")
	if err != nil {
		t.Errorf("Error creating bridge: %s", err)
	}
	//create ns
	err = pkg.CreateNamespace(namespace)
	if err != nil {
		t.Errorf("Error creating namespace: %s", err)
	}

	//create veth pair
	iface1, err := pkg.CreateVethPair()
	if err != nil {
		t.Errorf("Error creating veth pair: %s", err)
	}

	//add veth1 to ns
	err = pkg.AddIfaceToNamespace(namespace, iface1)
	if err != nil {
		t.Errorf("Error adding interface to namespace: %s", err)
	}

	//set ip to veth1 in ns
	err = pkg.SetIpInNamespace("192.168.137.2/24", iface1, namespace)
	if err != nil {
		t.Errorf("Error setting ip to interface in namespace: %s", err)
	}
	//set veth1 up
	err = pkg.UpIfaceInNamespace(namespace, iface1)
	if err != nil {
		t.Errorf("Error setting interface up in namespace: %s", err)
	}

	//ser gateway
	err = pkg.SetDefaultGatewayInNamespace("192.168.137.254", iface1, namespace)
	if err != nil {
		t.Errorf("Error setting default gateway in namespace: %s", err)
	}

	//enable nat for namespace
	err = pkg.EnableNat("192.168.137.2")
	if err != nil {
		t.Errorf("Error enabling nat in namespace: %s", err)
	}

	//delete all
	exec.Command("sudo", "ip", "link", "delete", "lab-bridge").Run()
	exec.Command("sudo", "ip", "netns", "delete", namespace).Run()
	exec.Command("sudo", "ip", "link", "delete", iface1).Run()

}
