// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestPerplexAddCommutative(t *testing.T) {
	f := func(x, y *Perplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Perplex).Add(x, y)
		r := new(Perplex).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestPerplexMulCommutative(t *testing.T) {
	f := func(x, y *Perplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Perplex).Mul(x, y)
		r := new(Perplex).Mul(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestPerplexNegConjCommutative(t *testing.T) {
	f := func(x *Perplex) bool {
		// t.Logf("x = %v", x)
		l, r := new(Perplex), new(Perplex)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestPerplexSubAntiCommutative(t *testing.T) {
	f := func(x, y *Perplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Perplex), new(Perplex)
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

func TestPerplexAddAssociative(t *testing.T) {
	f := func(x, y, z *Perplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Perplex), new(Perplex)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestPerplexMulAssociative(t *testing.T) {
	f := func(x, y, z *Perplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Perplex), new(Perplex)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestPerplexAddZero(t *testing.T) {
	zero := new(Perplex)
	f := func(x *Perplex) bool {
		// t.Logf("x = %v", x)
		l := new(Perplex).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestPerplexMulOne(t *testing.T) {
	one := &Perplex{
		l: *big.NewRat(1, 1),
	}
	f := func(x *Perplex) bool {
		// t.Logf("x = %v", x)
		l := new(Perplex).Mul(x, one)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestPerplexMulInvOne(t *testing.T) {
	one := &Perplex{
		l: *big.NewRat(1, 1),
	}
	f := func(x *Perplex) bool {
		// t.Logf("x = %v", x)
		l := new(Perplex)
		l.Mul(x, l.Inv(x))
		return l.Equals(one)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestPerplexAddNegSub(t *testing.T) {
	f := func(x, y *Perplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Perplex), new(Perplex)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestPerplexAddScalDouble(t *testing.T) {
	f := func(x *Perplex) bool {
		// t.Logf("x = %v", x)
		l, r := new(Perplex), new(Perplex)
		l.Add(x, x)
		r.Scal(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestPerplexNegInvolutive(t *testing.T) {
	f := func(x *Perplex) bool {
		// t.Logf("x = %v", x)
		l := new(Perplex)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestPerplexConjInvolutive(t *testing.T) {
	f := func(x *Perplex) bool {
		// t.Logf("x = %v", x)
		l := new(Perplex)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestPerplexMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *Perplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Perplex), new(Perplex)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(Perplex).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestPerplexMulInvAntiDistributive(t *testing.T) {
	f := func(x, y *Perplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Perplex), new(Perplex)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(Perplex).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestPerplexAddConjDistributive(t *testing.T) {
	f := func(x, y *Perplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Perplex), new(Perplex)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(Perplex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestPerplexSubConjDistributive(t *testing.T) {
	f := func(x, y *Perplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Perplex), new(Perplex)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(Perplex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestPerplexAddScalDistributive(t *testing.T) {
	f := func(x, y *Perplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Perplex), new(Perplex)
		l.Scal(l.Add(x, y), a)
		r.Add(r.Scal(x, a), new(Perplex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestPerplexSubScalDistributive(t *testing.T) {
	f := func(x, y *Perplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Perplex), new(Perplex)
		l.Scal(l.Sub(x, y), a)
		r.Sub(r.Scal(x, a), new(Perplex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestPerplexAddMulDistributive(t *testing.T) {
	f := func(x, y, z *Perplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Perplex), new(Perplex)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(Perplex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestPerplexSubMulDistributive(t *testing.T) {
	f := func(x, y, z *Perplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Perplex), new(Perplex)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(Perplex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
