package main

import (
	"fmt"
	"portscanner/lib"
	"runtime"
)

var (
	ipList   []string
	portList []int
	runNum   int
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	runNum = 1000
	for s := 1; s <= 255; s++ {
		ipList = append(ipList, fmt.Sprintf("192.168.10.%d", s))
	}
	for s := 1; s < 10000; s++ {
		portList = append(portList, s)
	}
}

func main() {
	lib.IPPortsScanner(ipList, portList)
}
