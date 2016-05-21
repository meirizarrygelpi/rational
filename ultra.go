// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"fmt"
	"math/big"
	"strings"
)

var symbUltra = [8]string{"", "α", "β", "γ", "δ", "ε", "ζ", "η"}

// An Ultra represents a rational ultra number.
type Ultra struct {
	l, r *Supra
}

// L returns the left Cayley-Dickson part of z, a pointer to a Supra value.
func (z *Ultra) L() *Supra {
	return z.l
}

// R returns the right Cayley-Dickson part of z, a pointer to a Supra value.
func (z *Ultra) R() *Supra {
	return z.r
}

// SetL sets the left Cayley-Dickson part of z equal to a.
func (z *Ultra) SetL(a *Supra) {
	z.l = a
}

// SetR sets the right Cayley-Dickson part of z equal to b.
func (z *Ultra) SetR(b *Supra) {
	z.r = b
}

// Cartesian returns the eight Cartesian components of z.
func (z *Ultra) Cartesian() (a, b, c, d, e, f, g, h *big.Rat) {
	a, b, c, d = z.L().Cartesian()
	e, f, g, h = z.R().Cartesian()
	return
}

// String returns the string representation of an Ultra value.
//
// If z corresponds to a + bα + cβ + dγ + eδ + fε + gζ + hη, then the string
// is "(a+bα+cβ+dγ+eδ+fε+gζ+hη)", similar to complex128 values.
func (z *Ultra) String() string {
	v := make([]*big.Rat, 8)
	v[0], v[1], v[2], v[3] = z.L().Cartesian()
	v[4], v[5], v[6], v[7] = z.R().Cartesian()
	a := make([]string, 17)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", v[0])
	i := 1
	for j := 2; j < 16; j = j + 2 {
		if v[i].Sign() == -1 {
			a[j] = fmt.Sprintf("%v", v[i])
		} else {
			a[j] = fmt.Sprintf("+%v", v[i])
		}
		a[j+1] = symbUltra[i]
		i++
	}
	a[16] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Ultra) Equals(y *Ultra) bool {
	if !z.L().Equals(y.L()) || !z.R().Equals(y.R()) {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Ultra) Copy(y *Ultra) *Ultra {
	z.SetL(y.L())
	z.SetR(y.R())
	return z
}

// NewUltra returns a pointer to a Ultra value made from eight given pointers
// to big.Rat values.
func NewUltra(a, b, c, d, e, f, g, h *big.Rat) *Ultra {
	z := new(Ultra)
	z.SetL(NewSupra(a, b, c, d))
	z.SetR(NewSupra(e, f, g, h))
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Ultra) Scal(y *Ultra, a *big.Rat) *Ultra {
	z.SetL(new(Supra).Scal(y.L(), a))
	z.SetR(new(Supra).Scal(y.R(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Ultra) Neg(y *Ultra) *Ultra {
	z.SetL(new(Supra).Neg(y.L()))
	z.SetR(new(Supra).Neg(y.R()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Ultra) Conj(y *Ultra) *Ultra {
	z.SetL(new(Supra).Conj(y.L()))
	z.SetR(new(Supra).Neg(y.R()))
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Ultra) Add(x, y *Ultra) *Ultra {
	z.SetL(new(Supra).Add(x.L(), y.L()))
	z.SetR(new(Supra).Add(x.R(), y.R()))
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Ultra) Sub(x, y *Ultra) *Ultra {
	z.SetL(new(Supra).Sub(x.L(), y.L()))
	z.SetR(new(Supra).Sub(x.R(), y.R()))
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
	a, b := x.L(), x.R()
	c, d := y.L(), y.R()
	s, t, u := new(Supra), new(Supra), new(Supra)
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
func (z *Ultra) Commutator(x, y *Ultra) *Ultra {
	return z.Sub(
		z.Mul(x, y),
		new(Ultra).Mul(y, x),
	)
}

// Associator sets z equal to the associator of w, x, and y, and returns z.
func (z *Ultra) Associator(w, x, y *Ultra) *Ultra {
	t := new(Ultra)
	return z.Sub(
		z.Mul(z.Mul(w, x), y),
		t.Mul(w, t.Mul(x, y)),
	)
}

// Quad returns the quadrance of z, a pointer to a big.Rat value.
func (z *Ultra) Quad() *big.Rat {
	return z.L().Quad()
}

// IsZeroDiv returns true if z is a zero divisor.
func (z *Ultra) IsZeroDiv() bool {
	return z.L().IsZeroDiv()
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Ultra) Inv(y *Ultra) *Ultra {
	return z.Scal(z.Conj(y), new(big.Rat).Inv(y.Quad()))
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Ultra) Quo(x, y *Ultra) *Ultra {
	return z.Mul(x, z.Inv(y))
}
