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

var symbTriPerplex = [8]string{"", "s", "T", "U", "V", "W", "X", "Y"}

// A TriPerplex represents a rational triperplex number.
type TriPerplex struct {
	l, r BiPerplex
}

// Real returns the (rational) real part of z.
func (z *TriPerplex) Real() *big.Rat {
	return (&z.l).Real()
}

// Rats returns the eight rational components of z.
func (z *TriPerplex) Rats() (*big.Rat, *big.Rat, *big.Rat, *big.Rat,
	*big.Rat, *big.Rat, *big.Rat, *big.Rat) {
	return &z.l.l.l, &z.l.l.r, &z.l.r.l, &z.l.r.r,
		&z.r.l.l, &z.r.l.r, &z.r.r.l, &z.r.r.r
}

// String returns the string representation of a TriPerplex value.
//
// If z corresponds to a + bs + cT + dU + eV + fW + gX + hY, then the string is
// "(a+bs+cT+dU+eV+fW+gX+hY)", similar to complex128 values.
func (z *TriPerplex) String() string {
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
		a[j+1] = symbTriPerplex[i]
		i++
	}
	a[16] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *TriPerplex) Equals(y *TriPerplex) bool {
	if !z.l.Equals(&y.l) || !z.r.Equals(&y.r) {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *TriPerplex) Set(y *TriPerplex) *TriPerplex {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewTriPerplex returns a *TriPerplex with value a+bs+cT+dU+eV+fW+gX+hY.
func NewTriPerplex(a, b, c, d, e, f, g, h *big.Rat) *TriPerplex {
	z := new(TriPerplex)
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
func (z *TriPerplex) Scal(y *TriPerplex, a *big.Rat) *TriPerplex {
	z.l.Scal(&y.l, a)
	z.r.Scal(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *TriPerplex) Neg(y *TriPerplex) *TriPerplex {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *TriPerplex) Conj(y *TriPerplex) *TriPerplex {
	z.l.Set(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *TriPerplex) Add(x, y *TriPerplex) *TriPerplex {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *TriPerplex) Sub(x, y *TriPerplex) *TriPerplex {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
// 		Mul(s, s) = Mul(T, T) = Mul(U, U) = +1
// 		Mul(V, V) = Mul(W, W) = Mul(X, X) = Mul(Y, Y) = +1
// 		Mul(s, T) = Mul(T, s) = U
// 		Mul(s, U) = Mul(U, s) = T
// 		Mul(s, V) = Mul(V, s) = W
// 		Mul(s, W) = Mul(W, s) = V
// 		Mul(s, X) = Mul(X, s) = Y
// 		Mul(s, Y) = Mul(Y, s) = X
// 		Mul(T, U) = Mul(U, T) = s
// 		Mul(T, V) = Mul(V, T) = X
// 		Mul(T, W) = Mul(W, T) = Y
// 		Mul(T, X) = Mul(X, T) = V
// 		Mul(T, Y) = Mul(Y, T) = W
// 		Mul(U, V) = Mul(V, U) = Y
// 		Mul(U, W) = Mul(W, U) = X
// 		Mul(U, X) = Mul(X, U) = W
// 		Mul(U, Y) = Mul(Y, U) = V
// 		Mul(V, W) = Mul(W, V) = s
// 		Mul(V, X) = Mul(X, V) = T
// 		Mul(V, Y) = Mul(Y, V) = U
// 		Mul(W, X) = Mul(X, W) = U
// 		Mul(W, Y) = Mul(Y, W) = T
// 		Mul(X, Y) = Mul(Y, X) = s
// This binary operation is commutative and associative.
func (z *TriPerplex) Mul(x, y *TriPerplex) *TriPerplex {
	a := new(BiPerplex).Set(&x.l)
	b := new(BiPerplex).Set(&x.r)
	c := new(BiPerplex).Set(&y.l)
	d := new(BiPerplex).Set(&y.r)
	temp := new(BiPerplex)
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

// Quad returns the quadrance of z. If z = a+bs+cT+dU+eV+fW+gX+hY, then the
// quadrance is
// 		a² - b² + c² - d² + 2(ab + cd)i
// Note that this is a biperplex number.
func (z *TriPerplex) Quad() *BiPerplex {
	quad := new(BiPerplex)
	quad.Mul(&z.l, &z.l)
	return quad.Sub(quad, new(BiPerplex).Mul(&z.r, &z.r))
}

// Norm returns the norm of z. If z = a+bi+cJ+dS, then the norm is
// 		(a² - b² + c² - d²)² + 4(ab + cd)²
// This can also be written as
// 		((a - d)² + (b + c)²)((a + d)² + (b - c)²)
// The norm is always non-negative.
func (z *TriPerplex) Norm() *big.Rat {
	return z.Quad().Quad().Quad()
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *TriPerplex) IsZeroDivisor() bool {
	return z.Quad().IsZeroDivisor()
}

// Inv sets z equal to the inverse of y, and returns z. If y is a zero divisor,
// then Inv panics.
func (z *TriPerplex) Inv(y *TriPerplex) *TriPerplex {
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
func (z *TriPerplex) Quo(x, y *TriPerplex) *TriPerplex {
	if y.IsZeroDivisor() {
		panic("denominator is zero divisor")
	}
	return z.Mul(z.Inv(y), x)
}

// CrossRatio sets z equal to the cross-ratio of v, w, x, and y:
// 		Inv(w - x) * (v - x) * Inv(v - y) * (w - y)
// Then it returns z.
func (z *TriPerplex) CrossRatio(v, w, x, y *TriPerplex) *TriPerplex {
	temp := new(TriPerplex)
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
func (z *TriPerplex) Möbius(y, a, b, c, d *TriPerplex) *TriPerplex {
	z.Mul(a, y)
	z.Add(z, b)
	temp := new(TriPerplex)
	temp.Mul(c, y)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// Generate returns a random TriPerplex value for quick.Check testing.
func (z *TriPerplex) Generate(rand *rand.Rand, size int) reflect.Value {
	randomTriPerplex := &TriPerplex{
		*NewBiPerplex(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
		*NewBiPerplex(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
	}
	return reflect.ValueOf(randomTriPerplex)
}
