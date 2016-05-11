// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"fmt"
	"math/big"
	"strings"
)

var symbKlein = [8]string{"", "i₁", "i₂", "i₃", "s₄", "s₅", "s₆", "s₇"}

// A Klein represents a rational Klein octonion.
type Klein struct {
	re, sp *Hamilton
}

// Re returns the Cayley-Dickson real part of z, a pointer to a Hamilton value.
func (z *Klein) Re() *Hamilton {
	return z.re
}

// Sp returns the Cayley-Dickson imaginary part of z, a pointer to a Hamilton
// value.
func (z *Klein) Sp() *Hamilton {
	return z.sp
}

// SetRe sets the Klein-Dickson real part of z equal to a.
func (z *Klein) SetRe(a *Hamilton) {
	z.re = a
}

// SetSp sets the Klein-Dickson imaginary part of z equal to b.
func (z *Klein) SetSp(b *Hamilton) {
	z.sp = b
}

// Cartesian returns the eight Cartesian components of z.
func (z *Klein) Cartesian() (a, b, c, d, e, f, g, h *big.Rat) {
	a, b, c, d = z.Re().Cartesian()
	e, f, g, h = z.Sp().Cartesian()
	return
}

// String returns the string representation of a Klein value.
//
// If z corresponds to a + bi₁ + ci₂ + di₃ + es₄ + fs₅ + gs₆ + hs₇, then the
// string is"(a+bi₁+ci₂+di₃+es₄+fs₅+gs₆+hs₇)", similar to complex128 values.
func (z *Klein) String() string {
	v := make([]*big.Rat, 8)
	v[0], v[1], v[2], v[3] = z.Re().Cartesian()
	v[4], v[5], v[6], v[7] = z.Sp().Cartesian()
	a := make([]string, 17)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", v[0])
	i := 1
	for j := 2; j < 16; j = j + 2 {
		if v[i].Sign() == -1 {
			a[j] = fmt.Sprintf("%v", v[i])
		} else {
			a[j] = fmt.Sprintf("+%v", v[i])
		}
		a[j+1] = symbKlein[i]
		i++
	}
	a[16] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Klein) Equals(y *Klein) bool {
	if !z.Re().Equals(y.Re()) || !z.Sp().Equals(y.Sp()) {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Klein) Copy(y *Klein) *Klein {
	z.SetRe(y.Re())
	z.SetSp(y.Sp())
	return z
}

// NewKlein returns a pointer to a Klein value made from eight given pointers
// to big.Rat values.
func NewKlein(a, b, c, d, e, f, g, h *big.Rat) *Klein {
	z := new(Klein)
	z.SetRe(NewHamilton(a, b, c, d))
	z.SetSp(NewHamilton(e, f, g, h))
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Klein) Scal(y *Klein, a *big.Rat) *Klein {
	z.SetRe(new(Hamilton).Scal(y.Re(), a))
	z.SetSp(new(Hamilton).Scal(y.Sp(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Klein) Neg(y *Klein) *Klein {
	z.SetRe(new(Hamilton).Neg(y.Re()))
	z.SetSp(new(Hamilton).Neg(y.Sp()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Klein) Conj(y *Klein) *Klein {
	z.SetRe(new(Hamilton).Conj(y.Re()))
	z.SetSp(new(Hamilton).Neg(y.Sp()))
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Klein) Add(x, y *Klein) *Klein {
	z.SetRe(new(Hamilton).Add(x.Re(), y.Re()))
	z.SetSp(new(Hamilton).Add(x.Sp(), y.Sp()))
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Klein) Sub(x, y *Klein) *Klein {
	z.SetRe(new(Hamilton).Sub(x.Re(), y.Re()))
	z.SetSp(new(Hamilton).Sub(x.Sp(), y.Sp()))
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
// 		Mul(i₁, i₁) = Mul(i₂, i₂) = Mul(i₃, i₃) = -1
// 		Mul(s₄, s₄) = Mul(s₅, s₅) = Mul(s₆, s₆) = Mul(s₇, s₇) = +1
// 		Mul(i₁, i₂) = -Mul(i₂, i₁) = +i₃
// 		Mul(i₁, i₃) = -Mul(i₃, i₁) = -i₂
// 		Mul(i₁, s₄) = -Mul(s₄, i₁) = +s₅
// 		Mul(i₁, s₅) = -Mul(s₅, i₁) = -s₄
// 		Mul(i₁, s₆) = -Mul(s₆, i₁) = -s₇
// 		Mul(i₁, s₇) = -Mul(s₇, i₁) = +s₆
// 		Mul(i₂, i₃) = -Mul(i₃, i₂) = +i₁
// 		Mul(i₂, s₄) = -Mul(s₄, i₂) = +s₆
// 		Mul(i₂, s₅) = -Mul(s₅, i₂) = +s₇
// 		Mul(i₂, s₆) = -Mul(s₆, i₂) = -s₄
// 		Mul(i₂, s₇) = -Mul(s₇, i₂) = -s₅
// 		Mul(i₃, s₄) = -Mul(s₄, i₃) = +s₇
// 		Mul(i₃, s₅) = -Mul(s₅, i₃) = -s₆
// 		Mul(i₃, s₆) = -Mul(s₆, i₃) = +s₅
// 		Mul(i₃, s₇) = -Mul(s₇, i₃) = -s₄
// 		Mul(s₄, s₅) = -Mul(s₅, s₄) = -i₁
// 		Mul(s₄, s₆) = -Mul(s₆, s₄) = -i₂
// 		Mul(s₄, s₇) = -Mul(s₇, s₄) = -i₃
// 		Mul(s₅, s₆) = -Mul(s₆, s₅) = +i₃
// 		Mul(s₅, s₇) = -Mul(s₇, s₅) = -i₂
// 		Mul(s₆, s₇) = -Mul(s₇, s₆) = +i₁
// This binary operation is noncommutative and nonassociative.
func (z *Klein) Mul(x, y *Klein) *Klein {
	p := new(Klein).Copy(x)
	q := new(Klein).Copy(y)
	z.SetRe(new(Hamilton).Add(
		new(Hamilton).Mul(p.Re(), q.Re()),
		new(Hamilton).Mul(new(Hamilton).Conj(q.Sp()), p.Sp()),
	))
	z.SetSp(new(Hamilton).Add(
		new(Hamilton).Mul(q.Sp(), p.Re()),
		new(Hamilton).Mul(p.Sp(), new(Hamilton).Conj(q.Re())),
	))
	return z
}

// Commutator sets z equal to the commutator of x and y, and returns z.
func (z *Klein) Commutator(x, y *Klein) *Klein {
	return z.Sub(
		new(Klein).Mul(x, y),
		new(Klein).Mul(y, x),
	)
}

// Associator sets z equal to the associator of w, x, and y, and returns z.
func (z *Klein) Associator(w, x, y *Klein) *Klein {
	return z.Sub(
		new(Klein).Mul(new(Klein).Mul(w, x), y),
		new(Klein).Mul(w, new(Klein).Mul(x, y)),
	)
}

// Quad returns the quadrance of z, a pointer to a big.Rat value.
func (z *Klein) Quad() *big.Rat {
	return new(big.Rat).Sub(
		z.Re().Quad(),
		z.Sp().Quad(),
	)
}

// IsZeroDiv returns true if z is a zero divisor.
func (z *Klein) IsZeroDiv() bool {
	return z.Re().Quad().Cmp(z.Sp().Quad()) == 0
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Klein) Inv(y *Klein) *Klein {
	if y.IsZeroDiv() {
		panic("inverse of zero divisor")
	}
	return z.Scal(new(Klein).Conj(y), new(big.Rat).Inv(y.Quad()))
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Klein) Quo(x, y *Klein) *Klein {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Mul(x, new(Klein).Inv(y))
}
