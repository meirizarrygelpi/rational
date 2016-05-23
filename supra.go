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

var symbSupra = [4]string{"", "α", "β", "γ"}

// A Supra represents a rational supra number.
type Supra struct {
	l, r Infra
}

// L returns the left Cayley-Dickson part of z, a pointer to a Infra value.
func (z *Supra) L() *Infra {
	return &z.l
}

// R returns the right Cayley-Dickson part of z, a pointer to a Infra value.
func (z *Supra) R() *Infra {
	return &z.r
}

// SetL sets the left Cayley-Dickson part of z equal to a.
func (z *Supra) SetL(a *Infra) {
	z.l = *a
}

// SetR sets the right Cayley-Dickson part of z equal to b.
func (z *Supra) SetR(b *Infra) {
	z.r = *b
}

// Cartesian returns the four Cartesian components of z.
func (z *Supra) Cartesian() (a, b, c, d *big.Rat) {
	a, b = z.L().Cartesian()
	c, d = z.R().Cartesian()
	return
}

// String returns the string representation of a Supra value.
//
// If z corresponds to a + bα + cβ + dγ, then the string is "(a+bα+cβ+dγ)",
// similar to complex128 values.
func (z *Supra) String() string {
	v := make([]*big.Rat, 4)
	v[0], v[1] = z.L().Cartesian()
	v[2], v[3] = z.R().Cartesian()
	a := make([]string, 9)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", v[0].RatString())
	i := 1
	for j := 2; j < 8; j = j + 2 {
		if v[i].Sign() == -1 {
			a[j] = fmt.Sprintf("%v", v[i].RatString())
		} else {
			a[j] = fmt.Sprintf("+%v", v[i].RatString())
		}
		a[j+1] = symbSupra[i]
		i++
	}
	a[8] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Supra) Equals(y *Supra) bool {
	if !z.L().Equals(y.L()) || !z.R().Equals(y.R()) {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Supra) Copy(y *Supra) *Supra {
	z.SetL(y.L())
	z.SetR(y.R())
	return z
}

// NewSupra returns a pointer to a Supra value made from four given pointers to
// big.Rat values.
func NewSupra(a, b, c, d *big.Rat) *Supra {
	z := new(Supra)
	z.SetL(NewInfra(a, b))
	z.SetR(NewInfra(c, d))
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Supra) Scal(y *Supra, a *big.Rat) *Supra {
	z.SetL(new(Infra).Scal(y.L(), a))
	z.SetR(new(Infra).Scal(y.R(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Supra) Neg(y *Supra) *Supra {
	z.SetL(new(Infra).Neg(y.L()))
	z.SetR(new(Infra).Neg(y.R()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Supra) Conj(y *Supra) *Supra {
	z.SetL(new(Infra).Conj(y.L()))
	z.SetR(new(Infra).Neg(y.R()))
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Supra) Add(x, y *Supra) *Supra {
	z.SetL(new(Infra).Add(x.L(), y.L()))
	z.SetR(new(Infra).Add(x.R(), y.R()))
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Supra) Sub(x, y *Supra) *Supra {
	z.SetL(new(Infra).Sub(x.L(), y.L()))
	z.SetR(new(Infra).Sub(x.R(), y.R()))
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
// 		Mul(α, α) = Mul(β, β) = Mul(γ, γ) = 0
// 		Mul(α, β) = -Mul(β, α) = γ
// 		Mul(β, γ) = Mul(γ, β) = 0
// 		Mul(γ, α) = Mul(α, γ) = 0
// This binary operation is noncommutative but associative.
func (z *Supra) Mul(x, y *Supra) *Supra {
	a := new(Infra).Copy(x.L())
	b := new(Infra).Copy(x.R())
	c := new(Infra).Copy(y.L())
	d := new(Infra).Copy(y.R())
	s, t, u := new(Infra), new(Infra), new(Infra)
	z.SetL(
		s.Mul(a, c),
	)
	z.SetR(t.Add(
		t.Mul(d, a),
		u.Mul(b, u.Conj(c)),
	))
	return z
}

// Commutator sets z equal to the commutator of x and y, and returns z.
func (z *Supra) Commutator(x, y *Supra) *Supra {
	return z.Sub(
		z.Mul(x, y),
		new(Supra).Mul(y, x),
	)
}

// Quad returns the quadrance of z, a pointer to a big.Rat value.
func (z *Supra) Quad() *big.Rat {
	return z.L().Quad()
}

// IsZeroDiv returns true if z is a zero divisor.
func (z *Supra) IsZeroDiv() bool {
	return z.L().IsZeroDiv()
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Supra) Inv(y *Supra) *Supra {
	if y.IsZeroDiv() {
		panic("inverse of zero divisor")
	}
	return z.Scal(z.Conj(y), new(big.Rat).Inv(y.Quad()))
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Supra) Quo(x, y *Supra) *Supra {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Mul(x, z.Inv(y))
}

// Generate returns a random Supra value for quick.Check testing.
func (z *Supra) Generate(rand *rand.Rand, size int) reflect.Value {
	randomSupra := &Supra{
		*NewInfra(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
		*NewInfra(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
	}
	return reflect.ValueOf(randomSupra)
}
