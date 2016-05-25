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

var symbInfraCockle = [8]string{"", "i", "t", "u", "ρ", "σ", "τ", "υ"}

// An InfraCockle represents a rational infra-complex number.
type InfraCockle struct {
	l, r Cockle
}

// L returns the left Cayley-Dickson part of z, a pointer to a Cockle value.
func (z *InfraCockle) L() *Cockle {
	return &z.l
}

// R returns the right Cayley-Dickson part of z, a pointer to a Cockle value.
func (z *InfraCockle) R() *Cockle {
	return &z.r
}

// SetL sets the left Cayley-Dickson part of z equal to a.
func (z *InfraCockle) SetL(a *Cockle) {
	z.l = *a
}

// SetR sets the right Cayley-Dickson part of z equal to b.
func (z *InfraCockle) SetR(b *Cockle) {
	z.r = *b
}

// Cartesian returns the eight Cartesian components of z.
func (z *InfraCockle) Cartesian() (a, b, c, d, e, f, g, h *big.Rat) {
	a, b, c, d = z.L().Cartesian()
	e, f, g, h = z.R().Cartesian()
	return
}

// String returns the string representation of an InfraCockle value.
//
// If z corresponds to a + bi + ct + du + eρ + fσ + gτ + hυ, then the string
// is"(a+bi+ct+du+eρ+fσ+gτ+hυ)", similar to complex128 values.
func (z *InfraCockle) String() string {
	v := make([]*big.Rat, 8)
	v[0], v[1], v[2], v[3] = z.L().Cartesian()
	v[4], v[5], v[6], v[7] = z.R().Cartesian()
	a := make([]string, 17)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", v[0].RatString())
	i := 1
	for j := 2; j < 16; j = j + 2 {
		if v[i].Sign() < 0 {
			a[j] = fmt.Sprintf("%v", v[i].RatString())
		} else {
			a[j] = fmt.Sprintf("+%v", v[i].RatString())
		}
		a[j+1] = symbInfraCockle[i]
		i++
	}
	a[16] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *InfraCockle) Equals(y *InfraCockle) bool {
	if !z.L().Equals(y.L()) || !z.R().Equals(y.R()) {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *InfraCockle) Copy(y *InfraCockle) *InfraCockle {
	z.SetL(y.L())
	z.SetR(y.R())
	return z
}

// NewInfraCockle returns a pointer to an InfraCockle value made from eight
// given pointers to big.Rat values.
func NewInfraCockle(a, b, c, d, e, f, g, h *big.Rat) *InfraCockle {
	z := new(InfraCockle)
	z.SetL(NewCockle(a, b, c, d))
	z.SetR(NewCockle(e, f, g, h))
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *InfraCockle) Scal(y *InfraCockle, a *big.Rat) *InfraCockle {
	z.SetL(new(Cockle).Scal(y.L(), a))
	z.SetR(new(Cockle).Scal(y.R(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *InfraCockle) Neg(y *InfraCockle) *InfraCockle {
	z.SetL(new(Cockle).Neg(y.L()))
	z.SetR(new(Cockle).Neg(y.R()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *InfraCockle) Conj(y *InfraCockle) *InfraCockle {
	z.SetL(new(Cockle).Conj(y.L()))
	z.SetR(new(Cockle).Neg(y.R()))
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *InfraCockle) Add(x, y *InfraCockle) *InfraCockle {
	z.SetL(new(Cockle).Add(x.L(), y.L()))
	z.SetR(new(Cockle).Add(x.R(), y.R()))
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *InfraCockle) Sub(x, y *InfraCockle) *InfraCockle {
	z.SetL(new(Cockle).Sub(x.L(), y.L()))
	z.SetR(new(Cockle).Sub(x.R(), y.R()))
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
//		Mul(i, i) = -1
// 		Mul(t, t) = Mul(u, u) = +1
// 		Mul(ρ, ρ) = Mul(σ, σ) = Mul(τ, τ) = Mul(υ, υ) = 0
// 		Mul(i, t) = -Mul(t, i) = +u
// 		Mul(i, u) = -Mul(u, i) = -t
// 		Mul(i, ρ) = -Mul(ρ, i) = +σ
// 		Mul(i, σ) = -Mul(σ, i) = -ρ
// 		Mul(i, τ) = -Mul(τ, i) = -υ
// 		Mul(i, υ) = -Mul(υ, i) = +τ
// 		Mul(t, u) = -Mul(u, t) = -i
// 		Mul(t, ρ) = -Mul(ρ, t) = +τ
// 		Mul(t, σ) = -Mul(σ, t) = +υ
// 		Mul(t, τ) = -Mul(τ, t) = +ρ
// 		Mul(t, υ) = -Mul(υ, t) = +σ
// 		Mul(u, ρ) = -Mul(ρ, u) = +υ
// 		Mul(u, σ) = -Mul(σ, u) = -τ
// 		Mul(u, τ) = -Mul(τ, u) = -σ
// 		Mul(u, υ) = -Mul(υ, u) = +ρ
// 		Mul(ρ, σ) = Mul(σ, ρ) = 0
// 		Mul(ρ, τ) = Mul(τ, ρ) = 0
// 		Mul(ρ, υ) = Mul(υ, ρ) = 0
// 		Mul(σ, τ) = Mul(τ, σ) = 0
// 		Mul(σ, υ) = Mul(υ, σ) = 0
// 		Mul(τ, υ) = Mul(υ, τ) = 0
// This binary operation is noncommutative but associative.
func (z *InfraCockle) Mul(x, y *InfraCockle) *InfraCockle {
	a := new(Cockle).Copy(x.L())
	b := new(Cockle).Copy(x.R())
	c := new(Cockle).Copy(y.L())
	d := new(Cockle).Copy(y.R())
	s, t, u := new(Cockle), new(Cockle), new(Cockle)
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
func (z *InfraCockle) Commutator(x, y *InfraCockle) *InfraCockle {
	return z.Sub(
		z.Mul(x, y),
		new(InfraCockle).Mul(y, x),
	)
}

// Associator sets z equal to the associator of w, x, and y, and returns z.
func (z *InfraCockle) Associator(w, x, y *InfraCockle) *InfraCockle {
	t := new(InfraCockle)
	return z.Sub(
		z.Mul(z.Mul(w, x), y),
		t.Mul(w, t.Mul(x, y)),
	)
}

// Quad returns the quadrance of z, a pointer to a big.Rat value.
func (z *InfraCockle) Quad() *big.Rat {
	return z.L().Quad()
}

// IsZeroDiv returns true if z is a zero divisor. This is equivalent to z being
// nilpotent.
func (z *InfraCockle) IsZeroDiv() bool {
	return z.L().IsZeroDiv()
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *InfraCockle) Inv(y *InfraCockle) *InfraCockle {
	if y.IsZeroDiv() {
		panic("inverse of zero divisor")
	}
	return z.Scal(z.Conj(y), new(big.Rat).Inv(y.Quad()))
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *InfraCockle) Quo(x, y *InfraCockle) *InfraCockle {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Mul(x, z.Inv(y))
}

// Generate returns a random InfraCockle value for quick.Check testing.
func (z *InfraCockle) Generate(rand *rand.Rand, size int) reflect.Value {
	randomInfraCockle := &InfraCockle{
		*NewCockle(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
		*NewCockle(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
	}
	return reflect.ValueOf(randomInfraCockle)
}
