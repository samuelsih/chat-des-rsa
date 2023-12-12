package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/samuelsih/chat-des-rsa/des"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer conn.Close()

	go prompt(conn)
	go recv(conn)

	wg.Wait()
}

func prompt(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if ok := scanner.Scan(); !ok {
			if err := scanner.Err(); err != nil {
				fmt.Println("Error scanner:", err)
				return
			}
		}

		msg := fmt.Sprintf("%s", scanner.Text())
		encryptedTxt := des.Encrypt(msg, des.EncryptionBase64)

		_, err := conn.Write([]byte(encryptedTxt + "\n"))
		if err != nil {
			fmt.Println("Error conn.Write:", err)
			return
		}
	}
}

func recv(listener net.Conn) {
	defer wg.Done()
	reader := bufio.NewReader(listener)

	for {
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}

		if(strings.Contains(response, "No client in server")) {
			fmt.Println("No second client.")
			continue
		}

		decryptedTxt := des.Decrypt(response, des.DecryptionBase64)

		fmt.Println(sanitize(decryptedTxt))
	}
}


func sanitize(s string) string {
	trimSpaced := strings.TrimSpace(s)
	trimRight := strings.TrimRight(trimSpaced, "\r\n")
	return strings.TrimRight(trimRight, "\n")
}
