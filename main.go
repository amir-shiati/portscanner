package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	helpers "./helpers"
	ps "./portscanner"
	"github.com/akamensky/argparse"
)

var protocol string = "tcp"
var common bool = false
var rangeScan bool = false
var duration int = 10

func main() {
	parser := argparse.NewParser("Port scanner", "")
	ip := parser.String("i", "ip", &argparse.Options{Required: true, Help: "\n *Required* \n specifies an ip that you like to scan. \n must be a valid ip"})
	port := parser.String("p", "port", &argparse.Options{Required: false, Help: "\n specifies a port that you like to scan.\n must be a number between '0' and '65535'"})
	sProtocol := parser.String("l", "protocol", &argparse.Options{Required: false, Help: "\n specifies the protocol. \n must be either 'tcp' or 'udp' \n default value is 'tcp'"})
	fCommon := parser.String("c", "common", &argparse.Options{Required: false, Help: "\n if 'true' scans common ports: '0' to '1024' \n default value is 'false' "})
	fRangeScane := parser.String("r", "range", &argparse.Options{Required: false, Help: "\n if 'true' asks for ranges \n default value is 'false' "})
	sDuration := parser.String("d", "duration", &argparse.Options{Required: false, Help: "\n duration of connection. \n must be a value between '7' and '300'\n default value is '10' "})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	if *fCommon != "" {
		common = helpers.StringToBool(*fCommon)
	}
	if *fRangeScane != "" {
		rangeScan = helpers.StringToBool(*fRangeScane)
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

	toDo(*ip, *port, protocol, common, rangeScan, duration)
}

func toDo(ip, port, dProtocol string, dCommon, dRangeScan bool, dDuration int) {
	if !dCommon && !rangeScan && port != "" {
		ps.ScanAPort(ip, port, protocol, duration)
		return
	}
	if dCommon && !rangeScan {
		ps.ScanRange(0, 1024, ip, port, protocol, duration)
		return
	}
	if !dCommon && rangeScan {
		from := readNum("Enter the firts port number")
		to := readNum("Enter the seconde port number")
		ps.ScanRange(from, to, ip, port, protocol, duration)
		return
	}

	fmt.Println("Wronge arguments! use --help for more information.")
}

func readNum(message string) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message + ": ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return helpers.StringToInt(text)
}
