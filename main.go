package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func scanPort(protocol, hostname string, port int, wg *sync.WaitGroup) {
	defer wg.Done()

	address := fmt.Sprintf("%s:%d", hostname, port)
	conn, err := net.DialTimeout(protocol, address, 5*time.Second)

	if err != nil {
		return
	}

	defer conn.Close()

	fmt.Printf("Port %d: Open\n", port)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <hostname> <start_port-end_port>")
		return
	}

	hostname := os.Args[1]
	ports := os.Args[2]

	startPort, err := strconv.Atoi(ports[:strings.IndexByte(ports, '-')])

	if err != nil {
		fmt.Println("Invalid start port number")
		return
	}

	endPort, err := strconv.Atoi(ports[strings.IndexByte(ports, '-')+1:])

	if err != nil {
		fmt.Println("Invalid end port number")
		return
	}

	var wg sync.WaitGroup

	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go scanPort("tcp", hostname, port, &wg)
	}

	wg.Wait()
}
