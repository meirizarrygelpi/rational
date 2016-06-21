// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestDualComplexAddCommutative(t *testing.T) {
	f := func(x, y *DualComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(DualComplex).Add(x, y)
		r := new(DualComplex).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualComplexMulCommutative(t *testing.T) {
	f := func(x, y *DualComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(DualComplex).Mul(x, y)
		r := new(DualComplex).Mul(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualComplexNegConjCommutative(t *testing.T) {
	f := func(x *DualComplex) bool {
		// t.Logf("x = %v", x)
		l, r := new(DualComplex), new(DualComplex)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestDualComplexSubAntiCommutative(t *testing.T) {
	f := func(x, y *DualComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(DualComplex), new(DualComplex)
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

func TestDualComplexAddAssociative(t *testing.T) {
	f := func(x, y, z *DualComplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(DualComplex), new(DualComplex)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualComplexMulAssociative(t *testing.T) {
	f := func(x, y, z *DualComplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(DualComplex), new(DualComplex)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestDualComplexAddZero(t *testing.T) {
	zero := new(DualComplex)
	f := func(x *DualComplex) bool {
		// t.Logf("x = %v", x)
		l := new(DualComplex).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualComplexMulOne(t *testing.T) {
	one := &Complex{
		l: *big.NewRat(1, 1),
	}
	zero := new(Complex)
	f := func(x *DualComplex) bool {
		// t.Logf("x = %v", x)
		l := new(DualComplex).Mul(x, &DualComplex{*one, *zero})
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualComplexMulInvOne(t *testing.T) {
	one := &Complex{
		l: *big.NewRat(1, 1),
	}
	zero := new(Complex)
	f := func(x *DualComplex) bool {
		// t.Logf("x = %v", x)
		l := new(DualComplex)
		l.Mul(x, l.Inv(x))
		return l.Equals(&DualComplex{*one, *zero})
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualComplexAddNegSub(t *testing.T) {
	f := func(x, y *DualComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(DualComplex), new(DualComplex)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualComplexAddScalDouble(t *testing.T) {
	f := func(x *DualComplex) bool {
		// t.Logf("x = %v", x)
		l, r := new(DualComplex), new(DualComplex)
		l.Add(x, x)
		r.Scal(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestDualComplexInvInvolutive(t *testing.T) {
	f := func(x *DualComplex) bool {
		// t.Logf("x = %v", x)
		l := new(DualComplex)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualComplexNegInvolutive(t *testing.T) {
	f := func(x *DualComplex) bool {
		// t.Logf("x = %v", x)
		l := new(DualComplex)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualComplexConjInvolutive(t *testing.T) {
	f := func(x *DualComplex) bool {
		// t.Logf("x = %v", x)
		l := new(DualComplex)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestDualComplexMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *DualComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(DualComplex), new(DualComplex)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(DualComplex).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualComplexMulInvAntiDistributive(t *testing.T) {
	f := func(x, y *DualComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(DualComplex), new(DualComplex)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(DualComplex).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestDualComplexAddConjDistributive(t *testing.T) {
	f := func(x, y *DualComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(DualComplex), new(DualComplex)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(DualComplex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualComplexSubConjDistributive(t *testing.T) {
	f := func(x, y *DualComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(DualComplex), new(DualComplex)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(DualComplex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualComplexAddScalDistributive(t *testing.T) {
	f := func(x, y *DualComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(DualComplex), new(DualComplex)
		l.Scal(l.Add(x, y), a)
		r.Add(r.Scal(x, a), new(DualComplex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualComplexSubScalDistributive(t *testing.T) {
	f := func(x, y *DualComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(DualComplex), new(DualComplex)
		l.Scal(l.Sub(x, y), a)
		r.Sub(r.Scal(x, a), new(DualComplex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualComplexAddMulDistributive(t *testing.T) {
	f := func(x, y, z *DualComplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(DualComplex), new(DualComplex)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(DualComplex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualComplexSubMulDistributive(t *testing.T) {
	f := func(x, y, z *DualComplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(DualComplex), new(DualComplex)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(DualComplex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Positivity

func TestDualComplexQuadPositive(t *testing.T) {
	f := func(x *DualComplex) bool {
		// t.Logf("x = %v", x)
		return x.Quad().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func TestDualComplexComposition(t *testing.T) {
	f := func(x, y *DualComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(DualComplex)
		a, b := new(big.Rat), new(big.Rat)
		p.Mul(x, y)
		a.Set(p.Quad())
		b.Mul(x.Quad(), y.Quad())
		return a.Cmp(b) == 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
