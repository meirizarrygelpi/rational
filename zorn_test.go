// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestZornAddCommutative(t *testing.T) {
	f := func(x, y *Zorn) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Zorn).Add(x, y)
		r := new(Zorn).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestZornNegConjCommutative(t *testing.T) {
	f := func(x *Zorn) bool {
		// t.Logf("x = %v", x)
		l, r := new(Zorn), new(Zorn)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Non-commutativity

func TestZornMulNonCommutative(t *testing.T) {
	f := func(x, y *Zorn) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Zorn).Commutator(x, y)
		zero := new(Zorn)
		return !l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestZornSubAntiCommutative(t *testing.T) {
	f := func(x, y *Zorn) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Zorn), new(Zorn)
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

func TestZornAddAssociative(t *testing.T) {
	f := func(x, y, z *Zorn) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Zorn), new(Zorn)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Non-associativity

func TestZornMulNonAssociative(t *testing.T) {
	f := func(x, y, z *Zorn) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l := new(Zorn).Associator(x, y, z)
		zero := new(Zorn)
		return !l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestZornAddZero(t *testing.T) {
	zero := new(Zorn)
	f := func(x *Zorn) bool {
		// t.Logf("x = %v", x)
		l := new(Zorn).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestZornMulOne(t *testing.T) {
	one := &Hamilton{
		l: Complex{
			l: *big.NewRat(1, 1),
		},
	}
	zero := new(Hamilton)
	f := func(x *Zorn) bool {
		// t.Logf("x = %v", x)
		l := new(Zorn).Mul(x, &Zorn{*one, *zero})
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestZornMulInvOne(t *testing.T) {
	one := &Hamilton{
		l: Complex{
			l: *big.NewRat(1, 1),
		},
	}
	zero := new(Hamilton)
	f := func(x *Zorn) bool {
		// t.Logf("x = %v", x)
		l := new(Zorn)
		l.Mul(x, l.Inv(x))
		return l.Equals(&Zorn{*one, *zero})
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestZornAddNegSub(t *testing.T) {
	f := func(x, y *Zorn) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Zorn), new(Zorn)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestZornAddScalDouble(t *testing.T) {
	f := func(x *Zorn) bool {
		// t.Logf("x = %v", x)
		l, r := new(Zorn), new(Zorn)
		l.Add(x, x)
		r.Scal(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestZornInvInvolutive(t *testing.T) {
	f := func(x *Zorn) bool {
		// t.Logf("x = %v", x)
		l := new(Zorn)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestZornNegInvolutive(t *testing.T) {
	f := func(x *Zorn) bool {
		// t.Logf("x = %v", x)
		l := new(Zorn)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestZornConjInvolutive(t *testing.T) {
	f := func(x *Zorn) bool {
		// t.Logf("x = %v", x)
		l := new(Zorn)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestZornMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *Zorn) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Zorn), new(Zorn)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(Zorn).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestZornMulInvAntiDistributive(t *testing.T) {
	f := func(x, y *Zorn) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Zorn), new(Zorn)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(Zorn).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestZornAddConjDistributive(t *testing.T) {
	f := func(x, y *Zorn) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Zorn), new(Zorn)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(Zorn).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestZornSubConjDistributive(t *testing.T) {
	f := func(x, y *Zorn) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Zorn), new(Zorn)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(Zorn).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestZornAddScalDistributive(t *testing.T) {
	f := func(x, y *Zorn) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Zorn), new(Zorn)
		l.Scal(l.Add(x, y), a)
		r.Add(r.Scal(x, a), new(Zorn).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestZornSubScalDistributive(t *testing.T) {
	f := func(x, y *Zorn) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Zorn), new(Zorn)
		l.Scal(l.Sub(x, y), a)
		r.Sub(r.Scal(x, a), new(Zorn).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestZornAddMulDistributive(t *testing.T) {
	f := func(x, y, z *Zorn) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Zorn), new(Zorn)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(Zorn).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestZornSubMulDistributive(t *testing.T) {
	f := func(x, y, z *Zorn) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Zorn), new(Zorn)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(Zorn).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
