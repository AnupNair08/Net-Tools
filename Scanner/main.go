package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/go-ping/ping"
)

var hosts []string

func callPing(IP string) bool {
	var flag int
	flag = 0
	pinger, err := ping.NewPinger(IP)
	pinger.SetPrivileged(true)
	if err != nil {
		fmt.Println(err)
	}
	pinger.Count = 1
	pinger.OnRecv = func(pkt *ping.Packet) {
		flag = 1
	}
	var d time.Duration = 1500000000
	pinger.Timeout = d
	err = pinger.Run()
	if err != nil {
		fmt.Println("Ping failed")
	}
	if flag == 0 {
		return false
	}
	hosts = append(hosts, IP)
	return true
}

//TODO
// Fetch IP, subnet automatically and run loops dynmaically
// Add multithreading to get over performance problem
func main() {
	var wg sync.WaitGroup
	// op, err := exec.Command("ipconfig").Output()
	fmt.Println("Scanning.....")
	for i := 200; i < 210; i++ {
		callPing("192.168.1." + strconv.Itoa(i))
	}

	wg.Wait()
	fmt.Printf("Hosts up: %d\n", len(hosts))
	fmt.Println(hosts)
}
