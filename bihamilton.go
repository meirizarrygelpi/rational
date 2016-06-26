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

var symbBiHamilton = [8]string{"", "i", "j", "k", "H", "S", "T", "U"}

// A BiHamilton represents a rational Hamilton biquaternion.
type BiHamilton struct {
	l, r Hamilton
}

// Real returns the (rational) real part of z.
func (z *BiHamilton) Real() *big.Rat {
	return (&z.l).Real()
}

// Rats returns the eight rational components of z.
func (z *BiHamilton) Rats() (*big.Rat, *big.Rat, *big.Rat, *big.Rat,
	*big.Rat, *big.Rat, *big.Rat, *big.Rat) {
	return &z.l.l.l, &z.l.l.r, &z.l.r.l, &z.l.r.r,
		&z.r.l.l, &z.r.l.r, &z.r.r.l, &z.r.r.r
}

// String returns the string representation of a BiHamilton value.
//
// If z corresponds to a + bi + cj + dk + eH + fS + gT + hU, then the string is
// "(a+bi+cj+dk+eH+fV+gW+hL)", similar to complex128 values.
func (z *BiHamilton) String() string {
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
		a[j+1] = symbBiHamilton[i]
		i++
	}
	a[16] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *BiHamilton) Equals(y *BiHamilton) bool {
	if !z.l.Equals(&y.l) || !z.r.Equals(&y.r) {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *BiHamilton) Set(y *BiHamilton) *BiHamilton {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewBiHamilton returns a *BiHamilton with value a+bi+cj+dk+eH+fS+gT+hU.
func NewBiHamilton(a, b, c, d, e, f, g, h *big.Rat) *BiHamilton {
	z := new(BiHamilton)
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
func (z *BiHamilton) Scal(y *BiHamilton, a *big.Rat) *BiHamilton {
	z.l.Scal(&y.l, a)
	z.r.Scal(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *BiHamilton) Neg(y *BiHamilton) *BiHamilton {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *BiHamilton) Conj(y *BiHamilton) *BiHamilton {
	z.l.Set(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *BiHamilton) Add(x, y *BiHamilton) *BiHamilton {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *BiHamilton) Sub(x, y *BiHamilton) *BiHamilton {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
// 		Mul(i, i) = Mul(j, j) = Mul(k, k) = Mul(H, H) = -1
// 		Mul(S, S) = Mul(T, T) = Mul(U, U) = +1
// 		Mul(i, j) = -Mul(j, i) = +k
// 		Mul(i, k) = -Mul(k, i) = -j
// 		Mul(i, H) = Mul(H, i) = +S
// 		Mul(i, S) = Mul(S, i) = -H
// 		Mul(i, T) = -Mul(T, i) = +U
// 		Mul(i, U) = -Mul(U, i) = -T
// 		Mul(j, k) = -Mul(k, j) = +i
// 		Mul(j, H) = Mul(H, j) = +T
// 		Mul(j, S) = -Mul(S, j) = -U
// 		Mul(j, T) = Mul(T, j) = -H
// 		Mul(j, U) = -Mul(U, j) = +S
// 		Mul(k, H) = Mul(H, k) = +U
// 		Mul(k, S) = -Mul(S, k) = +T
// 		Mul(k, T) = -Mul(T, k) = -S
// 		Mul(k, U) = Mul(U, k) = -H
// 		Mul(H, S) = Mul(S, H) = -i
// 		Mul(H, T) = Mul(T, H) = -j
// 		Mul(H, U) = Mul(U, H) = -k
// 		Mul(S, T) = -Mul(T, S) = -k
// 		Mul(S, U) = -Mul(U, S) = +j
// 		Mul(T, U) = -Mul(U, T) = -i
// This binary operation is noncommutative but associative.
func (z *BiHamilton) Mul(x, y *BiHamilton) *BiHamilton {
	a := new(Hamilton).Set(&x.l)
	b := new(Hamilton).Set(&x.r)
	c := new(Hamilton).Set(&y.l)
	d := new(Hamilton).Set(&y.r)
	temp := new(Hamilton)
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

// Commutator sets z equal to the commutator of x and y:
// 		Mul(x, y) - Mul(y, x)
// Then it returns z.
func (z *BiHamilton) Commutator(x, y *BiHamilton) *BiHamilton {
	return z.Sub(
		z.Mul(x, y),
		new(BiHamilton).Mul(y, x),
	)
}

// Quad returns the quadrance of z. If z = a+bi+cj+dk+eH+fS+gT+hU, then the
// quadrance is
// 		a² - b² - c² - d² + e² - f² - g² - h² +
// 		2(ab + ef)i + 2(ac + eg)j +2(ad + eh)k
// Note that this is a Hamilton quaternion.
func (z *BiHamilton) Quad() *Hamilton {
	quad := new(Hamilton)
	quad.Mul(&z.l, &z.l)
	return quad.Add(quad, new(Hamilton).Mul(&z.r, &z.r))
}

// Norm returns the norm of z. If z = a+bi+cJ+dS, then the norm is
// 		(a² - b² + c² - d²)² + 4(ab + cd)²
// This can also be written as
// 		((a - d)² + (b + c)²)((a + d)² + (b - c)²)
// The norm is always non-negative.
func (z *BiHamilton) Norm() *big.Rat {
	return z.Quad().Quad()
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *BiHamilton) IsZeroDivisor() bool {
	zero := new(Hamilton)
	return zero.Equals(z.Quad())
}

// Inv sets z equal to the inverse of y, and returns z. If y is a zero divisor,
// then Inv panics.
func (z *BiHamilton) Inv(y *BiHamilton) *BiHamilton {
	if y.IsZeroDivisor() {
		panic("inverse of zero divisor")
	}
	quad := y.Quad()
	quad.Inv(quad)
	z.Conj(y)
	z.l.Mul(quad, &z.l)
	z.r.Mul(quad, &z.r)
	return z
}

// Quo sets z equal to the quotient of x and y. If y is a zero divisor, then
// Quo panics.
func (z *BiHamilton) Quo(x, y *BiHamilton) *BiHamilton {
	if y.IsZeroDivisor() {
		panic("denominator is zero divisor")
	}
	return z.Mul(z.Inv(y), x)
}

// CrossRatio sets z equal to the cross-ratio of v, w, x, and y:
// 		Inv(w - x) * (v - x) * Inv(v - y) * (w - y)
// Then it returns z.
func (z *BiHamilton) CrossRatio(v, w, x, y *BiHamilton) *BiHamilton {
	temp := new(BiHamilton)
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
func (z *BiHamilton) Möbius(y, a, b, c, d *BiHamilton) *BiHamilton {
	z.Mul(a, y)
	z.Add(z, b)
	temp := new(BiHamilton)
	temp.Mul(c, y)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// Generate returns a random BiHamilton value for quick.Check testing.
func (z *BiHamilton) Generate(rand *rand.Rand, size int) reflect.Value {
	randomBiHamilton := &BiHamilton{
		*NewHamilton(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
		*NewHamilton(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
	}
	return reflect.ValueOf(randomBiHamilton)
}
