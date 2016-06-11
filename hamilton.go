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

// Cartesian returns the four rational Cartesian components of z.
func (z *Hamilton) Cartesian() (*big.Rat, *big.Rat, *big.Rat, *big.Rat) {
	return &z.l.l, &z.l.r, &z.r.l, &z.r.r
}

// String returns the string representation of a Hamilton value.
//
// If z corresponds to a + bi + cj + dk, then the string is"(a+bi+cj+dk)",
// similar to complex128 values.
func (z *Hamilton) String() string {
	v := make([]*big.Rat, 4)
	v[0], v[1] = z.l.Cartesian()
	v[2], v[3] = z.r.Cartesian()
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
	if !z.l.Equals(&y.l) || !z.r.Equals(&y.r) {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *Hamilton) Set(y *Hamilton) *Hamilton {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewHamilton returns a pointer to the Hamilton value a+bi+cj+dk.
func NewHamilton(a, b, c, d *big.Rat) *Hamilton {
	z := new(Hamilton)
	z.l.l.Set(a)
	z.l.r.Set(b)
	z.r.l.Set(c)
	z.r.r.Set(d)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Hamilton) Scal(y *Hamilton, a *big.Rat) *Hamilton {
	z.l.Scal(&y.l, a)
	z.r.Scal(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Hamilton) Neg(y *Hamilton) *Hamilton {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Hamilton) Conj(y *Hamilton) *Hamilton {
	z.l.Conj(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *Hamilton) Add(x, y *Hamilton) *Hamilton {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *Hamilton) Sub(x, y *Hamilton) *Hamilton {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
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
	a := new(Complex).Set(&x.l)
	b := new(Complex).Set(&x.r)
	c := new(Complex).Set(&y.l)
	d := new(Complex).Set(&y.r)
	temp := new(Complex)
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

// Commutator sets z equal to the commutator of x and y:
// 		Mul(x, y) - Mul(y, x)
// Then it returns z.
func (z *Hamilton) Commutator(x, y *Hamilton) *Hamilton {
	return z.Sub(
		z.Mul(x, y),
		new(Hamilton).Mul(y, x),
	)
}

// Quad returns the quadrance of z. If z = a+bi+cj+dk, then the quadrance is
// 		Mul(a, a) + Mul(b, b) + Mul(c, c) + Mul(d, d)
// This is always non-negative.
func (z *Hamilton) Quad() *big.Rat {
	return new(big.Rat).Add(
		z.l.Quad(),
		z.r.Quad(),
	)
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Hamilton) Inv(y *Hamilton) *Hamilton {
	if zero := new(Hamilton); y.Equals(zero) {
		panic("inverse of zero")
	}
	a := y.Quad()
	a.Inv(a)
	return z.Scal(z.Conj(y), a)
}

// QuoL sets z equal to the left quotient of x and y:
// 		Mul(Inv(y), x)
// Then it returns z.
func (z *Hamilton) QuoL(x, y *Hamilton) *Hamilton {
	if zero := new(Hamilton); y.Equals(zero) {
		panic("denominator is zero")
	}
	return z.Mul(z.Inv(y), x)
}

// QuoR sets z equal to the right quotient of x and y:
// 		Mul(x, Inv(y))
// Then it returns z.
func (z *Hamilton) QuoR(x, y *Hamilton) *Hamilton {
	if zero := new(Hamilton); y.Equals(zero) {
		panic("denominator is zero")
	}
	return z.Mul(x, z.Inv(y))
}

// Lipschitz sets z equal to the Lipschitz integer a+bi+cj+dk, and returns z.
func (z *Hamilton) Lipschitz(a, b, c, d *big.Int) *Hamilton {
	z.l.Gauss(a, b)
	z.r.Gauss(c, d)
	return z
}

// Hurwitz sets z equal to the Hurwitz integer (a+½)+(b+½)i+(c+½)j+(d+½)k,
// and returns z.
func (z *Hamilton) Hurwitz(a, b, c, d *big.Int) *Hamilton {
	z.Lipschitz(a, b, c, d)
	half := big.NewRat(1, 2)
	return z.Add(z, NewHamilton(half, half, half, half))
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
