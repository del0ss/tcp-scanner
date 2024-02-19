package main

import (
	"fmt"
	"net"
	"sync"
)

func TCPScanner(network, address string, numsOfPort int) {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 50) 

	for i := 0; i <= numsOfPort; i++ {
		semaphore <- struct{}{} 
		wg.Add(1)
		go func(j int) {
			defer func() {
				<-semaphore 
				wg.Done()
			}()

			conn, err := net.DialTimeout(network, fmt.Sprintf("%s:%d", address, j))
			if err != nil {
				fmt.Println("[-] Port Closed", j)
				return
			}
			defer conn.Close()
			fmt.Println("[+] Port open", j)
		}(i)
	}

	wg.Wait()
}

func main() {
	TCPScanner("tcp", "127.0.0.1", 100)
}
