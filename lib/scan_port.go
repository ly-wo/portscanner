package lib

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func IPPortsScanner(ipList []string, portList []int) {
	all := len(ipList) * len(portList)
	wg = sync.WaitGroup{}
	wg.Add(all)

	for _, ip := range ipList {
		for _, port := range portList {
			go func(ip string, port int) {
				defer func() {
					wg.Done()
				}()

				ipPort := fmt.Sprintf("%s:%d", ip, port)
				if conn, ok := net.DialTimeout("tcp", ipPort, time.Second*3); ok == nil {
					conn.Close()
					fmt.Println(ipPort, " 结果:", ok == nil)
				}

			}(ip, port)
		}
	}

	wg.Wait()
}
