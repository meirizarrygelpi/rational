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

// Cartesian returns the four rational Cartesian components of z.
func (z *InfraPerplex) Cartesian() (*big.Rat, *big.Rat, *big.Rat, *big.Rat) {
	return &z.l.l, &z.l.r, &z.r.l, &z.r.r
}

// String returns the string representation of an InfraPerplex value.
//
// If z corresponds to a + bs + cτ + dυ, then the string is"(a+bs+cτ+dυ)",
// similar to complex128 values.
func (z *InfraPerplex) String() string {
	v := make([]*big.Rat, 4)
	v[0], v[1] = z.l.Cartesian()
	v[2], v[3] = z.r.Cartesian()
	a := make([]string, 9)
	a[0] = "("
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
	a[8] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *InfraPerplex) Equals(y *InfraPerplex) bool {
	if !z.l.Equals(&y.l) || !z.r.Equals(&y.r) {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *InfraPerplex) Copy(y *InfraPerplex) *InfraPerplex {
	z.l.Copy(&y.l)
	z.r.Copy(&y.r)
	return z
}

// NewInfraPerplex returns a pointer to an InfraPerplex value made from four
// given pointers to big.Rat values.
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

// Add sets z equal to the sum of x and y, and returns z.
func (z *InfraPerplex) Add(x, y *InfraPerplex) *InfraPerplex {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
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
	a := new(Perplex).Copy(&x.l)
	b := new(Perplex).Copy(&x.r)
	c := new(Perplex).Copy(&y.l)
	d := new(Perplex).Copy(&y.r)
	temp := new(Perplex)
	z.l.Mul(a, c)
	z.r.Add(
		z.r.Mul(d, a),
		temp.Mul(b, temp.Conj(c)),
	)
	return z
}

// Commutator sets z equal to the commutator of x and y, and returns z.
func (z *InfraPerplex) Commutator(x, y *InfraPerplex) *InfraPerplex {
	return z.Sub(
		z.Mul(x, y),
		new(InfraPerplex).Mul(y, x),
	)
}

// Quad returns the quadrance of z. If z = a+bs+cτ+dυ, then the quadrance is
//		Mul(a, a) - Mul(b, b)
// This can be positive, negative, or zero.
func (z *InfraPerplex) Quad() *big.Rat {
	return z.l.Quad()
}

// IsZeroDiv returns true if z is a zero divisor. This is equivalent to z being
// nilpotent.
func (z *InfraPerplex) IsZeroDiv() bool {
	return z.l.IsZeroDiv()
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *InfraPerplex) Inv(y *InfraPerplex) *InfraPerplex {
	if y.IsZeroDiv() {
		panic("inverse of zero divisor")
	}
	return z.Scal(z.Conj(y), new(big.Rat).Inv(y.Quad()))
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *InfraPerplex) Quo(x, y *InfraPerplex) *InfraPerplex {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Mul(x, z.Inv(y))
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
