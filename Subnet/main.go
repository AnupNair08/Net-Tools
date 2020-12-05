package main

import (
	"fmt"
	"strconv"
	"strings"
)

func populate() map[int][]int {
	sheet := make(map[int][]int)
	subnet := 32
	for i := 1; i <= 128; i *= 2 {
		sheet[subnet] = append(sheet[subnet], i)
		sheet[subnet] = append(sheet[subnet], 256-i)
		subnet--
	}
	return sheet
}

func parse(query string) {

	datasheet := populate()
	querySplit := strings.Split(query, "/")
	netmask := querySplit[1]
	ip := querySplit[0]
	ipadd := strings.Split(ip, ".")[3]
	netmaskInt, err := strconv.Atoi(netmask)
	ipaddInt, err2 := strconv.Atoi(ipadd)

	if err != nil || err2 != nil {
		fmt.Println("Error")
		return
	}
	maskArr := datasheet[netmaskInt]
	nextNetworkID := 0
	for i := 0; i < ipaddInt; i += maskArr[0] {
		nextNetworkID += maskArr[0]
	}
	nsp := querySplit[0][:len(querySplit[0])-3]

	fmt.Printf("Next Network ID : %s.%d\n", nsp, nextNetworkID)
	fmt.Printf("Network ID : %s.%d\n", nsp, nextNetworkID-maskArr[0])
	fmt.Printf("Broadcast IP : %s.%d\n", nsp, nextNetworkID-1)
	fmt.Printf("First IP : %s.%d\n", nsp, nextNetworkID-maskArr[0]+1)
	fmt.Printf("Last IP : %s.%d\n", nsp, nextNetworkID-1)
	fmt.Printf("Range of IP address : %d\n", maskArr[0])
	fmt.Printf("Range of usable IP address : %d\n", maskArr[0]-2)
	fmt.Printf("CIDR/Subnet : %d\n", netmaskInt)
	return
}

func main() {
	// A simple subnet calculator for calculating attributes of the networks having a netmask in the /24 or below range
	var query string
	fmt.Println("Enter a network id in the form a.b.c.d/n")
	fmt.Scan(&query)
	parse(query)
	return
}
