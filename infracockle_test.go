// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestInfraCockleAddCommutative(t *testing.T) {
	f := func(x, y *InfraCockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(InfraCockle).Add(x, y)
		r := new(InfraCockle).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraCockleNegConjCommutative(t *testing.T) {
	f := func(x *InfraCockle) bool {
		// t.Logf("x = %v", x)
		l, r := new(InfraCockle), new(InfraCockle)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Non-commutativity

func TestInfraCockleMulNonCommutative(t *testing.T) {
	f := func(x, y *InfraCockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(InfraCockle).Commutator(x, y)
		zero := new(InfraCockle)
		return !l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestInfraCockleSubAntiCommutative(t *testing.T) {
	f := func(x, y *InfraCockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(InfraCockle), new(InfraCockle)
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

func TestInfraCockleAddAssociative(t *testing.T) {
	f := func(x, y, z *InfraCockle) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(InfraCockle), new(InfraCockle)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Non-associativity

func TestInfraCockleMulNonAssociative(t *testing.T) {
	f := func(x, y, z *InfraCockle) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l := new(InfraCockle).Associator(x, y, z)
		zero := new(InfraCockle)
		return !l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestInfraCockleAddZero(t *testing.T) {
	zero := new(InfraCockle)
	f := func(x *InfraCockle) bool {
		// t.Logf("x = %v", x)
		l := new(InfraCockle).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraCockleMulOne(t *testing.T) {
	one := new(Cockle)
	one.SetL(NewComplex(
		big.NewRat(1, 1),
		big.NewRat(0, 1),
	))
	zero := new(Cockle)
	f := func(x *InfraCockle) bool {
		// t.Logf("x = %v", x)
		l := new(InfraCockle).Mul(x, &InfraCockle{*one, *zero})
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraCockleMulInvOne(t *testing.T) {
	one := new(Cockle)
	one.SetL(NewComplex(
		big.NewRat(1, 1),
		big.NewRat(0, 1),
	))
	zero := new(Cockle)
	f := func(x *InfraCockle) bool {
		// t.Logf("x = %v", x)
		l := new(InfraCockle)
		l.Mul(x, l.Inv(x))
		return l.Equals(&InfraCockle{*one, *zero})
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraCockleAddNegSub(t *testing.T) {
	f := func(x, y *InfraCockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(InfraCockle), new(InfraCockle)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraCockleAddScalDouble(t *testing.T) {
	f := func(x *InfraCockle) bool {
		// t.Logf("x = %v", x)
		l, r := new(InfraCockle), new(InfraCockle)
		l.Add(x, x)
		r.Scal(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestInfraCockleInvInvolutive(t *testing.T) {
	f := func(x *InfraCockle) bool {
		// t.Logf("x = %v", x)
		l := new(InfraCockle)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraCockleNegInvolutive(t *testing.T) {
	f := func(x *InfraCockle) bool {
		// t.Logf("x = %v", x)
		l := new(InfraCockle)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraCockleConjInvolutive(t *testing.T) {
	f := func(x *InfraCockle) bool {
		// t.Logf("x = %v", x)
		l := new(InfraCockle)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestInfraCockleMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *InfraCockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(InfraCockle), new(InfraCockle)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(InfraCockle).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraCockleMulInvAntiDistributive(t *testing.T) {
	f := func(x, y *InfraCockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(InfraCockle), new(InfraCockle)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(InfraCockle).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestInfraCockleAddConjDistributive(t *testing.T) {
	f := func(x, y *InfraCockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(InfraCockle), new(InfraCockle)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(InfraCockle).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraCockleSubConjDistributive(t *testing.T) {
	f := func(x, y *InfraCockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(InfraCockle), new(InfraCockle)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(InfraCockle).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraCockleAddScalDistributive(t *testing.T) {
	f := func(x, y *InfraCockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(InfraCockle), new(InfraCockle)
		l.Scal(l.Add(x, y), a)
		r.Add(r.Scal(x, a), new(InfraCockle).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraCockleSubScalDistributive(t *testing.T) {
	f := func(x, y *InfraCockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(InfraCockle), new(InfraCockle)
		l.Scal(l.Sub(x, y), a)
		r.Sub(r.Scal(x, a), new(InfraCockle).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraCockleAddMulDistributive(t *testing.T) {
	f := func(x, y, z *InfraCockle) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(InfraCockle), new(InfraCockle)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(InfraCockle).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraCockleSubMulDistributive(t *testing.T) {
	f := func(x, y, z *InfraCockle) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(InfraCockle), new(InfraCockle)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(InfraCockle).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
