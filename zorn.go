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

var symbZorn = [8]string{"", "i", "j", "k", "r", "s", "t", "u"}

// A Zorn represents a rational Zorn octonion.
type Zorn struct {
	l, r Hamilton
}

// Real returns the (rational) real part of z.
func (z *Zorn) Real() *big.Rat {
	return (&z.l).Real()
}

// Rats returns the eight rational components of z.
func (z *Zorn) Rats() (*big.Rat, *big.Rat, *big.Rat, *big.Rat,
	*big.Rat, *big.Rat, *big.Rat, *big.Rat) {
	return &z.l.l.l, &z.l.l.r, &z.l.r.l, &z.l.r.r,
		&z.r.l.l, &z.r.l.r, &z.r.r.l, &z.r.r.r
}

// String returns the string representation of a Zorn value.
//
// If z corresponds to a + bi + cj + dk + er + fs + gt + hu, then the
// string is"(a+bi+cj+dk+er+fs+gt+hu)", similar to complex128 values.
func (z *Zorn) String() string {
	v := make([]*big.Rat, 8)
	v[0], v[1], v[2], v[3] = z.l.Rats()
	v[4], v[5], v[6], v[7] = z.r.Rats()
	a := make([]string, 17)
	a[0] = leftBracket
	a[1] = fmt.Sprintf("%v", v[0].RatString())
	i := 1
	for j := 2; j < 16; j = j + 2 {
		if v[i].Sign() == -1 {
			a[j] = fmt.Sprintf("%v", v[i].RatString())
		} else {
			a[j] = fmt.Sprintf("+%v", v[i].RatString())
		}
		a[j+1] = symbZorn[i]
		i++
	}
	a[16] = rightBracket
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Zorn) Equals(y *Zorn) bool {
	if !z.l.Equals(&y.l) || !z.r.Equals(&y.r) {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *Zorn) Set(y *Zorn) *Zorn {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewZorn returns a pointer to the Zorn value a+bi+cj+dk+er+fs+gt+hu.
func NewZorn(a, b, c, d, e, f, g, h *big.Rat) *Zorn {
	z := new(Zorn)
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
func (z *Zorn) Scal(y *Zorn, a *big.Rat) *Zorn {
	z.l.Scal(&y.l, a)
	z.r.Scal(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Zorn) Neg(y *Zorn) *Zorn {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Zorn) Conj(y *Zorn) *Zorn {
	z.l.Conj(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *Zorn) Add(x, y *Zorn) *Zorn {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *Zorn) Sub(x, y *Zorn) *Zorn {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
// 		Mul(i, i) = Mul(j, j) = Mul(k, k) = -1
// 		Mul(r, r) = Mul(s, s) = Mul(t, t) = Mul(u, u) = +1
// 		Mul(i, j) = -Mul(j, i) = +k
// 		Mul(i, k) = -Mul(k, i) = -j
// 		Mul(i, r) = -Mul(r, i) = +s
// 		Mul(i, s) = -Mul(s, i) = -r
// 		Mul(i, t) = -Mul(t, i) = -u
// 		Mul(i, u) = -Mul(u, i) = +t
// 		Mul(j, k) = -Mul(k, j) = +i
// 		Mul(j, r) = -Mul(r, j) = +t
// 		Mul(j, s) = -Mul(s, j) = +u
// 		Mul(j, t) = -Mul(t, j) = -r
// 		Mul(j, u) = -Mul(u, j) = -s
// 		Mul(k, r) = -Mul(r, k) = +u
// 		Mul(k, s) = -Mul(s, k) = -t
// 		Mul(k, t) = -Mul(t, k) = +s
// 		Mul(k, u) = -Mul(u, k) = -r
// 		Mul(r, s) = -Mul(s, r) = -i
// 		Mul(r, t) = -Mul(t, r) = -j
// 		Mul(r, u) = -Mul(u, r) = -k
// 		Mul(s, t) = -Mul(t, s) = +k
// 		Mul(s, u) = -Mul(u, s) = -j
// 		Mul(t, u) = -Mul(u, t) = +i
// This binary operation is noncommutative and nonassociative.
func (z *Zorn) Mul(x, y *Zorn) *Zorn {
	a := new(Hamilton).Set(&x.l)
	b := new(Hamilton).Set(&x.r)
	c := new(Hamilton).Set(&y.l)
	d := new(Hamilton).Set(&y.r)
	temp := new(Hamilton)
	z.l.Add(
		z.l.Mul(a, c),
		temp.Mul(temp.Conj(d), b),
	)
	z.r.Add(
		z.r.Mul(d, a),
		temp.Mul(b, temp.Conj(c)),
	)
	return z
}

// Commutator sets z equal to the commutator of x and y:
// 		Mul(x, y) - Mul(y, x)
// Then it returns z.
func (z *Zorn) Commutator(x, y *Zorn) *Zorn {
	return z.Sub(
		z.Mul(x, y),
		new(Zorn).Mul(y, x),
	)
}

// Associator sets z equal to the associator of w, x, and y:
// 		Mul(Mul(w, x), y) - Mul(w, Mul(x, y))
// Then it returns z.
func (z *Zorn) Associator(w, x, y *Zorn) *Zorn {
	temp := new(Zorn)
	return z.Sub(
		z.Mul(z.Mul(w, x), y),
		temp.Mul(w, temp.Mul(x, y)),
	)
}

// Quad returns the quadrance of z. If z = a+bi+cj+dk+er+fs+gt+hu, then the
// quadrance is
//		a² + b² + c² + d² - e² - f² - g² - h²
// This can be positive, negative, or zero.
func (z *Zorn) Quad() *big.Rat {
	return new(big.Rat).Sub(
		z.l.Quad(),
		z.r.Quad(),
	)
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *Zorn) IsZeroDivisor() bool {
	return z.l.Quad().Cmp(z.r.Quad()) == 0
}

// Inv sets z equal to the inverse of y, and returns z. If y is a zero divisor,
// then Inv panics.
func (z *Zorn) Inv(y *Zorn) *Zorn {
	if y.IsZeroDivisor() {
		panic("inverse of zero divisor")
	}
	a := y.Quad()
	a.Inv(a)
	return z.Scal(z.Conj(y), a)
}

// QuoL sets z equal to the left quotient of x and y:
// 		Mul(Inv(y), x)
// Then it returns z. If y is a zero divisor, then QuoL panics.
func (z *Zorn) QuoL(x, y *Zorn) *Zorn {
	if y.IsZeroDivisor() {
		panic("denominator is zero divisor")
	}
	return z.Mul(z.Inv(y), x)
}

// QuoR sets z equal to the right quotient of x and y:
// 		Mul(x, Inv(y))
// Then it returns z. If y is a zero divisor, then QuoR panics.
func (z *Zorn) QuoR(x, y *Zorn) *Zorn {
	if y.IsZeroDivisor() {
		panic("denominator is zero divisor")
	}
	return z.Mul(x, z.Inv(y))
}

// Generate returns a random Zorn value for quick.Check testing.
func (z *Zorn) Generate(rand *rand.Rand, size int) reflect.Value {
	randomZorn := &Zorn{
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
	return reflect.ValueOf(randomZorn)
}
