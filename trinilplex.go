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

var symbTriNilplex = [8]string{"", "α", "Γ", "Λ", "Σ", "Φ", "Ψ", "Ω"}

// A TriNilplex represents a rational tricomplex number.
type TriNilplex struct {
	l, r Hyper
}

// Real returns the (rational) real part of z.
func (z *TriNilplex) Real() *big.Rat {
	return (&z.l).Real()
}

// Rats returns the eight rational components of z.
func (z *TriNilplex) Rats() (*big.Rat, *big.Rat, *big.Rat, *big.Rat,
	*big.Rat, *big.Rat, *big.Rat, *big.Rat) {
	return &z.l.l.l, &z.l.l.r, &z.l.r.l, &z.l.r.r,
		&z.r.l.l, &z.r.l.r, &z.r.r.l, &z.r.r.r
}

// String returns the string representation of a TriNilplex value.
//
// If z corresponds to a + bα + cΓ + dΛ + eΣ + fΦ + gΨ + hΩ, then the string is
// "(a+bα+cΓ+dΛ+eΣ+fΦ+gΨ+hΩ)", similar to complex128 values.
func (z *TriNilplex) String() string {
	v := make([]*big.Rat, 8)
	v[0], v[1], v[2], v[3] = z.l.Rats()
	v[4], v[5], v[6], v[7] = z.r.Rats()
	a := make([]string, 17)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", v[0].RatString())
	i := 1
	for j := 2; j < 16; j = j + 2 {
		if v[i].Sign() == -1 {
			a[j] = fmt.Sprintf("%v", v[i].RatString())
		} else {
			a[j] = fmt.Sprintf("+%v", v[i].RatString())
		}
		a[j+1] = symbTriNilplex[i]
		i++
	}
	a[16] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *TriNilplex) Equals(y *TriNilplex) bool {
	if !z.l.Equals(&y.l) || !z.r.Equals(&y.r) {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *TriNilplex) Set(y *TriNilplex) *TriNilplex {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewTriNilplex returns a *TriNilplex with value a+bα+cΓ+dΛ+eΣ+fΦ+gΨ+hΩ.
func NewTriNilplex(a, b, c, d, e, f, g, h *big.Rat) *TriNilplex {
	z := new(TriNilplex)
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
func (z *TriNilplex) Scal(y *TriNilplex, a *big.Rat) *TriNilplex {
	z.l.Scal(&y.l, a)
	z.r.Scal(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *TriNilplex) Neg(y *TriNilplex) *TriNilplex {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *TriNilplex) Conj(y *TriNilplex) *TriNilplex {
	z.l.Set(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *TriNilplex) Add(x, y *TriNilplex) *TriNilplex {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *TriNilplex) Sub(x, y *TriNilplex) *TriNilplex {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
// 		Mul(α, α) = Mul(Γ, Γ) = Mul(Λ, Λ) = 0
// 		Mul(Σ, Σ) = Mul(Φ, Φ) = Mul(Ψ, Ψ) = Mul(Ω, Ω) = 0
// 		Mul(α, Γ) = Mul(Γ, α) = Λ
// 		Mul(α, Λ) = Mul(Λ, α) = 0
// 		Mul(α, Σ) = Mul(Σ, α) = Φ
// 		Mul(α, Φ) = Mul(Φ, α) = 0
// 		Mul(α, Ψ) = Mul(Ψ, α) = Ω
// 		Mul(α, Ω) = Mul(Ω, α) = 0
// 		Mul(Γ, Λ) = Mul(Λ, Γ) = 0
// 		Mul(Γ, Σ) = Mul(Σ, Γ) = Ψ
// 		Mul(Γ, Φ) = Mul(Φ, Γ) = Ω
// 		Mul(Γ, Ψ) = Mul(Ψ, Γ) = 0
// 		Mul(Γ, Ω) = Mul(Ω, Γ) = 0
// 		Mul(Λ, Σ) = Mul(Σ, Λ) = Ω
// 		Mul(Λ, Φ) = Mul(Φ, Λ) = 0
// 		Mul(Λ, Ψ) = Mul(Ψ, Λ) = 0
// 		Mul(Λ, Ω) = Mul(Ω, Λ) = 0
// 		Mul(Σ, Φ) = Mul(Φ, Σ) = 0
// 		Mul(Σ, Ψ) = Mul(Ψ, Σ) = 0
// 		Mul(Σ, Ω) = Mul(Ω, Σ) = 0
// 		Mul(Φ, Ψ) = Mul(Ψ, Φ) = 0
// 		Mul(Φ, Ω) = Mul(Ω, Φ) = 0
// 		Mul(Ψ, Ω) = Mul(Ω, Ψ) = 0
// This binary operation is commutative and associative.
func (z *TriNilplex) Mul(x, y *TriNilplex) *TriNilplex {
	a := new(Hyper).Set(&x.l)
	b := new(Hyper).Set(&x.r)
	c := new(Hyper).Set(&y.l)
	d := new(Hyper).Set(&y.r)
	temp := new(Hyper)
	z.l.Mul(a, c)
	z.r.Add(
		z.r.Mul(d, a),
		temp.Mul(b, c),
	)
	return z
}

// Quad returns the quadrance of z. If z = a+bi+cΓ+dΛ, then the quadrance is
// 		a² - b² + c² - d² + 2(ab + cd)i
// Note that this is a bicomplex number.
func (z *TriNilplex) Quad() *Hyper {
	quad := new(Hyper)
	return quad.Mul(&z.l, &z.l)
}

// Norm returns the norm of z. If z = a+bα+cΓ+dΛ+eΣ+fΦ+gΨ+hΩ, then the norm is
// 		((a²)²)²
// The norm is always non-negative.
func (z *TriNilplex) Norm() *big.Rat {
	return z.Quad().Quad().Quad()
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *TriNilplex) IsZeroDivisor() bool {
	return z.Quad().IsZeroDivisor()
}

// Inv sets z equal to the inverse of y, and returns z. If y is a zero divisor,
// then Inv panics.
func (z *TriNilplex) Inv(y *TriNilplex) *TriNilplex {
	if y.IsZeroDivisor() {
		panic("inverse of zero divisor")
	}
	quad := y.Quad()
	quad.Inv(quad)
	z.Conj(y)
	z.l.Mul(&z.l, quad)
	z.r.Mul(&z.r, quad)
	return z
}

// Quo sets z equal to the quotient of x and y. If y is a zero divisor, then
// Quo panics.
func (z *TriNilplex) Quo(x, y *TriNilplex) *TriNilplex {
	if y.IsZeroDivisor() {
		panic("denominator is zero divisor")
	}
	return z.Mul(z.Inv(y), x)
}

// CrossRatio sets z equal to the cross-ratio of v, w, x, and y:
// 		Inv(w - x) * (v - x) * Inv(v - y) * (w - y)
// Then it returns z.
func (z *TriNilplex) CrossRatio(v, w, x, y *TriNilplex) *TriNilplex {
	temp := new(TriNilplex)
	z.Sub(w, x)
	z.Inv(z)
	temp.Sub(v, x)
	z.Mul(z, temp)
	temp.Sub(v, y)
	temp.Inv(temp)
	z.Mul(z, temp)
	temp.Sub(w, y)
	return z.Mul(z, temp)
}

// Möbius sets z equal to the Möbius (fractional linear) transform of y:
// 		(a*y + b) * Inv(c*y + d)
// Then it returns z.
func (z *TriNilplex) Möbius(y, a, b, c, d *TriNilplex) *TriNilplex {
	z.Mul(a, y)
	z.Add(z, b)
	temp := new(TriNilplex)
	temp.Mul(c, y)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// Generate returns a random TriNilplex value for quick.Check testing.
func (z *TriNilplex) Generate(rand *rand.Rand, size int) reflect.Value {
	randomTriNilplex := &TriNilplex{
		*NewHyper(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
		*NewHyper(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
	}
	return reflect.ValueOf(randomTriNilplex)
}
