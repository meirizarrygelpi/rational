// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"fmt"
	"math/big"
	"strings"
)

var symbInfraPerplex = [4]string{"", "s", "σ", "τ"}

// An InfraPerplex represents a rational infra-perplex number.
type InfraPerplex struct {
	l, r *Perplex
}

// L returns the left Cayley-Dickson part of z, a pointer to a Perplex value.
func (z *InfraPerplex) L() *Perplex {
	return z.l
}

// R returns the right Cayley-Dickson part of z, a pointer to a Perplex value.
func (z *InfraPerplex) R() *Perplex {
	return z.r
}

// SetL sets the left Cayley-Dickson part of z equal to a.
func (z *InfraPerplex) SetL(a *Perplex) {
	z.l = a
}

// SetR sets the right Cayley-Dickson part of z equal to b.
func (z *InfraPerplex) SetR(b *Perplex) {
	z.r = b
}

// Cartesian returns the four Cartesian components of z.
func (z *InfraPerplex) Cartesian() (a, b, c, d *big.Rat) {
	a, b = z.L().Cartesian()
	c, d = z.R().Cartesian()
	return
}

// String returns the string representation of an InfraPerplex value.
//
// If z corresponds to a + bs + cσ + dτ, then the string is"(a+bs+cσ+dτ)",
// similar to complex128 values.
func (z *InfraPerplex) String() string {
	v := make([]*big.Rat, 4)
	v[0], v[1] = z.L().Cartesian()
	v[2], v[3] = z.R().Cartesian()
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
	if !z.L().Equals(y.L()) || !z.R().Equals(y.R()) {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *InfraPerplex) Copy(y *InfraPerplex) *InfraPerplex {
	z.SetL(y.L())
	z.SetR(y.R())
	return z
}

// NewInfraPerplex returns a pointer to an InfraPerplex value made from four
// given pointers to big.Rat values.
func NewInfraPerplex(a, b, c, d *big.Rat) *InfraPerplex {
	z := new(InfraPerplex)
	z.SetL(NewPerplex(a, b))
	z.SetR(NewPerplex(c, d))
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *InfraPerplex) Scal(y *InfraPerplex, a *big.Rat) *InfraPerplex {
	z.SetL(new(Perplex).Scal(y.L(), a))
	z.SetR(new(Perplex).Scal(y.R(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *InfraPerplex) Neg(y *InfraPerplex) *InfraPerplex {
	z.SetL(new(Perplex).Neg(y.L()))
	z.SetR(new(Perplex).Neg(y.R()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *InfraPerplex) Conj(y *InfraPerplex) *InfraPerplex {
	z.SetL(new(Perplex).Conj(y.L()))
	z.SetR(new(Perplex).Neg(y.R()))
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *InfraPerplex) Add(x, y *InfraPerplex) *InfraPerplex {
	z.SetL(new(Perplex).Add(x.L(), y.L()))
	z.SetR(new(Perplex).Add(x.R(), y.R()))
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *InfraPerplex) Sub(x, y *InfraPerplex) *InfraPerplex {
	z.SetL(new(Perplex).Sub(x.L(), y.L()))
	z.SetR(new(Perplex).Sub(x.R(), y.R()))
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
// 		Mul(s, s) = +1
// 		Mul(σ, σ) = Mul(τ, τ) = 0
// 		Mul(σ, τ) = Mul(τ, σ) = 0
// 		Mul(s, σ) = -Mul(σ, s) = τ
// 		Mul(s, τ) = -Mul(τ, s) = σ
// This binary operation is noncommutative but associative.
func (z *InfraPerplex) Mul(x, y *InfraPerplex) *InfraPerplex {
	a, b := x.L(), x.R()
	c, d := y.L(), y.R()
	s, t, u := new(Perplex), new(Perplex), new(Perplex)
	z.SetL(
		s.Mul(a, c),
	)
	z.SetR(t.Add(
		t.Mul(d, a),
		u.Mul(b, u.Conj(c)),
	))
	return z
}

// Commutator sets z equal to the commutator of x and y, and returns z.
func (z *InfraPerplex) Commutator(x, y *InfraPerplex) *InfraPerplex {
	return z.Sub(
		z.Mul(x, y),
		new(InfraPerplex).Mul(y, x),
	)
}

// Quad returns the quadrance of z, a pointer to a big.Rat value.
func (z *InfraPerplex) Quad() *big.Rat {
	return z.L().Quad()
}

// IsZeroDiv returns true if z is a zero divisor. This is equivalent to z being
// nilpotent.
func (z *InfraPerplex) IsZeroDiv() bool {
	return z.L().IsZeroDiv()
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
