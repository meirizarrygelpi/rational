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

var symbBiComplex = [4]string{"", "i", "h", "s"}

// A BiComplex represents a rational bicomplex number.
type BiComplex struct {
	l, r Complex
}

// Real returns the (rational) real part of z.
func (z *BiComplex) Real() *big.Rat {
	return (&z.l).Real()
}

// Rats returns the four rational components of z.
func (z *BiComplex) Rats() (*big.Rat, *big.Rat, *big.Rat, *big.Rat) {
	return &z.l.l, &z.l.r, &z.r.l, &z.r.r
}

// String returns the string representation of a BiComplex value.
//
// If z corresponds to a + bi + ch + ds, then the string is "(a+bi+ch+ds)",
// similar to complex128 values.
func (z *BiComplex) String() string {
	v := make([]*big.Rat, 4)
	v[0], v[1] = z.l.Rats()
	v[2], v[3] = z.r.Rats()
	a := make([]string, 9)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", v[0].RatString())
	i := 1
	for j := 2; j < 8; j = j + 2 {
		if v[i].Sign() == -1 {
			a[j] = fmt.Sprintf("%v", v[i].RatString())
		} else {
			a[j] = fmt.Sprintf("+%v", v[i].RatString())
		}
		a[j+1] = symbBiComplex[i]
		i++
	}
	a[8] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *BiComplex) Equals(y *BiComplex) bool {
	if !z.l.Equals(&y.l) || !z.r.Equals(&y.r) {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *BiComplex) Set(y *BiComplex) *BiComplex {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewBiComplex returns a *BiComplex with value a+bh+ch+ds.
func NewBiComplex(a, b, c, d *big.Rat) *BiComplex {
	z := new(BiComplex)
	z.l.l.Set(a)
	z.l.r.Set(b)
	z.r.l.Set(c)
	z.r.r.Set(d)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *BiComplex) Scal(y *BiComplex, a *big.Rat) *BiComplex {
	z.l.Scal(&y.l, a)
	z.r.Scal(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *BiComplex) Neg(y *BiComplex) *BiComplex {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the bicomplex conjugate of y, and returns z.
func (z *BiComplex) Conj(y *BiComplex) *BiComplex {
	z.l.Set(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Star sets z equal to the star conjugate of y, and returns z.
func (z *BiComplex) Star(y *BiComplex) *BiComplex {
	z.l.Conj(&y.l)
	z.r.Conj(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *BiComplex) Add(x, y *BiComplex) *BiComplex {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *BiComplex) Sub(x, y *BiComplex) *BiComplex {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
// 		Mul(i, i) = Mul(h, h) = -1
// 		Mul(s, s) = +1
// 		Mul(i, h) = Mul(h, i) = s
// 		Mul(h, s) = Mul(s, h) = -i
// 		Mul(s, i) = Mul(i, s) = -h
// This binary operation is commutative and associative.
func (z *BiComplex) Mul(x, y *BiComplex) *BiComplex {
	a := new(Complex).Set(&x.l)
	b := new(Complex).Set(&x.r)
	c := new(Complex).Set(&y.l)
	d := new(Complex).Set(&y.r)
	temp := new(Complex)
	z.l.Sub(
		z.l.Mul(a, c),
		temp.Mul(b, d),
	)
	z.r.Add(
		z.r.Mul(a, d),
		temp.Mul(b, c),
	)
	return z
}

// Norm returns the complex norm of z.
func (z *BiComplex) Norm() *Complex {
	norm := new(Complex)
	norm.Mul(&z.l, &z.l)
	norm.Add(norm, new(Complex).Mul(&z.r, &z.r))
	return norm
}

// Quad returns the quadrance of z. This is always non-negative.
func (z *BiComplex) Quad() *big.Rat {
	return z.Norm().Quad()
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *BiComplex) IsZeroDivisor() bool {
	zero := new(Complex)
	return zero.Equals(z.Norm())
}

// Inv sets z equal to the inverse of y, and returns z. If y is a zero divisor,
// then Inv panics.
func (z *BiComplex) Inv(y *BiComplex) *BiComplex {
	if y.IsZeroDivisor() {
		panic("inverse of zero divisor")
	}
	p := new(BiComplex)
	p.Set(y)
	quad := p.Quad()
	quad.Inv(quad)
	temp := new(BiComplex)
	z.Conj(p)
	z.Mul(z, temp.Star(p))
	z.Mul(z, temp.Conj(temp.Star(p)))
	return z.Scal(z, quad)
}

// Quo sets z equal to the quotient of x and y. If y is a zero divisor, then
// Quo panics.
func (z *BiComplex) Quo(x, y *BiComplex) *BiComplex {
	if y.IsZeroDivisor() {
		panic("denominator is zero divisor")
	}
	return z.Mul(z.Inv(y), x)
}

// CrossRatio sets z equal to the cross-ratio of v, w, x, and y:
// 		Inv(w - x) * (v - x) * Inv(v - y) * (w - y)
// Then it returns z.
func (z *BiComplex) CrossRatio(v, w, x, y *BiComplex) *BiComplex {
	temp := new(BiComplex)
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
func (z *BiComplex) Möbius(y, a, b, c, d *BiComplex) *BiComplex {
	z.Mul(a, y)
	z.Add(z, b)
	temp := new(BiComplex)
	temp.Mul(c, y)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// Generate returns a random BiComplex value for quick.Check testing.
func (z *BiComplex) Generate(rand *rand.Rand, size int) reflect.Value {
	randomBiComplex := &BiComplex{
		*NewComplex(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
		*NewComplex(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
	}
	return reflect.ValueOf(randomBiComplex)
}
