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

// A Perplex represents a rational split-complex number.
type Perplex struct {
	l, r big.Rat
}

// Cartesian returns the two rational Cartesian components of z.
func (z *Perplex) Cartesian() (*big.Rat, *big.Rat) {
	return &z.l, &z.r
}

// String returns the string version of a Perplex value.
//
// If z corresponds to a + bs, then the string is "(a+bs)", similar to
// complex128 values.
func (z *Perplex) String() string {
	a := make([]string, 5)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", z.l.RatString())
	if z.r.Sign() == -1 {
		a[2] = fmt.Sprintf("%v", z.r.RatString())
	} else {
		a[2] = fmt.Sprintf("+%v", z.r.RatString())
	}
	a[3] = "s"
	a[4] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Perplex) Equals(y *Perplex) bool {
	if z.l.Cmp(&y.l) != 0 || z.r.Cmp(&y.r) != 0 {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *Perplex) Set(y *Perplex) *Perplex {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewPerplex returns a pointer to a Perplex value made from two given pointers
// to big.Rat values.
func NewPerplex(a, b *big.Rat) *Perplex {
	z := new(Perplex)
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Perplex) Scal(y *Perplex, a *big.Rat) *Perplex {
	z.l.Mul(&y.l, a)
	z.r.Mul(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Perplex) Neg(y *Perplex) *Perplex {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Perplex) Conj(y *Perplex) *Perplex {
	z.l.Set(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Perplex) Add(x, y *Perplex) *Perplex {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Perplex) Sub(x, y *Perplex) *Perplex {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rule is:
// 		Mul(s, s) = +1
// This binary operation is commutative and associative.
func (z *Perplex) Mul(x, y *Perplex) *Perplex {
	a := new(big.Rat).Set(&x.l)
	b := new(big.Rat).Set(&x.r)
	c := new(big.Rat).Set(&y.l)
	d := new(big.Rat).Set(&y.r)
	temp := new(big.Rat)
	z.l.Add(
		z.l.Mul(a, c),
		temp.Mul(d, b),
	)
	z.r.Add(
		z.r.Mul(d, a),
		temp.Mul(b, c),
	)
	return z
}

// Quad returns the quadrance of z. If z = a+bs, then the quadrance is
// 		Mul(a, a) - Mul(b, b)
// This can be positive, negative, or zero.
func (z *Perplex) Quad() *big.Rat {
	quad := new(big.Rat)
	return quad.Sub(
		quad.Mul(&z.l, &z.l),
		new(big.Rat).Mul(&z.r, &z.r),
	)
}

// IsZeroDiv returns true if z is a zero divisor.
func (z *Perplex) IsZeroDiv() bool {
	if z.l.Cmp(&z.r) == 0 {
		return true
	}
	if z.l.Cmp(new(big.Rat).Neg(&z.r)) == 0 {
		return true
	}
	return false
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Perplex) Inv(y *Perplex) *Perplex {
	if y.IsZeroDiv() {
		panic("inverse of zero divisor")
	}
	a := y.Quad()
	a.Inv(a)
	return z.Scal(z.Conj(y), a)
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Perplex) Quo(x, y *Perplex) *Perplex {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Mul(x, z.Inv(y))
}

// Idempotent sets z equal to a pointer to an idempotent Perplex.
func (z *Perplex) Idempotent(sign int) *Perplex {
	z.l.SetFrac64(1, 2)
	if sign < 0 {
		z.r.SetFrac64(-1, 2)
		return z
	}
	z.r.SetFrac64(1, 2)
	return z
}

// CrossRatio sets z equal to the cross ratio
// 		Inv(w - x) * (v - x) * Inv(v - y) * (w - y)
// Then it returns z.
func (z *Perplex) CrossRatio(v, w, x, y *Perplex) *Perplex {
	temp := new(Perplex)
	z.Sub(w, x)
	z.Inv(z)
	temp.Sub(v, x)
	z.Mul(z, temp)
	temp.Sub(v, y)
	temp.Inv(temp)
	z.Mul(z, temp)
	temp.Sub(w, y)
	z.Mul(z, temp)
	return z
}

// Möbius sets z equal to the Möbius (fractional linear) transform
// 		(a*y + b) * Inv(c*y + d)
// Then it returns z.
func (z *Perplex) Möbius(y, a, b, c, d *Perplex) *Perplex {
	z.Mul(a, y)
	z.Add(z, b)
	temp := new(Perplex)
	temp.Mul(c, y)
	temp.Add(temp, d)
	temp.Inv(temp)
	z.Mul(z, temp)
	return z
}

// Generate returns a random Perplex value for quick.Check testing.
func (z *Perplex) Generate(rand *rand.Rand, size int) reflect.Value {
	randomPerplex := &Perplex{
		*big.NewRat(rand.Int63(), rand.Int63()),
		*big.NewRat(rand.Int63(), rand.Int63()),
	}
	return reflect.ValueOf(randomPerplex)
}
