// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"fmt"
	"math/big"
	"strings"
)

var symbCayley = [8]string{"", "i₁", "i₂", "i₃", "i₄", "i₅", "i₆", "i₇"}

// A Cayley represents a rational Cayley octonion.
type Cayley struct {
	re, im *Hamilton
}

// Re returns the Cayley-Dickson real part of z, a pointer to a Hamilton value.
func (z *Cayley) Re() *Hamilton {
	return z.re
}

// Im returns the Cayley-Dickson imaginary part of z, a pointer to a Hamilton
// value.
func (z *Cayley) Im() *Hamilton {
	return z.im
}

// SetRe sets the Cayley-Dickson real part of z equal to a.
func (z *Cayley) SetRe(a *Hamilton) {
	z.re = a
}

// SetIm sets the Cayley-Dickson imaginary part of z equal to b.
func (z *Cayley) SetIm(b *Hamilton) {
	z.im = b
}

// Cartesian returns the eight Cartesian components of z.
func (z *Cayley) Cartesian() (a, b, c, d, e, f, g, h *big.Rat) {
	a, b, c, d = z.Re().Cartesian()
	e, f, g, h = z.Im().Cartesian()
	return
}

// String returns the string representation of a Cayley value.
//
// If z corresponds to a + bi₁ + ci₂ + di₃ + ei₄ + fi₅ + gi₆ + hi₇, then the
// string is"(a+bi₁+ci₂+di₃+ei₄+fi₅+gi₆+hi₇)", similar to complex128 values.
func (z *Cayley) String() string {
	v := make([]*big.Rat, 8)
	v[0], v[1], v[2], v[3] = z.Re().Cartesian()
	v[4], v[5], v[6], v[7] = z.Im().Cartesian()
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
		a[j+1] = symbCayley[i]
		i++
	}
	a[16] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Cayley) Equals(y *Cayley) bool {
	if !z.Re().Equals(y.Re()) || !z.Im().Equals(y.Im()) {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Cayley) Copy(y *Cayley) *Cayley {
	z.SetRe(y.Re())
	z.SetIm(y.Im())
	return z
}

// NewCayley returns a pointer to a Cayley value made from eight given pointers
// to big.Rat values.
func NewCayley(a, b, c, d, e, f, g, h *big.Rat) *Cayley {
	z := new(Cayley)
	z.SetRe(NewHamilton(a, b, c, d))
	z.SetIm(NewHamilton(e, f, g, h))
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Cayley) Scal(y *Cayley, a *big.Rat) *Cayley {
	z.SetRe(new(Hamilton).Scal(y.Re(), a))
	z.SetIm(new(Hamilton).Scal(y.Im(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Cayley) Neg(y *Cayley) *Cayley {
	z.SetRe(new(Hamilton).Neg(y.Re()))
	z.SetIm(new(Hamilton).Neg(y.Im()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Cayley) Conj(y *Cayley) *Cayley {
	z.SetRe(new(Hamilton).Conj(y.Re()))
	z.SetIm(new(Hamilton).Neg(y.Im()))
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Cayley) Add(x, y *Cayley) *Cayley {
	z.SetRe(new(Hamilton).Add(x.Re(), y.Re()))
	z.SetIm(new(Hamilton).Add(x.Im(), y.Im()))
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Cayley) Sub(x, y *Cayley) *Cayley {
	z.SetRe(new(Hamilton).Sub(x.Re(), y.Re()))
	z.SetIm(new(Hamilton).Sub(x.Im(), y.Im()))
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
// 		Mul(i₁, i₁) = Mul(i₂, i₂) = Mul(i₃, i₃) = -1
// 		Mul(i₄, i₄) = Mul(i₅, i₅) = Mul(i₆, i₆) = Mul(i₇, i₇) = -1
// 		Mul(i₁, i₂) = -Mul(i₂, i₁) = +i₃
// 		Mul(i₁, i₃) = -Mul(i₃, i₁) = -i₂
// 		Mul(i₁, i₄) = -Mul(i₄, i₁) = +i₅
// 		Mul(i₁, i₅) = -Mul(i₅, i₁) = -i₄
// 		Mul(i₁, i₆) = -Mul(i₆, i₁) = -i₇
// 		Mul(i₁, i₇) = -Mul(i₇, i₁) = +i₆
// 		Mul(i₂, i₃) = -Mul(i₃, i₂) = +i₁
// 		Mul(i₂, i₄) = -Mul(i₄, i₂) = +i₆
// 		Mul(i₂, i₅) = -Mul(i₅, i₂) = +i₇
// 		Mul(i₂, i₆) = -Mul(i₆, i₂) = -i₄
// 		Mul(i₂, i₇) = -Mul(i₇, i₂) = -i₅
// 		Mul(i₃, i₄) = -Mul(i₄, i₃) = +i₇
// 		Mul(i₃, i₅) = -Mul(i₅, i₃) = -i₆
// 		Mul(i₃, i₆) = -Mul(i₆, i₃) = +i₅
// 		Mul(i₃, i₇) = -Mul(i₇, i₃) = -i₄
// 		Mul(i₄, i₅) = -Mul(i₅, i₄) = +i₁
// 		Mul(i₄, i₆) = -Mul(i₆, i₄) = +i₂
// 		Mul(i₄, i₇) = -Mul(i₇, i₄) = +i₃
// 		Mul(i₅, i₆) = -Mul(i₆, i₅) = -i₃
// 		Mul(i₅, i₇) = -Mul(i₇, i₅) = +i₂
// 		Mul(i₆, i₇) = -Mul(i₇, i₆) = -i₁
// This binary operation is noncommutative and nonassociative.
func (z *Cayley) Mul(x, y *Cayley) *Cayley {
	p := new(Cayley).Copy(x)
	q := new(Cayley).Copy(y)
	z.SetRe(new(Hamilton).Sub(
		new(Hamilton).Mul(p.Re(), q.Re()),
		new(Hamilton).Mul(new(Hamilton).Conj(q.Im()), p.Im()),
	))
	z.SetIm(new(Hamilton).Add(
		new(Hamilton).Mul(q.Im(), p.Re()),
		new(Hamilton).Mul(p.Im(), new(Hamilton).Conj(q.Re())),
	))
	return z
}

// Commutator sets z equal to the commutator of x and y, and returns z.
func (z *Cayley) Commutator(x, y *Cayley) *Cayley {
	return z.Sub(
		new(Cayley).Mul(x, y),
		new(Cayley).Mul(y, x),
	)
}

// Associator sets z equal to the associator of w, x, and y, and returns z.
func (z *Cayley) Associator(w, x, y *Cayley) *Cayley {
	return z.Sub(
		new(Cayley).Mul(new(Cayley).Mul(w, x), y),
		new(Cayley).Mul(w, new(Cayley).Mul(x, y)),
	)
}

// Quad returns the quadrance of z, a pointer to a big.Rat value.
func (z *Cayley) Quad() *big.Rat {
	return new(big.Rat).Add(
		z.Re().Quad(),
		z.Im().Quad(),
	)
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Cayley) Inv(y *Cayley) *Cayley {
	return z.Scal(new(Cayley).Conj(y), new(big.Rat).Inv(y.Quad()))
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Cayley) Quo(x, y *Cayley) *Cayley {
	return z.Mul(x, new(Cayley).Inv(y))
}
