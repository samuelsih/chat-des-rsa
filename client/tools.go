package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"net"
	"os"
)

type RSAExchange struct {
	PublicKey string `json:"public_key"`
	N         string `json:"n"`
}

func generateInitialPandQ() (*big.Int, *big.Int) {
	num1, err := rand.Prime(rand.Reader, 8)
	if err != nil {
		fmt.Println("Error generating num1:", err)
		os.Exit(1)
	}

	num2, err := rand.Prime(rand.Reader, 8)
	if err != nil {
		fmt.Println("Error generating num2:", err)
		os.Exit(1)
	}

	if num1.Cmp(num2) == -1 {
		return num1, num2
	}

	return num2, num1
}

func sendOurRSA(conn net.Conn, publicKey string, n string) {
	rsa := RSAExchange{PublicKey: publicKey, N: n}
	msg, err := json.Marshal(&rsa)
	if err != nil {
		fmt.Println("Error exchange:", err)
		os.Exit(1)
	}

	_, err = conn.Write([]byte("PUBKEY " + string(msg) + "\n"))
	if err != nil {
		fmt.Println("Error conn.Write:", err)
		os.Exit(1)
	}
}

func decodeRSAPayload(payload string) RSAExchange {
	var result RSAExchange
	json.Unmarshal([]byte(payload), &result)
	return result
}
