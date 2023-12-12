package rsa

import (
	"math/big"
	"testing"
)

func TestRSAFlow(t *testing.T) {
	p := big.NewInt(53)
	q := big.NewInt(59)

	public, private, n := Generate(p, q)

	msg := "Saya hantu"

	encrypted := Encrypt(msg, public, n)
	decrypted := Decrypt(encrypted, private, n)

	if msg != decrypted {
		t.Fatalf("RSAFlow: Expected %s, Got %s", msg, decrypted)
	}
}
