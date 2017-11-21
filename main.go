package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("ip is empty")
		return
	}
	ip := os.Args[1]
	parseIP := net.ParseIP(ip)
	if parseIP == nil {
		fmt.Println("ip err")
		return
	}
	ipv4 := parseIP.To4()
	fmt.Println(binary.BigEndian.Uint32(ipv4))
}
