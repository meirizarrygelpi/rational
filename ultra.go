// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"fmt"
	"math/big"
	"strings"
)

var symbUltra = [8]string{"", "ε₁", "ε₂", "ε₃", "ε₄", "ε₅", "ε₆", "ε₇"}

// A Ultra represents a rational ultra number.
type Ultra struct {
	re, du *Supra
}

// Re returns the real part of z, a pointer to a Supra value.
func (z *Ultra) Re() *Supra {
	return z.re
}

// Du returns the dual part of z, a pointer to a Supra value.
func (z *Ultra) Du() *Supra {
	return z.du
}

// SetRe sets the real part of z equal to a.
func (z *Ultra) SetRe(a *Supra) {
	z.re = a
}

// SetDu sets the dual part of z equal to b.
func (z *Ultra) SetDu(b *Supra) {
	z.du = b
}

// Cartesian returns the four Cartesian components of z.
func (z *Ultra) Cartesian() (a, b, c, d, e, f, g, h *big.Rat) {
	a, b, c, d = z.Re().Cartesian()
	e, f, g, h = z.Du().Cartesian()
	return
}

// String returns the string representation of a Ultra value.
//
// If z corresponds to a + bε₁ + cε₂ + dε₃, then the string is "(a+bε₁+cε₂+dε₃)",
// similar to complex128 values.
func (z *Ultra) String() string {
	v := make([]*big.Rat, 8)
	v[0], v[1], v[2], v[3] = z.Re().Cartesian()
	v[4], v[5], v[6], v[7] = z.Du().Cartesian()
	a := make([]string, 17)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", v[0])
	i := 1
	for j := 2; j < 16; j = j + 2 {
		if v[i].Sign() == -1 {
			a[j] = fmt.Sprintf("%v", v[i])
		} else {
			a[j] = fmt.Sprintf("+%v", v[i])
		}
		a[j+1] = symbUltra[i]
		i++
	}
	a[16] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Ultra) Equals(y *Ultra) bool {
	if !z.Re().Equals(y.Re()) || !z.Du().Equals(y.Du()) {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Ultra) Copy(y *Ultra) *Ultra {
	z.SetRe(y.Re())
	z.SetDu(y.Du())
	return z
}

// NewUltra returns a pointer to a Ultra value made from four given
// pointers to Supra values.
func NewUltra(a, b, c, d, e, f, g, h *big.Rat) *Ultra {
	z := new(Ultra)
	z.SetRe(NewSupra(a, b, c, d))
	z.SetDu(NewSupra(e, f, g, h))
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Ultra) Scal(y *Ultra, a *big.Rat) *Ultra {
	z.SetRe(new(Supra).Scal(y.Re(), a))
	z.SetDu(new(Supra).Scal(y.Du(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Ultra) Neg(y *Ultra) *Ultra {
	z.SetRe(new(Supra).Neg(y.Re()))
	z.SetDu(new(Supra).Neg(y.Du()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Ultra) Conj(y *Ultra) *Ultra {
	z.SetRe(new(Supra).Conj(y.Re()))
	z.SetDu(new(Supra).Neg(y.Du()))
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Ultra) Add(x, y *Ultra) *Ultra {
	z.SetRe(new(Supra).Add(x.Re(), y.Re()))
	z.SetDu(new(Supra).Add(x.Du(), y.Du()))
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Ultra) Sub(x, y *Ultra) *Ultra {
	z.SetRe(new(Supra).Sub(x.Re(), y.Re()))
	z.SetDu(new(Supra).Sub(x.Du(), y.Du()))
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
// 		Mul(ε₁, ε₁) = Mul(ε₂, ε₂) = Mul(ε₃, ε₃) = 0
// 		Mul(ε₄, ε₄) = Mul(ε₅, ε₅) = Mul(ε₆, ε₆) = Mul(ε₇, ε₇) = 0
// 		Mul(ε₁, ε₂) = -Mul(ε₂, ε₁) = +ε₃
// 		Mul(ε₁, ε₃) = Mul(ε₃, ε₁) = 0
// 		Mul(ε₁, ε₄) = -Mul(ε₄, ε₁) = +ε₅
// 		Mul(ε₁, ε₅) = Mul(ε₅, ε₁) = 0
// 		Mul(ε₁, ε₆) = -Mul(ε₆, ε₁) = -ε₇
// 		Mul(ε₁, ε₇) = -Mul(ε₇, ε₁) = +ε₆
// 		Mul(ε₂, ε₃) = Mul(ε₃, ε₂) = 0
// 		Mul(ε₂, ε₄) = -Mul(ε₄, ε₂) = +ε₆
// 		Mul(ε₂, ε₅) = -Mul(ε₅, ε₂) = +ε₇
// 		Mul(ε₂, ε₆) = Mul(ε₆, ε₂) = 0
// 		Mul(ε₂, ε₇) = Mul(ε₇, ε₂) = 0
// 		Mul(ε₃, ε₄) = -Mul(ε₄, ε₃) = +ε₇
// 		Mul(ε₃, ε₅) = Mul(ε₅, ε₃) = 0
// 		Mul(ε₃, ε₆) = Mul(ε₆, ε₃) = 0
// 		Mul(ε₃, ε₇) = Mul(ε₇, ε₃) = 0
// 		Mul(ε₄, ε₅) = Mul(ε₅, ε₄) = 0
// 		Mul(ε₄, ε₆) = Mul(ε₆, ε₄) = 0
// 		Mul(ε₄, ε₇) = Mul(ε₇, ε₄) = 0
// 		Mul(ε₅, ε₆) = Mul(ε₆, ε₅) = 0
// 		Mul(ε₅, ε₇) = Mul(ε₇, ε₅) = 0
// 		Mul(ε₆, ε₇) = Mul(ε₇, ε₆) = 0
// This binary operation is noncommutative and nonassociative.
func (z *Ultra) Mul(x, y *Ultra) *Ultra {
	p := new(Ultra).Copy(x)
	q := new(Ultra).Copy(y)
	z.SetRe(new(Supra).Mul(p.Re(), q.Re()))
	z.SetDu(new(Supra).Add(
		new(Supra).Mul(q.Du(), p.Re()),
		new(Supra).Mul(p.Du(), new(Supra).Conj(q.Re())),
	))
	return z
}

// Commutator sets z equal to the commutator of x and y, and returns z.
func (z *Ultra) Commutator(x, y *Ultra) *Ultra {
	return z.Sub(
		new(Ultra).Mul(x, y),
		new(Ultra).Mul(y, x),
	)
}

// Associator sets z equal to the associator of w, x, and y, and returns z.
func (z *Ultra) Associator(w, x, y *Ultra) *Ultra {
	return z.Sub(
		new(Ultra).Mul(new(Ultra).Mul(w, x), y),
		new(Ultra).Mul(w, new(Ultra).Mul(x, y)),
	)
}

// Quad returns the quadrance of z, a pointer to a big.Rat value.
func (z *Ultra) Quad() *big.Rat {
	return z.Re().Quad()
}

// IsZeroDiv returns true if z is a zero divisor.
func (z *Ultra) IsZeroDiv() bool {
	return z.Re().IsZeroDiv()
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Ultra) Inv(y *Ultra) *Ultra {
	return z.Scal(new(Ultra).Conj(y), new(big.Rat).Inv(y.Quad()))
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Ultra) Quo(x, y *Ultra) *Ultra {
	return z.Mul(x, new(Ultra).Inv(y))
}
