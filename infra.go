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

// An Infra represents a rational infra number.
type Infra struct {
	l, r big.Rat
}

// L returns the left Cayley-Dickson part of z, a pointer to a big.Rat value.
// This coincides with the real part of z.
func (z *Infra) L() *big.Rat {
	return &z.l
}

// R returns the right Cayley-Dickson part of z, a pointer to a big.Rat value.
// This coincides with the dual part of z.
func (z *Infra) R() *big.Rat {
	return &z.r
}

// SetL sets the left Cayley-Dickson part of z equal to a.
func (z *Infra) SetL(a *big.Rat) {
	z.l = *a
}

// SetR sets the right Cayley-Dickson part of z equal to b.
func (z *Infra) SetR(b *big.Rat) {
	z.r = *b
}

// Cartesian returns the two Cartesian components of z.
func (z *Infra) Cartesian() (a, b *big.Rat) {
	a = z.L()
	b = z.R()
	return
}

// String returns the string version of a Infra value.
//
// If z corresponds to a + bα, then the string is "(a+bα)", similar to
// complex128 values.
func (z *Infra) String() string {
	a := make([]string, 5)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", z.L().RatString())
	if z.R().Sign() == -1 {
		a[2] = fmt.Sprintf("%v", z.R().RatString())
	} else {
		a[2] = fmt.Sprintf("+%v", z.R().RatString())
	}
	a[3] = "α"
	a[4] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Infra) Equals(y *Infra) bool {
	if z.L().Cmp(y.L()) != 0 || z.R().Cmp(y.R()) != 0 {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Infra) Copy(y *Infra) *Infra {
	z.SetL(y.L())
	z.SetR(y.R())
	return z
}

// NewInfra returns a pointer to a Infra value made from two given pointers to
// big.Rat values.
func NewInfra(a, b *big.Rat) *Infra {
	z := new(Infra)
	z.SetL(a)
	z.SetR(b)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Infra) Scal(y *Infra, a *big.Rat) *Infra {
	z.SetL(new(big.Rat).Mul(y.L(), a))
	z.SetR(new(big.Rat).Mul(y.R(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Infra) Neg(y *Infra) *Infra {
	z.SetL(new(big.Rat).Neg(y.L()))
	z.SetR(new(big.Rat).Neg(y.R()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Infra) Conj(y *Infra) *Infra {
	z.SetL(y.L())
	z.SetR(new(big.Rat).Neg(y.R()))
	return z
}

// Add sets z to the sum of x and y, and returns z.
func (z *Infra) Add(x, y *Infra) *Infra {
	z.SetL(new(big.Rat).Add(x.L(), y.L()))
	z.SetR(new(big.Rat).Add(x.R(), y.R()))
	return z
}

// Sub sets z to the difference of x and y, and returns z.
func (z *Infra) Sub(x, y *Infra) *Infra {
	z.SetL(new(big.Rat).Sub(x.L(), y.L()))
	z.SetR(new(big.Rat).Sub(x.R(), y.R()))
	return z
}

// Mul sets z to the product of x and y, and returns z.
//
// The multiplication rule is:
// 		Mul(α, α) = 0
// This binary operation is commutative and associative.
func (z *Infra) Mul(x, y *Infra) *Infra {
	a := new(big.Rat).Set(x.L())
	b := new(big.Rat).Set(x.R())
	c := new(big.Rat).Set(y.L())
	d := new(big.Rat).Set(y.R())
	s, t, u := new(big.Rat), new(big.Rat), new(big.Rat)
	z.SetL(
		s.Mul(a, c),
	)
	z.SetR(t.Add(
		t.Mul(d, a),
		u.Mul(b, c),
	))
	return z
}

// Quad returns the quadrance of z, a pointer to a big.Rat value.
func (z *Infra) Quad() *big.Rat {
	return new(big.Rat).Mul(z.L(), z.L())
}

// IsZeroDiv returns true if z is a zero divisor. This is equivalent to z being
// nilpotent.
func (z *Infra) IsZeroDiv() bool {
	a := z.L()
	return a.Num().Cmp(big.NewInt(0)) == 0
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Infra) Inv(y *Infra) *Infra {
	if y.IsZeroDiv() {
		panic("inverse of zero divisor")
	}
	return z.Scal(z.Conj(y), new(big.Rat).Inv(y.Quad()))
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Infra) Quo(x, y *Infra) *Infra {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Mul(x, z.Inv(y))
}

// Generate a random Infra value for quick.Check testing.
func (z *Infra) Generate(rand *rand.Rand, size int) reflect.Value {
	randomInfra := &Infra{
		*big.NewRat(rand.Int63(), rand.Int63()),
		*big.NewRat(rand.Int63(), rand.Int63()),
	}
	return reflect.ValueOf(randomInfra)
}
