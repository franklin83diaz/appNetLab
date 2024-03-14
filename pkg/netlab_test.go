package pkg_test

//test for getIfaces
import (
	"appnetlab/pkg"
	"os/exec"
	"testing"
)

func TestGetIfaces(t *testing.T) {
	exec.Command("sudo", "ip", "link", "add", "vth1", "type", "veth").Run()
	ifaces := pkg.GetIfaces()
	if len(ifaces) == 0 {
		t.Errorf("No interfaces found")
	}
	exec.Command("sudo", "ip", "link", "delete", "vth1").Run()
}

// test for nextIfaces
func TestNextIfaces(t *testing.T) {
	exec.Command("sudo", "ip", "link", "add", "vth1", "type", "veth").Run()
	iface := pkg.NextIfaces()
	if iface != "vth2" {
		t.Errorf("Expected vth2, got %s", iface)
	}
	exec.Command("sudo", "ip", "link", "delete", "vth1").Run()
}
