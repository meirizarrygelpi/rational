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

var symbTriComplex = [8]string{"", "i", "J", "S", "K", "V", "W", "L"}

// A TriComplex represents a rational tricomplex number.
type TriComplex struct {
	l, r BiComplex
}

// Real returns the (rational) real part of z.
func (z *TriComplex) Real() *big.Rat {
	return (&z.l).Real()
}

// Rats returns the eight rational components of z.
func (z *TriComplex) Rats() (*big.Rat, *big.Rat, *big.Rat, *big.Rat,
	*big.Rat, *big.Rat, *big.Rat, *big.Rat) {
	return &z.l.l.l, &z.l.l.r, &z.l.r.l, &z.l.r.r,
		&z.r.l.l, &z.r.l.r, &z.r.r.l, &z.r.r.r
}

// String returns the string representation of a TriComplex value.
//
// If z corresponds to a + bi + cJ + dS + eK + fV + gW + hL, then the string is
// "(a+bi+cJ+dS+eK+fV+gW+hL)", similar to complex128 values.
func (z *TriComplex) String() string {
	v := make([]*big.Rat, 8)
	v[0], v[1], v[2], v[3] = z.l.Rats()
	v[4], v[5], v[6], v[7] = z.r.Rats()
	a := make([]string, 17)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", v[0].RatString())
	i := 1
	for j := 2; j < 16; j = j + 2 {
		if v[i].Sign() == -1 {
			a[j] = fmt.Sprintf("%v", v[i].RatString())
		} else {
			a[j] = fmt.Sprintf("+%v", v[i].RatString())
		}
		a[j+1] = symbTriComplex[i]
		i++
	}
	a[16] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *TriComplex) Equals(y *TriComplex) bool {
	if !z.l.Equals(&y.l) || !z.r.Equals(&y.r) {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *TriComplex) Set(y *TriComplex) *TriComplex {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewTriComplex returns a *TriComplex with value a+bi+cJ+dS+eK+fV+gW+hL.
func NewTriComplex(a, b, c, d, e, f, g, h *big.Rat) *TriComplex {
	z := new(TriComplex)
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
func (z *TriComplex) Scal(y *TriComplex, a *big.Rat) *TriComplex {
	z.l.Scal(&y.l, a)
	z.r.Scal(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *TriComplex) Neg(y *TriComplex) *TriComplex {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *TriComplex) Conj(y *TriComplex) *TriComplex {
	z.l.Set(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *TriComplex) Add(x, y *TriComplex) *TriComplex {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *TriComplex) Sub(x, y *TriComplex) *TriComplex {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
// 		Mul(i, i) = Mul(J, J) = Mul(K, K) = Mul(L, L) = -1
// 		Mul(S, S) = Mul(V, V) = Mul(W, W) = +1
// 		Mul(i, J) = Mul(J, i) = +S
// 		Mul(i, S) = Mul(S, i) = -J
// 		Mul(i, K) = Mul(K, i) = +V
// 		Mul(i, V) = Mul(V, i) = -K
// 		Mul(i, W) = Mul(W, i) = +L
// 		Mul(i, L) = Mul(L, i) = -W
// 		Mul(J, S) = Mul(S, J) = -i
// 		Mul(J, K) = Mul(K, J) = +W
// 		Mul(J, V) = Mul(V, J) = +L
// 		Mul(J, W) = Mul(W, J) = -K
// 		Mul(J, L) = Mul(L, J) = -V
// 		Mul(S, K) = Mul(K, S) = +L
// 		Mul(S, V) = Mul(V, S) = -W
// 		Mul(S, W) = Mul(W, S) = -V
// 		Mul(S, L) = Mul(L, S) = +K
// 		Mul(K, V) = Mul(V, K) = -i
// 		Mul(K, W) = Mul(W, K) = -J
// 		Mul(K, L) = Mul(L, K) = -S
// 		Mul(V, W) = Mul(W, V) = -S
// 		Mul(V, L) = Mul(L, V) = +J
// 		Mul(W, L) = Mul(L, W) = +i
// This binary operation is commutative and associative.
func (z *TriComplex) Mul(x, y *TriComplex) *TriComplex {
	a := new(BiComplex).Set(&x.l)
	b := new(BiComplex).Set(&x.r)
	c := new(BiComplex).Set(&y.l)
	d := new(BiComplex).Set(&y.r)
	temp := new(BiComplex)
	z.l.Sub(
		z.l.Mul(a, c),
		temp.Mul(b, d),
	)
	z.r.Add(
		z.r.Mul(a, d),
		temp.Mul(b, c),
	)
	return z
}

// Quad returns the quadrance of z. If z = a+bi+cJ+dS, then the quadrance is
// 		a² - b² + c² - d² + 2(ab + cd)i
// Note that this is a bicomplex number.
func (z *TriComplex) Quad() *BiComplex {
	quad := new(BiComplex)
	quad.Mul(&z.l, &z.l)
	return quad.Add(quad, new(BiComplex).Mul(&z.r, &z.r))
}

// Norm returns the norm of z. If z = a+bi+cJ+dS, then the norm is
// 		(a² - b² + c² - d²)² + 4(ab + cd)²
// This can also be written as
// 		((a - d)² + (b + c)²)((a + d)² + (b - c)²)
// The norm is always non-negative.
func (z *TriComplex) Norm() *big.Rat {
	return z.Quad().Quad().Quad()
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *TriComplex) IsZeroDivisor() bool {
	return z.Quad().IsZeroDivisor()
}

// Inv sets z equal to the inverse of y, and returns z. If y is a zero divisor,
// then Inv panics.
func (z *TriComplex) Inv(y *TriComplex) *TriComplex {
	if y.IsZeroDivisor() {
		panic("inverse of zero divisor")
	}
	quad := y.Quad()
	quad.Inv(quad)
	z.Conj(y)
	z.l.Mul(&z.l, quad)
	z.r.Mul(&z.r, quad)
	return z
}

// Quo sets z equal to the quotient of x and y. If y is a zero divisor, then
// Quo panics.
func (z *TriComplex) Quo(x, y *TriComplex) *TriComplex {
	if y.IsZeroDivisor() {
		panic("denominator is zero divisor")
	}
	return z.Mul(z.Inv(y), x)
}

// CrossRatio sets z equal to the cross-ratio of v, w, x, and y:
// 		Inv(w - x) * (v - x) * Inv(v - y) * (w - y)
// Then it returns z.
func (z *TriComplex) CrossRatio(v, w, x, y *TriComplex) *TriComplex {
	temp := new(TriComplex)
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
func (z *TriComplex) Möbius(y, a, b, c, d *TriComplex) *TriComplex {
	z.Mul(a, y)
	z.Add(z, b)
	temp := new(TriComplex)
	temp.Mul(c, y)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// Generate returns a random TriComplex value for quick.Check testing.
func (z *TriComplex) Generate(rand *rand.Rand, size int) reflect.Value {
	randomTriComplex := &TriComplex{
		*NewBiComplex(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
		*NewBiComplex(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
	}
	return reflect.ValueOf(randomTriComplex)
}
