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
	body [2]*big.Rat
}

// Real returns the real part of z, a pointer to a big.Rat value.
func (z *Perplex) Real() *big.Rat {
	return z.body[0]
}

// Imag returns the imaginary part of z, a pointer to a big.Rat value.
func (z *Perplex) Imag() *big.Rat {
	return z.body[1]
}

// SetReal sets the real part of z equal to a.
func (z *Perplex) SetReal(a *big.Rat) {
	z.body[0] = a
}

// SetImag sets the imaginary part of z equal to b.
func (z *Perplex) SetImag(b *big.Rat) {
	z.body[1] = b
}

// Cartesian returns the two Cartesian components of z.
func (z *Perplex) Cartesian() (a, b *big.Rat) {
	a, b = z.body[0], z.body[1]
	return
}

// String returns the string version of a Perplex value. If z = a + bs, then
// the string is "(a+bs)", similar to complex128 values.
func (z *Perplex) String() string {
	a := make([]string, 5)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", z.Real())
	if z.Imag().Sign() == -1 {
		a[2] = fmt.Sprintf("%v", z.Imag())
	} else {
		a[2] = fmt.Sprintf("+%v", z.Imag())
	}
	a[3] = "s"
	a[4] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Perplex) Equals(y *Perplex) bool {
	if z.Real().Cmp(y.Real()) != 0 || z.Imag().Cmp(y.Imag()) != 0 {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Perplex) Copy(y *Perplex) *Perplex {
	z.SetReal(y.Real())
	z.SetImag(y.Imag())
	return z
}

// NewPerplex returns a pointer to a Perplex value made from two given pointers
// to big.Rat values.
func NewPerplex(a, b *big.Rat) *Perplex {
	z := new(Perplex)
	z.SetReal(a)
	z.SetImag(b)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Perplex) Scal(y *Perplex, a *big.Rat) *Perplex {
	z.SetReal(new(big.Rat).Mul(y.Real(), a))
	z.SetImag(new(big.Rat).Mul(y.Imag(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Perplex) Neg(y *Perplex) *Perplex {
	z.SetReal(new(big.Rat).Neg(y.Real()))
	z.SetImag(new(big.Rat).Neg(y.Imag()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Perplex) Conj(y *Perplex) *Perplex {
	z.SetReal(y.Real())
	z.SetImag(new(big.Rat).Neg(y.Imag()))
	return z
}

// Add sets z to the sum of x and y, and returns z.
func (z *Perplex) Add(x, y *Perplex) *Perplex {
	z.SetReal(new(big.Rat).Add(x.Real(), y.Real()))
	z.SetImag(new(big.Rat).Add(x.Imag(), y.Imag()))
	return z
}

// Sub sets z to the difference of x and y, and returns z.
func (z *Perplex) Sub(x, y *Perplex) *Perplex {
	z.SetReal(new(big.Rat).Sub(x.Real(), y.Real()))
	z.SetImag(new(big.Rat).Sub(x.Imag(), y.Imag()))
	return z
}

// Mul sets z to the product of x and y, and returns z.
func (z *Perplex) Mul(x, y *Perplex) *Perplex {
	p := new(Perplex).Copy(x)
	q := new(Perplex).Copy(y)
	z.SetReal(new(big.Rat).Add(
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
func (z *Perplex) Quad() *big.Rat {
	return new(big.Rat).Sub(
		new(big.Rat).Mul(z.Real(), z.Real()),
		new(big.Rat).Mul(z.Imag(), z.Imag()),
	)
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Perplex) Inv(y *Perplex) *Perplex {
	return z.Scal(z.Conj(y), new(big.Rat).Inv(y.Quad()))
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Perplex) Quo(x, y *Perplex) *Perplex {
	return z.Mul(x, z.Inv(y))
}
