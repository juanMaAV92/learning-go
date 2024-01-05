package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

// go run main.go --site scanme.nmap.org
var site = flag.String("site", "scanme.nmap.org", "site to scan")

func main() {
	flag.Parse()
	var startTime = time.Now()
	var wg sync.WaitGroup
	for port := 1; port <= 65535; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", port)
		}(port)
	}
	wg.Wait()
	fmt.Printf("Done in %d seconds", time.Duration(time.Since(startTime).Seconds()))
}
