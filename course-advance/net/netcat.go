package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// go run main.go --port 3090 --host localhost
var (
	port = flag.Int("port", 3090, "port to scan")
	host = flag.String("host", "localhost", "url to scan")
)

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatalf("Error dialing %s:%d: %s", *host, *port, err)
	}

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()

	CopyContents(conn, os.Stdin)
	conn.Close()
	<-done
}

func CopyContents(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatalf("Error copying contents: %s", err)
	}
}
