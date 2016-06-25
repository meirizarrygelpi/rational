// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestTriNilplexAddCommutative(t *testing.T) {
	f := func(x, y *TriNilplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(TriNilplex).Add(x, y)
		r := new(TriNilplex).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriNilplexMulCommutative(t *testing.T) {
	f := func(x, y *TriNilplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(TriNilplex).Mul(x, y)
		r := new(TriNilplex).Mul(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriNilplexNegConjCommutative(t *testing.T) {
	f := func(x *TriNilplex) bool {
		// t.Logf("x = %v", x)
		l, r := new(TriNilplex), new(TriNilplex)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestTriNilplexSubAntiCommutative(t *testing.T) {
	f := func(x, y *TriNilplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(TriNilplex), new(TriNilplex)
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

func TestTriNilplexAddAssociative(t *testing.T) {
	f := func(x, y, z *TriNilplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(TriNilplex), new(TriNilplex)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriNilplexMulAssociative(t *testing.T) {
	f := func(x, y, z *TriNilplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(TriNilplex), new(TriNilplex)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestTriNilplexAddZero(t *testing.T) {
	zero := new(TriNilplex)
	f := func(x *TriNilplex) bool {
		// t.Logf("x = %v", x)
		l := new(TriNilplex).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriNilplexMulOne(t *testing.T) {
	one := &Hyper{
		l: Infra{
			l: *big.NewRat(1, 1),
		},
	}
	zero := new(Hyper)
	f := func(x *TriNilplex) bool {
		// t.Logf("x = %v", x)
		l := new(TriNilplex).Mul(x, &TriNilplex{*one, *zero})
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriNilplexMulInvOne(t *testing.T) {
	one := &Hyper{
		l: Infra{
			l: *big.NewRat(1, 1),
		},
	}
	zero := new(Hyper)
	f := func(x *TriNilplex) bool {
		// t.Logf("x = %v", x)
		l := new(TriNilplex)
		l.Mul(x, l.Inv(x))
		return l.Equals(&TriNilplex{*one, *zero})
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriNilplexAddNegSub(t *testing.T) {
	f := func(x, y *TriNilplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(TriNilplex), new(TriNilplex)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriNilplexAddScalDouble(t *testing.T) {
	f := func(x *TriNilplex) bool {
		// t.Logf("x = %v", x)
		l, r := new(TriNilplex), new(TriNilplex)
		l.Add(x, x)
		r.Scal(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestTriNilplexInvInvolutive(t *testing.T) {
	f := func(x *TriNilplex) bool {
		// t.Logf("x = %v", x)
		l := new(TriNilplex)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriNilplexNegInvolutive(t *testing.T) {
	f := func(x *TriNilplex) bool {
		// t.Logf("x = %v", x)
		l := new(TriNilplex)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriNilplexConjInvolutive(t *testing.T) {
	f := func(x *TriNilplex) bool {
		// t.Logf("x = %v", x)
		l := new(TriNilplex)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestTriNilplexMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *TriNilplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(TriNilplex), new(TriNilplex)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(TriNilplex).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriNilplexMulInvAntiDistributive(t *testing.T) {
	f := func(x, y *TriNilplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(TriNilplex), new(TriNilplex)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(TriNilplex).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestTriNilplexAddConjDistributive(t *testing.T) {
	f := func(x, y *TriNilplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(TriNilplex), new(TriNilplex)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(TriNilplex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriNilplexSubConjDistributive(t *testing.T) {
	f := func(x, y *TriNilplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(TriNilplex), new(TriNilplex)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(TriNilplex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriNilplexAddScalDistributive(t *testing.T) {
	f := func(x, y *TriNilplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(TriNilplex), new(TriNilplex)
		l.Scal(l.Add(x, y), a)
		r.Add(r.Scal(x, a), new(TriNilplex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriNilplexSubScalDistributive(t *testing.T) {
	f := func(x, y *TriNilplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(TriNilplex), new(TriNilplex)
		l.Scal(l.Sub(x, y), a)
		r.Sub(r.Scal(x, a), new(TriNilplex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriNilplexAddMulDistributive(t *testing.T) {
	f := func(x, y, z *TriNilplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(TriNilplex), new(TriNilplex)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(TriNilplex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriNilplexSubMulDistributive(t *testing.T) {
	f := func(x, y, z *TriNilplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(TriNilplex), new(TriNilplex)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(TriNilplex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Positivity

func TestTriNilplexNormPositive(t *testing.T) {
	f := func(x *TriNilplex) bool {
		// t.Logf("x = %v", x)
		return x.Norm().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func TestTriNilplexComposition(t *testing.T) {
	f := func(x, y *TriNilplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(TriNilplex)
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
