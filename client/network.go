package main

import (
	"net"
	"strings"
)

func getNetworkAddrs(prefix string) ([]string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	var addrList []string
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				if len(prefix) > 0 {
					if strings.HasPrefix(ipnet.IP.String(), prefix) {
						addrList = append(addrList, ipnet.IP.String())
					}
				} else {
					addrList = append(addrList, ipnet.IP.String())
				}
			}
		}
	}

	return addrList, nil
}
