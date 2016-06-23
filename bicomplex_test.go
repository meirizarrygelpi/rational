// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestBiComplexAddCommutative(t *testing.T) {
	f := func(x, y *BiComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(BiComplex).Add(x, y)
		r := new(BiComplex).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiComplexMulCommutative(t *testing.T) {
	f := func(x, y *BiComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(BiComplex).Mul(x, y)
		r := new(BiComplex).Mul(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiComplexNegConjCommutative(t *testing.T) {
	f := func(x *BiComplex) bool {
		// t.Logf("x = %v", x)
		l, r := new(BiComplex), new(BiComplex)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestBiComplexSubAntiCommutative(t *testing.T) {
	f := func(x, y *BiComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(BiComplex), new(BiComplex)
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

func TestBiComplexAddAssociative(t *testing.T) {
	f := func(x, y, z *BiComplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(BiComplex), new(BiComplex)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiComplexMulAssociative(t *testing.T) {
	f := func(x, y, z *BiComplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(BiComplex), new(BiComplex)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestBiComplexAddZero(t *testing.T) {
	zero := new(BiComplex)
	f := func(x *BiComplex) bool {
		// t.Logf("x = %v", x)
		l := new(BiComplex).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiComplexMulOne(t *testing.T) {
	one := &Complex{
		l: *big.NewRat(1, 1),
	}
	zero := new(Complex)
	f := func(x *BiComplex) bool {
		// t.Logf("x = %v", x)
		l := new(BiComplex).Mul(x, &BiComplex{*one, *zero})
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiComplexMulInvOne(t *testing.T) {
	one := &Complex{
		l: *big.NewRat(1, 1),
	}
	zero := new(Complex)
	f := func(x *BiComplex) bool {
		// t.Logf("x = %v", x)
		l := new(BiComplex)
		l.Mul(x, l.Inv(x))
		return l.Equals(&BiComplex{*one, *zero})
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiComplexAddNegSub(t *testing.T) {
	f := func(x, y *BiComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(BiComplex), new(BiComplex)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiComplexAddScalDouble(t *testing.T) {
	f := func(x *BiComplex) bool {
		// t.Logf("x = %v", x)
		l, r := new(BiComplex), new(BiComplex)
		l.Add(x, x)
		r.Scal(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestBiComplexInvInvolutive(t *testing.T) {
	f := func(x *BiComplex) bool {
		// t.Logf("x = %v", x)
		l := new(BiComplex)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiComplexNegInvolutive(t *testing.T) {
	f := func(x *BiComplex) bool {
		// t.Logf("x = %v", x)
		l := new(BiComplex)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiComplexConjInvolutive(t *testing.T) {
	f := func(x *BiComplex) bool {
		// t.Logf("x = %v", x)
		l := new(BiComplex)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestBiComplexMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *BiComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(BiComplex), new(BiComplex)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(BiComplex).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiComplexMulInvAntiDistributive(t *testing.T) {
	f := func(x, y *BiComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(BiComplex), new(BiComplex)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(BiComplex).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestBiComplexAddConjDistributive(t *testing.T) {
	f := func(x, y *BiComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(BiComplex), new(BiComplex)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(BiComplex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiComplexSubConjDistributive(t *testing.T) {
	f := func(x, y *BiComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(BiComplex), new(BiComplex)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(BiComplex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiComplexAddScalDistributive(t *testing.T) {
	f := func(x, y *BiComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(BiComplex), new(BiComplex)
		l.Scal(l.Add(x, y), a)
		r.Add(r.Scal(x, a), new(BiComplex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiComplexSubScalDistributive(t *testing.T) {
	f := func(x, y *BiComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(BiComplex), new(BiComplex)
		l.Scal(l.Sub(x, y), a)
		r.Sub(r.Scal(x, a), new(BiComplex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiComplexAddMulDistributive(t *testing.T) {
	f := func(x, y, z *BiComplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(BiComplex), new(BiComplex)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(BiComplex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiComplexSubMulDistributive(t *testing.T) {
	f := func(x, y, z *BiComplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(BiComplex), new(BiComplex)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(BiComplex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Positivity

func TestBiComplexNormPositive(t *testing.T) {
	f := func(x *BiComplex) bool {
		// t.Logf("x = %v", x)
		return x.Norm().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func TestBiComplexComposition(t *testing.T) {
	f := func(x, y *BiComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(BiComplex)
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
