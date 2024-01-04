package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

type Client chan<- string

var (
	incomingClients = make(chan Client)
	leavingClients  = make(chan Client)
	messages        = make(chan string)
)

var (
	host = flag.String("host", "localhost", "host where the server will run")
	port = flag.Int("port", 3090, "port where the server will listen")
)

func HandlerConnection(conn net.Conn) {
	defer conn.Close()

	messageClientChannel := make(chan string)
	go ClientWriter(conn, messageClientChannel)

	who := conn.RemoteAddr().String()
	messageClientChannel <- "Welcome to the server, You are " + who + "\n"
	messages <- who + " has arrived \n"
	incomingClients <- messageClientChannel

	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		messages <- who + " : " + inputMessage.Text() + "\n"
	}

	leavingClients <- messageClientChannel
	messages <- who + " has left\n"
}

func ClientWriter(conn net.Conn, messageClientChannel <-chan string) {
	for msg := range messageClientChannel {
		fmt.Fprintln(conn, msg)
	}
}

func Broadcaster() {
	clients := make(map[Client]bool)
	for {
		select {
		case msg := <-messages:
			for client := range clients {
				client <- msg
			}
		case newClient := <-incomingClients:
			clients[newClient] = true
		case leavingClient := <-leavingClients:
			delete(clients, leavingClient)
			close(leavingClient)
		}
	}
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}

	go Broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}
		go HandlerConnection(conn)
	}

}
