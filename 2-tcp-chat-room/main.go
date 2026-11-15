package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	clientsChan := make(chan net.Conn)
	listener, err := net.Listen("tcp", ":8090")

	if err != nil {
		log.Fatal("Error listening: ", err)
	}

	defer listener.Close()

	go func() {
		for {
			conn, err := listener.Accept()

			if err != nil {
				log.Println("Error accepting conn:", err)
			}
			clientsChan <- conn
		}
	}()

	clients := []net.Conn{}
	for {
		client, ok := <-clientsChan
		clients = append(clients, client)
		if ok {
			fmt.Println(client)
			broadcast(clients, client, "New Connection")
			go func() {
				for {
					reader := bufio.NewReader(client)
					clientMessage, err := reader.ReadString('\n')
					if err != nil {
						log.Printf("Read error: %v", err)
						return
					}
					broadcast(clients, client, clientMessage)
				}
			}()
		}
	}
}

func broadcast(clients []net.Conn, from net.Conn, message string) {
	messageText := fmt.Sprintln(from, ":\t", message)
	for _, cl := range clients {
		if cl != from {
			cl.Write([]byte(messageText))
		}
	}
}
