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

// An InfraCockle represents a rational infra-Cockle quaternion.
type InfraCockle struct {
	l, r Cockle
}

// Real returns the (rational) real part of z.
func (z *InfraCockle) Real() *big.Rat {
	return (&z.l).Real()
}

// Cartesian returns the eight rational Cartesian components of z.
func (z *InfraCockle) Cartesian() (*big.Rat, *big.Rat, *big.Rat, *big.Rat,
	*big.Rat, *big.Rat, *big.Rat, *big.Rat) {
	return &z.l.l.l, &z.l.l.r, &z.l.r.l, &z.l.r.r,
		&z.r.l.l, &z.r.l.r, &z.r.r.l, &z.r.r.r
}

// String returns the string representation of an InfraCockle value.
//
// If z corresponds to a + bi + ct + du + eρ + fσ + gτ + hυ, then the string
// is"(a+bi+ct+du+eρ+fσ+gτ+hυ)", similar to complex128 values.
func (z *InfraCockle) String() string {
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
		a[j+1] = symbInfraCockle[i]
		i++
	}
	a[16] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *InfraCockle) Equals(y *InfraCockle) bool {
	if !z.l.Equals(&y.l) || !z.r.Equals(&y.r) {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *InfraCockle) Set(y *InfraCockle) *InfraCockle {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewInfraCockle returns a pointer to the InfraCockle value
// a+bi+ct+du+eρ+fσ+gτ+hυ.
func NewInfraCockle(a, b, c, d, e, f, g, h *big.Rat) *InfraCockle {
	z := new(InfraCockle)
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
func (z *InfraCockle) Scal(y *InfraCockle, a *big.Rat) *InfraCockle {
	z.l.Scal(&y.l, a)
	z.r.Scal(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *InfraCockle) Neg(y *InfraCockle) *InfraCockle {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *InfraCockle) Conj(y *InfraCockle) *InfraCockle {
	z.l.Conj(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *InfraCockle) Add(x, y *InfraCockle) *InfraCockle {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *InfraCockle) Sub(x, y *InfraCockle) *InfraCockle {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
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
// This binary operation is noncommutative and nonassociative.
func (z *InfraCockle) Mul(x, y *InfraCockle) *InfraCockle {
	a := new(Cockle).Set(&x.l)
	b := new(Cockle).Set(&x.r)
	c := new(Cockle).Set(&y.l)
	d := new(Cockle).Set(&y.r)
	temp := new(Cockle)
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
func (z *InfraCockle) Commutator(x, y *InfraCockle) *InfraCockle {
	return z.Sub(
		z.Mul(x, y),
		new(InfraCockle).Mul(y, x),
	)
}

// Associator sets z equal to the associator of w, x, and y:
// 		Mul(Mul(w, x), y) - Mul(w, Mul(x, y))
// Then it returns z.
func (z *InfraCockle) Associator(w, x, y *InfraCockle) *InfraCockle {
	temp := new(InfraCockle)
	return z.Sub(
		z.Mul(z.Mul(w, x), y),
		temp.Mul(w, temp.Mul(x, y)),
	)
}

// Quad returns the quadrance of z. If z = a+bi+ct+du+eρ+fσ+gτ+hυ, then the
// quadrance is
// 		Mul(a, a) + Mul(b, b) - Mul(c, c) - Mul(d, d)
// This can be positive, negative, or zero.
func (z *InfraCockle) Quad() *big.Rat {
	return z.l.Quad()
}

// IsZeroDiv returns true if z is a zero divisor. This is equivalent to z being
// nilpotent.
func (z *InfraCockle) IsZeroDiv() bool {
	return z.l.IsZeroDiv()
}

// Inv sets z equal to the inverse of y, and returns z. If y is a zero divisor,
// then Inv panics.
func (z *InfraCockle) Inv(y *InfraCockle) *InfraCockle {
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
func (z *InfraCockle) QuoL(x, y *InfraCockle) *InfraCockle {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Mul(z.Inv(y), x)
}

// QuoR sets z equal to the right quotient of x and y:
// 		Mul(x, Inv(y))
// Then it returns z. If y is a zero divisor, then QuoR panics.
func (z *InfraCockle) QuoR(x, y *InfraCockle) *InfraCockle {
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
