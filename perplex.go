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

// Real returns the (rational) real part of z.
func (z *Perplex) Real() *big.Rat {
	return &z.l
}

// Rats returns the two rational components of z.
func (z *Perplex) Rats() (*big.Rat, *big.Rat) {
	return &z.l, &z.r
}

// String returns the string version of a Perplex value.
//
// If z corresponds to a + bs, then the string is "(a+bs)", similar to
// complex128 values.
func (z *Perplex) String() string {
	a := make([]string, 5)
	a[0] = leftBracket
	a[1] = fmt.Sprintf("%v", z.l.RatString())
	if z.r.Sign() == -1 {
		a[2] = fmt.Sprintf("%v", z.r.RatString())
	} else {
		a[2] = fmt.Sprintf("+%v", z.r.RatString())
	}
	a[3] = "s"
	a[4] = rightBracket
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

// NewPerplex returns a pointer to the Perplex value a+bs.
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

// Add sets z equal to x+y, and returns z.
func (z *Perplex) Add(x, y *Perplex) *Perplex {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
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
// 		a² - b²
// This can be positive, negative, or zero.
func (z *Perplex) Quad() *big.Rat {
	quad := new(big.Rat)
	return quad.Sub(
		quad.Mul(&z.l, &z.l),
		new(big.Rat).Mul(&z.r, &z.r),
	)
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *Perplex) IsZeroDivisor() bool {
	if z.l.Cmp(&z.r) == 0 {
		return true
	}
	if z.l.Cmp(new(big.Rat).Neg(&z.r)) == 0 {
		return true
	}
	return false
}

// Inv sets z equal to the inverse of y, and returns z. If y is a zero divisor,
// then Inv panics.
func (z *Perplex) Inv(y *Perplex) *Perplex {
	if y.IsZeroDivisor() {
		panic("inverse of zero divisor")
	}
	a := y.Quad()
	a.Inv(a)
	return z.Scal(z.Conj(y), a)
}

// Quo sets z equal to the quotient of x and y, and returns z. If y is a zero
// divisor, then Quo panics.
func (z *Perplex) Quo(x, y *Perplex) *Perplex {
	if y.IsZeroDivisor() {
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

// CrossRatio sets z equal to the cross-ratio of v, w, x, and y:
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
	return z.Mul(z, temp)
}

// Möbius sets z equal to the Möbius (fractional linear) transform of y:
// 		(a*y + b) * Inv(c*y + d)
// Then it returns z.
func (z *Perplex) Möbius(y, a, b, c, d *Perplex) *Perplex {
	z.Mul(a, y)
	z.Add(z, b)
	temp := new(Perplex)
	temp.Mul(c, y)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// Plus sets z equal to y shifted by the rational a, and returns z.
func (z *Perplex) Plus(y *Perplex, a *big.Rat) *Perplex {
	z.l.Add(&y.l, a)
	z.r.Set(&y.r)
	return z
}

// PolyEval sets z equal to poly evaluated at y, and returns z.
func (z *Perplex) PolyEval(y *Perplex, poly Laurent) *Perplex {
	neg, nonneg := poly.Degrees()
	n := len(neg)
	nn := len(nonneg)
	rank := n + nn
	temp := new(Perplex)
	if rank == 0 {
		z.Set(temp)
		return z
	}
	// zero degree
	if c, ok := poly[0]; ok {
		z.Plus(z, c)
	}
	pow := new(Perplex)
	// negative degrees
	if n > 0 {
		inv := new(Perplex)
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

// Dot returns the (rational) dot product of z and y.
func (z *Perplex) Dot(y *Perplex) *big.Rat {
	dot := new(big.Rat)
	temp := new(big.Rat)
	dot.Mul(&z.l, &y.l)
	return dot.Sub(dot, temp.Mul(&z.r, &z.r))
}

// Generate returns a random Perplex value for quick.Check testing.
func (z *Perplex) Generate(rand *rand.Rand, size int) reflect.Value {
	randomPerplex := &Perplex{
		*big.NewRat(rand.Int63(), rand.Int63()),
		*big.NewRat(rand.Int63(), rand.Int63()),
	}
	return reflect.ValueOf(randomPerplex)
}
