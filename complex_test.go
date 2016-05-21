package rational

import (
	"math/big"
	"testing"
	"testing/quick"
)

func TestComplexAddCommutative(t *testing.T) {
	f := func(x, y *Complex) bool {
		l := new(Complex).Add(x, y)
		r := new(Complex).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestComplexAddAssociative(t *testing.T) {
	f := func(x, y, z *Complex) bool {
		l, r := new(Complex), new(Complex)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestComplexAddZero(t *testing.T) {
	zero := &Complex{
		l: big.NewRat(0, 1),
		r: big.NewRat(0, 1),
	}
	f := func(x *Complex) bool {
		l := new(Complex).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestComplexMulCommutative(t *testing.T) {
	f := func(x, y *Complex) bool {
		l := new(Complex).Mul(x, y)
		r := new(Complex).Mul(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestComplexMulAssociative(t *testing.T) {
	f := func(x, y, z *Complex) bool {
		l, r := new(Complex), new(Complex)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestComplexMulOne(t *testing.T) {
	one := &Complex{
		l: big.NewRat(1, 1),
		r: big.NewRat(0, 1),
	}
	f := func(x *Complex) bool {
		l := new(Complex).Mul(x, one)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestComplexInvInvolutive(t *testing.T) {
	f := func(x *Complex) bool {
		l := new(Complex)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestComplexNegInvolutive(t *testing.T) {
	f := func(x *Complex) bool {
		l := new(Complex)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestComplexConjInvolutive(t *testing.T) {
	f := func(x *Complex) bool {
		l := new(Complex)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestComplexNegConjCommutative(t *testing.T) {
	f := func(x *Complex) bool {
		l, r := new(Complex), new(Complex)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestComplexMulInvOne(t *testing.T) {
	one := &Complex{
		l: big.NewRat(1, 1),
		r: big.NewRat(0, 1),
	}
	f := func(x *Complex) bool {
		l := new(Complex)
		l.Mul(x, l.Inv(x))
		return l.Equals(one)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestComplexMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *Complex) bool {
		l, r := new(Complex), new(Complex)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(Complex).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestComplexMulInvAntiDistributive(t *testing.T) {
	f := func(x, y *Complex) bool {
		l, r := new(Complex), new(Complex)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(Complex).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
