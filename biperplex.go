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

var symbBiPerplex = [4]string{"", "s", "r", "q"}

// A BiPerplex represents a rational biperplex number.
type BiPerplex struct {
	l, r Perplex
}

// Real returns the (rational) real part of z.
func (z *BiPerplex) Real() *big.Rat {
	return (&z.l).Real()
}

// Rats returns the four rational components of z.
func (z *BiPerplex) Rats() (*big.Rat, *big.Rat, *big.Rat, *big.Rat) {
	return &z.l.l, &z.l.r, &z.r.l, &z.r.r
}

// String returns the string representation of a BiPerplex value.
//
// If z corresponds to a + bs + cr + dq, then the string is "(a+bs+cr+dq)",
// similar to complex128 values.
func (z *BiPerplex) String() string {
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
		a[j+1] = symbBiPerplex[i]
		i++
	}
	a[8] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *BiPerplex) Equals(y *BiPerplex) bool {
	if !z.l.Equals(&y.l) || !z.r.Equals(&y.r) {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *BiPerplex) Set(y *BiPerplex) *BiPerplex {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewBiPerplex returns a *BiPerplex with value a+bs+cr+dq.
func NewBiPerplex(a, b, c, d *big.Rat) *BiPerplex {
	z := new(BiPerplex)
	z.l.l.Set(a)
	z.l.r.Set(b)
	z.r.l.Set(c)
	z.r.r.Set(d)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *BiPerplex) Scal(y *BiPerplex, a *big.Rat) *BiPerplex {
	z.l.Scal(&y.l, a)
	z.r.Scal(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *BiPerplex) Neg(y *BiPerplex) *BiPerplex {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the biperplex conjugate of y, and returns z.
func (z *BiPerplex) Conj(y *BiPerplex) *BiPerplex {
	z.l.Set(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Star sets z equal to the star conjugate of y, and returns z.
func (z *BiPerplex) Star(y *BiPerplex) *BiPerplex {
	z.l.Conj(&y.l)
	z.r.Conj(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *BiPerplex) Add(x, y *BiPerplex) *BiPerplex {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *BiPerplex) Sub(x, y *BiPerplex) *BiPerplex {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
// 		Mul(s, s) = Mul(r, r) = Mul(q, q) = +1
// 		Mul(s, r) = Mul(r, s) = q
// 		Mul(r, q) = Mul(q, r) = s
// 		Mul(q, s) = Mul(s, q) = r
// This binary operation is commutative and associative.
func (z *BiPerplex) Mul(x, y *BiPerplex) *BiPerplex {
	a := new(Perplex).Set(&x.l)
	b := new(Perplex).Set(&x.r)
	c := new(Perplex).Set(&y.l)
	d := new(Perplex).Set(&y.r)
	temp := new(Perplex)
	z.l.Add(
		z.l.Mul(a, c),
		temp.Mul(b, d),
	)
	z.r.Add(
		z.r.Mul(a, d),
		temp.Mul(b, c),
	)
	return z
}

// Norm returns the perplex norm of z.
func (z *BiPerplex) Norm() *Perplex {
	norm := new(Perplex)
	norm.Mul(&z.l, &z.l)
	norm.Sub(norm, new(Perplex).Mul(&z.r, &z.r))
	return norm
}

// Quad returns the quadrance of z. This can be positive, negative, or zero.
func (z *BiPerplex) Quad() *big.Rat {
	return z.Norm().Quad()
}

// IsZeroDiv returns true if z is a zero divisor.
func (z *BiPerplex) IsZeroDiv() bool {
	return z.Norm().IsZeroDiv()
}

// Inv sets z equal to the inverse of y, and returns z. If y is a zero divisor,
// then Inv panics.
func (z *BiPerplex) Inv(y *BiPerplex) *BiPerplex {
	if y.IsZeroDiv() {
		panic("inverse of zero divisor")
	}
	p := new(BiPerplex)
	p.Set(y)
	quad := p.Quad()
	quad.Inv(quad)
	temp := new(BiPerplex)
	z.Conj(p)
	z.Mul(z, temp.Star(p))
	z.Mul(z, temp.Conj(temp.Star(p)))
	return z.Scal(z, quad)
}

// Quo sets z equal to the quotient of x and y. If y is a zero divisor, then
// Quo panics.
func (z *BiPerplex) Quo(x, y *BiPerplex) *BiPerplex {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Mul(z.Inv(y), x)
}

// CrossRatio sets z equal to the cross-ratio of v, w, x, and y:
// 		Inv(w - x) * (v - x) * Inv(v - y) * (w - y)
// Then it returns z.
func (z *BiPerplex) CrossRatio(v, w, x, y *BiPerplex) *BiPerplex {
	temp := new(BiPerplex)
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
func (z *BiPerplex) Möbius(y, a, b, c, d *BiPerplex) *BiPerplex {
	z.Mul(a, y)
	z.Add(z, b)
	temp := new(BiPerplex)
	temp.Mul(c, y)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// Generate returns a random BiPerplex value for quick.Check testing.
func (z *BiPerplex) Generate(rand *rand.Rand, size int) reflect.Value {
	randomBiPerplex := &BiPerplex{
		*NewPerplex(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
		*NewPerplex(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
	}
	return reflect.ValueOf(randomBiPerplex)
}