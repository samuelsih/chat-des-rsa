package rsa

import (
	"math/big"
	"testing"
)

func TestGenerate(t *testing.T) {
	p := big.NewInt(53)
	q := big.NewInt(59)

	public, private, n := Generate(p, q)

	expectedPublic := big.NewInt(3)
	expectedPrivate := big.NewInt(2011)
	expectedN := big.NewInt(3127)

	if public.Cmp(expectedPublic) != 0 {
		t.Fatalf("Public: Expected %s, got %s", expectedPublic.String(), public.String())
	}

	if private.Cmp(expectedPrivate) != 0 {
		t.Fatalf("Private: Expected %s, got %s", expectedPrivate.String(), private.String())
	}

	if n.Cmp(expectedN) != 0 {
		t.Fatalf("N: Expected %s, got %s", expectedN.String(), n.String())
	}
}
