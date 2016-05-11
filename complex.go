// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"fmt"
	"math/big"
	"strings"
)

// A Complex represents a rational complex number.
type Complex struct {
	re, im *big.Rat
}

// Re returns the real part of z, a pointer to a big.Rat value.
func (z *Complex) Re() *big.Rat {
	return z.re
}

// Im returns the imaginary part of z, a pointer to a big.Rat value.
func (z *Complex) Im() *big.Rat {
	return z.im
}

// SetRe sets the real part of z equal to a.
func (z *Complex) SetRe(a *big.Rat) {
	z.re = a
}

// SetIm sets the imaginary part of z equal to b.
func (z *Complex) SetIm(b *big.Rat) {
	z.im = b
}

// Cartesian returns the two Cartesian components of z.
func (z *Complex) Cartesian() (a, b *big.Rat) {
	a = z.Re()
	b = z.Im()
	return
}

// String returns the string version of a Complex value. If z = a + bi, then
// the string is "(a+bi)", similar to complex128 values.
func (z *Complex) String() string {
	a := make([]string, 5)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", z.Re())
	if z.Im().Sign() == -1 {
		a[2] = fmt.Sprintf("%v", z.Im())
	} else {
		a[2] = fmt.Sprintf("+%v", z.Im())
	}
	a[3] = "i"
	a[4] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Complex) Equals(y *Complex) bool {
	if z.Re().Cmp(y.Re()) != 0 || z.Im().Cmp(y.Im()) != 0 {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Complex) Copy(y *Complex) *Complex {
	z.SetRe(y.Re())
	z.SetIm(y.Im())
	return z
}

// NewComplex returns a pointer to a Complex value made from four given int64
// values.
func NewComplex(a, b, c, d int64) *Complex {
	z := new(Complex)
	z.SetRe(big.NewRat(a, b))
	z.SetIm(big.NewRat(c, d))
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Complex) Scal(y *Complex, a *big.Rat) *Complex {
	z.SetRe(new(big.Rat).Mul(y.Re(), a))
	z.SetIm(new(big.Rat).Mul(y.Im(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Complex) Neg(y *Complex) *Complex {
	z.SetRe(new(big.Rat).Neg(y.Re()))
	z.SetIm(new(big.Rat).Neg(y.Im()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Complex) Conj(y *Complex) *Complex {
	z.SetRe(y.Re())
	z.SetIm(new(big.Rat).Neg(y.Im()))
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Complex) Add(x, y *Complex) *Complex {
	z.SetRe(new(big.Rat).Add(x.Re(), y.Re()))
	z.SetIm(new(big.Rat).Add(x.Im(), y.Im()))
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Complex) Sub(x, y *Complex) *Complex {
	z.SetRe(new(big.Rat).Sub(x.Re(), y.Re()))
	z.SetIm(new(big.Rat).Sub(x.Im(), y.Im()))
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rule is:
// 		Mul(i, i) = -1
// This binary operation is commutative and associative.
func (z *Complex) Mul(x, y *Complex) *Complex {
	p := new(Complex).Copy(x)
	q := new(Complex).Copy(y)
	z.SetRe(new(big.Rat).Sub(
		new(big.Rat).Mul(p.Re(), q.Re()),
		new(big.Rat).Mul(p.Im(), q.Im()),
	))
	z.SetIm(new(big.Rat).Add(
		new(big.Rat).Mul(p.Re(), q.Im()),
		new(big.Rat).Mul(p.Im(), q.Re()),
	))
	return z
}

// Quad returns the quadrance of z, a pointer to a big.Rat value.
func (z *Complex) Quad() *big.Rat {
	return new(big.Rat).Add(
		new(big.Rat).Mul(z.Re(), z.Re()),
		new(big.Rat).Mul(z.Im(), z.Im()),
	)
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Complex) Inv(y *Complex) *Complex {
	return z.Scal(new(Complex).Conj(y), new(big.Rat).Inv(y.Quad()))
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Complex) Quo(x, y *Complex) *Complex {
	return z.Mul(x, new(Complex).Inv(y))
}
