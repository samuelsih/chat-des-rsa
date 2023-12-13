package rsa

import (
	"fmt"
	"math/big"
	"strings"
)

func Encrypt(ciphertext string, publicKey *big.Int, n *big.Int) string {
	var sb strings.Builder

	for _, text := range ciphertext {
		num := big.NewInt(int64(text))

		//  C = M^e mod |n|
		encrypted := new(big.Int).Exp(num, publicKey, n)

		toWrite := fmt.Sprintf("%010d", encrypted)
		sb.WriteString(toWrite)
	}

	return sb.String()
}

func Decrypt(ciphertext string, privateKey *big.Int, n *big.Int) string {
	var sb strings.Builder

	chunks := split(ciphertext, 10)
	for _, chunk := range chunks {
		encryptedNumber, _ := new(big.Int).SetString(chunk, 10)

		// M = C^d mod |n|
		decrypted := new(big.Int).Exp(encryptedNumber, privateKey, n)
		ascii := rune(decrypted.Int64())
		sb.WriteString(string(ascii))
	}

	return sb.String()
}

func split(s string, size int) []string {
	var chunks []string

	for i := 0; i < len(s); i += size {
		end := i + size

		if end > len(s) {
			end = len(s)
		}

		chunks = append(chunks, s[i:end])
	}

	return chunks
}
