package rational

import (
	"fmt"
	"math/big"
	"strings"
)

var symbCockle = [4]string{"", "i₁", "κ₂", "κ₃"}

// A Cockle represents a rational Cockle quaternion.
type Cockle struct {
	body [2]*Complex
}

// Re returns the Cayley-Dickson real part of z, a pointer to a Complex value.
func (z *Cockle) Re() *Complex {
	return z.body[0]
}

// Sp returns the Cayley-Dickson split part of z, a pointer to a Complex
// value.
func (z *Cockle) Sp() *Complex {
	return z.body[1]
}

// SetRe sets the Cayley-Dickson real part of z equal to a.
func (z *Cockle) SetRe(a *Complex) {
	z.body[0] = a
}

// SetSp sets the Cayley-Dickson split part of z equal to b.
func (z *Cockle) SetSp(b *Complex) {
	z.body[1] = b
}

// Cartesian returns the four Cartesian components of z.
func (z *Cockle) Cartesian() (a, b, c, d *big.Rat) {
	a, b = z.Re().Cartesian()
	c, d = z.Sp().Cartesian()
	return
}

// String returns the string representation of a Cockle value. If z
// corresponds to the Cockle quaternion a + bi₁ + cκ₂ + dκ₃, then the string
// is "(a+bi₁+cκ₂+dκ₃)", similar to complex128 values.
func (z *Cockle) String() string {
	v := make([]*big.Rat, 4)
	v[0], v[1] = z.Re().Cartesian()
	v[2], v[3] = z.Sp().Cartesian()
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
		a[j+1] = symbCockle[i]
		i++
	}
	a[8] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Cockle) Equals(y *Cockle) bool {
	if !z.Re().Equals(y.Re()) || !z.Sp().Equals(y.Sp()) {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Cockle) Copy(y *Cockle) *Cockle {
	z.SetRe(y.Re())
	z.SetSp(y.Sp())
	return z
}

// NewCockle returns a pointer to a Cockle value made from eight given
// int64 values.
func NewCockle(a, b, c, d, e, f, g, h int64) *Cockle {
	z := new(Cockle)
	z.SetRe(NewComplex(a, b, c, d))
	z.SetSp(NewComplex(e, f, g, h))
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Cockle) Scal(y *Cockle, a *big.Rat) *Cockle {
	z.SetRe(new(Complex).Scal(y.Re(), a))
	z.SetSp(new(Complex).Scal(y.Sp(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Cockle) Neg(y *Cockle) *Cockle {
	z.SetRe(new(Complex).Neg(y.Re()))
	z.SetSp(new(Complex).Neg(y.Sp()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Cockle) Conj(y *Cockle) *Cockle {
	z.SetRe(new(Complex).Conj(y.Re()))
	z.SetSp(new(Complex).Neg(y.Sp()))
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Cockle) Add(x, y *Cockle) *Cockle {
	z.SetRe(new(Complex).Add(x.Re(), y.Re()))
	z.SetSp(new(Complex).Add(x.Sp(), y.Sp()))
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Cockle) Sub(x, y *Cockle) *Cockle {
	z.SetRe(new(Complex).Sub(x.Re(), y.Re()))
	z.SetSp(new(Complex).Sub(x.Sp(), y.Sp()))
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rule is:
// 		Mul(i₁, i₁) = -1
// 		Mul(κ₂, κ₂) = -1
// 		Mul(κ₃, κ₃) = -1
// 		Mul(i₁, κ₂) = -Mul(κ₂, i₁) = +κ₃
// 		Mul(κ₂, κ₃) = -Mul(κ₃, κ₂) = -i₁
// 		Mul(κ₃, i₁) = -Mul(i₁, κ₃) = +κ₂
// This binary operation is noncommutative but associative.
func (z *Cockle) Mul(x, y *Cockle) *Cockle {
	p := new(Cockle).Copy(x)
	q := new(Cockle).Copy(y)
	z.SetRe(new(Complex).Add(
		new(Complex).Mul(p.Re(), q.Re()),
		new(Complex).Mul(new(Complex).Conj(q.Sp()), p.Sp()),
	))
	z.SetSp(new(Complex).Add(
		new(Complex).Mul(p.Re(), q.Sp()),
		new(Complex).Mul(p.Sp(), new(Complex).Conj(q.Re())),
	))
	return z
}

// Commutator sets z equal to the commutator of x and y, and returns z.
func (z *Cockle) Commutator(x, y *Cockle) *Cockle {
	return z.Sub(
		new(Cockle).Mul(x, y),
		new(Cockle).Mul(y, x),
	)
}

// Quad returns the quadrance of z, a pointer to a big.Rat value.
func (z *Cockle) Quad() *big.Rat {
	return new(big.Rat).Sub(
		z.Re().Quad(),
		z.Sp().Quad(),
	)
}

// IsZeroDiv returns true if z is a zero divisor.
func (z *Cockle) IsZeroDiv() bool {
	return z.Re().Quad().Cmp(z.Sp().Quad()) == 0
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Cockle) Inv(y *Cockle) *Cockle {
	if y.IsZeroDiv() {
		panic("zero divisor inverse")
	}
	return z.Scal(new(Cockle).Conj(y), new(big.Rat).Inv(y.Quad()))
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Cockle) Quo(x, y *Cockle) *Cockle {
	if y.IsZeroDiv() {
		panic("zero divisor denominator")
	}
	return z.Mul(x, new(Cockle).Inv(y))
}
