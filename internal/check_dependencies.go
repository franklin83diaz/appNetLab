package internal

import "appnetlab/pkg"

func CheckDependencies() {
	// check dependencies
	exists := pkg.CheckBinaryInPath("tc")
	if !exists {
		panic("tc not found, please install iproute2")
	}
	exists = pkg.CheckBinaryInPath("tcpdump")
	if !exists {
		panic("tcpdump not found, please install tcpdump")
	}
}
