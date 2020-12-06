package main

import (
	"fmt"
	"math"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/go-ping/ping"
)

var hosts []string

type LocalIP struct {
	ipaddr  string
	netmask float64
}

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

func getHosts(subnet net.IPMask) float64 {
	a, b, c, d := float64(subnet[3]), float64(subnet[2]), float64(subnet[1]), float64(subnet[0])
	k := math.Logb(256-a) + math.Logb(256-b) + math.Logb(256-c) + math.Logb(256-d)
	return math.Pow(2.0, k)
}

func getLocalIP() LocalIP {
	var res LocalIP
	query, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println(err)
		return res
	} else {
		addr := query.LocalAddr().(*net.UDPAddr)
		// fmt.Println(net.IPMask(addr.IP))
		ip := strconv.Itoa(int(addr.IP[0])) + "." + strconv.Itoa(int(addr.IP[1])) + "." + strconv.Itoa(int(addr.IP[2]))
		res.ipaddr = ip
		res.netmask = getHosts(addr.IP.DefaultMask())
		return res
	}

}

//TODO
// Add multithreading to get over performance problem
func main() {
	var wg sync.WaitGroup
	// op, err := exec.Command("ipconfig").Output()
	local := getLocalIP()
	selfIP := local.ipaddr
	hostNo := int(local.netmask)
	fmt.Println(hostNo, selfIP)
	fmt.Println("Scanning.....")
	k := 256 - hostNo
	for i := 0; i < 10; i++ {
		callPing(selfIP + "." + strconv.Itoa(k))
		k++
	}

	wg.Wait()
	fmt.Printf("Hosts up: %d\n", len(hosts))
	fmt.Println(hosts)
}
