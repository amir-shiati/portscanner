package helpers

import (
	"net"
	"strings"
)

//IsPortValid : checks if a port is valid or not
func IsPortValid(port int) bool {
	return port > 0 && port < 65535
}

//IsIPValid : checks if a ip is valid or not.
func IsIPValid(ip string) bool {
	return net.ParseIP(ip) != nil
}

//IsProtocolValid : checks if protocol is valid or not.
func IsProtocolValid(protocol string) bool {
	protocol = strings.ToLower(protocol)
	return protocol == "tcp" || protocol == "udp"
}

//IsRangeValid : checks if given range is valid or not.
func IsRangeValid(from, to int) bool {
	return from <= to && from >= 0 && to <= 65535
}

//IsDurationValid : checks if duration is valid or not.
func IsDurationValid(duration int) bool {
	return duration >= 7 && duration <= 300
}
