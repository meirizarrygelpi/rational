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

var symbInfraPerplex = [4]string{"", "s", "τ", "υ"}

// An InfraPerplex represents a rational infra-perplex number.
type InfraPerplex struct {
	l, r Perplex
}

// Real returns the (rational) real part of z.
func (z *InfraPerplex) Real() *big.Rat {
	return (&z.l).Real()
}

// Rats returns the four rational components of z.
func (z *InfraPerplex) Rats() (*big.Rat, *big.Rat, *big.Rat, *big.Rat) {
	return &z.l.l, &z.l.r, &z.r.l, &z.r.r
}

// String returns the string representation of an InfraPerplex value.
//
// If z corresponds to a + bs + cτ + dυ, then the string is"(a+bs+cτ+dυ)",
// similar to complex128 values.
func (z *InfraPerplex) String() string {
	v := make([]*big.Rat, 4)
	v[0], v[1] = z.l.Rats()
	v[2], v[3] = z.r.Rats()
	a := make([]string, 9)
	a[0] = leftBracket
	a[1] = fmt.Sprintf("%v", v[0].RatString())
	i := 1
	for j := 2; j < 8; j = j + 2 {
		if v[i].Sign() < 0 {
			a[j] = fmt.Sprintf("%v", v[i].RatString())
		} else {
			a[j] = fmt.Sprintf("+%v", v[i].RatString())
		}
		a[j+1] = symbInfraPerplex[i]
		i++
	}
	a[8] = rightBracket
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *InfraPerplex) Equals(y *InfraPerplex) bool {
	if !z.l.Equals(&y.l) || !z.r.Equals(&y.r) {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *InfraPerplex) Set(y *InfraPerplex) *InfraPerplex {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewInfraPerplex returns a pointer to the InfraPerplex value a+bs+cτ+dυ.
func NewInfraPerplex(a, b, c, d *big.Rat) *InfraPerplex {
	z := new(InfraPerplex)
	z.l.l.Set(a)
	z.l.r.Set(b)
	z.r.l.Set(c)
	z.r.r.Set(d)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *InfraPerplex) Scal(y *InfraPerplex, a *big.Rat) *InfraPerplex {
	z.l.Scal(&y.l, a)
	z.r.Scal(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *InfraPerplex) Neg(y *InfraPerplex) *InfraPerplex {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *InfraPerplex) Conj(y *InfraPerplex) *InfraPerplex {
	z.l.Conj(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *InfraPerplex) Add(x, y *InfraPerplex) *InfraPerplex {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *InfraPerplex) Sub(x, y *InfraPerplex) *InfraPerplex {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
// 		Mul(s, s) = +1
// 		Mul(τ, τ) = Mul(υ, υ) = 0
// 		Mul(τ, υ) = Mul(υ, τ) = 0
// 		Mul(s, τ) = -Mul(τ, s) = υ
// 		Mul(s, υ) = -Mul(υ, s) = τ
// This binary operation is noncommutative but associative.
func (z *InfraPerplex) Mul(x, y *InfraPerplex) *InfraPerplex {
	a := new(Perplex).Set(&x.l)
	b := new(Perplex).Set(&x.r)
	c := new(Perplex).Set(&y.l)
	d := new(Perplex).Set(&y.r)
	temp := new(Perplex)
	z.l.Mul(a, c)
	z.r.Add(
		z.r.Mul(d, a),
		temp.Mul(b, temp.Conj(c)),
	)
	return z
}

// Commutator sets z equal to the commutator of x and y:
// 		Mul(x, y) - Mul(y, x)
// Then it returns z.
func (z *InfraPerplex) Commutator(x, y *InfraPerplex) *InfraPerplex {
	return z.Sub(
		z.Mul(x, y),
		new(InfraPerplex).Mul(y, x),
	)
}

// Quad returns the quadrance of z. If z = a+bs+cτ+dυ, then the quadrance is
//		a² - b²
// This can be positive, negative, or zero.
func (z *InfraPerplex) Quad() *big.Rat {
	return z.l.Quad()
}

// IsZeroDivisor returns true if z is a zero divisor. This is equivalent to z
// being nilpotent.
func (z *InfraPerplex) IsZeroDivisor() bool {
	return z.l.IsZeroDivisor()
}

// Inv sets z equal to the inverse of y, and returns z. If y is a zero divisor,
// then Inv panics.
func (z *InfraPerplex) Inv(y *InfraPerplex) *InfraPerplex {
	if y.IsZeroDivisor() {
		panic("inverse of zero divisor")
	}
	a := y.Quad()
	a.Inv(a)
	return z.Scal(z.Conj(y), a)
}

// QuoL sets z equal to the left quotient of x and y:
// 		Mul(Inv(y), x)
// Then it returns z. If y is a zero divisor, then QuoL panics.
func (z *InfraPerplex) QuoL(x, y *InfraPerplex) *InfraPerplex {
	if y.IsZeroDivisor() {
		panic("denominator is zero divisor")
	}
	return z.Mul(z.Inv(y), x)
}

// QuoR sets z equal to the right quotient of x and y:
// 		Mul(x, Inv(y))
// Then it returns z. If y is a zero divisor, then QuoR panics.
func (z *InfraPerplex) QuoR(x, y *InfraPerplex) *InfraPerplex {
	if y.IsZeroDivisor() {
		panic("denominator is zero divisor")
	}
	return z.Mul(x, z.Inv(y))
}

// CrossRatioL sets z equal to the left cross-ratio of v, w, x, and y:
// 		Inv(w - x) * (v - x) * Inv(v - y) * (w - y)
// Then it returns z.
func (z *InfraPerplex) CrossRatioL(v, w, x, y *InfraPerplex) *InfraPerplex {
	temp := new(InfraPerplex)
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

// CrossRatioR sets z equal to the right cross-ratio of v, w, x, and y:
// 		(v - x) * Inv(w - x) * (w - y) * Inv(v - y)
// Then it returns z.
func (z *InfraPerplex) CrossRatioR(v, w, x, y *InfraPerplex) *InfraPerplex {
	temp := new(InfraPerplex)
	z.Sub(v, x)
	temp.Sub(w, x)
	temp.Inv(temp)
	z.Mul(z, temp)
	temp.Sub(w, y)
	z.Mul(z, temp)
	temp.Sub(v, y)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// MöbiusL sets z equal to the left Möbius (fractional linear) transform of y:
// 		Inv(y*c + d) * (y*a + b)
// Then it returns z.
func (z *InfraPerplex) MöbiusL(y, a, b, c, d *InfraPerplex) *InfraPerplex {
	z.Mul(y, a)
	z.Add(z, b)
	temp := new(InfraPerplex)
	temp.Mul(y, c)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(temp, z)
}

// MöbiusR sets z equal to the right Möbius (fractional linear) transform of y:
// 		(a*y + b) * Inv(c*y + d)
// Then it returns z.
func (z *InfraPerplex) MöbiusR(y, a, b, c, d *InfraPerplex) *InfraPerplex {
	z.Mul(a, y)
	z.Add(z, b)
	temp := new(InfraPerplex)
	temp.Mul(c, y)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// Dot returns the (rational) dot product of z and y.
func (z *InfraPerplex) Dot(y *InfraPerplex) *big.Rat {
	return z.l.Dot(&y.l)
}

// Generate returns a random InfraPerplex value for quick.Check testing.
func (z *InfraPerplex) Generate(rand *rand.Rand, size int) reflect.Value {
	randomInfraPerplex := &InfraPerplex{
		*NewPerplex(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
		*NewPerplex(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
	}
	return reflect.ValueOf(randomInfraPerplex)
}
