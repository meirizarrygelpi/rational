// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestInfraAddCommutative(t *testing.T) {
	f := func(x, y *Infra) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Infra).Add(x, y)
		r := new(Infra).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraMulCommutative(t *testing.T) {
	f := func(x, y *Infra) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Infra).Mul(x, y)
		r := new(Infra).Mul(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraNegConjCommutative(t *testing.T) {
	f := func(x *Infra) bool {
		// t.Logf("x = %v", x)
		l, r := new(Infra), new(Infra)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestInfraSubAntiCommutative(t *testing.T) {
	f := func(x, y *Infra) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Infra), new(Infra)
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

func TestInfraAddAssociative(t *testing.T) {
	f := func(x, y, z *Infra) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Infra), new(Infra)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraMulAssociative(t *testing.T) {
	f := func(x, y, z *Infra) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Infra), new(Infra)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestInfraAddZero(t *testing.T) {
	zero := new(Infra)
	f := func(x *Infra) bool {
		// t.Logf("x = %v", x)
		l := new(Infra).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraMulOne(t *testing.T) {
	one := &Infra{
		l: *big.NewRat(1, 1),
	}
	f := func(x *Infra) bool {
		// t.Logf("x = %v", x)
		l := new(Infra).Mul(x, one)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraMulInvOne(t *testing.T) {
	one := &Infra{
		l: *big.NewRat(1, 1),
	}
	f := func(x *Infra) bool {
		// t.Logf("x = %v", x)
		l := new(Infra)
		l.Mul(x, l.Inv(x))
		return l.Equals(one)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraAddNegSub(t *testing.T) {
	f := func(x, y *Infra) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Infra), new(Infra)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraAddScalDouble(t *testing.T) {
	f := func(x *Infra) bool {
		// t.Logf("x = %v", x)
		l, r := new(Infra), new(Infra)
		l.Add(x, x)
		r.Scal(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestInfraNegInvolutive(t *testing.T) {
	f := func(x *Infra) bool {
		// t.Logf("x = %v", x)
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
		// t.Logf("x = %v", x)
		l := new(Infra)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestInfraMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *Infra) bool {
		// t.Logf("x = %v, y = %v", x, y)
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
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Infra), new(Infra)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(Infra).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestInfraAddConjDistributive(t *testing.T) {
	f := func(x, y *Infra) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Infra), new(Infra)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(Infra).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraSubConjDistributive(t *testing.T) {
	f := func(x, y *Infra) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Infra), new(Infra)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(Infra).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraAddScalDistributive(t *testing.T) {
	f := func(x, y *Infra) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Infra), new(Infra)
		l.Scal(l.Add(x, y), a)
		r.Add(r.Scal(x, a), new(Infra).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraSubScalDistributive(t *testing.T) {
	f := func(x, y *Infra) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Infra), new(Infra)
		l.Scal(l.Sub(x, y), a)
		r.Sub(r.Scal(x, a), new(Infra).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraAddMulDistributive(t *testing.T) {
	f := func(x, y, z *Infra) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Infra), new(Infra)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(Infra).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraSubMulDistributive(t *testing.T) {
	f := func(x, y, z *Infra) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Infra), new(Infra)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(Infra).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Positivity

func TestInfraQuadPositive(t *testing.T) {
	f := func(x *Infra) bool {
		// t.Logf("x = %v", x)
		return x.Quad().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
