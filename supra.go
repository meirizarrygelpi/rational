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

// Real returns the (rational) real part of z.
func (z *Supra) Real() *big.Rat {
	return (&z.l).Real()
}

// Rats returns the four rational components of z.
func (z *Supra) Rats() (*big.Rat, *big.Rat, *big.Rat, *big.Rat) {
	return &z.l.l, &z.l.r, &z.r.l, &z.r.r
}

// String returns the string representation of a Supra value.
//
// If z corresponds to a + bα + cβ + dγ, then the string is "(a+bα+cβ+dγ)",
// similar to complex128 values.
func (z *Supra) String() string {
	v := make([]*big.Rat, 4)
	v[0], v[1] = z.l.Rats()
	v[2], v[3] = z.r.Rats()
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
	if !z.l.Equals(&y.l) || !z.r.Equals(&y.r) {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *Supra) Set(y *Supra) *Supra {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewSupra returns a pointer to the Supra value a+bα+cβ+dγ.
func NewSupra(a, b, c, d *big.Rat) *Supra {
	z := new(Supra)
	z.l.l.Set(a)
	z.l.r.Set(b)
	z.r.l.Set(c)
	z.r.r.Set(d)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Supra) Scal(y *Supra, a *big.Rat) *Supra {
	z.l.Scal(&y.l, a)
	z.r.Scal(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Supra) Neg(y *Supra) *Supra {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Supra) Conj(y *Supra) *Supra {
	z.l.Conj(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *Supra) Add(x, y *Supra) *Supra {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *Supra) Sub(x, y *Supra) *Supra {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
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
	a := new(Infra).Set(&x.l)
	b := new(Infra).Set(&x.r)
	c := new(Infra).Set(&y.l)
	d := new(Infra).Set(&y.r)
	temp := new(Infra)
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
func (z *Supra) Commutator(x, y *Supra) *Supra {
	return z.Sub(
		z.Mul(x, y),
		new(Supra).Mul(y, x),
	)
}

// Quad returns the quadrance of z. If z = a+bα+cβ+dγ, then the quadrance is
// 		Mul(a, a)
// This is always non-negative.
func (z *Supra) Quad() *big.Rat {
	return z.l.Quad()
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *Supra) IsZeroDivisor() bool {
	return z.l.IsZeroDivisor()
}

// Inv sets z equal to the inverse of y, and returns z. If y is a zero divisor,
// then Inv panics.
func (z *Supra) Inv(y *Supra) *Supra {
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
func (z *Supra) QuoL(x, y *Supra) *Supra {
	if y.IsZeroDivisor() {
		panic("denominator is zero divisor")
	}
	return z.Mul(z.Inv(y), x)
}

// QuoR sets z equal to the right quotient of x and y:
// 		Mul(x, Inv(y))
// Then it returns z. If y is a zero divisor, then QuoR panics.
func (z *Supra) QuoR(x, y *Supra) *Supra {
	if y.IsZeroDivisor() {
		panic("denominator is zero divisor")
	}
	return z.Mul(x, z.Inv(y))
}

// CrossRatioL sets z equal to the left cross-ratio of v, w, x, and y:
// 		Inv(w - x) * (v - x) * Inv(v - y) * (w - y)
// Then it returns z.
func (z *Supra) CrossRatioL(v, w, x, y *Supra) *Supra {
	temp := new(Supra)
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

// CrossRatioR sets z equal to the right cross-ratio of v, w, x, and y:
// 		(v - x) * Inv(w - x) * (w - y) * Inv(v - y)
// Then it returns z.
func (z *Supra) CrossRatioR(v, w, x, y *Supra) *Supra {
	temp := new(Supra)
	z.Sub(v, x)
	temp.Sub(w, x)
	temp.Inv(temp)
	z.Mul(z, temp)
	temp.Sub(w, y)
	z.Mul(z, temp)
	temp.Sub(v, y)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// MöbiusL sets z equal to the left Möbius (fractional linear) transform of y:
// 		Inv(y*c + d) * (y*a + b)
// Then it returns z.
func (z *Supra) MöbiusL(y, a, b, c, d *Supra) *Supra {
	z.Mul(y, a)
	z.Add(z, b)
	temp := new(Supra)
	temp.Mul(y, c)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(temp, z)
}

// MöbiusR sets z equal to the right Möbius (fractional linear) transform of y:
// 		(a*y + b) * Inv(c*y + d)
// Then it returns z.
func (z *Supra) MöbiusR(y, a, b, c, d *Supra) *Supra {
	z.Mul(a, y)
	z.Add(z, b)
	temp := new(Supra)
	temp.Mul(c, y)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(z, temp)
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
