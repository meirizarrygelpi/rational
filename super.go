// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"fmt"
	"math/big"
	"strings"
)

var symbSuper = [4]string{"", "ε₁", "ε₂", "ε₃"}

// A Super represents a rational super dual number.
type Super struct {
	re, du *Dual
}

// Re returns the real part of z, a pointer to a Dual value.
func (z *Super) Re() *Dual {
	return z.re
}

// Du returns the dual part of z, a pointer to a Dual value.
func (z *Super) Du() *Dual {
	return z.du
}

// SetRe sets the real part of z equal to a.
func (z *Super) SetRe(a *Dual) {
	z.re = a
}

// SetDu sets the dual part of z equal to b.
func (z *Super) SetDu(b *Dual) {
	z.du = b
}

// Cartesian returns the four Cartesian components of z.
func (z *Super) Cartesian() (a, b, c, d *big.Rat) {
	a, b = z.Re().Cartesian()
	c, d = z.Du().Cartesian()
	return
}

// String returns the string representation of a Super value.
//
// If z corresponds to a + bε₁ + cε₂ + dε₃, then the string is "(a+bε₁+cε₂+dε₃)",
// similar to complex128 values.
func (z *Super) String() string {
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
		a[j+1] = symbSuper[i]
		i++
	}
	a[8] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Super) Equals(y *Super) bool {
	if !z.Re().Equals(y.Re()) || !z.Du().Equals(y.Du()) {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Super) Copy(y *Super) *Super {
	z.SetRe(y.Re())
	z.SetDu(y.Du())
	return z
}

// NewSuper returns a pointer to a Super value made from four given
// pointers to Dual values.
func NewSuper(a, b, c, d *big.Rat) *Super {
	z := new(Super)
	z.SetRe(NewDual(a, b))
	z.SetDu(NewDual(c, d))
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Super) Scal(y *Super, a *big.Rat) *Super {
	z.SetRe(new(Dual).Scal(y.Re(), a))
	z.SetDu(new(Dual).Scal(y.Du(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Super) Neg(y *Super) *Super {
	z.SetRe(new(Dual).Neg(y.Re()))
	z.SetDu(new(Dual).Neg(y.Du()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Super) Conj(y *Super) *Super {
	z.SetRe(new(Dual).Conj(y.Re()))
	z.SetDu(new(Dual).Neg(y.Du()))
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Super) Add(x, y *Super) *Super {
	z.SetRe(new(Dual).Add(x.Re(), y.Re()))
	z.SetDu(new(Dual).Add(x.Du(), y.Du()))
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Super) Sub(x, y *Super) *Super {
	z.SetRe(new(Dual).Sub(x.Re(), y.Re()))
	z.SetDu(new(Dual).Sub(x.Du(), y.Du()))
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
func (z *Super) Mul(x, y *Super) *Super {
	p := new(Super).Copy(x)
	q := new(Super).Copy(y)
	z.SetRe(new(Dual).Mul(p.Re(), q.Re()))
	z.SetDu(new(Dual).Add(
		new(Dual).Mul(p.Re(), q.Du()),
		new(Dual).Mul(p.Du(), new(Dual).Conj(q.Re())),
	))
	return z
}

// Commutator sets z equal to the commutator of x and y, and returns z.
func (z *Super) Commutator(x, y *Super) *Super {
	return z.Sub(
		new(Super).Mul(x, y),
		new(Super).Mul(y, x),
	)
}

// Quad returns the quadrance of z, a pointer to a big.Rat value.
func (z *Super) Quad() *big.Rat {
	return new(big.Rat).Add(
		z.Re().Quad(),
		z.Du().Quad(),
	)
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Super) Inv(y *Super) *Super {
	return z.Scal(new(Super).Conj(y), new(big.Rat).Inv(y.Quad()))
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Super) Quo(x, y *Super) *Super {
	return z.Mul(x, new(Super).Inv(y))
}
