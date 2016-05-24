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

// L returns the left Cayley-Dickson part of z, a pointer to a Hamilton value.
func (z *Zorn) L() *Hamilton {
	return &z.l
}

// R returns the right Cayley-Dickson part of z, a pointer to a Hamilton value.
func (z *Zorn) R() *Hamilton {
	return &z.r
}

// SetL sets the left Cayley-Dickson part of z equal to a.
func (z *Zorn) SetL(a *Hamilton) {
	z.l = *a
}

// SetR sets the right Cayley-Dickson part of z equal to b.
func (z *Zorn) SetR(b *Hamilton) {
	z.r = *b
}

// Cartesian returns the eight Cartesian components of z.
func (z *Zorn) Cartesian() (a, b, c, d, e, f, g, h *big.Rat) {
	a, b, c, d = z.L().Cartesian()
	e, f, g, h = z.R().Cartesian()
	return
}

// String returns the string representation of a Zorn value.
//
// If z corresponds to a + bi + cj + dk + er + fs + gt + hu, then the
// string is"(a+bi+cj+dk+er+fs+gt+hu)", similar to complex128 values.
func (z *Zorn) String() string {
	v := make([]*big.Rat, 8)
	v[0], v[1], v[2], v[3] = z.L().Cartesian()
	v[4], v[5], v[6], v[7] = z.R().Cartesian()
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
		a[j+1] = symbZorn[i]
		i++
	}
	a[16] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Zorn) Equals(y *Zorn) bool {
	if !z.L().Equals(y.L()) || !z.R().Equals(y.R()) {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Zorn) Copy(y *Zorn) *Zorn {
	z.SetL(y.L())
	z.SetR(y.R())
	return z
}

// NewZorn returns a pointer to a Zorn value made from eight given pointers
// to big.Rat values.
func NewZorn(a, b, c, d, e, f, g, h *big.Rat) *Zorn {
	z := new(Zorn)
	z.SetL(NewHamilton(a, b, c, d))
	z.SetR(NewHamilton(e, f, g, h))
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Zorn) Scal(y *Zorn, a *big.Rat) *Zorn {
	z.SetL(new(Hamilton).Scal(y.L(), a))
	z.SetR(new(Hamilton).Scal(y.R(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Zorn) Neg(y *Zorn) *Zorn {
	z.SetL(new(Hamilton).Neg(y.L()))
	z.SetR(new(Hamilton).Neg(y.R()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Zorn) Conj(y *Zorn) *Zorn {
	z.SetL(new(Hamilton).Conj(y.L()))
	z.SetR(new(Hamilton).Neg(y.R()))
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Zorn) Add(x, y *Zorn) *Zorn {
	z.SetL(new(Hamilton).Add(x.L(), y.L()))
	z.SetR(new(Hamilton).Add(x.R(), y.R()))
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Zorn) Sub(x, y *Zorn) *Zorn {
	z.SetL(new(Hamilton).Sub(x.L(), y.L()))
	z.SetR(new(Hamilton).Sub(x.R(), y.R()))
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
	a := new(Hamilton).Copy(x.L())
	b := new(Hamilton).Copy(x.R())
	c := new(Hamilton).Copy(y.L())
	d := new(Hamilton).Copy(y.R())
	s, t, u := new(Hamilton), new(Hamilton), new(Hamilton)
	z.SetL(s.Add(
		s.Mul(a, c),
		u.Mul(u.Conj(d), b),
	))
	z.SetR(t.Add(
		t.Mul(d, a),
		u.Mul(b, u.Conj(c)),
	))
	return z
}

// Commutator sets z equal to the commutator of x and y, and returns z.
func (z *Zorn) Commutator(x, y *Zorn) *Zorn {
	return z.Sub(
		z.Mul(x, y),
		new(Zorn).Mul(y, x),
	)
}

// Associator sets z equal to the associator of w, x, and y, and returns z.
func (z *Zorn) Associator(w, x, y *Zorn) *Zorn {
	t := new(Zorn)
	return z.Sub(
		z.Mul(z.Mul(w, x), y),
		t.Mul(w, t.Mul(x, y)),
	)
}

// Quad returns the quadrance of z, a pointer to a big.Rat value.
func (z *Zorn) Quad() *big.Rat {
	return new(big.Rat).Sub(
		z.L().Quad(),
		z.R().Quad(),
	)
}

// IsZeroDiv returns true if z is a zero divisor.
func (z *Zorn) IsZeroDiv() bool {
	return z.L().Quad().Cmp(z.R().Quad()) == 0
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Zorn) Inv(y *Zorn) *Zorn {
	if y.IsZeroDiv() {
		panic("inverse of zero divisor")
	}
	return z.Scal(z.Conj(y), new(big.Rat).Inv(y.Quad()))
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Zorn) Quo(x, y *Zorn) *Zorn {
	if y.IsZeroDiv() {
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
