// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"fmt"
	"math/big"
	"math/rand"
	"reflect"
	"strings"
)

var symbDualComplex = [4]string{"", "i", "Γ", "iΓ"}

// A DualComplex represents a rational dual-complex number.
type DualComplex struct {
	l, r Complex
}

// Real returns the (rational) real part of z.
func (z *DualComplex) Real() *big.Rat {
	return (&z.l).Real()
}

// Rats returns the four rational components of z.
func (z *DualComplex) Rats() (*big.Rat, *big.Rat, *big.Rat, *big.Rat) {
	return &z.l.l, &z.l.r, &z.r.l, &z.r.r
}

// String returns the string representation of a DualComplex value.
func (z *DualComplex) String() string {
	v := make([]*big.Rat, 4)
	v[0], v[1] = z.l.Rats()
	v[2], v[3] = z.r.Rats()
	a := make([]string, 9)
	a[0] = leftBracket
	a[1] = fmt.Sprintf("%v", v[0].RatString())
	i := 1
	for j := 2; j < 8; j = j + 2 {
		if v[i].Sign() == -1 {
			a[j] = fmt.Sprintf("%v", v[i].RatString())
		} else {
			a[j] = fmt.Sprintf("+%v", v[i].RatString())
		}
		a[j+1] = symbDualComplex[i]
		i++
	}
	a[8] = rightBracket
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *DualComplex) Equals(y *DualComplex) bool {
	if !z.l.Equals(&y.l) || !z.r.Equals(&y.r) {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *DualComplex) Set(y *DualComplex) *DualComplex {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewDualComplex returns a *DualComplex with value a+bi+cΓ+diΓ.
func NewDualComplex(a, b, c, d *big.Rat) *DualComplex {
	z := new(DualComplex)
	z.l.l.Set(a)
	z.l.r.Set(b)
	z.r.l.Set(c)
	z.r.r.Set(d)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *DualComplex) Scal(y *DualComplex, a *big.Rat) *DualComplex {
	z.l.Scal(&y.l, a)
	z.r.Scal(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *DualComplex) Neg(y *DualComplex) *DualComplex {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *DualComplex) Conj(y *DualComplex) *DualComplex {
	z.l.Set(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Star sets z equal to the star conjugate of y, and returns z.
func (z *DualComplex) Star(y *DualComplex) *DualComplex {
	z.l.Conj(&y.l)
	z.r.Conj(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *DualComplex) Add(x, y *DualComplex) *DualComplex {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *DualComplex) Sub(x, y *DualComplex) *DualComplex {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
// 		Mul(i, i) = -1
// 		Mul(Γ, Γ) = 0
// 		Mul(i, Γ) = Mul(Γ, i)
// This binary operation is commutative and associative.
func (z *DualComplex) Mul(x, y *DualComplex) *DualComplex {
	a := new(Complex).Set(&x.l)
	b := new(Complex).Set(&x.r)
	c := new(Complex).Set(&y.l)
	d := new(Complex).Set(&y.r)
	temp := new(Complex)
	z.l.Mul(a, c)
	z.r.Add(
		z.r.Mul(a, d),
		temp.Mul(b, c),
	)
	return z
}

// Quad returns the quadrance of z. If z = a+bi+cΓ+diΓ, then the quadrance is
// 		a² - b² + 2abi
// Note that this is a complex number.
func (z *DualComplex) Quad() *Complex {
	quad := new(Complex)
	return quad.Mul(&z.l, &z.l)
}

// Norm returns the norm of z. If z = a+bi+cΓ+diΓ, then the norm is
// 		(a² + b²)²
// This is always non-negative.
func (z *DualComplex) Norm() *big.Rat {
	return z.Quad().Quad()
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *DualComplex) IsZeroDivisor() bool {
	zero := new(Complex)
	return zero.Equals(z.Quad())
}

// Inv sets z equal to the inverse of y, and returns z. If y is a zero divisor,
// then Inv panics.
func (z *DualComplex) Inv(y *DualComplex) *DualComplex {
	if y.IsZeroDivisor() {
		panic("inverse of zero divisor")
	}
	quad := y.Quad()
	quad.Inv(quad)
	z.Conj(y)
	z.l.Mul(&z.l, quad)
	z.r.Mul(&z.r, quad)
	return z
}

// Quo sets z equal to the quotient of x and y. If y is a zero divisor, then
// Quo panics.
func (z *DualComplex) Quo(x, y *DualComplex) *DualComplex {
	if y.IsZeroDivisor() {
		panic("denominator is zero divisor")
	}
	return z.Mul(z.Inv(y), x)
}

// CrossRatio sets z equal to the cross-ratio of v, w, x, and y:
// 		Inv(w - x) * (v - x) * Inv(v - y) * (w - y)
// Then it returns z.
func (z *DualComplex) CrossRatio(v, w, x, y *DualComplex) *DualComplex {
	temp := new(DualComplex)
	z.Sub(w, x)
	z.Inv(z)
	temp.Sub(v, x)
	z.Mul(z, temp)
	temp.Sub(v, y)
	temp.Inv(temp)
	z.Mul(z, temp)
	temp.Sub(w, y)
	return z.Mul(z, temp)
}

// Möbius sets z equal to the Möbius (fractional linear) transform of y:
// 		(a*y + b) * Inv(c*y + d)
// Then it returns z.
func (z *DualComplex) Möbius(y, a, b, c, d *DualComplex) *DualComplex {
	z.Mul(a, y)
	z.Add(z, b)
	temp := new(DualComplex)
	temp.Mul(c, y)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// Generate returns a random DualComplex value for quick.Check testing.
func (z *DualComplex) Generate(rand *rand.Rand, size int) reflect.Value {
	randomDualComplex := &DualComplex{
		*NewComplex(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
		*NewComplex(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
	}
	return reflect.ValueOf(randomDualComplex)
}
