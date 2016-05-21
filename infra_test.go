package rational

import (
	"math/big"
	"testing"
	"testing/quick"
)

func TestInfraAddCommutative(t *testing.T) {
	f := func(x, y *Infra) bool {
		l := new(Infra).Add(x, y)
		r := new(Infra).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraAddAssociative(t *testing.T) {
	f := func(x, y, z *Infra) bool {
		l, r := new(Infra), new(Infra)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraAddZero(t *testing.T) {
	zero := &Infra{
		l: big.NewRat(0, 1),
		r: big.NewRat(0, 1),
	}
	f := func(x *Infra) bool {
		l := new(Infra).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraMulCommutative(t *testing.T) {
	f := func(x, y *Infra) bool {
		l := new(Infra).Mul(x, y)
		r := new(Infra).Mul(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraMulAssociative(t *testing.T) {
	f := func(x, y, z *Infra) bool {
		l, r := new(Infra), new(Infra)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraMulOne(t *testing.T) {
	one := &Infra{
		l: big.NewRat(1, 1),
		r: big.NewRat(0, 1),
	}
	f := func(x *Infra) bool {
		l := new(Infra).Mul(x, one)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraInvInvolutive(t *testing.T) {
	f := func(x *Infra) bool {
		l := new(Infra)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraNegInvolutive(t *testing.T) {
	f := func(x *Infra) bool {
		l := new(Infra)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraConjInvolutive(t *testing.T) {
	f := func(x *Infra) bool {
		l := new(Infra)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraNegConjCommutative(t *testing.T) {
	f := func(x *Infra) bool {
		l, r := new(Infra), new(Infra)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraMulInvOne(t *testing.T) {
	one := &Infra{
		l: big.NewRat(1, 1),
		r: big.NewRat(0, 1),
	}
	f := func(x *Infra) bool {
		l := new(Infra)
		l.Mul(x, l.Inv(x))
		return l.Equals(one)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *Infra) bool {
		l, r := new(Infra), new(Infra)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(Infra).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraMulInvAntiDistributive(t *testing.T) {
	f := func(x, y *Infra) bool {
		l, r := new(Infra), new(Infra)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(Infra).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
