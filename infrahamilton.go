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

// An InfraHamilton represents a rational infra-Hamilton quaternion.
type InfraHamilton struct {
	l, r Hamilton
}

// Real returns the (rational) real part of z.
func (z *InfraHamilton) Real() *big.Rat {
	return (&z.l).Real()
}

// Rats returns the eight components of z.
func (z *InfraHamilton) Rats() (*big.Rat, *big.Rat, *big.Rat, *big.Rat,
	*big.Rat, *big.Rat, *big.Rat, *big.Rat) {
	return &z.l.l.l, &z.l.l.r, &z.l.r.l, &z.l.r.r,
		&z.r.l.l, &z.r.l.r, &z.r.r.l, &z.r.r.r
}

// String returns the string representation of an InfraHamilton value.
//
// If z corresponds to a + bi + cj + dk + eα + fβ + gγ + hδ, then the string
// is"(a+bi+cj+dk+eα+fβ+gγ+hδ)", similar to complex128 values.
func (z *InfraHamilton) String() string {
	v := make([]*big.Rat, 8)
	v[0], v[1], v[2], v[3] = z.l.Rats()
	v[4], v[5], v[6], v[7] = z.r.Rats()
	a := make([]string, 17)
	a[0] = leftBracket
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
	a[16] = rightBracket
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *InfraHamilton) Equals(y *InfraHamilton) bool {
	if !z.l.Equals(&y.l) || !z.r.Equals(&y.r) {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *InfraHamilton) Set(y *InfraHamilton) *InfraHamilton {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewInfraHamilton returns a pointer to the InfraHamilton value
// a+bi+cj+dk+eα+fβ+gγ+hδ.
func NewInfraHamilton(a, b, c, d, e, f, g, h *big.Rat) *InfraHamilton {
	z := new(InfraHamilton)
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
func (z *InfraHamilton) Scal(y *InfraHamilton, a *big.Rat) *InfraHamilton {
	z.l.Scal(&y.l, a)
	z.r.Scal(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *InfraHamilton) Neg(y *InfraHamilton) *InfraHamilton {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *InfraHamilton) Conj(y *InfraHamilton) *InfraHamilton {
	z.l.Conj(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *InfraHamilton) Add(x, y *InfraHamilton) *InfraHamilton {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *InfraHamilton) Sub(x, y *InfraHamilton) *InfraHamilton {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
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
// This binary operation is noncommutative and nonassociative.
func (z *InfraHamilton) Mul(x, y *InfraHamilton) *InfraHamilton {
	a := new(Hamilton).Set(&x.l)
	b := new(Hamilton).Set(&x.r)
	c := new(Hamilton).Set(&y.l)
	d := new(Hamilton).Set(&y.r)
	temp := new(Hamilton)
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
func (z *InfraHamilton) Commutator(x, y *InfraHamilton) *InfraHamilton {
	return z.Sub(
		z.Mul(x, y),
		new(InfraHamilton).Mul(y, x),
	)
}

// Associator sets z equal to the associator of w, x, and y:
// 		Mul(Mul(w, x), y) - Mul(w, Mul(x, y))
// Then it returns z.
func (z *InfraHamilton) Associator(w, x, y *InfraHamilton) *InfraHamilton {
	temp := new(InfraHamilton)
	return z.Sub(
		z.Mul(z.Mul(w, x), y),
		temp.Mul(w, temp.Mul(x, y)),
	)
}

// Quad returns the quadrance of z. If z = a+bi+cj+dk+eα+fβ+gγ+hδ, then the
// quadrance is
//		a² + b² + c² + d²
// This is always non-negative.
func (z *InfraHamilton) Quad() *big.Rat {
	return z.l.Quad()
}

// IsZeroDivisor returns true if z is a zero divisor. This is equivalent to z being
// nilpotent.
func (z *InfraHamilton) IsZeroDivisor() bool {
	zero := new(Hamilton)
	return z.l.Equals(zero)
}

// Inv sets z equal to the inverse of y, and returns z. If y is a zero divisor,
// then Inv panics.
func (z *InfraHamilton) Inv(y *InfraHamilton) *InfraHamilton {
	if y.IsZeroDivisor() {
		panic("inverse of zero divisor")
	}
	a := y.Quad()
	a.Inv(a)
	return z.Scal(z.Conj(y), a)
}

// QuoL sets z equal to the left quotient of x and y:
// 		Mul(Inv(y), x)
// Then it returns z. If y is a zero divisor, then QuoL panics.
func (z *InfraHamilton) QuoL(x, y *InfraHamilton) *InfraHamilton {
	if y.IsZeroDivisor() {
		panic("denominator is zero divisor")
	}
	return z.Mul(z.Inv(y), x)
}

// QuoR sets z equal to the right quotient of x and y:
// 		Mul(x, Inv(y))
// Then it returns z. If y is a zero divisor, then QuoR panics.
func (z *InfraHamilton) QuoR(x, y *InfraHamilton) *InfraHamilton {
	if y.IsZeroDivisor() {
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
