package main

import (
	"fmt"
	"net"
	"sync"
)

func TCPScanner(network, address string, numsOfPort int) {
	var wg sync.WaitGroup

	for i := 0; i <= numsOfPort; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			conn, err := net.Dial(network, fmt.Sprintf(address+":%d", j))
			if err != nil {
				fmt.Println("Port is closed", j)
				return
			}
			fmt.Println("Connection successful on port ", j)
			conn.Close()

		}(i)
	}
	wg.Wait()
}
