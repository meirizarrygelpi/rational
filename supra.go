// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"fmt"
	"math/big"
	"strings"
)

var symbSupra = [4]string{"", "ε₁", "ε₂", "ε₃"}

// A Supra represents a rational supra number.
type Supra struct {
	re, du *Infra
}

// Re returns the real part of z, a pointer to a Infra value.
func (z *Supra) Re() *Infra {
	return z.re
}

// Du returns the dual part of z, a pointer to a Infra value.
func (z *Supra) Du() *Infra {
	return z.du
}

// SetRe sets the real part of z equal to a.
func (z *Supra) SetRe(a *Infra) {
	z.re = a
}

// SetDu sets the dual part of z equal to b.
func (z *Supra) SetDu(b *Infra) {
	z.du = b
}

// Cartesian returns the four Cartesian components of z.
func (z *Supra) Cartesian() (a, b, c, d *big.Rat) {
	a, b = z.Re().Cartesian()
	c, d = z.Du().Cartesian()
	return
}

// String returns the string representation of a Supra value.
//
// If z corresponds to a + bε₁ + cε₂ + dε₃, then the string is "(a+bε₁+cε₂+dε₃)",
// similar to complex128 values.
func (z *Supra) String() string {
	v := make([]*big.Rat, 4)
	v[0], v[1] = z.Re().Cartesian()
	v[2], v[3] = z.Du().Cartesian()
	a := make([]string, 9)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", v[0])
	i := 1
	for j := 2; j < 8; j = j + 2 {
		if v[i].Sign() == -1 {
			a[j] = fmt.Sprintf("%v", v[i])
		} else {
			a[j] = fmt.Sprintf("+%v", v[i])
		}
		a[j+1] = symbSupra[i]
		i++
	}
	a[8] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Supra) Equals(y *Supra) bool {
	if !z.Re().Equals(y.Re()) || !z.Du().Equals(y.Du()) {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Supra) Copy(y *Supra) *Supra {
	z.SetRe(y.Re())
	z.SetDu(y.Du())
	return z
}

// NewSupra returns a pointer to a Supra value made from four given
// pointers to Infra values.
func NewSupra(a, b, c, d *big.Rat) *Supra {
	z := new(Supra)
	z.SetRe(NewInfra(a, b))
	z.SetDu(NewInfra(c, d))
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Supra) Scal(y *Supra, a *big.Rat) *Supra {
	z.SetRe(new(Infra).Scal(y.Re(), a))
	z.SetDu(new(Infra).Scal(y.Du(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Supra) Neg(y *Supra) *Supra {
	z.SetRe(new(Infra).Neg(y.Re()))
	z.SetDu(new(Infra).Neg(y.Du()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Supra) Conj(y *Supra) *Supra {
	z.SetRe(new(Infra).Conj(y.Re()))
	z.SetDu(new(Infra).Neg(y.Du()))
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Supra) Add(x, y *Supra) *Supra {
	z.SetRe(new(Infra).Add(x.Re(), y.Re()))
	z.SetDu(new(Infra).Add(x.Du(), y.Du()))
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Supra) Sub(x, y *Supra) *Supra {
	z.SetRe(new(Infra).Sub(x.Re(), y.Re()))
	z.SetDu(new(Infra).Sub(x.Du(), y.Du()))
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
// 		Mul(ε₁, ε₁) = Mul(ε₂, ε₂) = Mul(ε₃, ε₃) = 0
// 		Mul(ε₁, ε₂) = -Mul(ε₂, ε₁) = ε₃
// 		Mul(ε₂, ε₃) = Mul(ε₃, ε₂) = 0
// 		Mul(ε₃, ε₁) = Mul(ε₁, ε₃) = 0
// This binary operation is noncommutative but associative.
func (z *Supra) Mul(x, y *Supra) *Supra {
	p := new(Supra).Copy(x)
	q := new(Supra).Copy(y)
	z.SetRe(new(Infra).Mul(p.Re(), q.Re()))
	z.SetDu(new(Infra).Add(
		new(Infra).Mul(q.Du(), p.Re()),
		new(Infra).Mul(p.Du(), new(Infra).Conj(q.Re())),
	))
	return z
}

// Commutator sets z equal to the commutator of x and y, and returns z.
func (z *Supra) Commutator(x, y *Supra) *Supra {
	return z.Sub(
		new(Supra).Mul(x, y),
		new(Supra).Mul(y, x),
	)
}

// Quad returns the quadrance of z, a pointer to a big.Rat value.
func (z *Supra) Quad() *big.Rat {
	return z.Re().Quad()
}

// IsZeroDiv returns true if z is a zero divisor.
func (z *Supra) IsZeroDiv() bool {
	return z.Re().IsZeroDiv()
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Supra) Inv(y *Supra) *Supra {
	return z.Scal(new(Supra).Conj(y), new(big.Rat).Inv(y.Quad()))
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Supra) Quo(x, y *Supra) *Supra {
	return z.Mul(x, new(Supra).Inv(y))
}
