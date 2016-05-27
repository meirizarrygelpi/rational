// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestInfraPerplexAddCommutative(t *testing.T) {
	f := func(x, y *InfraPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(InfraPerplex).Add(x, y)
		r := new(InfraPerplex).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraPerplexNegConjCommutative(t *testing.T) {
	f := func(x *InfraPerplex) bool {
		// t.Logf("x = %v", x)
		l, r := new(InfraPerplex), new(InfraPerplex)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Non-commutativity

func TestInfraPerplexMulNonCommutative(t *testing.T) {
	f := func(x, y *InfraPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(InfraPerplex).Commutator(x, y)
		zero := new(InfraPerplex)
		return !l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestInfraPerplexSubAntiCommutative(t *testing.T) {
	f := func(x, y *InfraPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(InfraPerplex), new(InfraPerplex)
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

func TestInfraPerplexAddAssociative(t *testing.T) {
	f := func(x, y, z *InfraPerplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(InfraPerplex), new(InfraPerplex)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraPerplexMulAssociative(t *testing.T) {
	f := func(x, y, z *InfraPerplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(InfraPerplex), new(InfraPerplex)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestInfraPerplexAddZero(t *testing.T) {
	zero := new(InfraPerplex)
	f := func(x *InfraPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(InfraPerplex).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraPerplexMulOne(t *testing.T) {
	one := &Perplex{
		l: *big.NewRat(1, 1),
	}
	zero := new(Perplex)
	f := func(x *InfraPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(InfraPerplex).Mul(x, &InfraPerplex{*one, *zero})
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraPerplexMulInvOne(t *testing.T) {
	one := &Perplex{
		l: *big.NewRat(1, 1),
	}
	zero := new(Perplex)
	f := func(x *InfraPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(InfraPerplex)
		l.Mul(x, l.Inv(x))
		return l.Equals(&InfraPerplex{*one, *zero})
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraPerplexAddNegSub(t *testing.T) {
	f := func(x, y *InfraPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(InfraPerplex), new(InfraPerplex)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraPerplexAddScalDouble(t *testing.T) {
	f := func(x *InfraPerplex) bool {
		// t.Logf("x = %v", x)
		l, r := new(InfraPerplex), new(InfraPerplex)
		l.Add(x, x)
		r.Scal(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestInfraPerplexInvInvolutive(t *testing.T) {
	f := func(x *InfraPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(InfraPerplex)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraPerplexNegInvolutive(t *testing.T) {
	f := func(x *InfraPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(InfraPerplex)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraPerplexConjInvolutive(t *testing.T) {
	f := func(x *InfraPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(InfraPerplex)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestInfraPerplexMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *InfraPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(InfraPerplex), new(InfraPerplex)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(InfraPerplex).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraPerplexMulInvAntiDistributive(t *testing.T) {
	f := func(x, y *InfraPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(InfraPerplex), new(InfraPerplex)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(InfraPerplex).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestInfraPerplexAddConjDistributive(t *testing.T) {
	f := func(x, y *InfraPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(InfraPerplex), new(InfraPerplex)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(InfraPerplex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraPerplexSubConjDistributive(t *testing.T) {
	f := func(x, y *InfraPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(InfraPerplex), new(InfraPerplex)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(InfraPerplex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraPerplexAddScalDistributive(t *testing.T) {
	f := func(x, y *InfraPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(InfraPerplex), new(InfraPerplex)
		l.Scal(l.Add(x, y), a)
		r.Add(r.Scal(x, a), new(InfraPerplex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraPerplexSubScalDistributive(t *testing.T) {
	f := func(x, y *InfraPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(InfraPerplex), new(InfraPerplex)
		l.Scal(l.Sub(x, y), a)
		r.Sub(r.Scal(x, a), new(InfraPerplex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraPerplexAddMulDistributive(t *testing.T) {
	f := func(x, y, z *InfraPerplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(InfraPerplex), new(InfraPerplex)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(InfraPerplex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraPerplexSubMulDistributive(t *testing.T) {
	f := func(x, y, z *InfraPerplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(InfraPerplex), new(InfraPerplex)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(InfraPerplex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func TestInfraPerplexComposition(t *testing.T) {
	f := func(x, y *InfraPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(InfraPerplex)
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
