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

// An Infra represents a rational infra number.
type Infra struct {
	l, r big.Rat
}

// Cartesian returns the two rational Cartesian components of z.
func (z *Infra) Cartesian() (*big.Rat, *big.Rat) {
	return &z.l, &z.r
}

// String returns the string version of a Infra value.
//
// If z corresponds to a + bα, then the string is "(a+bα)", similar to
// complex128 values.
func (z *Infra) String() string {
	a := make([]string, 5)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", z.l.RatString())
	if z.r.Sign() == -1 {
		a[2] = fmt.Sprintf("%v", z.r.RatString())
	} else {
		a[2] = fmt.Sprintf("+%v", z.r.RatString())
	}
	a[3] = "α"
	a[4] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Infra) Equals(y *Infra) bool {
	if z.l.Cmp(&y.l) != 0 || z.r.Cmp(&y.r) != 0 {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *Infra) Set(y *Infra) *Infra {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewInfra returns a pointer to the Infra value a+bα.
func NewInfra(a, b *big.Rat) *Infra {
	z := new(Infra)
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Infra) Scal(y *Infra, a *big.Rat) *Infra {
	z.l.Mul(&y.l, a)
	z.r.Mul(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Infra) Neg(y *Infra) *Infra {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Infra) Conj(y *Infra) *Infra {
	z.l.Set(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *Infra) Add(x, y *Infra) *Infra {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *Infra) Sub(x, y *Infra) *Infra {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z to the product of x and y, and returns z.
//
// The multiplication rule is:
// 		Mul(α, α) = 0
// This binary operation is commutative and associative.
func (z *Infra) Mul(x, y *Infra) *Infra {
	a := new(big.Rat).Set(&x.l)
	b := new(big.Rat).Set(&x.r)
	c := new(big.Rat).Set(&y.l)
	d := new(big.Rat).Set(&y.r)
	temp := new(big.Rat)
	z.l.Mul(a, c)
	z.r.Add(
		z.r.Mul(d, a),
		temp.Mul(b, c),
	)
	return z
}

// Quad returns the quadrance of z. If z = a+bα, then the quadrance is
// 		Mul(a, a)
// This is always non-negative.
func (z *Infra) Quad() *big.Rat {
	return new(big.Rat).Mul(&z.l, &z.l)
}

// IsZeroDiv returns true if z is a zero divisor. This is equivalent to z being
// nilpotent.
func (z *Infra) IsZeroDiv() bool {
	zero := new(big.Int)
	return z.l.Num().Cmp(zero) == 0
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Infra) Inv(y *Infra) *Infra {
	if y.IsZeroDiv() {
		panic("inverse of zero divisor")
	}
	a := y.Quad()
	a.Inv(a)
	return z.Scal(z.Conj(y), a)
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Infra) Quo(x, y *Infra) *Infra {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Mul(x, z.Inv(y))
}

// CrossRatio sets z equal to the cross ratio
// 		Inv(w - x) * (v - x) * Inv(v - y) * (w - y)
// Then it returns z.
func (z *Infra) CrossRatio(v, w, x, y *Infra) *Infra {
	temp := new(Infra)
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

// Möbius sets z equal to the Möbius (fractional linear) transform
// 		(a*y + b) * Inv(c*y + d)
// Then it returns z.
func (z *Infra) Möbius(y, a, b, c, d *Infra) *Infra {
	z.Mul(a, y)
	z.Add(z, b)
	temp := new(Infra)
	temp.Mul(c, y)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// PlusReal sets z equal to y shifted by the rational a, and returns z.
func (z *Infra) PlusReal(y *Infra, a *big.Rat) *Infra {
	z.l.Add(&y.l, a)
	z.r.Set(&y.r)
	return z
}

// PolyEval sets z equal to poly evaluated at y, and returns z.
func (z *Infra) PolyEval(y *Infra, poly Laurent) *Infra {
	neg, nonneg := poly.Degrees()
	n := len(neg)
	nn := len(nonneg)
	rank := n + nn
	temp := new(Infra)
	if rank == 0 {
		z.Set(temp)
		return z
	}
	// zero degree
	if c, ok := poly[0]; ok {
		z.PlusReal(z, c)
	}
	pow := new(Infra)
	// negative degrees
	if n > 0 {
		inv := new(Infra)
		inv.Inv(y)
		pow.Set(inv)
		for d := int64(-1); d > neg[n-1]-1; d-- {
			if c, ok := poly[d]; ok {
				temp.Scal(pow, c)
				z.Add(z, temp)
			}
			pow.Mul(pow, inv)
		}
	}
	// positive degrees
	if nn > 0 {
		pow.Set(y)
		for d := int64(1); d < nonneg[nn-1]+1; d++ {
			if c, ok := poly[d]; ok {
				temp.Scal(pow, c)
				z.Add(z, temp)
			}
			pow.Mul(pow, y)
		}
	}
	return z
}

// Generate returns a random Infra value for quick.Check testing.
func (z *Infra) Generate(rand *rand.Rand, size int) reflect.Value {
	randomInfra := &Infra{
		*big.NewRat(rand.Int63(), rand.Int63()),
		*big.NewRat(rand.Int63(), rand.Int63()),
	}
	return reflect.ValueOf(randomInfra)
}
