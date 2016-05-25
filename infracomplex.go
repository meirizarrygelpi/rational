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

// L returns the left Cayley-Dickson part of z, a pointer to a Complex value.
func (z *InfraComplex) L() *Complex {
	return &z.l
}

// R returns the right Cayley-Dickson part of z, a pointer to a Complex value.
func (z *InfraComplex) R() *Complex {
	return &z.r
}

// SetL sets the left Cayley-Dickson part of z equal to a.
func (z *InfraComplex) SetL(a *Complex) {
	z.l = *a
}

// SetR sets the right Cayley-Dickson part of z equal to b.
func (z *InfraComplex) SetR(b *Complex) {
	z.r = *b
}

// Cartesian returns the four Cartesian components of z.
func (z *InfraComplex) Cartesian() (a, b, c, d *big.Rat) {
	a, b = z.L().Cartesian()
	c, d = z.R().Cartesian()
	return
}

// String returns the string representation of an InfraComplex value.
//
// If z corresponds to a + bi + cβ + dγ, then the string is"(a+bi+cβ+dγ)",
// similar to complex128 values.
func (z *InfraComplex) String() string {
	v := make([]*big.Rat, 4)
	v[0], v[1] = z.L().Cartesian()
	v[2], v[3] = z.R().Cartesian()
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
	if !z.L().Equals(y.L()) || !z.R().Equals(y.R()) {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *InfraComplex) Copy(y *InfraComplex) *InfraComplex {
	z.SetL(y.L())
	z.SetR(y.R())
	return z
}

// NewInfraComplex returns a pointer to an InfraComplex value made from four
// given pointers to big.Rat values.
func NewInfraComplex(a, b, c, d *big.Rat) *InfraComplex {
	z := new(InfraComplex)
	z.SetL(NewComplex(a, b))
	z.SetR(NewComplex(c, d))
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *InfraComplex) Scal(y *InfraComplex, a *big.Rat) *InfraComplex {
	z.SetL(new(Complex).Scal(y.L(), a))
	z.SetR(new(Complex).Scal(y.R(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *InfraComplex) Neg(y *InfraComplex) *InfraComplex {
	z.SetL(new(Complex).Neg(y.L()))
	z.SetR(new(Complex).Neg(y.R()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *InfraComplex) Conj(y *InfraComplex) *InfraComplex {
	z.SetL(new(Complex).Conj(y.L()))
	z.SetR(new(Complex).Neg(y.R()))
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *InfraComplex) Add(x, y *InfraComplex) *InfraComplex {
	z.SetL(new(Complex).Add(x.L(), y.L()))
	z.SetR(new(Complex).Add(x.R(), y.R()))
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *InfraComplex) Sub(x, y *InfraComplex) *InfraComplex {
	z.SetL(new(Complex).Sub(x.L(), y.L()))
	z.SetR(new(Complex).Sub(x.R(), y.R()))
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
	a := new(Complex).Copy(x.L())
	b := new(Complex).Copy(x.R())
	c := new(Complex).Copy(y.L())
	d := new(Complex).Copy(y.R())
	s, t, u := new(Complex), new(Complex), new(Complex)
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
func (z *InfraComplex) Commutator(x, y *InfraComplex) *InfraComplex {
	return z.Sub(
		z.Mul(x, y),
		new(InfraComplex).Mul(y, x),
	)
}

// Quad returns the quadrance of z, a pointer to a big.Rat value.
func (z *InfraComplex) Quad() *big.Rat {
	return z.L().Quad()
}

// IsZeroDiv returns true if z is a zero divisor. This is equivalent to z being
// nilpotent.
func (z *InfraComplex) IsZeroDiv() bool {
	zero := new(Complex)
	return z.L().Equals(zero)
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *InfraComplex) Inv(y *InfraComplex) *InfraComplex {
	if y.IsZeroDiv() {
		panic("inverse of zero divisor")
	}
	return z.Scal(z.Conj(y), new(big.Rat).Inv(y.Quad()))
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *InfraComplex) Quo(x, y *InfraComplex) *InfraComplex {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Mul(x, z.Inv(y))
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
