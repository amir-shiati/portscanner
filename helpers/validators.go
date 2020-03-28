package helpers

import (
	"net"
	"strings"
)

//IsPortValid : checks if a port is valid or not
func IsPortValid(port int) bool {
	if port < 0 || port > 65535 {
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

//IsRangeValid : checks if given range is valid or not.
func IsRangeValid(from, to int) bool {
	if from > to || from < 0 || to > 65535 {
		return false
	}
	return true
}

//IsDurationValid : checks if duration is valid or not.
func IsDurationValid(duration int) bool {
	if duration < 7 || duration > 300 {
		return false
	}
	return true
}
