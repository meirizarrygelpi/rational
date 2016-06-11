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

var symbSupraPerplex = [8]string{"", "s", "ρ", "σ", "τ", "υ", "φ", "ψ"}

// An SupraPerplex represents a rational supra-perplex number.
type SupraPerplex struct {
	l, r InfraPerplex
}

// Cartesian returns the eight rational Cartesian components of z.
func (z *SupraPerplex) Cartesian() (*big.Rat, *big.Rat, *big.Rat, *big.Rat,
	*big.Rat, *big.Rat, *big.Rat, *big.Rat) {
	return &z.l.l.l, &z.l.l.r, &z.l.r.l, &z.l.r.r,
		&z.r.l.l, &z.r.l.r, &z.r.r.l, &z.r.r.r
}

// String returns the string representation of an SupraPerplex value.
//
// If z corresponds to a + bs + cρ + dσ + eτ + fυ + gφ + hψ, then the string
// is "(a+bs+cρ+dσ+eτ+fυ+gφ+hψ)", similar to complex128 values.
func (z *SupraPerplex) String() string {
	v := make([]*big.Rat, 8)
	v[0], v[1], v[2], v[3] = z.l.Cartesian()
	v[4], v[5], v[6], v[7] = z.r.Cartesian()
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
		a[j+1] = symbSupraPerplex[i]
		i++
	}
	a[16] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *SupraPerplex) Equals(y *SupraPerplex) bool {
	if !z.l.Equals(&y.l) || !z.r.Equals(&y.r) {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *SupraPerplex) Set(y *SupraPerplex) *SupraPerplex {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewSupraPerplex returns a pointer to the SupraPerplex value
// a+bs+cρ+dσ+eτ+fυ+gφ+hψ.
func NewSupraPerplex(a, b, c, d, e, f, g, h *big.Rat) *SupraPerplex {
	z := new(SupraPerplex)
	z.l.l.l.Set(a)
	z.l.l.r.Set(b)
	z.l.r.l.Set(c)
	z.l.r.r.Set(d)
	z.r.l.l.Set(e)
	z.r.l.r.Set(f)
	z.r.r.l.Set(g)
	z.r.r.r.Set(h)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *SupraPerplex) Scal(y *SupraPerplex, a *big.Rat) *SupraPerplex {
	z.l.Scal(&y.l, a)
	z.r.Scal(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *SupraPerplex) Neg(y *SupraPerplex) *SupraPerplex {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *SupraPerplex) Conj(y *SupraPerplex) *SupraPerplex {
	z.l.Conj(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *SupraPerplex) Add(x, y *SupraPerplex) *SupraPerplex {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *SupraPerplex) Sub(x, y *SupraPerplex) *SupraPerplex {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
//		Mul(s, s) = +1
// 		Mul(ρ, ρ) = Mul(σ, σ) = Mul(τ, τ) = 0
// 		Mul(υ, υ) = Mul(φ, φ) = Mul(ψ, ψ) = 0
// 		Mul(s, ρ) = -Mul(ρ, s) = +σ
// 		Mul(s, σ) = -Mul(σ, s) = +ρ
// 		Mul(s, τ) = -Mul(τ, s) = +υ
// 		Mul(s, υ) = -Mul(υ, s) = +τ
// 		Mul(s, φ) = -Mul(φ, s) = -ψ
// 		Mul(s, ψ) = -Mul(ψ, s) = -φ
// 		Mul(ρ, σ) = Mul(σ, ρ) = 0
// 		Mul(ρ, τ) = -Mul(τ, ρ) = +φ
// 		Mul(ρ, υ) = -Mul(υ, ρ) = +ψ
// 		Mul(ρ, φ) = Mul(φ, ρ) = 0
// 		Mul(ρ, ψ) = Mul(ψ, ρ) = 0
// 		Mul(σ, τ) = -Mul(τ, σ) = +ψ
// 		Mul(σ, υ) = -Mul(υ, σ) = +φ
// 		Mul(σ, φ) = Mul(φ, σ) = 0
// 		Mul(σ, ψ) = Mul(ψ, σ) = 0
// 		Mul(τ, υ) = Mul(υ, τ) = 0
// 		Mul(τ, φ) = Mul(φ, τ) = 0
// 		Mul(τ, ψ) = Mul(ψ, τ) = 0
// 		Mul(υ, φ) = Mul(φ, υ) = 0
// 		Mul(υ, ψ) = Mul(ψ, υ) = 0
// 		Mul(φ, ψ) = Mul(ψ, φ) = 0
// This binary operation is noncommutative and nonassociative.
func (z *SupraPerplex) Mul(x, y *SupraPerplex) *SupraPerplex {
	a := new(InfraPerplex).Set(&x.l)
	b := new(InfraPerplex).Set(&x.r)
	c := new(InfraPerplex).Set(&y.l)
	d := new(InfraPerplex).Set(&y.r)
	temp := new(InfraPerplex)
	z.l.Mul(a, c)
	z.r.Add(
		z.r.Mul(d, a),
		temp.Mul(b, temp.Conj(c)),
	)
	return z
}

// Commutator sets z equal to the commutator of x and y:
// 		Mul(x, y) - Mul(y, x)
// Then it returns z.
func (z *SupraPerplex) Commutator(x, y *SupraPerplex) *SupraPerplex {
	return z.Sub(
		z.Mul(x, y),
		new(SupraPerplex).Mul(y, x),
	)
}

// Associator sets z equal to the associator of w, x, and y:
// 		Mul(Mul(w, x), y) - Mul(w, Mul(x, y))
// Then it returns z.
func (z *SupraPerplex) Associator(w, x, y *SupraPerplex) *SupraPerplex {
	temp := new(SupraPerplex)
	return z.Sub(
		z.Mul(z.Mul(w, x), y),
		temp.Mul(w, temp.Mul(x, y)),
	)
}

// Quad returns the quadrance of z. If z = a+bs+cρ+dσ+eτ+fυ+gφ+hψ, then the
// quadrance is
//		Mul(a, a) - Mul(b, b)
// This can be positive, negative, or zero.
func (z *SupraPerplex) Quad() *big.Rat {
	return z.l.Quad()
}

// IsZeroDiv returns true if z is a zero divisor. This is equivalent to z being
// nilpotent.
func (z *SupraPerplex) IsZeroDiv() bool {
	return z.l.IsZeroDiv()
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *SupraPerplex) Inv(y *SupraPerplex) *SupraPerplex {
	if y.IsZeroDiv() {
		panic("inverse of zero divisor")
	}
	a := y.Quad()
	a.Inv(a)
	return z.Scal(z.Conj(y), a)
}

// QuoL sets z equal to the left quotient of x and y:
// 		Mul(Inv(y), x)
// Then it returns z.
func (z *SupraPerplex) QuoL(x, y *SupraPerplex) *SupraPerplex {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Mul(z.Inv(y), x)
}

// QuoR sets z equal to the right quotient of x and y:
// 		Mul(x, Inv(y))
// Then it returns z.
func (z *SupraPerplex) QuoR(x, y *SupraPerplex) *SupraPerplex {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Mul(x, z.Inv(y))
}

// Generate returns a random SupraPerplex value for quick.Check testing.
func (z *SupraPerplex) Generate(rand *rand.Rand, size int) reflect.Value {
	randomSupraPerplex := &SupraPerplex{
		*NewInfraPerplex(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
		*NewInfraPerplex(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
	}
	return reflect.ValueOf(randomSupraPerplex)
}
