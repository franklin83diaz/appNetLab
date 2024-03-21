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
	err := pkg.CreateVethPair()
	if err != nil {
		t.Errorf("Error creating veth pair: %s", err)
	}
	exec.Command("sudo", "ip", "link", "delete", "lab-veth1").Run()
}
