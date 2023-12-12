package main

import (
	"bufio"
	"fmt"
	"net"
	"slices"
	"time"
)

func singleSend(conn net.Conn, msg string) error {
	writer := bufio.NewWriter(conn)
	_, err := writer.WriteString(msg)
	if err != nil {
		return err
	}

	if err := writer.Flush(); err != nil {
		return err
	}

	return nil
}

func detectDisconnect(conn net.Conn) {
	disconnect := make(chan struct{})
	ticker := time.NewTicker(time.Second * 2)

	defer close(disconnect)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ping(conn, disconnect)
		case <-disconnect:
			fmt.Println(conn.RemoteAddr().String(), "disconnect from server")

			index := slices.IndexFunc(clients, func(c net.Conn) bool {
				return c.RemoteAddr().String() == conn.RemoteAddr().String()
			})

			clients = slices.Delete(clients, index, len(clients)-1)
			return
		}
	}
}

func signalForSendPubKey() {
	full := make(chan struct{})
	ticker := time.NewTicker(time.Second * 2)

	defer close(full)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			clientFullSignal(full)
		case <-full:
			for _, client := range clients {
				singleSend(client, "START"+"\n")
			}

			return
		}
	}
}

func ping(conn net.Conn, disconnect chan<- struct{}) {
	go func() {
		n, err := conn.Write([]byte("PING\n"))
		if err != nil || n == 0 {
			disconnect <- struct{}{}
			conn.Close()
		}
	}()
}

func clientFullSignal(full chan<- struct{}) {
	go func() {
		if len(clients) == 2 {
			full <- struct{}{}
		}
	}()
}
