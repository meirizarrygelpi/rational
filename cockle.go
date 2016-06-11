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

var symbCockle = [4]string{"", "i", "t", "u"}

// A Cockle represents a rational Cockle quaternion.
type Cockle struct {
	l, r Complex
}

// Cartesian returns the four rational Cartesian components of z.
func (z *Cockle) Cartesian() (*big.Rat, *big.Rat, *big.Rat, *big.Rat) {
	return &z.l.l, &z.l.r, &z.r.l, &z.r.r
}

// String returns the string representation of a Cockle value.
// If z corresponds to a + bi + ct + du, then the string is "(a+bi+ct+du)",
// similar to complex128 values.
func (z *Cockle) String() string {
	v := make([]*big.Rat, 4)
	v[0], v[1] = z.l.Cartesian()
	v[2], v[3] = z.r.Cartesian()
	a := make([]string, 9)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", v[0].RatString())
	i := 1
	for j := 2; j < 8; j = j + 2 {
		if v[i].Sign() == -1 {
			a[j] = fmt.Sprintf("%v", v[i].RatString())
		} else {
			a[j] = fmt.Sprintf("+%v", v[i].RatString())
		}
		a[j+1] = symbCockle[i]
		i++
	}
	a[8] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Cockle) Equals(y *Cockle) bool {
	if !z.l.Equals(&y.l) || !z.r.Equals(&y.r) {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *Cockle) Set(y *Cockle) *Cockle {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// NewCockle returns a pointer to the Cockle value a+bi+ct+du.
func NewCockle(a, b, c, d *big.Rat) *Cockle {
	z := new(Cockle)
	z.l.l.Set(a)
	z.l.r.Set(b)
	z.r.l.Set(c)
	z.r.r.Set(d)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Cockle) Scal(y *Cockle, a *big.Rat) *Cockle {
	z.l.Scal(&y.l, a)
	z.r.Scal(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Cockle) Neg(y *Cockle) *Cockle {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Cockle) Conj(y *Cockle) *Cockle {
	z.l.Conj(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *Cockle) Add(x, y *Cockle) *Cockle {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *Cockle) Sub(x, y *Cockle) *Cockle {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rules are:
// 		Mul(i, i) = -1
// 		Mul(t, t) = Mul(u, u) = +1
// 		Mul(i, t) = -Mul(t, i) = u
// 		Mul(u, t) = -Mul(t, u) = i
// 		Mul(u, i) = -Mul(i, u) = t
// This binary operation is noncommutative but associative.
func (z *Cockle) Mul(x, y *Cockle) *Cockle {
	a := new(Complex).Set(&x.l)
	b := new(Complex).Set(&x.r)
	c := new(Complex).Set(&y.l)
	d := new(Complex).Set(&y.r)
	temp := new(Complex)
	z.l.Add(
		z.l.Mul(a, c),
		temp.Mul(temp.Conj(d), b),
	)
	z.r.Add(
		z.r.Mul(d, a),
		temp.Mul(b, temp.Conj(c)),
	)
	return z
}

// Commutator sets z equal to the commutator of x and y
// 		Mul(x, y) - Mul(y, x)
// Then it returns z.
func (z *Cockle) Commutator(x, y *Cockle) *Cockle {
	return z.Sub(
		z.Mul(x, y),
		new(Cockle).Mul(y, x),
	)
}

// Quad returns the quadrance of z. If z = a+bi+ct+du, then the quadrance is
// 		Mul(a, a) + Mul(b, b) - Mul(c, c) - Mul(d, d)
// This can be positive, negative, or zero.
func (z *Cockle) Quad() *big.Rat {
	return new(big.Rat).Sub(
		z.l.Quad(),
		z.r.Quad(),
	)
}

// IsZeroDiv returns true if z is a zero divisor.
func (z *Cockle) IsZeroDiv() bool {
	return z.l.Quad().Cmp(z.r.Quad()) == 0
}

// Inv sets z equal to the inverse of y, and returns z.
func (z *Cockle) Inv(y *Cockle) *Cockle {
	if y.IsZeroDiv() {
		panic("inverse of zero divisor")
	}
	a := y.Quad()
	a.Inv(a)
	return z.Scal(z.Conj(y), a)
}

// QuoL sets z equal to the left quotient of x and y:
// 		Mul(Inv(y), x)
// Then it returns z.
func (z *Cockle) QuoL(x, y *Cockle) *Cockle {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Mul(z.Inv(y), x)
}

// QuoR sets z equal to the right quotient of x and y:
// 		Mul(x, Inv(y))
// Then it returns z.
func (z *Cockle) QuoR(x, y *Cockle) *Cockle {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Mul(x, z.Inv(y))
}

// IsNilpotent returns true if z raised to the n-th power vanishes.
func (z *Cockle) IsNilpotent(n int) bool {
	zero := new(Cockle)
	zeroRat := new(big.Rat)
	if z.Equals(zero) {
		return true
	}
	p := NewCockle(big.NewRat(1, 1), zeroRat, zeroRat, zeroRat)
	for i := 0; i < n; i++ {
		p.Mul(p, z)
		if p.Equals(zero) {
			return true
		}
	}
	return false
}

// Generate returns a random Cockle value for quick.Check testing.
func (z *Cockle) Generate(rand *rand.Rand, size int) reflect.Value {
	randomCockle := &Cockle{
		*NewComplex(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
		*NewComplex(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
	}
	return reflect.ValueOf(randomCockle)
}
