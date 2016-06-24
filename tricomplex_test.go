// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestTriComplexAddCommutative(t *testing.T) {
	f := func(x, y *TriComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(TriComplex).Add(x, y)
		r := new(TriComplex).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriComplexMulCommutative(t *testing.T) {
	f := func(x, y *TriComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(TriComplex).Mul(x, y)
		r := new(TriComplex).Mul(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriComplexNegConjCommutative(t *testing.T) {
	f := func(x *TriComplex) bool {
		// t.Logf("x = %v", x)
		l, r := new(TriComplex), new(TriComplex)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriComplexStarConjCommutative(t *testing.T) {
	f := func(x *TriComplex) bool {
		// t.Logf("x = %v", x)
		l, r := new(TriComplex), new(TriComplex)
		l.Star(l.Conj(x))
		r.Conj(r.Star(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestTriComplexSubAntiCommutative(t *testing.T) {
	f := func(x, y *TriComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(TriComplex), new(TriComplex)
		l.Sub(x, y)
		r.Sub(y, x)
		r.Neg(r)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Associativity

func TestTriComplexAddAssociative(t *testing.T) {
	f := func(x, y, z *TriComplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(TriComplex), new(TriComplex)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriComplexMulAssociative(t *testing.T) {
	f := func(x, y, z *TriComplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(TriComplex), new(TriComplex)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestTriComplexAddZero(t *testing.T) {
	zero := new(TriComplex)
	f := func(x *TriComplex) bool {
		// t.Logf("x = %v", x)
		l := new(TriComplex).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriComplexMulOne(t *testing.T) {
	one := &BiComplex{
		l: Complex{
			l: *big.NewRat(1, 1),
		},
	}
	zero := new(BiComplex)
	f := func(x *TriComplex) bool {
		// t.Logf("x = %v", x)
		l := new(TriComplex).Mul(x, &TriComplex{*one, *zero})
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriComplexMulInvOne(t *testing.T) {
	one := &BiComplex{
		l: Complex{
			l: *big.NewRat(1, 1),
		},
	}
	zero := new(BiComplex)
	f := func(x *TriComplex) bool {
		// t.Logf("x = %v", x)
		l := new(TriComplex)
		l.Mul(x, l.Inv(x))
		return l.Equals(&TriComplex{*one, *zero})
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriComplexAddNegSub(t *testing.T) {
	f := func(x, y *TriComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(TriComplex), new(TriComplex)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriComplexAddScalDouble(t *testing.T) {
	f := func(x *TriComplex) bool {
		// t.Logf("x = %v", x)
		l, r := new(TriComplex), new(TriComplex)
		l.Add(x, x)
		r.Scal(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestTriComplexInvInvolutive(t *testing.T) {
	f := func(x *TriComplex) bool {
		// t.Logf("x = %v", x)
		l := new(TriComplex)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriComplexNegInvolutive(t *testing.T) {
	f := func(x *TriComplex) bool {
		// t.Logf("x = %v", x)
		l := new(TriComplex)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriComplexConjInvolutive(t *testing.T) {
	f := func(x *TriComplex) bool {
		// t.Logf("x = %v", x)
		l := new(TriComplex)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriComplexStarInvolutive(t *testing.T) {
	f := func(x *TriComplex) bool {
		// t.Logf("x = %v", x)
		l := new(TriComplex)
		l.Star(l.Star(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestTriComplexMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *TriComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(TriComplex), new(TriComplex)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(TriComplex).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriComplexMulInvAntiDistributive(t *testing.T) {
	f := func(x, y *TriComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(TriComplex), new(TriComplex)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(TriComplex).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestTriComplexAddConjDistributive(t *testing.T) {
	f := func(x, y *TriComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(TriComplex), new(TriComplex)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(TriComplex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriComplexSubConjDistributive(t *testing.T) {
	f := func(x, y *TriComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(TriComplex), new(TriComplex)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(TriComplex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriComplexAddScalDistributive(t *testing.T) {
	f := func(x, y *TriComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(TriComplex), new(TriComplex)
		l.Scal(l.Add(x, y), a)
		r.Add(r.Scal(x, a), new(TriComplex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriComplexSubScalDistributive(t *testing.T) {
	f := func(x, y *TriComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(TriComplex), new(TriComplex)
		l.Scal(l.Sub(x, y), a)
		r.Sub(r.Scal(x, a), new(TriComplex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriComplexAddMulDistributive(t *testing.T) {
	f := func(x, y, z *TriComplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(TriComplex), new(TriComplex)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(TriComplex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriComplexSubMulDistributive(t *testing.T) {
	f := func(x, y, z *TriComplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(TriComplex), new(TriComplex)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(TriComplex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Positivity

func TestTriComplexNormPositive(t *testing.T) {
	f := func(x *TriComplex) bool {
		// t.Logf("x = %v", x)
		return x.Norm().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func TestTriComplexComposition(t *testing.T) {
	f := func(x, y *TriComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(TriComplex)
		a, b := new(big.Rat), new(big.Rat)
		p.Mul(x, y)
		a.Set(p.Norm())
		b.Mul(x.Norm(), y.Norm())
		return a.Cmp(b) == 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
