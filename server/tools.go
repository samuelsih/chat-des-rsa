package main

import (
	"bufio"
	"net"
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
