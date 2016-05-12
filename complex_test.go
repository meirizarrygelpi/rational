package rational

import (
	"math/big"
	"testing"
	"testing/quick"
)

func TestAddCommutative(t *testing.T) {
	f := func(x, y *Complex) bool {
		l := new(Complex).Add(x, y)
		r := new(Complex).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddAssociative(t *testing.T) {
	f := func(x, y, z *Complex) bool {
		l := new(Complex).Add(new(Complex).Add(x, y), z)
		r := new(Complex).Add(x, new(Complex).Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddZero(t *testing.T) {
	zero := &Complex{
		re: big.NewRat(0, 1),
		im: big.NewRat(0, 1),
	}
	f := func(x *Complex) bool {
		l := new(Complex).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulCommutative(t *testing.T) {
	f := func(x, y *Complex) bool {
		l := new(Complex).Mul(x, y)
		r := new(Complex).Mul(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulAssociative(t *testing.T) {
	f := func(x, y, z *Complex) bool {
		l := new(Complex).Mul(new(Complex).Mul(x, y), z)
		r := new(Complex).Mul(x, new(Complex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulOne(t *testing.T) {
	one := &Complex{
		re: big.NewRat(1, 1),
		im: big.NewRat(0, 1),
	}
	f := func(x *Complex) bool {
		l := new(Complex).Mul(x, one)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInvInvolutive(t *testing.T) {
	f := func(x *Complex) bool {
		l := new(Complex).Inv(new(Complex).Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegInvolutive(t *testing.T) {
	f := func(x *Complex) bool {
		l := new(Complex).Neg(new(Complex).Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestConjInvolutive(t *testing.T) {
	f := func(x *Complex) bool {
		l := new(Complex).Conj(new(Complex).Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegConjCommutative(t *testing.T) {
	f := func(x *Complex) bool {
		l := new(Complex).Neg(new(Complex).Conj(x))
		r := new(Complex).Conj(new(Complex).Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulInvOne(t *testing.T) {
	one := &Complex{
		re: big.NewRat(1, 1),
		im: big.NewRat(0, 1),
	}
	f := func(x *Complex) bool {
		l := new(Complex).Mul(x, new(Complex).Inv(x))
		return l.Equals(one)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
