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

// A Complex represents a rational complex number.
type Complex struct {
	l, r big.Rat
}

// L returns the left Cayley-Dickson part of z, a pointer to a big.Rat value.
// This coincides with the real part of z.
func (z *Complex) L() *big.Rat {
	return &z.l
}

// R returns the right Cayley-Dickson part of z, a pointer to a big.Rat value.
// This coincides with imaginary part of z.
func (z *Complex) R() *big.Rat {
	return &z.r
}

// SetL sets the left Cayley-Dickson part of z equal to a.
func (z *Complex) SetL(a *big.Rat) {
	z.l = *a
}

// SetR sets the right Cayley-Dickson part of z equal to b.
func (z *Complex) SetR(b *big.Rat) {
	z.r = *b
}

// Cartesian returns the two Cartesian components of z.
func (z *Complex) Cartesian() (a, b *big.Rat) {
	a = z.L()
	b = z.R()
	return
}

// String returns the string version of a Complex value.
//
// If z corresponds to a + bi, then the string is "(a+bi)", similar to
// complex128 values.
func (z *Complex) String() string {
	a := make([]string, 5)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", z.L().RatString())
	if z.R().Sign() < 0 {
		a[2] = fmt.Sprintf("%v", z.R().RatString())
	} else {
		a[2] = fmt.Sprintf("+%v", z.R().RatString())
	}
	a[3] = "i"
	a[4] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Complex) Equals(y *Complex) bool {
	if z.L().Cmp(y.L()) != 0 || z.R().Cmp(y.R()) != 0 {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Complex) Copy(y *Complex) *Complex {
	z.SetL(y.L())
	z.SetR(y.R())
	return z
}

// NewComplex returns a pointer to a Complex value made from two given pointers
// to big.Rat values.
func NewComplex(a, b *big.Rat) *Complex {
	z := new(Complex)
	z.SetL(a)
	z.SetR(b)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Complex) Scal(y *Complex, a *big.Rat) *Complex {
	z.SetL(new(big.Rat).Mul(y.L(), a))
	z.SetR(new(big.Rat).Mul(y.R(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Complex) Neg(y *Complex) *Complex {
	z.SetL(new(big.Rat).Neg(y.L()))
	z.SetR(new(big.Rat).Neg(y.R()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Complex) Conj(y *Complex) *Complex {
	z.SetL(y.L())
	z.SetR(new(big.Rat).Neg(y.R()))
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Complex) Add(x, y *Complex) *Complex {
	z.SetL(new(big.Rat).Add(x.L(), y.L()))
	z.SetR(new(big.Rat).Add(x.R(), y.R()))
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Complex) Sub(x, y *Complex) *Complex {
	z.SetL(new(big.Rat).Sub(x.L(), y.L()))
	z.SetR(new(big.Rat).Sub(x.R(), y.R()))
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rule is:
// 		Mul(i, i) = -1
// This binary operation is commutative and associative.
func (z *Complex) Mul(x, y *Complex) *Complex {
	a := new(big.Rat).Set(x.L())
	b := new(big.Rat).Set(x.R())
	c := new(big.Rat).Set(y.L())
	d := new(big.Rat).Set(y.R())
	s, t, u := new(big.Rat), new(big.Rat), new(big.Rat)
	z.SetL(s.Sub(
		s.Mul(a, c),
		u.Mul(d, b),
	))
	z.SetR(t.Add(
		t.Mul(d, a),
		u.Mul(b, c),
	))
	return z
}

// Quad returns the non-negative quadrance of z, a pointer to a big.Rat value.
func (z *Complex) Quad() *big.Rat {
	t := new(big.Rat)
	return t.Add(
		t.Mul(z.L(), z.L()),
		new(big.Rat).Mul(z.R(), z.R()),
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

// Gauss returns a Complex value corresponding to a Gaussian integer with real
// part equal to a and imaginary part equal to b.
func Gauss(a, b *big.Int) *Complex {
	z := new(Complex)
	z.SetL(new(big.Rat).SetInt(a))
	z.SetR(new(big.Rat).SetInt(b))
	return z
}

// Generate a random Complex value for quick.Check testing.
func (z *Complex) Generate(rand *rand.Rand, size int) reflect.Value {
	randomComplex := &Complex{
		*big.NewRat(rand.Int63(), rand.Int63()),
		*big.NewRat(rand.Int63(), rand.Int63()),
	}
	return reflect.ValueOf(randomComplex)
}
