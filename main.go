package main

const (
	network    = "tcp"
	address    = "127.0.0.1"
	numsOfPort = 65545
)

func main() {
	TCPScanner(network, address, numsOfPort)
}
