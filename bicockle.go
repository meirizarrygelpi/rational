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

var symbBiCockle = [8]string{"", "i", "t", "u", "H", "S", "J", "K"}

// A BiCockle represents a rational Cockle biquaternion.
type BiCockle struct {
	l, r Cockle
}

// Real returns the (rational) real part of z.
func (z *BiCockle) Real() *big.Rat {
	return (&z.l).Real()
}

// Rats returns the eight rational components of z.
func (z *BiCockle) Rats() (*big.Rat, *big.Rat, *big.Rat, *big.Rat,
	*big.Rat, *big.Rat, *big.Rat, *big.Rat) {
	return &z.l.l.l, &z.l.l.r, &z.l.r.l, &z.l.r.r,
		&z.r.l.l, &z.r.l.r, &z.r.r.l, &z.r.r.r
}

// String returns the string representation of a BiCockle value.
//
// If z corresponds to a + bi + cj + dk + eH + fS + gT + hU, then the string is
// "(a+bi+cj+dk+eH+fV+gW+hL)", similar to complex128 values.
func (z *BiCockle) String() string {
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
		a[j+1] = symbBiCockle[i]
		i++
	}
	a[16] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *BiCockle) Equals(y *BiCockle) bool {
	if !z.l.Equals(&y.l) || !z.r.Equals(&y.r) {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *BiCockle) Set(y *BiCockle) *BiCockle {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewBiCockle returns a *BiCockle with value a+bi+cj+dk+eH+fS+gT+hU.
func NewBiCockle(a, b, c, d, e, f, g, h *big.Rat) *BiCockle {
	z := new(BiCockle)
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
func (z *BiCockle) Scal(y *BiCockle, a *big.Rat) *BiCockle {
	z.l.Scal(&y.l, a)
	z.r.Scal(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *BiCockle) Neg(y *BiCockle) *BiCockle {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *BiCockle) Conj(y *BiCockle) *BiCockle {
	z.l.Conj(&y.l)
	z.r.Conj(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *BiCockle) Add(x, y *BiCockle) *BiCockle {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *BiCockle) Sub(x, y *BiCockle) *BiCockle {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
// 		Mul(i, i) = Mul(H, H) = Mul(J, J) = Mul(K, K) = -1
// 		Mul(S, S) = Mul(t, t) = Mul(u, u) = +1
// 		Mul(i, t) = -Mul(t, i) = +u
// 		Mul(i, u) = -Mul(u, i) = -t
// 		Mul(i, H) = Mul(H, i) = +S
// 		Mul(i, S) = Mul(S, i) = -H
// 		Mul(i, J) = -Mul(J, i) = +K
// 		Mul(i, K) = -Mul(K, i) = -J
// 		Mul(t, u) = -Mul(u, t) = -i
// 		Mul(t, H) = Mul(H, t) = +J
// 		Mul(t, S) = -Mul(S, t) = -K
// 		Mul(t, J) = Mul(J, t) = +H
// 		Mul(t, K) = -Mul(K, t) = -S
// 		Mul(u, H) = Mul(H, u) = +K
// 		Mul(u, S) = -Mul(S, u) = +J
// 		Mul(u, J) = -Mul(J, u) = +S
// 		Mul(u, K) = Mul(K, u) = +H
// 		Mul(H, S) = Mul(S, H) = -i
// 		Mul(H, J) = Mul(J, H) = -t
// 		Mul(H, K) = Mul(K, H) = -u
// 		Mul(S, J) = -Mul(J, S) = -u
// 		Mul(S, K) = -Mul(K, S) = +t
// 		Mul(J, K) = -Mul(K, J) = +i
// This binary operation is noncommutative but associative.
func (z *BiCockle) Mul(x, y *BiCockle) *BiCockle {
	a := new(Cockle).Set(&x.l)
	b := new(Cockle).Set(&x.r)
	c := new(Cockle).Set(&y.l)
	d := new(Cockle).Set(&y.r)
	temp := new(Cockle)
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
func (z *BiCockle) Commutator(x, y *BiCockle) *BiCockle {
	return z.Sub(
		z.Mul(x, y),
		new(BiCockle).Mul(y, x),
	)
}

// quad returns the quadrance of z. If z = a+bi+ct+du+eH+fS+gT+hU, then the
// quadrance is
//		a² + b² - c² - d² - e² - f² + g² + h² +
// 		2(ae + bf - cg - dh)H
// Note that this is a complex number with H serving as the imaginary unit.
func (z *BiCockle) quad() *Complex {
	q := new(Complex)
	q.l.Sub(z.l.Quad(), z.r.Quad())
	temp := new(big.Rat)
	q.r.Mul(&z.l.l.l, &z.r.l.l)
	q.r.Add(&q.r, temp.Mul(&z.l.l.r, &z.r.l.r))
	q.r.Sub(&q.r, temp.Mul(&z.l.r.l, &z.r.r.l))
	q.r.Sub(&q.r, temp.Mul(&z.l.r.r, &z.r.r.r))
	q.r.Add(&q.r, &q.r)
	return q
}

// Norm returns the norm of z. If z = a+bi+cj+dk+eH+fS+gT+hU, then the norm is
// 		(a² + b² - c² - d² - e² - f² + g² + h²)² +
// 		4(ae + bf - cg - dh)²
// The norm is always non-negative.
func (z *BiCockle) Norm() *big.Rat {
	return z.quad().Quad()
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *BiCockle) IsZeroDivisor() bool {
	zero := new(Complex)
	return zero.Equals(z.quad())
}

// Inv sets z equal to the inverse of y, and returns z. If y is a zero divisor,
// then Inv panics.
func (z *BiCockle) Inv(y *BiCockle) *BiCockle {
	if y.IsZeroDivisor() {
		panic("inverse of zero divisor")
	}
	p := new(BiCockle).Conj(y)
	q := y.quad()
	q.Inv(q)
	z.Conj(y)
	temp := new(Cockle)
	z.l.Scal(&p.l, &q.l)
	z.l.Sub(&z.l, temp.Scal(&p.r, &q.r))
	z.r.Scal(&p.l, &q.r)
	z.r.Add(&z.r, temp.Scal(&p.r, &q.l))
	return z
}

// QuoL sets z equal to the left quotient of x and y:
// 		Mul(Inv(y), x)
// Then it returns z. If y is zero, then QuoL panics.
func (z *BiCockle) QuoL(x, y *BiCockle) *BiCockle {
	if y.IsZeroDivisor() {
		panic("denominator is zero divisor")
	}
	return z.Mul(z.Inv(y), x)
}

// QuoR sets z equal to the right quotient of x and y:
// 		Mul(x, Inv(y))
// Then it returns z. If y is zero, then QuoR panics.
func (z *BiCockle) QuoR(x, y *BiCockle) *BiCockle {
	if y.IsZeroDivisor() {
		panic("denominator is zero divisor")
	}
	return z.Mul(x, z.Inv(y))
}

// CrossRatioL sets z equal to the left cross-ratio of v, w, x, and y:
// 		Inv(w - x) * (v - x) * Inv(v - y) * (w - y)
// Then it returns z.
func (z *BiCockle) CrossRatioL(v, w, x, y *BiCockle) *BiCockle {
	temp := new(BiCockle)
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
func (z *BiCockle) CrossRatioR(v, w, x, y *BiCockle) *BiCockle {
	temp := new(BiCockle)
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
func (z *BiCockle) MöbiusL(y, a, b, c, d *BiCockle) *BiCockle {
	z.Mul(y, a)
	z.Add(z, b)
	temp := new(BiCockle)
	temp.Mul(y, c)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(temp, z)
}

// MöbiusR sets z equal to the right Möbius (fractional linear) transform of y:
// 		(a*y + b) * Inv(c*y + d)
// Then it returns z.
func (z *BiCockle) MöbiusR(y, a, b, c, d *BiCockle) *BiCockle {
	z.Mul(a, y)
	z.Add(z, b)
	temp := new(BiCockle)
	temp.Mul(c, y)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// Generate returns a random BiCockle value for quick.Check testing.
func (z *BiCockle) Generate(rand *rand.Rand, size int) reflect.Value {
	randomBiCockle := &BiCockle{
		*NewCockle(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
		*NewCockle(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
	}
	return reflect.ValueOf(randomBiCockle)
}
