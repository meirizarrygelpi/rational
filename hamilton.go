package rational

import (
	"fmt"
	"math/big"
	"strings"
)

var symbHamilton = [4]string{"", "i₁", "i₂", "i₃"}

// A Hamilton represents a rational Hamilton quaternion.
type Hamilton struct {
	re, im *Complex
}

// Re returns the Cayley-Dickson real part of z, a pointer to a Complex value.
func (z *Hamilton) Re() *Complex {
	return z.re
}

// Im returns the Cayley-Dickson imaginary part of z, a pointer to a Complex
// value.
func (z *Hamilton) Im() *Complex {
	return z.im
}

// SetRe sets the Cayley-Dickson real part of z equal to a.
func (z *Hamilton) SetRe(a *Complex) {
	z.re = a
}

// SetIm sets the Cayley-Dickson imaginary part of z equal to b.
func (z *Hamilton) SetIm(b *Complex) {
	z.im = b
}

// Cartesian returns the four Cartesian components of z.
func (z *Hamilton) Cartesian() (a, b, c, d *big.Rat) {
	a, b = z.Re().Cartesian()
	c, d = z.Im().Cartesian()
	return
}

// String returns the string representation of a Hamilton value. If z
// corresponds to the Hamilton quaternion a + bi₁ + ci₂ + di₃, then the string
// is "(a+bi₁+ci₂+di₃)", similar to complex128 values.
func (z *Hamilton) String() string {
	v := make([]*big.Rat, 4)
	v[0], v[1] = z.Re().Cartesian()
	v[2], v[3] = z.Im().Cartesian()
	a := make([]string, 9)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", v[0])
	i := 1
	for j := 2; j < 8; j = j + 2 {
		if v[i].Sign() == -1 {
			a[j] = fmt.Sprintf("%v", v[i])
		} else {
			a[j] = fmt.Sprintf("+%v", v[i])
		}
		a[j+1] = symbHamilton[i]
		i++
	}
	a[8] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Hamilton) Equals(y *Hamilton) bool {
	if !z.Re().Equals(y.Re()) || !z.Im().Equals(y.Im()) {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Hamilton) Copy(y *Hamilton) *Hamilton {
	z.SetRe(y.Re())
	z.SetIm(y.Im())
	return z
}

// NewHamilton returns a pointer to a Hamilton value made from eight given
// int64 values.
func NewHamilton(a, b, c, d, e, f, g, h int64) *Hamilton {
	z := new(Hamilton)
	z.SetRe(NewComplex(a, b, c, d))
	z.SetIm(NewComplex(e, f, g, h))
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Hamilton) Scal(y *Hamilton, a *big.Rat) *Hamilton {
	z.SetRe(new(Complex).Scal(y.Re(), a))
	z.SetIm(new(Complex).Scal(y.Im(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Hamilton) Neg(y *Hamilton) *Hamilton {
	z.SetRe(new(Complex).Neg(y.Re()))
	z.SetIm(new(Complex).Neg(y.Im()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Hamilton) Conj(y *Hamilton) *Hamilton {
	z.SetRe(new(Complex).Conj(y.Re()))
	z.SetIm(new(Complex).Neg(y.Im()))
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Hamilton) Add(x, y *Hamilton) *Hamilton {
	z.SetRe(new(Complex).Add(x.Re(), y.Re()))
	z.SetIm(new(Complex).Add(x.Im(), y.Im()))
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Hamilton) Sub(x, y *Hamilton) *Hamilton {
	z.SetRe(new(Complex).Sub(x.Re(), y.Re()))
	z.SetIm(new(Complex).Sub(x.Im(), y.Im()))
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
// 		Mul(i₁, i₁) = Mul(i₂, i₂) = Mul(i₃, i₃) = -1
// 		Mul(i₁, i₂) = -Mul(i₂, i₁) = i₃
// 		Mul(i₂, i₃) = -Mul(i₃, i₂) = i₁
// 		Mul(i₃, i₁) = -Mul(i₁, i₃) = i₂
// This binary operation is noncommutative but associative.
func (z *Hamilton) Mul(x, y *Hamilton) *Hamilton {
	p := new(Hamilton).Copy(x)
	q := new(Hamilton).Copy(y)
	z.SetRe(new(Complex).Sub(
		new(Complex).Mul(p.Re(), q.Re()),
		new(Complex).Mul(new(Complex).Conj(q.Im()), p.Im()),
	))
	z.SetIm(new(Complex).Add(
		new(Complex).Mul(p.Re(), q.Im()),
		new(Complex).Mul(p.Im(), new(Complex).Conj(q.Re())),
	))
	return z
}

// Commutator sets z equal to the commutator of x and y, and returns z.
func (z *Hamilton) Commutator(x, y *Hamilton) *Hamilton {
	return z.Sub(
		new(Hamilton).Mul(x, y),
		new(Hamilton).Mul(y, x),
	)
}

// Quad returns the quadrance of z, a pointer to a big.Rat value.
func (z *Hamilton) Quad() *big.Rat {
	return new(big.Rat).Add(
		z.Re().Quad(),
		z.Im().Quad(),
	)
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Hamilton) Inv(y *Hamilton) *Hamilton {
	return z.Scal(new(Hamilton).Conj(y), new(big.Rat).Inv(y.Quad()))
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Hamilton) Quo(x, y *Hamilton) *Hamilton {
	return z.Mul(x, new(Hamilton).Inv(y))
}
