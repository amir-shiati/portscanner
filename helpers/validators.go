package helpers

import (
	"net"
	"strings"
)

//IsPortValid : checks if a port is valid or not
func IsPortValid(port int) bool {
	if port > 0 || port < 65535 {
		return false
	}
	return true
}

//IsIPValid : checks if a ip is valid or not.
func IsIPValid(ip string) bool {
	return net.ParseIP(ip) != nil
}

//IsProtocolValid : checks if protocol is valid or not.
func IsProtocolValid(protocol string) bool {
	protocol = strings.ToLower(protocol)
	if protocol == "tcp" || protocol == "udp" {
		return true
	}
	return false
}
