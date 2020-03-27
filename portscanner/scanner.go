package portscanner

import (
	"fmt"
	"net"
	"strconv"
	"time"

	helpers "../helpers"
)

//Result : result of a scan.
type Result struct {
	Port     int
	State    string
	Protocol string
}

//ScanAPort : scans only one port.
func ScanAPort(ip, port, protocol string) {
	fmt.Println("Scanning port")
	result := IsPortOpen(protocol, ip, helpers.StringToInt(port))
	fmt.Println("Results: \n", result)
}

//IsPortOpen : checks if a port is open or not , returns a Result type.
func IsPortOpen(protocol, hostName string, port int) Result {
	var result Result
	address := hostName + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 10*time.Second)
	result.Port = port
	result.Protocol = protocol
	if err != nil {
		result.State = "Closed"
		return result
	}
	defer conn.Close()
	result.State = "Open"
	return result
}
