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

// A Perplex represents a rational split-complex number.
type Perplex struct {
	l, r *big.Rat
}

// L returns the left Cayley-Dickson part of z, a pointer to a big.Rat value.
// This is equivalent to the real part of z.
func (z *Perplex) L() *big.Rat {
	return z.l
}

// R returns the right Cayley-Dickson part of z, a pointer to a big.Rat value.
// This is equivalent to the split part of z.
func (z *Perplex) R() *big.Rat {
	return z.r
}

// SetL sets the left Cayley-Dickson part of z equal to a.
func (z *Perplex) SetL(a *big.Rat) {
	z.l = a
}

// SetR sets the right Cayley-Dickson part of z equal to b.
func (z *Perplex) SetR(b *big.Rat) {
	z.r = b
}

// Cartesian returns the two Cartesian components of z.
func (z *Perplex) Cartesian() (a, b *big.Rat) {
	a, b = z.L(), z.R()
	return
}

// String returns the string version of a Perplex value.
//
// If z corresponds to a + bs, then the string is "(a+bs)", similar to
// complex128 values.
func (z *Perplex) String() string {
	a := make([]string, 5)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", z.L())
	if z.R().Sign() == -1 {
		a[2] = fmt.Sprintf("%v", z.R())
	} else {
		a[2] = fmt.Sprintf("+%v", z.R())
	}
	a[3] = "s"
	a[4] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Perplex) Equals(y *Perplex) bool {
	if z.L().Cmp(y.L()) != 0 || z.R().Cmp(y.R()) != 0 {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Perplex) Copy(y *Perplex) *Perplex {
	z.SetL(y.L())
	z.SetR(y.R())
	return z
}

// NewPerplex returns a pointer to a Perplex value made from two given pointers
// to big.Rat values.
func NewPerplex(a, b *big.Rat) *Perplex {
	z := new(Perplex)
	z.SetL(a)
	z.SetR(b)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Perplex) Scal(y *Perplex, a *big.Rat) *Perplex {
	z.SetL(new(big.Rat).Mul(y.L(), a))
	z.SetR(new(big.Rat).Mul(y.R(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Perplex) Neg(y *Perplex) *Perplex {
	z.SetL(new(big.Rat).Neg(y.L()))
	z.SetR(new(big.Rat).Neg(y.R()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Perplex) Conj(y *Perplex) *Perplex {
	z.SetL(y.L())
	z.SetR(new(big.Rat).Neg(y.R()))
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Perplex) Add(x, y *Perplex) *Perplex {
	z.SetL(new(big.Rat).Add(x.L(), y.L()))
	z.SetR(new(big.Rat).Add(x.R(), y.R()))
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Perplex) Sub(x, y *Perplex) *Perplex {
	z.SetL(new(big.Rat).Sub(x.L(), y.L()))
	z.SetR(new(big.Rat).Sub(x.R(), y.R()))
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rule is:
// 		Mul(s, s) = +1
// This binary operation is commutative and associative.
func (z *Perplex) Mul(x, y *Perplex) *Perplex {
	a, b := x.L(), x.R()
	c, d := y.L(), y.R()
	s, t, u := new(big.Rat), new(big.Rat), new(big.Rat)
	z.SetL(s.Add(
		s.Mul(a, c),
		u.Mul(d, b),
	))
	z.SetR(t.Add(
		t.Mul(d, a),
		u.Mul(b, c),
	))
	return z
}

// Quad returns the quadrance of z, a pointer to a big.Rat value.
func (z *Perplex) Quad() *big.Rat {
	t := new(big.Rat)
	return t.Sub(
		t.Mul(z.L(), z.L()),
		new(big.Rat).Mul(z.R(), z.R()),
	)
}

// IsZeroDiv returns true if z is a zero divisor.
func (z *Perplex) IsZeroDiv() bool {
	if z.L().Cmp(z.R()) == 0 {
		return true
	}
	if z.L().Cmp(new(big.Rat).Neg(z.R())) == 0 {
		return true
	}
	return false
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Perplex) Inv(y *Perplex) *Perplex {
	if y.IsZeroDiv() {
		panic("inverse of zero divisor")
	}
	return z.Scal(z.Conj(y), new(big.Rat).Inv(y.Quad()))
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Perplex) Quo(x, y *Perplex) *Perplex {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Mul(x, z.Inv(y))
}

// Idempotent sets z equal to a pointer to an idempotent Perplex.
func (z *Perplex) Idempotent(sign int) *Perplex {
	z.SetL(big.NewRat(1, 2))
	if sign < 0 {
		z.SetR(big.NewRat(-1, 2))
		return z
	}
	z.SetR(big.NewRat(1, 2))
	return z
}

// Generate a random Perplex value for quick.Check.
func (z *Perplex) Generate(rand *rand.Rand, size int) reflect.Value {
	randomPerplex := &Perplex{
		l: big.NewRat(rand.Int63(), rand.Int63()),
		r: big.NewRat(rand.Int63(), rand.Int63()),
	}
	return reflect.ValueOf(randomPerplex)
}
