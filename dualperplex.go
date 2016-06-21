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

var symbDualPerplex = [4]string{"", "s", "κ", "λ"}

// A DualPerplex represents a rational dual perplex number.
type DualPerplex struct {
	l, r Perplex
}

// Real returns the (rational) real part of z.
func (z *DualPerplex) Real() *big.Rat {
	return (&z.l).Real()
}

// Rats returns the four rational components of z.
func (z *DualPerplex) Rats() (*big.Rat, *big.Rat, *big.Rat, *big.Rat) {
	return &z.l.l, &z.l.r, &z.r.l, &z.r.r
}

// String returns the string representation of a DualPerplex value.
//
// If z corresponds to a + bs + cκ + dλ, then the string is "(a+bs+cκ+dλ)",
// similar to complex128 values.
func (z *DualPerplex) String() string {
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
		a[j+1] = symbDualPerplex[i]
		i++
	}
	a[8] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *DualPerplex) Equals(y *DualPerplex) bool {
	if !z.l.Equals(&y.l) || !z.r.Equals(&y.r) {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *DualPerplex) Set(y *DualPerplex) *DualPerplex {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewDualPerplex returns a *DualPerplex with value a+bs+cκ+dλ.
func NewDualPerplex(a, b, c, d *big.Rat) *DualPerplex {
	z := new(DualPerplex)
	z.l.l.Set(a)
	z.l.r.Set(b)
	z.r.l.Set(c)
	z.r.r.Set(d)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *DualPerplex) Scal(y *DualPerplex, a *big.Rat) *DualPerplex {
	z.l.Scal(&y.l, a)
	z.r.Scal(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *DualPerplex) Neg(y *DualPerplex) *DualPerplex {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the dual-perplex conjugate of y, and returns z.
func (z *DualPerplex) Conj(y *DualPerplex) *DualPerplex {
	z.l.Set(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Star sets z equal to the star conjugate of y, and returns z.
func (z *DualPerplex) Star(y *DualPerplex) *DualPerplex {
	z.l.Conj(&y.l)
	z.r.Conj(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *DualPerplex) Add(x, y *DualPerplex) *DualPerplex {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *DualPerplex) Sub(x, y *DualPerplex) *DualPerplex {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
// 		Mul(s, s) = +1
// 		Mul(κ, κ) = Mul(λ, λ) = 0
// 		Mul(s, κ) = Mul(κ, s) = λ
// 		Mul(κ, λ) = Mul(λ, κ) = 0
// 		Mul(λ, s) = Mul(s, λ) = κ
// This binary operation is commutative and associative.
func (z *DualPerplex) Mul(x, y *DualPerplex) *DualPerplex {
	a := new(Perplex).Set(&x.l)
	b := new(Perplex).Set(&x.r)
	c := new(Perplex).Set(&y.l)
	d := new(Perplex).Set(&y.r)
	temp := new(Perplex)
	z.l.Mul(a, c)
	z.r.Add(
		z.r.Mul(a, d),
		temp.Mul(b, c),
	)
	return z
}

// Norm returns the perplex norm of z.
func (z *DualPerplex) Norm() *Perplex {
	norm := new(Perplex)
	return norm.Mul(&z.l, &z.l)
}

// Quad returns the quadrance of z. This is always non-negative.
func (z *DualPerplex) Quad() *big.Rat {
	return z.Norm().Quad()
}

// IsZeroDiv returns true if z is a zero divisor.
func (z *DualPerplex) IsZeroDiv() bool {
	return z.Norm().IsZeroDiv()
}

// Inv sets z equal to the inverse of y, and returns z. If y is a zero divisor,
// then Inv panics.
func (z *DualPerplex) Inv(y *DualPerplex) *DualPerplex {
	if y.IsZeroDiv() {
		panic("inverse of zero divisor")
	}
	p := new(DualPerplex)
	p.Set(y)
	quad := p.Quad()
	quad.Inv(quad)
	temp := new(DualPerplex)
	z.Conj(p)
	z.Mul(z, temp.Star(p))
	z.Mul(z, temp.Conj(temp.Star(p)))
	return z.Scal(z, quad)
}

// Quo sets z equal to the quotient of x and y. If y is a zero divisor, then
// Quo panics.
func (z *DualPerplex) Quo(x, y *DualPerplex) *DualPerplex {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Mul(z.Inv(y), x)
}

// CrossRatio sets z equal to the cross-ratio of v, w, x, and y:
// 		Inv(w - x) * (v - x) * Inv(v - y) * (w - y)
// Then it returns z.
func (z *DualPerplex) CrossRatio(v, w, x, y *DualPerplex) *DualPerplex {
	temp := new(DualPerplex)
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
func (z *DualPerplex) Möbius(y, a, b, c, d *DualPerplex) *DualPerplex {
	z.Mul(a, y)
	z.Add(z, b)
	temp := new(DualPerplex)
	temp.Mul(c, y)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// Generate returns a random DualPerplex value for quick.Check testing.
func (z *DualPerplex) Generate(rand *rand.Rand, size int) reflect.Value {
	randomDualPerplex := &DualPerplex{
		*NewPerplex(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
		*NewPerplex(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
	}
	return reflect.ValueOf(randomDualPerplex)
}