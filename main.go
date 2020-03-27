package main

import (
	"fmt"
	"os"

	ps "./portscanner"
	"github.com/akamensky/argparse"
)

var protocol string = "tcp"

func main() {
	parser := argparse.NewParser("Port scanner", "Scans ports")
	ip := parser.String("i", "ip", &argparse.Options{Required: true, Help: "specifies an ip that you like to scan. \n must be a valid ip"})
	port := parser.String("p", "port", &argparse.Options{Required: false, Help: "specifies a port that you like to scan.\n must be a number between '0' and '65535'"})
	sProtocol := parser.String("r", "protocol", &argparse.Options{Required: false, Help: "specifies the protocol. \n must be either 'tcp' or 'udp' "})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}
	if *sProtocol != "" {
		protocol = *sProtocol
	}
	ps.ScanAPort(*ip, *port, protocol)
}
