package main

import (
	"bufio"
	"fmt"
	"math/big"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/samuelsih/chat-des-rsa/des"
	"github.com/samuelsih/chat-des-rsa/rsa"
)

type Application struct {
	PublicKey      *big.Int
	PrivateKey     *big.Int
	N              *big.Int
	OtherPublicKey *big.Int
	OtherN         *big.Int
	wg             sync.WaitGroup
	conn           net.Conn
}

func New(p, q *big.Int) Application {
	publicKey, privateKey, n := rsa.Generate(p, q)
	return Application{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
		N:          n,
	}
}

func (a *Application) Start() {
	a.wg.Add(1)

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	a.conn = conn
	defer a.conn.Close()

	go a.Prompt()
	go a.Recv()

	a.wg.Wait()
}

func (a *Application) Prompt() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if ok := scanner.Scan(); !ok {
			if err := scanner.Err(); err != nil {
				fmt.Println("Error scanner:", err)
				return
			}
		}

		msg := fmt.Sprintf("%s", scanner.Text())
		encryptedWithPubKey := rsa.Encrypt(msg, a.OtherPublicKey, a.OtherN)
		encryptedTxt := des.Encrypt(encryptedWithPubKey, des.EncryptionBase64)

		_, err := a.conn.Write([]byte(encryptedTxt + "\n"))
		if err != nil {
			fmt.Println("Error conn.Write:", err)
			return
		}
	}
}

func (a *Application) Recv() {
	defer a.wg.Done()
	reader := bufio.NewReader(a.conn)

	for {
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}

		switch {
		case strings.Contains(response, "No client in server"):
			fmt.Println("No second client.")

		case strings.Contains(response, "PING"):
			continue

		case strings.Contains(response, "START"):
			sendOurRSA(a.conn, a.PublicKey.String(), a.N.String())

		case strings.Contains(response, "PUBKEY"):
			cmd := strings.Split(response, " ")
			other := decodeRSA(cmd[1])

			a.OtherPublicKey, _ = new(big.Int).SetString(other.PublicKey, 10)
			a.OtherN, _ = new(big.Int).SetString(other.N, 10)

		default:
			decryptedTxt := des.Decrypt(response, des.DecryptionBase64)
			sanitizedDecryptedTxt := a.sanitize(decryptedTxt)
			result := rsa.Decrypt(sanitizedDecryptedTxt, a.PrivateKey, a.N)
			fmt.Println("MESSAGE:", result)
		}
	}
}

func (a *Application) sanitize(s string) string {
	trimSpaced := strings.TrimSpace(s)
	trimRight := strings.TrimRight(trimSpaced, "\r\n")
	return strings.TrimRight(trimRight, "\n")
}

func main() {
	p, q := generateInitialPandQ()
	app := New(p, q)
	app.Start()
}
