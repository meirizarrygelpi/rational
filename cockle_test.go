// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestCockleAddCommutative(t *testing.T) {
	f := func(x, y *Cockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Cockle).Add(x, y)
		r := new(Cockle).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCockleNegConjCommutative(t *testing.T) {
	f := func(x *Cockle) bool {
		// t.Logf("x = %v", x)
		l, r := new(Cockle), new(Cockle)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Non-commutativity

func TestCockleMulNonCommutative(t *testing.T) {
	f := func(x, y *Cockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Cockle).Commutator(x, y)
		zero := new(Cockle)
		return !l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestCockleSubAntiCommutative(t *testing.T) {
	f := func(x, y *Cockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Cockle), new(Cockle)
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

func TestCockleAddAssociative(t *testing.T) {
	f := func(x, y, z *Cockle) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Cockle), new(Cockle)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCockleMulAssociative(t *testing.T) {
	f := func(x, y, z *Cockle) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Cockle), new(Cockle)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestCockleAddZero(t *testing.T) {
	zero := new(Cockle)
	f := func(x *Cockle) bool {
		// t.Logf("x = %v", x)
		l := new(Cockle).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCockleMulOne(t *testing.T) {
	one := &Complex{
		l: *big.NewRat(1, 1),
	}
	zero := new(Complex)
	f := func(x *Cockle) bool {
		// t.Logf("x = %v", x)
		l := new(Cockle).Mul(x, &Cockle{*one, *zero})
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCockleMulInvOne(t *testing.T) {
	one := &Complex{
		l: *big.NewRat(1, 1),
	}
	zero := new(Complex)
	f := func(x *Cockle) bool {
		// t.Logf("x = %v", x)
		l := new(Cockle)
		l.Mul(x, l.Inv(x))
		return l.Equals(&Cockle{*one, *zero})
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCockleAddNegSub(t *testing.T) {
	f := func(x, y *Cockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Cockle), new(Cockle)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCockleAddScalDouble(t *testing.T) {
	f := func(x *Cockle) bool {
		// t.Logf("x = %v", x)
		l, r := new(Cockle), new(Cockle)
		l.Add(x, x)
		r.Scal(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestCockleNegInvolutive(t *testing.T) {
	f := func(x *Cockle) bool {
		// t.Logf("x = %v", x)
		l := new(Cockle)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCockleConjInvolutive(t *testing.T) {
	f := func(x *Cockle) bool {
		// t.Logf("x = %v", x)
		l := new(Cockle)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestCockleMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *Cockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Cockle), new(Cockle)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(Cockle).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCockleMulInvAntiDistributive(t *testing.T) {
	f := func(x, y *Cockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Cockle), new(Cockle)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(Cockle).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestCockleAddConjDistributive(t *testing.T) {
	f := func(x, y *Cockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Cockle), new(Cockle)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(Cockle).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCockleSubConjDistributive(t *testing.T) {
	f := func(x, y *Cockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Cockle), new(Cockle)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(Cockle).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCockleAddScalDistributive(t *testing.T) {
	f := func(x, y *Cockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Cockle), new(Cockle)
		l.Scal(l.Add(x, y), a)
		r.Add(r.Scal(x, a), new(Cockle).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCockleSubScalDistributive(t *testing.T) {
	f := func(x, y *Cockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Cockle), new(Cockle)
		l.Scal(l.Sub(x, y), a)
		r.Sub(r.Scal(x, a), new(Cockle).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCockleAddMulDistributive(t *testing.T) {
	f := func(x, y, z *Cockle) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Cockle), new(Cockle)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(Cockle).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCockleSubMulDistributive(t *testing.T) {
	f := func(x, y, z *Cockle) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Cockle), new(Cockle)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(Cockle).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
