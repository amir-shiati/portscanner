package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	fmt.Println("scanning ports")
	fmt.Println("results: %t\n", isPortOpen("tcp", "185.189.112.194", 8002))
}

func isPortOpen(protocol, hostName string, port int) bool {
	address := hostName + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 10*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
