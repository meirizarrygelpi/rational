// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"fmt"
	"math/big"
	"strings"
)

var symbCayley = [8]string{"", "i", "j", "k", "m", "n", "p", "q"}

// A Cayley represents a rational Cayley octonion.
type Cayley struct {
	l, r *Hamilton
}

// L returns the left Cayley-Dickson part of z, a pointer to a Hamilton value.
func (z *Cayley) L() *Hamilton {
	return z.l
}

// R returns the right Cayley-Dickson part of z, a pointer to a Hamilton value.
func (z *Cayley) R() *Hamilton {
	return z.r
}

// SetL sets the left Cayley-Dickson part of z equal to a.
func (z *Cayley) SetL(a *Hamilton) {
	z.l = a
}

// SetR sets the right Cayley-Dickson part of z equal to b.
func (z *Cayley) SetR(b *Hamilton) {
	z.r = b
}

// Cartesian returns the eight Cartesian components of z.
func (z *Cayley) Cartesian() (a, b, c, d, e, f, g, h *big.Rat) {
	a, b, c, d = z.L().Cartesian()
	e, f, g, h = z.R().Cartesian()
	return
}

// String returns the string representation of a Cayley value.
//
// If z corresponds to a + bi + cj + dk + em + fn + gp + hq, then the
// string is"(a+bi+cj+dk+em+fn+gp+hq)", similar to complex128 values.
func (z *Cayley) String() string {
	v := make([]*big.Rat, 8)
	v[0], v[1], v[2], v[3] = z.L().Cartesian()
	v[4], v[5], v[6], v[7] = z.R().Cartesian()
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
	if !z.L().Equals(y.L()) || !z.R().Equals(y.R()) {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Cayley) Copy(y *Cayley) *Cayley {
	z.SetL(y.L())
	z.SetR(y.R())
	return z
}

// NewCayley returns a pointer to a Cayley value made from eight given pointers
// to big.Rat values.
func NewCayley(a, b, c, d, e, f, g, h *big.Rat) *Cayley {
	z := new(Cayley)
	z.SetL(NewHamilton(a, b, c, d))
	z.SetR(NewHamilton(e, f, g, h))
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Cayley) Scal(y *Cayley, a *big.Rat) *Cayley {
	z.SetL(new(Hamilton).Scal(y.L(), a))
	z.SetR(new(Hamilton).Scal(y.R(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Cayley) Neg(y *Cayley) *Cayley {
	z.SetL(new(Hamilton).Neg(y.L()))
	z.SetR(new(Hamilton).Neg(y.R()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Cayley) Conj(y *Cayley) *Cayley {
	z.SetL(new(Hamilton).Conj(y.L()))
	z.SetR(new(Hamilton).Neg(y.R()))
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Cayley) Add(x, y *Cayley) *Cayley {
	z.SetL(new(Hamilton).Add(x.L(), y.L()))
	z.SetR(new(Hamilton).Add(x.R(), y.R()))
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Cayley) Sub(x, y *Cayley) *Cayley {
	z.SetL(new(Hamilton).Sub(x.L(), y.L()))
	z.SetR(new(Hamilton).Sub(x.R(), y.R()))
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
	a, b := x.L(), x.R()
	c, d := y.L(), y.R()
	s, t, u := new(Hamilton), new(Hamilton), new(Hamilton)
	z.SetL(s.Sub(
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
func (z *Cayley) Commutator(x, y *Cayley) *Cayley {
	return z.Sub(
		z.Mul(x, y),
		new(Cayley).Mul(y, x),
	)
}

// Associator sets z equal to the associator of w, x, and y, and returns z.
func (z *Cayley) Associator(w, x, y *Cayley) *Cayley {
	t := new(Cayley)
	return z.Sub(
		z.Mul(z.Mul(w, x), y),
		t.Mul(w, t.Mul(x, y)),
	)
}

// Quad returns the quadrance of z, a pointer to a big.Rat value.
func (z *Cayley) Quad() *big.Rat {
	return new(big.Rat).Add(
		z.L().Quad(),
		z.R().Quad(),
	)
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Cayley) Inv(y *Cayley) *Cayley {
	return z.Scal(z.Conj(y), new(big.Rat).Inv(y.Quad()))
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Cayley) Quo(x, y *Cayley) *Cayley {
	return z.Mul(x, z.Inv(y))
}
