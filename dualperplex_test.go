// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestDualPerplexAddCommutative(t *testing.T) {
	f := func(x, y *DualPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(DualPerplex).Add(x, y)
		r := new(DualPerplex).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualPerplexMulCommutative(t *testing.T) {
	f := func(x, y *DualPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(DualPerplex).Mul(x, y)
		r := new(DualPerplex).Mul(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualPerplexNegConjCommutative(t *testing.T) {
	f := func(x *DualPerplex) bool {
		// t.Logf("x = %v", x)
		l, r := new(DualPerplex), new(DualPerplex)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestDualPerplexSubAntiCommutative(t *testing.T) {
	f := func(x, y *DualPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(DualPerplex), new(DualPerplex)
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

func TestDualPerplexAddAssociative(t *testing.T) {
	f := func(x, y, z *DualPerplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(DualPerplex), new(DualPerplex)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualPerplexMulAssociative(t *testing.T) {
	f := func(x, y, z *DualPerplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(DualPerplex), new(DualPerplex)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestDualPerplexAddZero(t *testing.T) {
	zero := new(DualPerplex)
	f := func(x *DualPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(DualPerplex).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualPerplexMulOne(t *testing.T) {
	one := &Perplex{
		l: *big.NewRat(1, 1),
	}
	zero := new(Perplex)
	f := func(x *DualPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(DualPerplex).Mul(x, &DualPerplex{*one, *zero})
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualPerplexMulInvOne(t *testing.T) {
	one := &Perplex{
		l: *big.NewRat(1, 1),
	}
	zero := new(Perplex)
	f := func(x *DualPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(DualPerplex)
		l.Mul(x, l.Inv(x))
		return l.Equals(&DualPerplex{*one, *zero})
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualPerplexAddNegSub(t *testing.T) {
	f := func(x, y *DualPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(DualPerplex), new(DualPerplex)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualPerplexAddScalDouble(t *testing.T) {
	f := func(x *DualPerplex) bool {
		// t.Logf("x = %v", x)
		l, r := new(DualPerplex), new(DualPerplex)
		l.Add(x, x)
		r.Scal(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestDualPerplexInvInvolutive(t *testing.T) {
	f := func(x *DualPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(DualPerplex)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualPerplexNegInvolutive(t *testing.T) {
	f := func(x *DualPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(DualPerplex)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualPerplexConjInvolutive(t *testing.T) {
	f := func(x *DualPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(DualPerplex)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestDualPerplexMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *DualPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(DualPerplex), new(DualPerplex)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(DualPerplex).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualPerplexMulInvAntiDistributive(t *testing.T) {
	f := func(x, y *DualPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(DualPerplex), new(DualPerplex)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(DualPerplex).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestDualPerplexAddConjDistributive(t *testing.T) {
	f := func(x, y *DualPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(DualPerplex), new(DualPerplex)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(DualPerplex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualPerplexSubConjDistributive(t *testing.T) {
	f := func(x, y *DualPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(DualPerplex), new(DualPerplex)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(DualPerplex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualPerplexAddScalDistributive(t *testing.T) {
	f := func(x, y *DualPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(DualPerplex), new(DualPerplex)
		l.Scal(l.Add(x, y), a)
		r.Add(r.Scal(x, a), new(DualPerplex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualPerplexSubScalDistributive(t *testing.T) {
	f := func(x, y *DualPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(DualPerplex), new(DualPerplex)
		l.Scal(l.Sub(x, y), a)
		r.Sub(r.Scal(x, a), new(DualPerplex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualPerplexAddMulDistributive(t *testing.T) {
	f := func(x, y, z *DualPerplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(DualPerplex), new(DualPerplex)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(DualPerplex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDualPerplexSubMulDistributive(t *testing.T) {
	f := func(x, y, z *DualPerplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(DualPerplex), new(DualPerplex)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(DualPerplex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Positivity

func TestDualPerplexNormPositive(t *testing.T) {
	f := func(x *DualPerplex) bool {
		// t.Logf("x = %v", x)
		return x.Norm().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func TestDualPerplexComposition(t *testing.T) {
	f := func(x, y *DualPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(DualPerplex)
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
