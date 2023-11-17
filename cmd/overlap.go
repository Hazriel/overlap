package main

import (
	"fmt"
	"os"

	"github.com/Hazriel/overlap/internal"
)

func main() {
	if len(os.Args) != 3 {
		printUsage()
		os.Exit(1)
	}

	ip1, err1 := internal.ParseString(os.Args[1])
	ip2, err2 := internal.ParseString(os.Args[2])

	if err1 != nil || err2 != nil {
		printUsage()
		os.Exit(2)
	}

	result := "different"

	if ip1.IsOnSameNetworkAs(ip2) {
		result = "same"
	}

	if ip1.IsSubnetOf(ip2) {
		result = "superset"
	}

	if ip2.IsSubnetOf(ip1) {
		result = "subset"
	}

	fmt.Println(result)
}

func printUsage() {
	fmt.Println("Usage: ./overlap <ip> <ip>")
	fmt.Println("<ip> format must be in CIDR format. E.g: 192.168.1.1/24")
}
