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

var symbInfraHamilton = [8]string{"", "i", "j", "k", "α", "β", "γ", "δ"}

// An InfraHamilton represents a rational infra-complex number.
type InfraHamilton struct {
	l, r Hamilton
}

// L returns the left Cayley-Dickson part of z, a pointer to a Hamilton value.
func (z *InfraHamilton) L() *Hamilton {
	return &z.l
}

// R returns the right Cayley-Dickson part of z, a pointer to a Hamilton value.
func (z *InfraHamilton) R() *Hamilton {
	return &z.r
}

// SetL sets the left Cayley-Dickson part of z equal to a.
func (z *InfraHamilton) SetL(a *Hamilton) {
	z.l = *a
}

// SetR sets the right Cayley-Dickson part of z equal to b.
func (z *InfraHamilton) SetR(b *Hamilton) {
	z.r = *b
}

// Cartesian returns the eight Cartesian components of z.
func (z *InfraHamilton) Cartesian() (a, b, c, d, e, f, g, h *big.Rat) {
	a, b, c, d = z.L().Cartesian()
	e, f, g, h = z.R().Cartesian()
	return
}

// String returns the string representation of an InfraHamilton value.
//
// If z corresponds to a + bi + cj + dk + eα + fβ + gγ + hδ, then the string
// is"(a+bi+cj+dk+eα+fβ+gγ+hδ)", similar to complex128 values.
func (z *InfraHamilton) String() string {
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
		a[j+1] = symbInfraHamilton[i]
		i++
	}
	a[16] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *InfraHamilton) Equals(y *InfraHamilton) bool {
	if !z.L().Equals(y.L()) || !z.R().Equals(y.R()) {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *InfraHamilton) Copy(y *InfraHamilton) *InfraHamilton {
	z.SetL(y.L())
	z.SetR(y.R())
	return z
}

// NewInfraHamilton returns a pointer to an InfraHamilton value made from eight
// given pointers to big.Rat values.
func NewInfraHamilton(a, b, c, d, e, f, g, h *big.Rat) *InfraHamilton {
	z := new(InfraHamilton)
	z.SetL(NewHamilton(a, b, c, d))
	z.SetR(NewHamilton(e, f, g, h))
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *InfraHamilton) Scal(y *InfraHamilton, a *big.Rat) *InfraHamilton {
	z.SetL(new(Hamilton).Scal(y.L(), a))
	z.SetR(new(Hamilton).Scal(y.R(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *InfraHamilton) Neg(y *InfraHamilton) *InfraHamilton {
	z.SetL(new(Hamilton).Neg(y.L()))
	z.SetR(new(Hamilton).Neg(y.R()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *InfraHamilton) Conj(y *InfraHamilton) *InfraHamilton {
	z.SetL(new(Hamilton).Conj(y.L()))
	z.SetR(new(Hamilton).Neg(y.R()))
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *InfraHamilton) Add(x, y *InfraHamilton) *InfraHamilton {
	z.SetL(new(Hamilton).Add(x.L(), y.L()))
	z.SetR(new(Hamilton).Add(x.R(), y.R()))
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *InfraHamilton) Sub(x, y *InfraHamilton) *InfraHamilton {
	z.SetL(new(Hamilton).Sub(x.L(), y.L()))
	z.SetR(new(Hamilton).Sub(x.R(), y.R()))
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
//		Mul(i, i) = Mul(j, j) = Mul(k, k) = -1
// 		Mul(α, α) = Mul(β, β) = Mul(γ, γ) = Mul(δ, δ) = 0
// 		Mul(i, j) = -Mul(j, i) = +k
// 		Mul(i, k) = -Mul(k, i) = -j
// 		Mul(i, α) = -Mul(α, i) = +β
// 		Mul(i, β) = -Mul(β, i) = -α
// 		Mul(i, γ) = -Mul(γ, i) = -δ
// 		Mul(i, δ) = -Mul(δ, i) = +γ
// 		Mul(j, k) = -Mul(k, j) = +i
// 		Mul(j, α) = -Mul(α, j) = +γ
// 		Mul(j, β) = -Mul(β, j) = +δ
// 		Mul(j, γ) = -Mul(γ, j) = -α
// 		Mul(j, δ) = -Mul(δ, j) = -β
// 		Mul(k, α) = -Mul(α, k) = +δ
// 		Mul(k, β) = -Mul(β, k) = -γ
// 		Mul(k, γ) = -Mul(γ, k) = +β
// 		Mul(k, δ) = -Mul(δ, k) = -α
// 		Mul(α, β) = Mul(β, α) = 0
// 		Mul(α, γ) = Mul(γ, α) = 0
// 		Mul(α, δ) = Mul(δ, α) = 0
// 		Mul(β, γ) = Mul(γ, β) = 0
// 		Mul(β, δ) = Mul(δ, β) = 0
// 		Mul(γ, δ) = Mul(δ, γ) = 0
// This binary operation is noncommutative but associative.
func (z *InfraHamilton) Mul(x, y *InfraHamilton) *InfraHamilton {
	a := new(Hamilton).Copy(x.L())
	b := new(Hamilton).Copy(x.R())
	c := new(Hamilton).Copy(y.L())
	d := new(Hamilton).Copy(y.R())
	s, t, u := new(Hamilton), new(Hamilton), new(Hamilton)
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
func (z *InfraHamilton) Commutator(x, y *InfraHamilton) *InfraHamilton {
	return z.Sub(
		z.Mul(x, y),
		new(InfraHamilton).Mul(y, x),
	)
}

// Associator sets z equal to the associator of w, x, and y, and returns z.
func (z *InfraHamilton) Associator(w, x, y *InfraHamilton) *InfraHamilton {
	t := new(InfraHamilton)
	return z.Sub(
		z.Mul(z.Mul(w, x), y),
		t.Mul(w, t.Mul(x, y)),
	)
}

// Quad returns the quadrance of z, a pointer to a big.Rat value.
func (z *InfraHamilton) Quad() *big.Rat {
	return z.L().Quad()
}

// IsZeroDiv returns true if z is a zero divisor. This is equivalent to z being
// nilpotent.
func (z *InfraHamilton) IsZeroDiv() bool {
	zero := new(Hamilton)
	return z.L().Equals(zero)
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *InfraHamilton) Inv(y *InfraHamilton) *InfraHamilton {
	if y.IsZeroDiv() {
		panic("inverse of zero divisor")
	}
	return z.Scal(z.Conj(y), new(big.Rat).Inv(y.Quad()))
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *InfraHamilton) Quo(x, y *InfraHamilton) *InfraHamilton {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Mul(x, z.Inv(y))
}

// Generate returns a random InfraHamilton value for quick.Check testing.
func (z *InfraHamilton) Generate(rand *rand.Rand, size int) reflect.Value {
	randomInfraHamilton := &InfraHamilton{
		*NewHamilton(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
		*NewHamilton(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
	}
	return reflect.ValueOf(randomInfraHamilton)
}
