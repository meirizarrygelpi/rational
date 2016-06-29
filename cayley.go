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

var symbCayley = [8]string{"", "i", "j", "k", "m", "n", "p", "q"}

// A Cayley represents a rational Cayley octonion.
type Cayley struct {
	l, r Hamilton
}

// Real returns the (rational) real part of z.
func (z *Cayley) Real() *big.Rat {
	return (&z.l).Real()
}

// Rats returns the eight rational components of z.
func (z *Cayley) Rats() (*big.Rat, *big.Rat, *big.Rat, *big.Rat,
	*big.Rat, *big.Rat, *big.Rat, *big.Rat) {
	return &z.l.l.l, &z.l.l.r, &z.l.r.l, &z.l.r.r,
		&z.r.l.l, &z.r.l.r, &z.r.r.l, &z.r.r.r
}

// String returns the string representation of a Cayley value.
//
// If z corresponds to a + bi + cj + dk + em + fn + gp + hq, then the
// string is"(a+bi+cj+dk+em+fn+gp+hq)", similar to complex128 values.
func (z *Cayley) String() string {
	v := make([]*big.Rat, 8)
	v[0], v[1], v[2], v[3] = z.l.Rats()
	v[4], v[5], v[6], v[7] = z.r.Rats()
	a := make([]string, 17)
	a[0] = leftBracket
	a[1] = fmt.Sprintf("%v", v[0].RatString())
	i := 1
	for j := 2; j < 16; j = j + 2 {
		if v[i].Sign() < 0 {
			a[j] = fmt.Sprintf("%v", v[i].RatString())
		} else {
			a[j] = fmt.Sprintf("+%v", v[i].RatString())
		}
		a[j+1] = symbCayley[i]
		i++
	}
	a[16] = rightBracket
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Cayley) Equals(y *Cayley) bool {
	if !z.l.Equals(&y.l) || !z.r.Equals(&y.r) {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *Cayley) Set(y *Cayley) *Cayley {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewCayley returns a pointer to the Cayley value a+bi+cj+dk+em+fn+gp+hq.
func NewCayley(a, b, c, d, e, f, g, h *big.Rat) *Cayley {
	z := new(Cayley)
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
func (z *Cayley) Scal(y *Cayley, a *big.Rat) *Cayley {
	z.l.Scal(&y.l, a)
	z.r.Scal(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Cayley) Neg(y *Cayley) *Cayley {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Cayley) Conj(y *Cayley) *Cayley {
	z.l.Conj(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *Cayley) Add(x, y *Cayley) *Cayley {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *Cayley) Sub(x, y *Cayley) *Cayley {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
// 		Mul(i, i) = Mul(j, j) = Mul(k, k) = -1
// 		Mul(m, m) = Mul(n, n) = Mul(p, p) = Mul(q, q) = -1
// 		Mul(i, j) = -Mul(j, i) = +k
// 		Mul(i, k) = -Mul(k, i) = -j
// 		Mul(i, m) = -Mul(m, i) = +n
// 		Mul(i, n) = -Mul(n, i) = -m
// 		Mul(i, p) = -Mul(p, i) = -q
// 		Mul(i, q) = -Mul(q, i) = +p
// 		Mul(j, k) = -Mul(k, j) = +i
// 		Mul(j, m) = -Mul(m, j) = +p
// 		Mul(j, n) = -Mul(n, j) = +q
// 		Mul(j, p) = -Mul(p, j) = -m
// 		Mul(j, q) = -Mul(q, j) = -n
// 		Mul(k, m) = -Mul(m, k) = +q
// 		Mul(k, n) = -Mul(n, k) = -p
// 		Mul(k, p) = -Mul(p, k) = +n
// 		Mul(k, q) = -Mul(q, k) = -m
// 		Mul(m, n) = -Mul(n, m) = +i
// 		Mul(m, p) = -Mul(p, m) = +j
// 		Mul(m, q) = -Mul(q, m) = +k
// 		Mul(n, p) = -Mul(p, n) = -k
// 		Mul(n, q) = -Mul(q, n) = +j
// 		Mul(p, q) = -Mul(q, p) = -i
// This binary operation is noncommutative and nonassociative.
func (z *Cayley) Mul(x, y *Cayley) *Cayley {
	a := new(Hamilton).Set(&x.l)
	b := new(Hamilton).Set(&x.r)
	c := new(Hamilton).Set(&y.l)
	d := new(Hamilton).Set(&y.r)
	temp := new(Hamilton)
	z.l.Sub(
		z.l.Mul(a, c),
		temp.Mul(temp.Conj(d), b),
	)
	z.r.Add(
		z.r.Mul(d, a),
		temp.Mul(b, temp.Conj(c)),
	)
	return z
}

// Commutator sets z equal to the commutator of x and y
// 		Mul(x, y) - Mul(y, x)
// Then it returns z.
func (z *Cayley) Commutator(x, y *Cayley) *Cayley {
	return z.Sub(
		z.Mul(x, y),
		new(Cayley).Mul(y, x),
	)
}

// Associator sets z equal to the associator of w, x, and y:
// 		Mul(Mul(w, x), y) - Mul(w, Mul(x, y))
// Then it returns z.
func (z *Cayley) Associator(w, x, y *Cayley) *Cayley {
	temp := new(Cayley)
	return z.Sub(
		z.Mul(z.Mul(w, x), y),
		temp.Mul(w, temp.Mul(x, y)),
	)
}

// Quad returns the quadrance of z. If z = a+bi+cj+dk+em+fn+gp+hq, then the
// quadrance is
//		a² + b² + c² + d² + e² + f² + g² + h²
// This is always non-negative.
func (z *Cayley) Quad() *big.Rat {
	return new(big.Rat).Add(
		z.l.Quad(),
		z.r.Quad(),
	)
}

// Inv sets z equal to the inverse of y, and returns z. If y is zero, then Inv
// panics.
func (z *Cayley) Inv(y *Cayley) *Cayley {
	if zero := new(Cayley); y.Equals(zero) {
		panic("inverse of zero")
	}
	a := y.Quad()
	a.Inv(a)
	return z.Scal(z.Conj(y), a)
}

// QuoL sets z equal to the left quotient of x and y:
// 		Mul(Inv(y), x)
// Then it returns z. If y is zero, then QuoL panics.
func (z *Cayley) QuoL(x, y *Cayley) *Cayley {
	if zero := new(Cayley); y.Equals(zero) {
		panic("denominator is zero")
	}
	return z.Mul(z.Inv(y), x)
}

// QuoR sets z equal to the right quotient of x and y:
// 		Mul(x, Inv(y))
// Then it returns z. If y is zero, then QuoR panics.
func (z *Cayley) QuoR(x, y *Cayley) *Cayley {
	if zero := new(Cayley); y.Equals(zero) {
		panic("denominator is zero")
	}
	return z.Mul(x, z.Inv(y))
}

// Graves sets z equal to the Gravesian integer a+bi+cj+dk+em+fn+gp+hq, and
// returns z.
func (z *Cayley) Graves(a, b, c, d, e, f, g, h *big.Int) *Cayley {
	z.l.Lipschitz(a, b, c, d)
	z.r.Lipschitz(e, f, g, h)
	return z
}

// Klein sets z equal to the Kleinian integer
// (a+½)+(b+½)i+(c+½)j+(d+½)k+(e+½)m+(f+½)n+(g+½)p+(h+½)q, and returns z.
func (z *Cayley) Klein(a, b, c, d, e, f, g, h *big.Int) *Cayley {
	z.Graves(a, b, c, d, e, f, g, h)
	half := big.NewRat(1, 2)
	z.l.l.l.Add(&z.l.l.l, half)
	z.l.l.r.Add(&z.l.l.r, half)
	z.l.r.l.Add(&z.l.r.l, half)
	z.l.r.r.Add(&z.l.r.r, half)
	z.r.l.l.Add(&z.r.l.l, half)
	z.r.l.r.Add(&z.r.l.r, half)
	z.r.r.l.Add(&z.r.r.l, half)
	z.r.r.r.Add(&z.r.r.r, half)
	return z
}

// Generate returns a random Cayley value for quick.Check testing.
func (z *Cayley) Generate(rand *rand.Rand, size int) reflect.Value {
	randomCayley := &Cayley{
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
	return reflect.ValueOf(randomCayley)
}
