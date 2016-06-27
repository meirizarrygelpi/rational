// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestBiCockleAddCommutative(t *testing.T) {
	f := func(x, y *BiCockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(BiCockle).Add(x, y)
		r := new(BiCockle).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiCockleNegConjCommutative(t *testing.T) {
	f := func(x *BiCockle) bool {
		// t.Logf("x = %v", x)
		l, r := new(BiCockle), new(BiCockle)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Non-commutativity

func TestBiCockleMulNonCommutative(t *testing.T) {
	f := func(x, y *BiCockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(BiCockle).Commutator(x, y)
		zero := new(BiCockle)
		return !l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestBiCockleSubAntiCommutative(t *testing.T) {
	f := func(x, y *BiCockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(BiCockle), new(BiCockle)
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

func TestBiCockleAddAssociative(t *testing.T) {
	f := func(x, y, z *BiCockle) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(BiCockle), new(BiCockle)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiCockleMulAssociative(t *testing.T) {
	f := func(x, y, z *BiCockle) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(BiCockle), new(BiCockle)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestBiCockleAddZero(t *testing.T) {
	zero := new(BiCockle)
	f := func(x *BiCockle) bool {
		// t.Logf("x = %v", x)
		l := new(BiCockle).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiCockleMulOne(t *testing.T) {
	one := &Cockle{
		l: Complex{
			l: *big.NewRat(1, 1),
		},
	}
	zero := new(Cockle)
	f := func(x *BiCockle) bool {
		// t.Logf("x = %v", x)
		l := new(BiCockle).Mul(x, &BiCockle{*one, *zero})
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiCockleMulInvOne(t *testing.T) {
	one := &Cockle{
		l: Complex{
			l: *big.NewRat(1, 1),
		},
	}
	zero := new(Cockle)
	f := func(x *BiCockle) bool {
		// t.Logf("x = %v", x)
		l := new(BiCockle)
		l.Mul(x, l.Inv(x))
		return l.Equals(&BiCockle{*one, *zero})
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiCockleAddNegSub(t *testing.T) {
	f := func(x, y *BiCockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(BiCockle), new(BiCockle)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiCockleAddScalDouble(t *testing.T) {
	f := func(x *BiCockle) bool {
		// t.Logf("x = %v", x)
		l, r := new(BiCockle), new(BiCockle)
		l.Add(x, x)
		r.Scal(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestBiCockleInvInvolutive(t *testing.T) {
	f := func(x *BiCockle) bool {
		// t.Logf("x = %v", x)
		l := new(BiCockle)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiCockleNegInvolutive(t *testing.T) {
	f := func(x *BiCockle) bool {
		// t.Logf("x = %v", x)
		l := new(BiCockle)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiCockleConjInvolutive(t *testing.T) {
	f := func(x *BiCockle) bool {
		// t.Logf("x = %v", x)
		l := new(BiCockle)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestBiCockleMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *BiCockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(BiCockle), new(BiCockle)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(BiCockle).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiCockleMulInvAntiDistributive(t *testing.T) {
	f := func(x, y *BiCockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(BiCockle), new(BiCockle)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(BiCockle).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestBiCockleAddConjDistributive(t *testing.T) {
	f := func(x, y *BiCockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(BiCockle), new(BiCockle)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(BiCockle).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiCockleSubConjDistributive(t *testing.T) {
	f := func(x, y *BiCockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(BiCockle), new(BiCockle)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(BiCockle).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiCockleAddScalDistributive(t *testing.T) {
	f := func(x, y *BiCockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(BiCockle), new(BiCockle)
		l.Scal(l.Add(x, y), a)
		r.Add(r.Scal(x, a), new(BiCockle).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiCockleSubScalDistributive(t *testing.T) {
	f := func(x, y *BiCockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(BiCockle), new(BiCockle)
		l.Scal(l.Sub(x, y), a)
		r.Sub(r.Scal(x, a), new(BiCockle).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiCockleAddMulDistributive(t *testing.T) {
	f := func(x, y, z *BiCockle) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(BiCockle), new(BiCockle)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(BiCockle).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiCockleSubMulDistributive(t *testing.T) {
	f := func(x, y, z *BiCockle) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(BiCockle), new(BiCockle)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(BiCockle).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Positivity

func TestBiCockleNormPositive(t *testing.T) {
	f := func(x *BiCockle) bool {
		// t.Logf("x = %v", x)
		return x.Norm().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func TestBiCockleComposition(t *testing.T) {
	f := func(x, y *BiCockle) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(BiCockle)
		a, b := new(big.Rat), new(big.Rat)
		p.Mul(x, y)
		a.Set(p.Norm())
		b.Mul(x.Norm(), y.Norm())
		return a.Cmp(b) == 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
