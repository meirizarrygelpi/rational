// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestTriPerplexAddCommutative(t *testing.T) {
	f := func(x, y *TriPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(TriPerplex).Add(x, y)
		r := new(TriPerplex).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriPerplexMulCommutative(t *testing.T) {
	f := func(x, y *TriPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(TriPerplex).Mul(x, y)
		r := new(TriPerplex).Mul(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriPerplexNegConjCommutative(t *testing.T) {
	f := func(x *TriPerplex) bool {
		// t.Logf("x = %v", x)
		l, r := new(TriPerplex), new(TriPerplex)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestTriPerplexSubAntiCommutative(t *testing.T) {
	f := func(x, y *TriPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(TriPerplex), new(TriPerplex)
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

func TestTriPerplexAddAssociative(t *testing.T) {
	f := func(x, y, z *TriPerplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(TriPerplex), new(TriPerplex)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriPerplexMulAssociative(t *testing.T) {
	f := func(x, y, z *TriPerplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(TriPerplex), new(TriPerplex)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestTriPerplexAddZero(t *testing.T) {
	zero := new(TriPerplex)
	f := func(x *TriPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(TriPerplex).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriPerplexMulOne(t *testing.T) {
	one := &BiPerplex{
		l: Perplex{
			l: *big.NewRat(1, 1),
		},
	}
	zero := new(BiPerplex)
	f := func(x *TriPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(TriPerplex).Mul(x, &TriPerplex{*one, *zero})
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriPerplexMulInvOne(t *testing.T) {
	one := &BiPerplex{
		l: Perplex{
			l: *big.NewRat(1, 1),
		},
	}
	zero := new(BiPerplex)
	f := func(x *TriPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(TriPerplex)
		l.Mul(x, l.Inv(x))
		return l.Equals(&TriPerplex{*one, *zero})
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriPerplexAddNegSub(t *testing.T) {
	f := func(x, y *TriPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(TriPerplex), new(TriPerplex)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriPerplexAddScalDouble(t *testing.T) {
	f := func(x *TriPerplex) bool {
		// t.Logf("x = %v", x)
		l, r := new(TriPerplex), new(TriPerplex)
		l.Add(x, x)
		r.Scal(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestTriPerplexInvInvolutive(t *testing.T) {
	f := func(x *TriPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(TriPerplex)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriPerplexNegInvolutive(t *testing.T) {
	f := func(x *TriPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(TriPerplex)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriPerplexConjInvolutive(t *testing.T) {
	f := func(x *TriPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(TriPerplex)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestTriPerplexMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *TriPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(TriPerplex), new(TriPerplex)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(TriPerplex).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriPerplexMulInvAntiDistributive(t *testing.T) {
	f := func(x, y *TriPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(TriPerplex), new(TriPerplex)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(TriPerplex).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestTriPerplexAddConjDistributive(t *testing.T) {
	f := func(x, y *TriPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(TriPerplex), new(TriPerplex)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(TriPerplex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriPerplexSubConjDistributive(t *testing.T) {
	f := func(x, y *TriPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(TriPerplex), new(TriPerplex)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(TriPerplex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriPerplexAddScalDistributive(t *testing.T) {
	f := func(x, y *TriPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(TriPerplex), new(TriPerplex)
		l.Scal(l.Add(x, y), a)
		r.Add(r.Scal(x, a), new(TriPerplex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriPerplexSubScalDistributive(t *testing.T) {
	f := func(x, y *TriPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(TriPerplex), new(TriPerplex)
		l.Scal(l.Sub(x, y), a)
		r.Sub(r.Scal(x, a), new(TriPerplex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriPerplexAddMulDistributive(t *testing.T) {
	f := func(x, y, z *TriPerplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(TriPerplex), new(TriPerplex)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(TriPerplex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTriPerplexSubMulDistributive(t *testing.T) {
	f := func(x, y, z *TriPerplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(TriPerplex), new(TriPerplex)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(TriPerplex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func TestTriPerplexComposition(t *testing.T) {
	f := func(x, y *TriPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(TriPerplex)
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
