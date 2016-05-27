// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestSupraPerplexAddCommutative(t *testing.T) {
	f := func(x, y *SupraPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(SupraPerplex).Add(x, y)
		r := new(SupraPerplex).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSupraPerplexNegConjCommutative(t *testing.T) {
	f := func(x *SupraPerplex) bool {
		// t.Logf("x = %v", x)
		l, r := new(SupraPerplex), new(SupraPerplex)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Non-commutativity

func TestSupraPerplexMulNonCommutative(t *testing.T) {
	f := func(x, y *SupraPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(SupraPerplex).Commutator(x, y)
		zero := new(SupraPerplex)
		return !l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestSupraPerplexSubAntiCommutative(t *testing.T) {
	f := func(x, y *SupraPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(SupraPerplex), new(SupraPerplex)
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

func TestSupraPerplexAddAssociative(t *testing.T) {
	f := func(x, y, z *SupraPerplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(SupraPerplex), new(SupraPerplex)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Non-associativity

func TestSupraPerplexMulNonAssociative(t *testing.T) {
	f := func(x, y, z *SupraPerplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l := new(SupraPerplex).Associator(x, y, z)
		zero := new(SupraPerplex)
		return !l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestSupraPerplexAddZero(t *testing.T) {
	zero := new(SupraPerplex)
	f := func(x *SupraPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(SupraPerplex).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSupraPerplexMulOne(t *testing.T) {
	one := &InfraPerplex{
		l: Perplex{
			l: *big.NewRat(1, 1),
		},
	}
	zero := new(InfraPerplex)
	f := func(x *SupraPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(SupraPerplex).Mul(x, &SupraPerplex{*one, *zero})
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSupraPerplexMulInvOne(t *testing.T) {
	one := &InfraPerplex{
		l: Perplex{
			l: *big.NewRat(1, 1),
		},
	}
	zero := new(InfraPerplex)
	f := func(x *SupraPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(SupraPerplex)
		l.Mul(x, l.Inv(x))
		return l.Equals(&SupraPerplex{*one, *zero})
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSupraPerplexAddNegSub(t *testing.T) {
	f := func(x, y *SupraPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(SupraPerplex), new(SupraPerplex)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSupraPerplexAddScalDouble(t *testing.T) {
	f := func(x *SupraPerplex) bool {
		// t.Logf("x = %v", x)
		l, r := new(SupraPerplex), new(SupraPerplex)
		l.Add(x, x)
		r.Scal(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestSupraPerplexNegInvolutive(t *testing.T) {
	f := func(x *SupraPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(SupraPerplex)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSupraPerplexConjInvolutive(t *testing.T) {
	f := func(x *SupraPerplex) bool {
		// t.Logf("x = %v", x)
		l := new(SupraPerplex)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestSupraPerplexMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *SupraPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(SupraPerplex), new(SupraPerplex)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(SupraPerplex).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSupraPerplexMulInvAntiDistributive(t *testing.T) {
	f := func(x, y *SupraPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(SupraPerplex), new(SupraPerplex)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(SupraPerplex).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestSupraPerplexAddConjDistributive(t *testing.T) {
	f := func(x, y *SupraPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(SupraPerplex), new(SupraPerplex)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(SupraPerplex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSupraPerplexSubConjDistributive(t *testing.T) {
	f := func(x, y *SupraPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(SupraPerplex), new(SupraPerplex)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(SupraPerplex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSupraPerplexAddScalDistributive(t *testing.T) {
	f := func(x, y *SupraPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(SupraPerplex), new(SupraPerplex)
		l.Scal(l.Add(x, y), a)
		r.Add(r.Scal(x, a), new(SupraPerplex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSupraPerplexSubScalDistributive(t *testing.T) {
	f := func(x, y *SupraPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(SupraPerplex), new(SupraPerplex)
		l.Scal(l.Sub(x, y), a)
		r.Sub(r.Scal(x, a), new(SupraPerplex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSupraPerplexAddMulDistributive(t *testing.T) {
	f := func(x, y, z *SupraPerplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(SupraPerplex), new(SupraPerplex)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(SupraPerplex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSupraPerplexSubMulDistributive(t *testing.T) {
	f := func(x, y, z *SupraPerplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(SupraPerplex), new(SupraPerplex)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(SupraPerplex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Left-alternativity

func TestSupraPerplexLeftAlternative(t *testing.T) {
	f := func(x, y *SupraPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(SupraPerplex)
		l.Associator(x, x, y)
		zero := new(SupraPerplex)
		return l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Right-alternativity

func TestSupraPerplexRightAlternative(t *testing.T) {
	f := func(x, y *SupraPerplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(SupraPerplex)
		l.Associator(x, y, y)
		zero := new(SupraPerplex)
		return l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
