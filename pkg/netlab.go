package pkg

import (
	"net"
	"strconv"
)

func GetIfaces() (list []string) {

	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, iface := range interfaces {
		//start with vth
		if len(iface.Name) > 3 && iface.Name[:3] == "vth" {
			list = append(list, iface.Name)
		}
	}
	return
}

func NextIfaces() string {

	s := []int{}

	ifaces := GetIfaces()

	//if there are no interfaces, return vth1
	if len(ifaces) == 0 {
		return "vth1"
	}

	//get the number of the interfaces
	for _, iface := range ifaces {
		n, _ := strconv.Atoi(iface[3:])
		s = append(s, n)
	}

	//get the max number
	max := s[0]
	for _, n := range s {
		if n > max {
			max = n
		}
	}

	return "vth" + strconv.Itoa(max+1)

}
