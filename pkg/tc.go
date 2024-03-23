package pkg

import (
	"fmt"
	"os/exec"
)

// // Traffic control
func SetBandwidth(iface string, rateKbit int) error {
	//set bandwidth

	cmd := exec.Command("sudo", "tc", "qdisc", "add", "dev", iface, "root", "handle", "1:", "htb", "default", "99")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set bandwidth: %w", err)
	}

	cmd = exec.Command("sudo", "tc", "class", "add", "dev", iface, "parent", "1:", "classid", "1:99", "htb", "rate", fmt.Sprintf("%dkbit", rateKbit), "ceil", fmt.Sprintf("%dkbit", rateKbit))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set bandwidth: %w", err)
	}
	return nil
}

func SetBandwidthInNamespace(namespace string, iface string, rateKbit int) error {

	//set bandwidth
	cmd := exec.Command("sudo", "ip", "netns", "exec", namespace, "tc", "qdisc", "add", "dev", iface, "root", "handle", "1:", "htb", "default", "99")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set bandwidth: %w", err)
	}

	cmd = exec.Command("sudo", "ip", "netns", "exec", namespace, "tc", "class", "add", "dev", iface, "parent", "1:", "classid", "1:99", "htb", "rate", fmt.Sprintf("%dkbit", rateKbit), "ceil", fmt.Sprintf("%dkbit", rateKbit))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set bandwidth: %w", err)
	}
	return nil
}
