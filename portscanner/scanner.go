package portscanner

import (
	"fmt"
	"net"
	"sort"
	"strconv"
	"sync"
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
func ScanAPort(ip, port, protocol string, dDuration int) {
	if !areInputsValid(ip, port, protocol, dDuration) {
		return
	}

	fmt.Println("Scanning port...")
	result := isPortOpen(protocol, ip, helpers.StringToInt(port), dDuration)
	fmt.Println("Results:")
	writeToConsole(result)
}

//ScanRange : scans a range of ports
func ScanRange(from, to int, ip, port, protocol string, dDuration int) {
	if !helpers.IsRangeValid(from, to) {
		fmt.Println("Ranges are not valid! use --help for more information.")
		return
	}
	if !areInputsValid(ip, port, protocol, dDuration) {
		return
	}
	fmt.Println("Scanning ports...")
	results := rangeScan(from, to, ip, port, protocol, dDuration)

	sort.Slice(results, func(i, j int) bool {
		return results[i].Port < results[j].Port
	})

	counter := 0
	for _, result := range results {
		if result.State == "Open" {
			counter++
			writeToConsole(result)
		}
	}

	if counter == 0 {
		fmt.Println("None of the ports were open!")
	}
}

func writeToConsole(result Result) {
	fmt.Print("Port:", result.Port)
	fmt.Print("   State:", result.State)
	fmt.Print("   Protocol:", result.Protocol+"\n")
}

func rangeScan(from, to int, ip, port, protocol string, dDuration int) []Result {
	var results []Result
	var waitgroup sync.WaitGroup
	for i := from; i <= to; i++ {
		waitgroup.Add(1)
		go func(counter int) {
			results = append(results, isPortOpen(protocol, ip, counter, dDuration))
			waitgroup.Done()
		}(i)
	}
	waitgroup.Wait()
	return results
}

func isPortOpen(protocol, hostName string, port, dDuration int) Result {
	var result Result
	address := hostName + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, time.Duration(dDuration)*time.Second)
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

func areInputsValid(ip, port, protocol string, duration int) bool {
	if !helpers.IsIPValid(ip) {
		fmt.Println("Ip is not valid! use --help for more information.")
		return false
	}

	if !helpers.IsPortValid(helpers.StringToInt(port)) {
		fmt.Println("Port number is not valid! use --help for more information.")
		return false
	}
	if !helpers.IsProtocolValid(protocol) {
		fmt.Println("Protocol number is not valid! use --help for more information.")
		return false
	}
	if !helpers.IsDurationValid(duration) {
		fmt.Println("Duration number is not valid! use --help for more information.")
		return false
	}
	return true
}
