package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func IsRoot() bool {
	return os.Geteuid() == 0
}

// run command in namespace
func RunCmdInNamespace(name string, command string) error {
	cmd := exec.Command("sudo", "ip", "netns", "exec", name, command)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run command in namespace: %w", err)
	}
	return nil
}

func ValidateIp(ip string) error {
	singleIp := strings.Split(ip, "/")
	//validate ip format
	bytesIp := strings.Split(singleIp[0], ".")
	if len(bytesIp) != 4 {
		return fmt.Errorf("ip format invalid")
	}

	//validate bytes ip
	for _, b := range bytesIp {
		//to int
		byteIp, err := strconv.Atoi(b)
		if err != nil {
			return fmt.Errorf("ip format invalid")
		}
		//validate range
		if byteIp < 0 || byteIp > 255 {
			return err
		}
	}

	//validate ip exists
	cmd := exec.Command("bash", "-c", "ip a | grep "+ip)
	if err := cmd.Run(); err == nil {
		return fmt.Errorf("ip already exists")
	}

	return nil
}

// check if binary is in path
func CheckBinaryInPath(binaryName string) bool {
	// Obtener el PATH del sistema.
	path := os.Getenv("PATH")
	paths := strings.Split(path, string(os.PathListSeparator))

	// Verificar si el binario está en alguno de los directorios del PATH.
	for _, p := range paths {
		fullPath := filepath.Join(p, binaryName)
		if _, err := os.Stat(fullPath); err == nil {
			return true
		}
	}
	return false
}

func GetIfacesLinkNs(namespace string) (string, error) {
	output, err := exec.Command("ip", "link").CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running 'ip link': %w", err)
	}

	lines := strings.Split(string(output), "\n")

	// get interfaces linked to namespace
	var ifaceFound string
	for i, line := range lines {

		if strings.Contains(line, "link-netns "+namespace) {
			ifaceFound = strings.Split(strings.Split(lines[i-1], ": ")[1], "@")[0]
			ifaceFound = strings.TrimSpace(ifaceFound)
			break
		}
	}

	return ifaceFound, nil
}
