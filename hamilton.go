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

var symbHamilton = [4]string{"", "i", "j", "k"}

// A Hamilton represents a rational Hamilton quaternion.
type Hamilton struct {
	l, r Complex
}

// L returns the left Cayley-Dickson part of z, a pointer to a Complex value.
func (z *Hamilton) L() *Complex {
	return &z.l
}

// R returns the right Cayley-Dickson part of z, a pointer to a Complex value.
func (z *Hamilton) R() *Complex {
	return &z.r
}

// SetL sets the left Cayley-Dickson part of z equal to a.
func (z *Hamilton) SetL(a *Complex) {
	z.l = *a
}

// SetR sets the right Cayley-Dickson part of z equal to b.
func (z *Hamilton) SetR(b *Complex) {
	z.r = *b
}

// Cartesian returns the four Cartesian components of z.
func (z *Hamilton) Cartesian() (a, b, c, d *big.Rat) {
	a, b = z.L().Cartesian()
	c, d = z.R().Cartesian()
	return
}

// String returns the string representation of a Hamilton value.
//
// If z corresponds to a + bi + cj + dk, then the string is"(a+bi+cj+dk)",
// similar to complex128 values.
func (z *Hamilton) String() string {
	v := make([]*big.Rat, 4)
	v[0], v[1] = z.L().Cartesian()
	v[2], v[3] = z.R().Cartesian()
	a := make([]string, 9)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", v[0].RatString())
	i := 1
	for j := 2; j < 8; j = j + 2 {
		if v[i].Sign() < 0 {
			a[j] = fmt.Sprintf("%v", v[i].RatString())
		} else {
			a[j] = fmt.Sprintf("+%v", v[i].RatString())
		}
		a[j+1] = symbHamilton[i]
		i++
	}
	a[8] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Hamilton) Equals(y *Hamilton) bool {
	if !z.L().Equals(y.L()) || !z.R().Equals(y.R()) {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Hamilton) Copy(y *Hamilton) *Hamilton {
	z.SetL(y.L())
	z.SetR(y.R())
	return z
}

// NewHamilton returns a pointer to a Hamilton value made from four given
// pointers to big.Rat values.
func NewHamilton(a, b, c, d *big.Rat) *Hamilton {
	z := new(Hamilton)
	z.SetL(NewComplex(a, b))
	z.SetR(NewComplex(c, d))
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Hamilton) Scal(y *Hamilton, a *big.Rat) *Hamilton {
	z.SetL(new(Complex).Scal(y.L(), a))
	z.SetR(new(Complex).Scal(y.R(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Hamilton) Neg(y *Hamilton) *Hamilton {
	z.SetL(new(Complex).Neg(y.L()))
	z.SetR(new(Complex).Neg(y.R()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Hamilton) Conj(y *Hamilton) *Hamilton {
	z.SetL(new(Complex).Conj(y.L()))
	z.SetR(new(Complex).Neg(y.R()))
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Hamilton) Add(x, y *Hamilton) *Hamilton {
	z.SetL(new(Complex).Add(x.L(), y.L()))
	z.SetR(new(Complex).Add(x.R(), y.R()))
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Hamilton) Sub(x, y *Hamilton) *Hamilton {
	z.SetL(new(Complex).Sub(x.L(), y.L()))
	z.SetR(new(Complex).Sub(x.R(), y.R()))
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
// 		Mul(i, i) = Mul(j, j) = Mul(k, k) = -1
// 		Mul(i, j) = -Mul(j, i) = k
// 		Mul(j, k) = -Mul(k, j) = i
// 		Mul(k, i) = -Mul(i, k) = j
// This binary operation is noncommutative but associative.
func (z *Hamilton) Mul(x, y *Hamilton) *Hamilton {
	a := new(Complex).Copy(x.L())
	b := new(Complex).Copy(x.R())
	c := new(Complex).Copy(y.L())
	d := new(Complex).Copy(y.R())
	s, t, u := new(Complex), new(Complex), new(Complex)
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
func (z *Hamilton) Commutator(x, y *Hamilton) *Hamilton {
	return z.Sub(
		z.Mul(x, y),
		new(Hamilton).Mul(y, x),
	)
}

// Quad returns the non-negative quadrance of z, a pointer to a big.Rat value.
func (z *Hamilton) Quad() *big.Rat {
	return new(big.Rat).Add(
		z.L().Quad(),
		z.R().Quad(),
	)
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Hamilton) Inv(y *Hamilton) *Hamilton {
	return z.Scal(z.Conj(y), new(big.Rat).Inv(y.Quad()))
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Hamilton) Quo(x, y *Hamilton) *Hamilton {
	return z.Mul(x, z.Inv(y))
}

// Lipschitz sets z equal to a Lipschitz integer made from four given pointers
// to big.Int values, and returns z.
func (z *Hamilton) Lipschitz(a, b, c, d *big.Int) *Hamilton {
	z.SetL(new(Complex).Gauss(a, b))
	z.SetR(new(Complex).Gauss(c, d))
	return z
}

// Hurwitz sets z equal to a Hurwitz integer made by adding 1/2 to each of the
// four given pointers to big.Int values, and returns z.
func (z *Hamilton) Hurwitz(a, b, c, d *big.Int) *Hamilton {
	z.Lipschitz(a, b, c, d)
	half := big.NewRat(1, 2)
	z.Add(z, NewHamilton(half, half, half, half))
	return z
}

// Generate returns a random Hamilton value for quick.Check testing.
func (z *Hamilton) Generate(rand *rand.Rand, size int) reflect.Value {
	randomHamilton := &Hamilton{
		*NewComplex(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
		*NewComplex(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
	}
	return reflect.ValueOf(randomHamilton)
}
