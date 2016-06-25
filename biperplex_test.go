// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestBiPerplexAddCommutative(t *testing.T) {
	f := func(x, y *BiPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(BiPerplex).Add(x, y)
		r := new(BiPerplex).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiPerplexMulCommutative(t *testing.T) {
	f := func(x, y *BiPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(BiPerplex).Mul(x, y)
		r := new(BiPerplex).Mul(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiPerplexNegConjCommutative(t *testing.T) {
	f := func(x *BiPerplex) bool {
		// t.Logf("x = %v", x)
		l, r := new(BiPerplex), new(BiPerplex)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestBiPerplexSubAntiCommutative(t *testing.T) {
	f := func(x, y *BiPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(BiPerplex), new(BiPerplex)
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

func TestBiPerplexAddAssociative(t *testing.T) {
	f := func(x, y, z *BiPerplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(BiPerplex), new(BiPerplex)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiPerplexMulAssociative(t *testing.T) {
	f := func(x, y, z *BiPerplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(BiPerplex), new(BiPerplex)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestBiPerplexAddZero(t *testing.T) {
	zero := new(BiPerplex)
	f := func(x *BiPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(BiPerplex).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiPerplexMulOne(t *testing.T) {
	one := &Perplex{
		l: *big.NewRat(1, 1),
	}
	zero := new(Perplex)
	f := func(x *BiPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(BiPerplex).Mul(x, &BiPerplex{*one, *zero})
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiPerplexMulInvOne(t *testing.T) {
	one := &Perplex{
		l: *big.NewRat(1, 1),
	}
	zero := new(Perplex)
	f := func(x *BiPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(BiPerplex)
		l.Mul(x, l.Inv(x))
		return l.Equals(&BiPerplex{*one, *zero})
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiPerplexAddNegSub(t *testing.T) {
	f := func(x, y *BiPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(BiPerplex), new(BiPerplex)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiPerplexAddScalDouble(t *testing.T) {
	f := func(x *BiPerplex) bool {
		// t.Logf("x = %v", x)
		l, r := new(BiPerplex), new(BiPerplex)
		l.Add(x, x)
		r.Scal(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestBiPerplexInvInvolutive(t *testing.T) {
	f := func(x *BiPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(BiPerplex)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiPerplexNegInvolutive(t *testing.T) {
	f := func(x *BiPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(BiPerplex)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiPerplexConjInvolutive(t *testing.T) {
	f := func(x *BiPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(BiPerplex)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestBiPerplexMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *BiPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(BiPerplex), new(BiPerplex)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(BiPerplex).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiPerplexMulInvAntiDistributive(t *testing.T) {
	f := func(x, y *BiPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(BiPerplex), new(BiPerplex)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(BiPerplex).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestBiPerplexAddConjDistributive(t *testing.T) {
	f := func(x, y *BiPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(BiPerplex), new(BiPerplex)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(BiPerplex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiPerplexSubConjDistributive(t *testing.T) {
	f := func(x, y *BiPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(BiPerplex), new(BiPerplex)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(BiPerplex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiPerplexAddScalDistributive(t *testing.T) {
	f := func(x, y *BiPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(BiPerplex), new(BiPerplex)
		l.Scal(l.Add(x, y), a)
		r.Add(r.Scal(x, a), new(BiPerplex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiPerplexSubScalDistributive(t *testing.T) {
	f := func(x, y *BiPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(BiPerplex), new(BiPerplex)
		l.Scal(l.Sub(x, y), a)
		r.Sub(r.Scal(x, a), new(BiPerplex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiPerplexAddMulDistributive(t *testing.T) {
	f := func(x, y, z *BiPerplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(BiPerplex), new(BiPerplex)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(BiPerplex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiPerplexSubMulDistributive(t *testing.T) {
	f := func(x, y, z *BiPerplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(BiPerplex), new(BiPerplex)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(BiPerplex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func TestBiPerplexComposition(t *testing.T) {
	f := func(x, y *BiPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(BiPerplex)
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
