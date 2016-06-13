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

var symbUltra = [8]string{"", "α", "β", "γ", "δ", "ε", "ζ", "η"}

// An Ultra represents a rational ultra number.
type Ultra struct {
	l, r Supra
}

// Real returns the (rational) real part of z.
func (z *Ultra) Real() *big.Rat {
	return (&z.l).Real()
}

// Cartesian returns the eight Cartesian components of z.
func (z *Ultra) Cartesian() (*big.Rat, *big.Rat, *big.Rat, *big.Rat,
	*big.Rat, *big.Rat, *big.Rat, *big.Rat) {
	return &z.l.l.l, &z.l.l.r, &z.l.r.l, &z.l.r.r,
		&z.r.l.l, &z.r.l.r, &z.r.r.l, &z.r.r.r
}

// String returns the string representation of an Ultra value.
//
// If z corresponds to a + bα + cβ + dγ + eδ + fε + gζ + hη, then the string
// is "(a+bα+cβ+dγ+eδ+fε+gζ+hη)", similar to complex128 values.
func (z *Ultra) String() string {
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
		a[j+1] = symbUltra[i]
		i++
	}
	a[16] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Ultra) Equals(y *Ultra) bool {
	if !z.l.Equals(&y.l) || !z.r.Equals(&y.r) {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *Ultra) Set(y *Ultra) *Ultra {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewUltra returns a pointer to the Ultra value a+bα+cβ+dγ+eδ+fε+gζ+hη.
func NewUltra(a, b, c, d, e, f, g, h *big.Rat) *Ultra {
	z := new(Ultra)
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
func (z *Ultra) Scal(y *Ultra, a *big.Rat) *Ultra {
	z.l.Scal(&y.l, a)
	z.r.Scal(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Ultra) Neg(y *Ultra) *Ultra {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Ultra) Conj(y *Ultra) *Ultra {
	z.l.Conj(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Ultra) Add(x, y *Ultra) *Ultra {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Ultra) Sub(x, y *Ultra) *Ultra {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
// 		Mul(α, α) = Mul(β, β) = Mul(γ, γ) = 0
// 		Mul(δ, δ) = Mul(ε, ε) = Mul(ζ, ζ) = Mul(η, η) = 0
// 		Mul(α, β) = -Mul(β, α) = +γ
// 		Mul(α, γ) = Mul(γ, α) = 0
// 		Mul(α, δ) = -Mul(δ, α) = +ε
// 		Mul(α, ε) = Mul(ε, α) = 0
// 		Mul(α, ζ) = -Mul(ζ, α) = -η
// 		Mul(α, η) = -Mul(η, α) = +ζ
// 		Mul(β, γ) = Mul(γ, β) = 0
// 		Mul(β, δ) = -Mul(δ, β) = +ζ
// 		Mul(β, ε) = -Mul(ε, β) = +η
// 		Mul(β, ζ) = Mul(ζ, β) = 0
// 		Mul(β, η) = Mul(η, β) = 0
// 		Mul(γ, δ) = -Mul(δ, γ) = +η
// 		Mul(γ, ε) = Mul(ε, γ) = 0
// 		Mul(γ, ζ) = Mul(ζ, γ) = 0
// 		Mul(γ, η) = Mul(η, γ) = 0
// 		Mul(δ, ε) = Mul(ε, δ) = 0
// 		Mul(δ, ζ) = Mul(ζ, δ) = 0
// 		Mul(δ, η) = Mul(η, δ) = 0
// 		Mul(ε, ζ) = Mul(ζ, ε) = 0
// 		Mul(ε, η) = Mul(η, ε) = 0
// 		Mul(ζ, η) = Mul(η, ζ) = 0
// This binary operation is noncommutative and nonassociative.
func (z *Ultra) Mul(x, y *Ultra) *Ultra {
	a := new(Supra).Set(&x.l)
	b := new(Supra).Set(&x.r)
	c := new(Supra).Set(&y.l)
	d := new(Supra).Set(&y.r)
	temp := new(Supra)
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
func (z *Ultra) Commutator(x, y *Ultra) *Ultra {
	return z.Sub(
		z.Mul(x, y),
		new(Ultra).Mul(y, x),
	)
}

// Associator sets z equal to the associator of w, x, and y:
// 		Mul(Mul(w, x), y) - Mul(w, Mul(x, y))
// Then it returns z.
func (z *Ultra) Associator(w, x, y *Ultra) *Ultra {
	temp := new(Ultra)
	return z.Sub(
		z.Mul(z.Mul(w, x), y),
		temp.Mul(w, temp.Mul(x, y)),
	)
}

// Quad returns the quadrance of z. If z = a+bα+cβ+dγ+eδ+fε+gζ+hη, then the
// quadrance is
//		Mul(a, a)
// This is always non-negative.
func (z *Ultra) Quad() *big.Rat {
	return z.l.Quad()
}

// IsZeroDiv returns true if z is a zero divisor.
func (z *Ultra) IsZeroDiv() bool {
	return z.l.IsZeroDiv()
}

// Inv sets z equal to the inverse of y, and returns z. If y is a zero divisor,
// then Inv panics.
func (z *Ultra) Inv(y *Ultra) *Ultra {
	if y.IsZeroDiv() {
		panic("inverse of zero divisor")
	}
	a := y.Quad()
	a.Inv(a)
	return z.Scal(z.Conj(y), a)
}

// QuoL sets z equal to the left quotient of x and y:
// 		Mul(Inv(y), x)
// Then it returns z. If y is a zero divisor, then QuoL panics.
func (z *Ultra) QuoL(x, y *Ultra) *Ultra {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Mul(z.Inv(y), x)
}

// QuoR sets z equal to the right quotient of x and y:
// 		Mul(x, Inv(y))
// Then it returns z. If y is a zero divisor, then QuoR panics.
func (z *Ultra) QuoR(x, y *Ultra) *Ultra {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Mul(x, z.Inv(y))
}

// Generate returns a random Ultra value for quick.Check testing.
func (z *Ultra) Generate(rand *rand.Rand, size int) reflect.Value {
	randomUltra := &Ultra{
		*NewSupra(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
		*NewSupra(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
	}
	return reflect.ValueOf(randomUltra)
}
