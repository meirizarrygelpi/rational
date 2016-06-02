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

// Cartesian returns the eight Cartesian components of z.
func (z *Cayley) Cartesian() (*big.Rat, *big.Rat, *big.Rat, *big.Rat,
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
	v[0], v[1], v[2], v[3] = z.l.Cartesian()
	v[4], v[5], v[6], v[7] = z.r.Cartesian()
	a := make([]string, 17)
	a[0] = "("
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
	a[16] = ")"
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

// NewCayley returns a pointer to a Cayley value made from eight given pointers
// to big.Rat values.
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

// Add sets z equal to the sum of x and y, and returns z.
func (z *Cayley) Add(x, y *Cayley) *Cayley {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
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

// Commutator sets z equal to the commutator of x and y, and returns z.
func (z *Cayley) Commutator(x, y *Cayley) *Cayley {
	return z.Sub(
		z.Mul(x, y),
		new(Cayley).Mul(y, x),
	)
}

// Associator sets z equal to the associator of w, x, and y, and returns z.
func (z *Cayley) Associator(w, x, y *Cayley) *Cayley {
	temp := new(Cayley)
	return z.Sub(
		z.Mul(z.Mul(w, x), y),
		temp.Mul(w, temp.Mul(x, y)),
	)
}

// Quad returns the quadrance of z. If z = a+bi+cj+dk+em+fn+gp+hq, then the
// quadrance is
//		Mul(a, a) + Mul(b, b) + Mul(c, c) + Mul(d, d) +
// 		Mul(e, e) + Mul(f, f) + Mul(g, g) + Mul(h, h)
// This is always non-negative.
func (z *Cayley) Quad() *big.Rat {
	return new(big.Rat).Add(
		z.l.Quad(),
		z.r.Quad(),
	)
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Cayley) Inv(y *Cayley) *Cayley {
	a := y.Quad()
	a.Inv(a)
	return z.Scal(z.Conj(y), a)
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Cayley) Quo(x, y *Cayley) *Cayley {
	return z.Mul(x, z.Inv(y))
}

// Graves sets z equal to the Graves integer a+bi+cj+dk+em+fn+gp+hq, and
// returns z.
func (z *Cayley) Graves(a, b, c, d, e, f, g, h *big.Int) *Cayley {
	z.l.l.l.SetInt(a)
	z.l.l.r.SetInt(b)
	z.l.r.l.SetInt(c)
	z.l.r.r.SetInt(d)
	z.r.l.l.SetInt(e)
	z.r.l.r.SetInt(f)
	z.r.r.l.SetInt(g)
	z.r.r.r.SetInt(h)
	return z
}

// Klein sets z equal to the Klein integer
// (a+1/2)+(b+1/2)i+(c+1/2)j+(d+1/2)k+(e+1/2)m+(f+1/2)n+(g+1/2)p+(h+1/2)q,
// and returns z.
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
