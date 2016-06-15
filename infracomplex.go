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

var symbInfraComplex = [4]string{"", "i", "β", "γ"}

// An InfraComplex represents a rational infra-complex number.
type InfraComplex struct {
	l, r Complex
}

// Real returns the (rational) real part of z.
func (z *InfraComplex) Real() *big.Rat {
	return (&z.l).Real()
}

// Cartesian returns the four rational Cartesian components of z.
func (z *InfraComplex) Cartesian() (*big.Rat, *big.Rat, *big.Rat, *big.Rat) {
	return &z.l.l, &z.l.r, &z.r.l, &z.r.r
}

// String returns the string representation of an InfraComplex value.
//
// If z corresponds to a + bi + cβ + dγ, then the string is"(a+bi+cβ+dγ)",
// similar to complex128 values.
func (z *InfraComplex) String() string {
	v := make([]*big.Rat, 4)
	v[0], v[1] = z.l.Cartesian()
	v[2], v[3] = z.r.Cartesian()
	a := make([]string, 9)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", v[0].RatString())
	i := 1
	for j := 2; j < 8; j = j + 2 {
		if v[i].Sign() < 0 {
			a[j] = fmt.Sprintf("%v", v[i].RatString())
		} else {
			a[j] = fmt.Sprintf("+%v", v[i].RatString())
		}
		a[j+1] = symbInfraComplex[i]
		i++
	}
	a[8] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *InfraComplex) Equals(y *InfraComplex) bool {
	if !z.l.Equals(&y.l) || !z.r.Equals(&y.r) {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *InfraComplex) Set(y *InfraComplex) *InfraComplex {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewInfraComplex returns a pointer to the InfraComplex value a+bi+cβ+dγ.
func NewInfraComplex(a, b, c, d *big.Rat) *InfraComplex {
	z := new(InfraComplex)
	z.l.l.Set(a)
	z.l.r.Set(b)
	z.r.l.Set(c)
	z.r.r.Set(d)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *InfraComplex) Scal(y *InfraComplex, a *big.Rat) *InfraComplex {
	z.l.Scal(&y.l, a)
	z.r.Scal(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *InfraComplex) Neg(y *InfraComplex) *InfraComplex {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *InfraComplex) Conj(y *InfraComplex) *InfraComplex {
	z.l.Conj(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *InfraComplex) Add(x, y *InfraComplex) *InfraComplex {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *InfraComplex) Sub(x, y *InfraComplex) *InfraComplex {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
// 		Mul(i, i) = -1
// 		Mul(β, β) = Mul(γ, γ) = 0
// 		Mul(β, γ) = Mul(γ, β) = 0
// 		Mul(i, β) = -Mul(β, i) = γ
// 		Mul(γ, i) = -Mul(i, γ) = β
// This binary operation is noncommutative but associative.
func (z *InfraComplex) Mul(x, y *InfraComplex) *InfraComplex {
	a := new(Complex).Set(&x.l)
	b := new(Complex).Set(&x.r)
	c := new(Complex).Set(&y.l)
	d := new(Complex).Set(&y.r)
	temp := new(Complex)
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
func (z *InfraComplex) Commutator(x, y *InfraComplex) *InfraComplex {
	return z.Sub(
		z.Mul(x, y),
		new(InfraComplex).Mul(y, x),
	)
}

// Quad returns the quadrance of z. If z = a+bi+cβ+dγ, then the quadrance is
//		Mul(a, a) + Mul(b, b)
// This is always non-negative.
func (z *InfraComplex) Quad() *big.Rat {
	return z.l.Quad()
}

// IsZeroDiv returns true if z is a zero divisor. This is equivalent to z being
// nilpotent.
func (z *InfraComplex) IsZeroDiv() bool {
	zero := new(Complex)
	return z.l.Equals(zero)
}

// Inv sets z equal to the inverse of y, and returns z. If y is a zero divisor,
// then Inv panics.
func (z *InfraComplex) Inv(y *InfraComplex) *InfraComplex {
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
func (z *InfraComplex) QuoL(x, y *InfraComplex) *InfraComplex {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Mul(z.Inv(y), x)
}

// QuoR sets z equal to the right quotient of x and y:
// 		Mul(x, Inv(y))
// Then it returns z. If y is a zero divisor, then QuoR panics.
func (z *InfraComplex) QuoR(x, y *InfraComplex) *InfraComplex {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Mul(x, z.Inv(y))
}

// CrossRatioL sets z equal to the left cross-ratio of v, w, x, and y:
// 		Inv(w - x) * (v - x) * Inv(v - y) * (w - y)
// Then it returns z.
func (z *InfraComplex) CrossRatioL(v, w, x, y *InfraComplex) *InfraComplex {
	temp := new(InfraComplex)
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
func (z *InfraComplex) CrossRatioR(v, w, x, y *InfraComplex) *InfraComplex {
	temp := new(InfraComplex)
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
func (z *InfraComplex) MöbiusL(y, a, b, c, d *InfraComplex) *InfraComplex {
	z.Mul(y, a)
	z.Add(z, b)
	temp := new(InfraComplex)
	temp.Mul(y, c)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(temp, z)
}

// MöbiusR sets z equal to the right Möbius (fractional linear) transform of y:
// 		(a*y + b) * Inv(c*y + d)
// Then it returns z.
func (z *InfraComplex) MöbiusR(y, a, b, c, d *InfraComplex) *InfraComplex {
	z.Mul(a, y)
	z.Add(z, b)
	temp := new(InfraComplex)
	temp.Mul(c, y)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// Generate returns a random InfraComplex value for quick.Check testing.
func (z *InfraComplex) Generate(rand *rand.Rand, size int) reflect.Value {
	randomInfraComplex := &InfraComplex{
		*NewComplex(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
		*NewComplex(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
	}
	return reflect.ValueOf(randomInfraComplex)
}
