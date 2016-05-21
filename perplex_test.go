package rational

import (
	"math/big"
	"testing"
	"testing/quick"
)

func TestPerplexAddCommutative(t *testing.T) {
	f := func(x, y *Perplex) bool {
		l := new(Perplex).Add(x, y)
		r := new(Perplex).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestPerplexAddAssociative(t *testing.T) {
	f := func(x, y, z *Perplex) bool {
		l, r := new(Perplex), new(Perplex)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestPerplexAddZero(t *testing.T) {
	zero := &Perplex{
		l: big.NewRat(0, 1),
		r: big.NewRat(0, 1),
	}
	f := func(x *Perplex) bool {
		l := new(Perplex).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestPerplexMulCommutative(t *testing.T) {
	f := func(x, y *Perplex) bool {
		l := new(Perplex).Mul(x, y)
		r := new(Perplex).Mul(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestPerplexMulAssociative(t *testing.T) {
	f := func(x, y, z *Perplex) bool {
		l, r := new(Perplex), new(Perplex)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestPerplexMulOne(t *testing.T) {
	one := &Perplex{
		l: big.NewRat(1, 1),
		r: big.NewRat(0, 1),
	}
	f := func(x *Perplex) bool {
		l := new(Perplex).Mul(x, one)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestPerplexInvInvolutive(t *testing.T) {
	f := func(x *Perplex) bool {
		l := new(Perplex)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestPerplexNegInvolutive(t *testing.T) {
	f := func(x *Perplex) bool {
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
		l := new(Perplex)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestPerplexNegConjCommutative(t *testing.T) {
	f := func(x *Perplex) bool {
		l, r := new(Perplex), new(Perplex)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestPerplexMulInvOne(t *testing.T) {
	one := &Perplex{
		l: big.NewRat(1, 1),
		r: big.NewRat(0, 1),
	}
	f := func(x *Perplex) bool {
		l := new(Perplex)
		l.Mul(x, l.Inv(x))
		return l.Equals(one)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestPerplexMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *Perplex) bool {
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
		l, r := new(Perplex), new(Perplex)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(Perplex).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
