package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"slices"
	"strings"
)

var clients []net.Conn

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	go signalForSendPubKey()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		fmt.Println("Connection established", conn.RemoteAddr().String())
		clients = append(clients, conn)
		go detectDisconnect(conn)
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error:", err)
				return
			}
		}

		if len(msg) <= 1 {
			return
		}

		fmt.Println("Message:", sanitize(msg))

		differentClientIndex := slices.IndexFunc(clients, func(c net.Conn) bool {
			return c.RemoteAddr().String() != conn.RemoteAddr().String()
		})

		if differentClientIndex == -1 {
			singleSend(conn, "No client in server\n")
			continue
		}

		connToSend := clients[differentClientIndex]
		singleSend(connToSend, msg)
	}
}

func sanitize(s string) string {
	trimSpaced := strings.TrimSpace(s)
	trimRight := strings.TrimRight(trimSpaced, "\r\n")
	return strings.TrimRight(trimRight, "\n")
}
