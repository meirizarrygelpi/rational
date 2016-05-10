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
	body [2]*big.Rat
}

// Real returns the real part of z, a pointer to a big.Rat value.
func (z *Complex) Real() *big.Rat {
	return z.body[0]
}

// Imag returns the imaginary part of z, a pointer to a big.Rat value.
func (z *Complex) Imag() *big.Rat {
	return z.body[1]
}

// SetReal sets the real part of z equal to a.
func (z *Complex) SetReal(a *big.Rat) {
	z.body[0] = a
}

// SetImag sets the imaginary part of z equal to b.
func (z *Complex) SetImag(b *big.Rat) {
	z.body[1] = b
}

// Cartesian returns the two Cartesian components of z.
func (z *Complex) Cartesian() (a, b *big.Rat) {
	a, b = z.body[0], z.body[1]
	return
}

// String returns the string version of a Complex value. If z = a + bi, then
// the string is "(a+bi)", similar to complex128 values.
func (z *Complex) String() string {
	a := make([]string, 5)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", z.Real())
	if z.Imag().Sign() == -1 {
		a[2] = fmt.Sprintf("%v", z.Imag())
	} else {
		a[2] = fmt.Sprintf("+%v", z.Imag())
	}
	a[3] = "i"
	a[4] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Complex) Equals(y *Complex) bool {
	if z.Real().Cmp(y.Real()) != 0 || z.Imag().Cmp(y.Imag()) != 0 {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Complex) Copy(y *Complex) *Complex {
	z.SetReal(y.Real())
	z.SetImag(y.Imag())
	return z
}

// NewComplex returns a pointer to a Complex value a/b + c/d i made from four
// given int64 values.
func NewComplex(a, b, c, d int64) *Complex {
	z := new(Complex)
	z.SetReal(big.NewRat(a, b))
	z.SetImag(big.NewRat(c, d))
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Complex) Scal(y *Complex, a *big.Rat) *Complex {
	z.SetReal(new(big.Rat).Mul(y.Real(), a))
	z.SetImag(new(big.Rat).Mul(y.Imag(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Complex) Neg(y *Complex) *Complex {
	z.SetReal(new(big.Rat).Neg(y.Real()))
	z.SetImag(new(big.Rat).Neg(y.Imag()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Complex) Conj(y *Complex) *Complex {
	z.SetReal(y.Real())
	z.SetImag(new(big.Rat).Neg(y.Imag()))
	return z
}

// Add sets z to the sum of x and y, and returns z.
func (z *Complex) Add(x, y *Complex) *Complex {
	z.SetReal(new(big.Rat).Add(x.Real(), y.Real()))
	z.SetImag(new(big.Rat).Add(x.Imag(), y.Imag()))
	return z
}

// Sub sets z to the difference of x and y, and returns z.
func (z *Complex) Sub(x, y *Complex) *Complex {
	z.SetReal(new(big.Rat).Sub(x.Real(), y.Real()))
	z.SetImag(new(big.Rat).Sub(x.Imag(), y.Imag()))
	return z
}

// Mul sets z to the product of x and y, and returns z.
func (z *Complex) Mul(x, y *Complex) *Complex {
	p := new(Complex).Copy(x)
	q := new(Complex).Copy(y)
	z.SetReal(new(big.Rat).Sub(
		new(big.Rat).Mul(p.Real(), q.Real()),
		new(big.Rat).Mul(p.Imag(), q.Imag()),
	))
	z.SetImag(new(big.Rat).Add(
		new(big.Rat).Mul(p.Real(), q.Imag()),
		new(big.Rat).Mul(p.Imag(), q.Real()),
	))
	return z
}

// Quad returns the quadrance of z, a pointer to a big.Rat value.
func (z *Complex) Quad() *big.Rat {
	return new(big.Rat).Add(
		new(big.Rat).Mul(z.Real(), z.Real()),
		new(big.Rat).Mul(z.Imag(), z.Imag()),
	)
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Complex) Inv(y *Complex) *Complex {
	return z.Scal(z.Conj(y), new(big.Rat).Inv(y.Quad()))
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Complex) Quo(x, y *Complex) *Complex {
	return z.Mul(x, z.Inv(y))
}
