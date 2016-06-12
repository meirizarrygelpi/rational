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

// A Complex represents a rational complex number.
type Complex struct {
	l, r big.Rat
}

// Cartesian returns the two rational Cartesian components of z.
func (z *Complex) Cartesian() (*big.Rat, *big.Rat) {
	return &z.l, &z.r
}

// String returns the string version of a Complex value.
//
// If z corresponds to a + bi, then the string is "(a+bi)", similar to
// complex128 values.
func (z *Complex) String() string {
	a := make([]string, 5)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", z.l.RatString())
	if z.r.Sign() < 0 {
		a[2] = fmt.Sprintf("%v", z.r.RatString())
	} else {
		a[2] = fmt.Sprintf("+%v", z.r.RatString())
	}
	a[3] = "i"
	a[4] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Complex) Equals(y *Complex) bool {
	if z.l.Cmp(&y.l) != 0 || z.r.Cmp(&y.r) != 0 {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *Complex) Set(y *Complex) *Complex {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewComplex returns a pointer to the Complex value a+bi.
func NewComplex(a, b *big.Rat) *Complex {
	z := new(Complex)
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Complex) Scal(y *Complex, a *big.Rat) *Complex {
	z.l.Mul(&y.l, a)
	z.r.Mul(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Complex) Neg(y *Complex) *Complex {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Complex) Conj(y *Complex) *Complex {
	z.l.Set(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *Complex) Add(x, y *Complex) *Complex {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *Complex) Sub(x, y *Complex) *Complex {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rule is:
// 		Mul(i, i) = -1
// This binary operation is commutative and associative.
func (z *Complex) Mul(x, y *Complex) *Complex {
	a := new(big.Rat).Set(&x.l)
	b := new(big.Rat).Set(&x.r)
	c := new(big.Rat).Set(&y.l)
	d := new(big.Rat).Set(&y.r)
	temp := new(big.Rat)
	z.l.Sub(
		z.l.Mul(a, c),
		temp.Mul(d, b),
	)
	z.r.Add(
		z.r.Mul(d, a),
		temp.Mul(b, c),
	)
	return z
}

// Quad returns the quadrance of z. If z = a+bi, then the quadrance is
// 		Mul(a, a) + Mul(b, b)
// This is always non-negative.
func (z *Complex) Quad() *big.Rat {
	quad := new(big.Rat)
	return quad.Add(
		quad.Mul(&z.l, &z.l),
		new(big.Rat).Mul(&z.r, &z.r),
	)
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Complex) Inv(y *Complex) *Complex {
	if zero := new(Complex); y.Equals(zero) {
		panic("zero inverse")
	}
	a := y.Quad()
	a.Inv(a)
	return z.Scal(z.Conj(y), a)
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Complex) Quo(x, y *Complex) *Complex {
	if zero := new(Complex); y.Equals(zero) {
		panic("zero denominator")
	}
	return z.Mul(x, z.Inv(y))
}

// Gauss sets z equal to the Gaussian integer a+bi, and returns z.
func (z *Complex) Gauss(a, b *big.Int) *Complex {
	z.l.SetInt(a)
	z.r.SetInt(b)
	return z
}

// CrossRatio sets z equal to the cross ratio
// 		Inv(w - x) * (v - x) * Inv(v - y) * (w - y)
// Then it returns z.
func (z *Complex) CrossRatio(v, w, x, y *Complex) *Complex {
	temp := new(Complex)
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
func (z *Complex) Möbius(y, a, b, c, d *Complex) *Complex {
	z.Mul(a, y)
	z.Add(z, b)
	temp := new(Complex)
	temp.Mul(c, y)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// PlusReal sets z equal to y shifted by the rational a, and returns z.
func (z *Complex) PlusReal(y *Complex, a *big.Rat) *Complex {
	z.l.Add(&y.l, a)
	z.r.Set(&y.r)
	return z
}

// PolyEval sets z equal to poly evaluated at y, and returns z.
func (z *Complex) PolyEval(y *Complex, poly Laurent) *Complex {
	neg, nonneg := poly.Degrees()
	n := len(neg)
	nn := len(nonneg)
	rank := n + nn
	temp := new(Complex)
	if rank == 0 {
		z.Set(temp)
		return z
	}
	// zero degree
	if c, ok := poly[0]; ok {
		z.PlusReal(z, c)
	}
	pow := new(Complex)
	// negative degrees
	if n > 0 {
		inv := new(Complex)
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

// Generate returns a random Complex value for quick.Check testing.
func (z *Complex) Generate(rand *rand.Rand, size int) reflect.Value {
	randomComplex := &Complex{
		*big.NewRat(rand.Int63(), rand.Int63()),
		*big.NewRat(rand.Int63(), rand.Int63()),
	}
	return reflect.ValueOf(randomComplex)
}
