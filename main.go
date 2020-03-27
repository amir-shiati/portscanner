package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("Port scanner", "Scans ports")
	ip := parser.String("i", "ip", &argparse.Options{Required: true, Help: "The ip that you like to scan."})
	port := parser.String("p", "port", &argparse.Options{Required: true, Help: "The port that you like to scan."})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}
	scanAPort(*ip, *port)
}

func scanAPort(ip, port string) {
	fmt.Println("Scanning port")
	fmt.Println("Results%t\n", isPortOpen("tcp", ip, port))
}

func isPortOpen(protocol, hostName, port string) bool {
	address := hostName + ":" + port
	conn, err := net.DialTimeout(protocol, address, 10*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
