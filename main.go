package main

import (
	"fmt"
	"os"

	helpers "./helpers"
	ps "./portscanner"
	"github.com/akamensky/argparse"
)

var protocol string = "tcp"
var common bool = false
var duration int = 10

func main() {
	parser := argparse.NewParser("Port scanner", "Scans ports")
	ip := parser.String("i", "ip", &argparse.Options{Required: true, Help: "\n specifies an ip that you like to scan. \n must be a valid ip"})
	port := parser.String("p", "port", &argparse.Options{Required: false, Help: "\n specifies a port that you like to scan.\n must be a number between '0' and '65535'"})
	sProtocol := parser.String("l", "protocol", &argparse.Options{Required: false, Help: "\n specifies the protocol. \n must be either 'tcp' or 'udp' \n default value is 'tcp'"})
	fCommon := parser.String("c", "common", &argparse.Options{Required: false, Help: "\n if 'true' scans common ports: '0' to '1024' \n default value is 'false' "})
	sDuration := parser.String("d", "duration", &argparse.Options{Required: false, Help: "\n duration of connection. \n must be a value between '7' and '300'\n default value is '10' "})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	if *fCommon != "" {
		common = helpers.StringToBool(*fCommon)
	}
	if *sProtocol != "" {
		protocol = *sProtocol
	}
	if *port == "" {
		*port = "0"
	}
	if *sDuration != "" {
		duration = helpers.StringToInt(*sDuration)
	}

	toDo(*ip, *port, protocol, common, duration)
}

func toDo(ip, port, dProtocol string, dCommon bool, dDuration int) {
	if !dCommon && port != "" {
		ps.ScanAPort(ip, port, protocol, duration)
		return
	}
	if dCommon {
		ps.ScanRange(0, 1024, ip, port, protocol, duration)
	}
}
