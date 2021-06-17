package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup
var hostRe = regexp.MustCompile(`^(25[0-5]|2[0-4]\d|[0-1]?\d?\d)(\.(25[0-5]|2[0-4]\d|[0-1]?\d?\d)){3}$`)

func main() {
	host, startPort, endPort := extractInput()

	for i := startPort; i <= endPort; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()

			address := fmt.Sprintf("%s:%d", host, port)
			conn, err := net.DialTimeout("tcp", address, 10*time.Second)

			if err != nil {
				return
			}

			conn.Close()
			fmt.Printf("%s opened\n", address)
		}(i)
	}

	wg.Wait()
}

func extractInput() (string, int, int) {
	if len(os.Args) < 3 {
		panic("must input host startPort endPort\n")
	}

	host := os.Args[1]
	if !hostRe.Match([]byte(host)) {
		panic(fmt.Sprintf("host %v must be ipv4 address\n", os.Args[1]))
	}

	startPort, err := strconv.Atoi(os.Args[2])
	if err != nil || startPort < 0 {
		panic(fmt.Sprintf("startPort %v must be positive number\n", os.Args[2]))
	}

	endPort, err := strconv.Atoi(os.Args[3])
	if err != nil || endPort < 0 {
		panic(fmt.Sprintf("startPort %v must be positive number\n", os.Args[3]))
	}

	if endPort <= startPort {
		panic(fmt.Sprintf("%d must greater than %d\n", endPort, startPort))
	}

	return host, startPort, endPort
}
