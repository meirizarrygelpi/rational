// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"fmt"
	"math/big"
	"strings"
)

// A Dual represents a rational dual number.
type Dual struct {
	re, du *big.Rat
}

// Re returns the real part of z, a pointer to a big.Rat value.
func (z *Dual) Re() *big.Rat {
	return z.re
}

// Du returns the dual part of z, a pointer to a big.Rat value.
func (z *Dual) Du() *big.Rat {
	return z.du
}

// SetRe sets the real part of z equal to a.
func (z *Dual) SetRe(a *big.Rat) {
	z.re = a
}

// SetDu sets the dual part of z equal to b.
func (z *Dual) SetDu(b *big.Rat) {
	z.du = b
}

// Cartesian returns the two Cartesian components of z.
func (z *Dual) Cartesian() (a, b *big.Rat) {
	a = z.Re()
	b = z.Du()
	return
}

// String returns the string version of a Dual value.
//
// If z corresponds to a + bε, then the string is "(a+bε)", similar to
// complex128 values.
func (z *Dual) String() string {
	a := make([]string, 5)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", z.Re())
	if z.Du().Sign() == -1 {
		a[2] = fmt.Sprintf("%v", z.Du())
	} else {
		a[2] = fmt.Sprintf("+%v", z.Du())
	}
	a[3] = "ε"
	a[4] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Dual) Equals(y *Dual) bool {
	if z.Re().Cmp(y.Re()) != 0 || z.Du().Cmp(y.Du()) != 0 {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Dual) Copy(y *Dual) *Dual {
	z.SetRe(y.Re())
	z.SetDu(y.Du())
	return z
}

// NewDual returns a pointer to a Dual value made from two given pointers to
// big.Rat values.
func NewDual(a, b *big.Rat) *Dual {
	z := new(Dual)
	z.SetRe(a)
	z.SetDu(b)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Dual) Scal(y *Dual, a *big.Rat) *Dual {
	z.SetRe(new(big.Rat).Mul(y.Re(), a))
	z.SetDu(new(big.Rat).Mul(y.Du(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Dual) Neg(y *Dual) *Dual {
	z.SetRe(new(big.Rat).Neg(y.Re()))
	z.SetDu(new(big.Rat).Neg(y.Du()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Dual) Conj(y *Dual) *Dual {
	z.SetRe(y.Re())
	z.SetDu(new(big.Rat).Neg(y.Du()))
	return z
}

// Add sets z to the sum of x and y, and returns z.
func (z *Dual) Add(x, y *Dual) *Dual {
	z.SetRe(new(big.Rat).Add(x.Re(), y.Re()))
	z.SetDu(new(big.Rat).Add(x.Du(), y.Du()))
	return z
}

// Sub sets z to the difference of x and y, and returns z.
func (z *Dual) Sub(x, y *Dual) *Dual {
	z.SetRe(new(big.Rat).Sub(x.Re(), y.Re()))
	z.SetDu(new(big.Rat).Sub(x.Du(), y.Du()))
	return z
}

// Mul sets z to the product of x and y, and returns z.
//
// The multiplication rule is:
// 		Mul(ε, ε) = 0
// This binary operation is commutative and associative.
func (z *Dual) Mul(x, y *Dual) *Dual {
	p := new(Dual).Copy(x)
	q := new(Dual).Copy(y)
	z.SetRe(
		new(big.Rat).Mul(p.Re(), q.Re()),
	)
	z.SetDu(new(big.Rat).Add(
		new(big.Rat).Mul(q.Du(), p.Re()),
		new(big.Rat).Mul(p.Du(), q.Re()),
	))
	return z
}

// Quad returns the quadrance of z, a pointer to a big.Rat value.
func (z *Dual) Quad() *big.Rat {
	return new(big.Rat).Mul(z.Re(), z.Re())
}

// IsZeroDiv returns true if z is a zero divisor. This is equivalent to z being
// nilpotent.
func (z *Dual) IsZeroDiv() bool {
	return z.Re().Num().Cmp(big.NewInt(0)) == 0
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Dual) Inv(y *Dual) *Dual {
	if y.IsZeroDiv() {
		panic("inverse of zero divisor")
	}
	return z.Scal(new(Dual).Conj(y), new(big.Rat).Inv(y.Quad()))
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Dual) Quo(x, y *Dual) *Dual {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Mul(x, new(Dual).Inv(y))
}
