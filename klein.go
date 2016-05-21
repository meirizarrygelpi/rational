// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"fmt"
	"math/big"
	"strings"
)

var symbKlein = [8]string{"", "i", "j", "k", "r", "s", "t", "u"}

// A Klein represents a rational Klein octonion.
type Klein struct {
	l, r *Hamilton
}

// L returns the left Cayley-Dickson part of z, a pointer to a Hamilton value.
func (z *Klein) L() *Hamilton {
	return z.l
}

// R returns the right Cayley-Dickson part of z, a pointer to a Hamilton value.
func (z *Klein) R() *Hamilton {
	return z.r
}

// SetL sets the left Cayley-Dickson part of z equal to a.
func (z *Klein) SetL(a *Hamilton) {
	z.l = a
}

// SetR sets the right Cayley-Dickson part of z equal to b.
func (z *Klein) SetR(b *Hamilton) {
	z.r = b
}

// Cartesian returns the eight Cartesian components of z.
func (z *Klein) Cartesian() (a, b, c, d, e, f, g, h *big.Rat) {
	a, b, c, d = z.L().Cartesian()
	e, f, g, h = z.R().Cartesian()
	return
}

// String returns the string representation of a Klein value.
//
// If z corresponds to a + bi + cj + dk + er + fs + gt + hu, then the
// string is"(a+bi+cj+dk+er+fs+gt+hu)", similar to complex128 values.
func (z *Klein) String() string {
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
		a[j+1] = symbKlein[i]
		i++
	}
	a[16] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Klein) Equals(y *Klein) bool {
	if !z.L().Equals(y.L()) || !z.R().Equals(y.R()) {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Klein) Copy(y *Klein) *Klein {
	z.SetL(y.L())
	z.SetR(y.R())
	return z
}

// NewKlein returns a pointer to a Klein value made from eight given pointers
// to big.Rat values.
func NewKlein(a, b, c, d, e, f, g, h *big.Rat) *Klein {
	z := new(Klein)
	z.SetL(NewHamilton(a, b, c, d))
	z.SetR(NewHamilton(e, f, g, h))
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Klein) Scal(y *Klein, a *big.Rat) *Klein {
	z.SetL(new(Hamilton).Scal(y.L(), a))
	z.SetR(new(Hamilton).Scal(y.R(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Klein) Neg(y *Klein) *Klein {
	z.SetL(new(Hamilton).Neg(y.L()))
	z.SetR(new(Hamilton).Neg(y.R()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Klein) Conj(y *Klein) *Klein {
	z.SetL(new(Hamilton).Conj(y.L()))
	z.SetR(new(Hamilton).Neg(y.R()))
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Klein) Add(x, y *Klein) *Klein {
	z.SetL(new(Hamilton).Add(x.L(), y.L()))
	z.SetR(new(Hamilton).Add(x.R(), y.R()))
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Klein) Sub(x, y *Klein) *Klein {
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
func (z *Klein) Mul(x, y *Klein) *Klein {
	a, b := x.L(), x.R()
	c, d := y.L(), y.R()
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
func (z *Klein) Commutator(x, y *Klein) *Klein {
	return z.Sub(
		z.Mul(x, y),
		new(Klein).Mul(y, x),
	)
}

// Associator sets z equal to the associator of w, x, and y, and returns z.
func (z *Klein) Associator(w, x, y *Klein) *Klein {
	t := new(Klein)
	return z.Sub(
		z.Mul(z.Mul(w, x), y),
		t.Mul(w, t.Mul(x, y)),
	)
}

// Quad returns the quadrance of z, a pointer to a big.Rat value.
func (z *Klein) Quad() *big.Rat {
	return new(big.Rat).Sub(
		z.L().Quad(),
		z.R().Quad(),
	)
}

// IsZeroDiv returns true if z is a zero divisor.
func (z *Klein) IsZeroDiv() bool {
	return z.L().Quad().Cmp(z.R().Quad()) == 0
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Klein) Inv(y *Klein) *Klein {
	if y.IsZeroDiv() {
		panic("inverse of zero divisor")
	}
	return z.Scal(z.Conj(y), new(big.Rat).Inv(y.Quad()))
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Klein) Quo(x, y *Klein) *Klein {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Mul(x, z.Inv(y))
}
