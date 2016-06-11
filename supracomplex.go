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

var symbSupraComplex = [8]string{"", "i", "α", "β", "γ", "δ", "ε", "ζ"}

// A SupraComplex represents a rational supra-complex number.
type SupraComplex struct {
	l, r InfraComplex
}

// Cartesian returns the eight rational Cartesian components of z.
func (z *SupraComplex) Cartesian() (*big.Rat, *big.Rat, *big.Rat, *big.Rat,
	*big.Rat, *big.Rat, *big.Rat, *big.Rat) {
	return &z.l.l.l, &z.l.l.r, &z.l.r.l, &z.l.r.r,
		&z.r.l.l, &z.r.l.r, &z.r.r.l, &z.r.r.r
}

// String returns the string representation of a SupraComplex value.
//
// If z corresponds to a + bi + cα + dβ + eγ + fδ + gε + hζ, then the string
// is"(a+bi+cα+dβ+eγ+fδ+gε+hζ)", similar to complex128 values.
func (z *SupraComplex) String() string {
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
		a[j+1] = symbSupraComplex[i]
		i++
	}
	a[16] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *SupraComplex) Equals(y *SupraComplex) bool {
	if !z.l.Equals(&y.l) || !z.r.Equals(&y.r) {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *SupraComplex) Set(y *SupraComplex) *SupraComplex {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewSupraComplex returns a pointer to the SupraComplex value
// a+bi+cα+dβ+eγ+fδ+gε+hζ.
func NewSupraComplex(a, b, c, d, e, f, g, h *big.Rat) *SupraComplex {
	z := new(SupraComplex)
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
func (z *SupraComplex) Scal(y *SupraComplex, a *big.Rat) *SupraComplex {
	z.l.Scal(&y.l, a)
	z.r.Scal(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *SupraComplex) Neg(y *SupraComplex) *SupraComplex {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *SupraComplex) Conj(y *SupraComplex) *SupraComplex {
	z.l.Conj(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *SupraComplex) Add(x, y *SupraComplex) *SupraComplex {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *SupraComplex) Sub(x, y *SupraComplex) *SupraComplex {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
//		Mul(i, i) = -1
// 		Mul(α, α) = Mul(β, β) = Mul(γ, γ) = 0
// 		Mul(δ, δ) = Mul(ε, ε) = Mul(ζ, ζ) = 0
// 		Mul(i, α) = -Mul(α, i) = +β
// 		Mul(i, β) = -Mul(β, i) = -α
// 		Mul(i, γ) = -Mul(γ, i) = +δ
// 		Mul(i, δ) = -Mul(δ, i) = -γ
// 		Mul(i, ε) = -Mul(ε, i) = -ζ
// 		Mul(i, ζ) = -Mul(ζ, i) = +ε
// 		Mul(α, β) = Mul(β, α) = 0
// 		Mul(α, γ) = -Mul(γ, α) = +ε
// 		Mul(α, δ) = -Mul(δ, α) = +ζ
// 		Mul(α, ε) = Mul(ε, α) = 0
// 		Mul(α, ζ) = Mul(ζ, α) = 0
// 		Mul(β, γ) = -Mul(γ, β) = +ζ
// 		Mul(β, δ) = -Mul(δ, β) = -ε
// 		Mul(β, ε) = Mul(ε, β) = 0
// 		Mul(β, ζ) = Mul(ζ, β) = 0
// 		Mul(γ, δ) = Mul(δ, γ) = 0
// 		Mul(γ, ε) = Mul(ε, γ) = 0
// 		Mul(γ, ζ) = Mul(ζ, γ) = 0
// 		Mul(δ, ε) = Mul(ε, δ) = 0
// 		Mul(δ, ζ) = Mul(ζ, δ) = 0
// 		Mul(ε, ζ) = Mul(ζ, ε) = 0
// This binary operation is noncommutative and nonassociative.
func (z *SupraComplex) Mul(x, y *SupraComplex) *SupraComplex {
	a := new(InfraComplex).Set(&x.l)
	b := new(InfraComplex).Set(&x.r)
	c := new(InfraComplex).Set(&y.l)
	d := new(InfraComplex).Set(&y.r)
	temp := new(InfraComplex)
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
func (z *SupraComplex) Commutator(x, y *SupraComplex) *SupraComplex {
	return z.Sub(
		z.Mul(x, y),
		new(SupraComplex).Mul(y, x),
	)
}

// Associator sets z equal to the associator of w, x, and y:
// 		Mul(Mul(w, x), y) - Mul(w, Mul(x, y))
// Then it returns z.
func (z *SupraComplex) Associator(w, x, y *SupraComplex) *SupraComplex {
	temp := new(SupraComplex)
	return z.Sub(
		z.Mul(z.Mul(w, x), y),
		temp.Mul(w, temp.Mul(x, y)),
	)
}

// Quad returns the quadrance of z. If z = a+bi+cα+dβ+eγ+fδ+gε+hζ, then the
// quadrance is
//		Mul(a, a) + Mul(b, b)
// This is always non-negative.
func (z *SupraComplex) Quad() *big.Rat {
	return z.l.Quad()
}

// IsZeroDiv returns true if z is a zero divisor. This is equivalent to z being
// nilpotent.
func (z *SupraComplex) IsZeroDiv() bool {
	zero := new(InfraComplex)
	return z.l.Equals(zero)
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *SupraComplex) Inv(y *SupraComplex) *SupraComplex {
	if y.IsZeroDiv() {
		panic("inverse of zero divisor")
	}
	a := y.Quad()
	a.Inv(a)
	return z.Scal(z.Conj(y), a)
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *SupraComplex) Quo(x, y *SupraComplex) *SupraComplex {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Mul(x, z.Inv(y))
}

// Generate returns a random SupraComplex value for quick.Check testing.
func (z *SupraComplex) Generate(rand *rand.Rand, size int) reflect.Value {
	randomSupraComplex := &SupraComplex{
		*NewInfraComplex(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
		*NewInfraComplex(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
	}
	return reflect.ValueOf(randomSupraComplex)
}
