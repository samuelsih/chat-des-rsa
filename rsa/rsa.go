package rsa

import (
	"math/big"
)

var zero = big.NewInt(0)
var one = big.NewInt(1)

func Generate(p *big.Int, q *big.Int) (*big.Int, *big.Int, *big.Int) {
	n := new(big.Int).Mul(p, q)
	phi := new(big.Int).Mul(
		new(big.Int).Sub(p, one),
		new(big.Int).Sub(q, one),
	)

	e := big.NewInt(2)

	for leftSideIsLessThanRightSide(e, phi) {
		if equal(gcd(e, phi), one) {
			break
		}

		e.Add(e, one) //e++
	}

	// d x e === 1 (mod ϕ(n))
	// bentuknya sama kaya bentuk modular invers --> ax === 1 (mod m) atau x === a^-1 (mod m)
	d := new(big.Int).ModInverse(e, phi)

	return e, d, n
}

func gcd(a, b *big.Int) *big.Int {
	for !equal(b, zero) {
		a, b = b, new(big.Int).Mod(a, b)
	}

	return a
}

func equal(a, b *big.Int) bool {
	return a.Cmp(b) == 0
}

func leftSideIsLessThanRightSide(a, b *big.Int) bool {
	return a.Cmp(b) == -1
}
