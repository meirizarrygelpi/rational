// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"fmt"
	"math/big"
	"strings"
)

// A Perplex represents a rational split-complex number.
type Perplex struct {
	re, sp *big.Rat
}

// Re returns the real part of z, a pointer to a big.Rat value.
func (z *Perplex) Re() *big.Rat {
	return z.re
}

// Sp returns the split part of z, a pointer to a big.Rat value.
func (z *Perplex) Sp() *big.Rat {
	return z.sp
}

// SetRe sets the real part of z equal to a.
func (z *Perplex) SetRe(a *big.Rat) {
	z.re = a
}

// SetSp sets the split part of z equal to b.
func (z *Perplex) SetSp(b *big.Rat) {
	z.sp = b
}

// Cartesian returns the two Cartesian components of z.
func (z *Perplex) Cartesian() (a, b *big.Rat) {
	a, b = z.Re(), z.Sp()
	return
}

// String returns the string version of a Perplex value. If z = a + bs, then
// the string is "(a+bs)", similar to complex128 values.
func (z *Perplex) String() string {
	a := make([]string, 5)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", z.Re())
	if z.Sp().Sign() == -1 {
		a[2] = fmt.Sprintf("%v", z.Sp())
	} else {
		a[2] = fmt.Sprintf("+%v", z.Sp())
	}
	a[3] = "s"
	a[4] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Perplex) Equals(y *Perplex) bool {
	if z.Re().Cmp(y.Re()) != 0 || z.Sp().Cmp(y.Sp()) != 0 {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Perplex) Copy(y *Perplex) *Perplex {
	z.SetRe(y.Re())
	z.SetSp(y.Sp())
	return z
}

// NewPerplex returns a pointer to a Perplex value made from four given int64
// values.
func NewPerplex(a, b, c, d int64) *Perplex {
	z := new(Perplex)
	z.SetRe(big.NewRat(a, b))
	z.SetSp(big.NewRat(c, d))
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Perplex) Scal(y *Perplex, a *big.Rat) *Perplex {
	z.SetRe(new(big.Rat).Mul(y.Re(), a))
	z.SetSp(new(big.Rat).Mul(y.Sp(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Perplex) Neg(y *Perplex) *Perplex {
	z.SetRe(new(big.Rat).Neg(y.Re()))
	z.SetSp(new(big.Rat).Neg(y.Sp()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Perplex) Conj(y *Perplex) *Perplex {
	z.SetRe(y.Re())
	z.SetSp(new(big.Rat).Neg(y.Sp()))
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Perplex) Add(x, y *Perplex) *Perplex {
	z.SetRe(new(big.Rat).Add(x.Re(), y.Re()))
	z.SetSp(new(big.Rat).Add(x.Sp(), y.Sp()))
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Perplex) Sub(x, y *Perplex) *Perplex {
	z.SetRe(new(big.Rat).Sub(x.Re(), y.Re()))
	z.SetSp(new(big.Rat).Sub(x.Sp(), y.Sp()))
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rule is:
// 		Mul(s, s) = +1
// This binary operation is commutative and associative.
func (z *Perplex) Mul(x, y *Perplex) *Perplex {
	p := new(Perplex).Copy(x)
	q := new(Perplex).Copy(y)
	z.SetRe(new(big.Rat).Add(
		new(big.Rat).Mul(p.Re(), q.Re()),
		new(big.Rat).Mul(p.Sp(), q.Sp()),
	))
	z.SetSp(new(big.Rat).Add(
		new(big.Rat).Mul(p.Re(), q.Sp()),
		new(big.Rat).Mul(p.Sp(), q.Re()),
	))
	return z
}

// Quad returns the quadrance of z, a pointer to a big.Rat value.
func (z *Perplex) Quad() *big.Rat {
	return new(big.Rat).Sub(
		new(big.Rat).Mul(z.Re(), z.Re()),
		new(big.Rat).Mul(z.Sp(), z.Sp()),
	)
}

// IsZeroDiv returns true if z is a zero divisor.
func (z *Perplex) IsZeroDiv() bool {
	if z.Re().Cmp(z.Sp()) == 0 {
		return true
	}
	if z.Re().Cmp(new(big.Rat).Neg(z.Sp())) == 0 {
		return true
	}
	return false
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Perplex) Inv(y *Perplex) *Perplex {
	if y.IsZeroDiv() {
		panic("zero divisor inverse")
	}
	return z.Scal(z.Conj(y), new(big.Rat).Inv(y.Quad()))
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Perplex) Quo(x, y *Perplex) *Perplex {
	if y.IsZeroDiv() {
		panic("zero divisor denominator")
	}
	return z.Mul(x, z.Inv(y))
}

// Idempotent sets z equal to a pointer to an idempotent Perplex.
func (z *Perplex) Idempotent(sign int) *Perplex {
	z.SetRe(big.NewRat(1, 2))
	if sign < 0 {
		z.SetSp(big.NewRat(-1, 2))
		return z
	}
	z.SetSp(big.NewRat(1, 2))
	return z
}
