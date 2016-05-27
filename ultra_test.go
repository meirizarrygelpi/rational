// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestUltraAddCommutative(t *testing.T) {
	f := func(x, y *Ultra) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Ultra).Add(x, y)
		r := new(Ultra).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestUltraNegConjCommutative(t *testing.T) {
	f := func(x *Ultra) bool {
		// t.Logf("x = %v", x)
		l, r := new(Ultra), new(Ultra)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Non-commutativity

func TestUltraMulNonCommutative(t *testing.T) {
	f := func(x, y *Ultra) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Ultra).Commutator(x, y)
		zero := new(Ultra)
		return !l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestUltraSubAntiCommutative(t *testing.T) {
	f := func(x, y *Ultra) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Ultra), new(Ultra)
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

func TestUltraAddAssociative(t *testing.T) {
	f := func(x, y, z *Ultra) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Ultra), new(Ultra)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Non-associativity

func TestUltraMulNonAssociative(t *testing.T) {
	f := func(x, y, z *Ultra) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l := new(Ultra).Associator(x, y, z)
		zero := new(Ultra)
		return !l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestUltraAddZero(t *testing.T) {
	zero := new(Ultra)
	f := func(x *Ultra) bool {
		// t.Logf("x = %v", x)
		l := new(Ultra).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestUltraMulOne(t *testing.T) {
	one := &Supra{
		l: Infra{
			l: *big.NewRat(1, 1),
		},
	}
	zero := new(Supra)
	f := func(x *Ultra) bool {
		// t.Logf("x = %v", x)
		l := new(Ultra).Mul(x, &Ultra{*one, *zero})
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestUltraMulInvOne(t *testing.T) {
	one := &Supra{
		l: Infra{
			l: *big.NewRat(1, 1),
		},
	}
	zero := new(Supra)
	f := func(x *Ultra) bool {
		// t.Logf("x = %v", x)
		l := new(Ultra)
		l.Mul(x, l.Inv(x))
		return l.Equals(&Ultra{*one, *zero})
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestUltraAddNegSub(t *testing.T) {
	f := func(x, y *Ultra) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Ultra), new(Ultra)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestUltraAddScalDouble(t *testing.T) {
	f := func(x *Ultra) bool {
		// t.Logf("x = %v", x)
		l, r := new(Ultra), new(Ultra)
		l.Add(x, x)
		r.Scal(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestUltraInvInvolutive(t *testing.T) {
	f := func(x *Ultra) bool {
		// t.Logf("x = %v", x)
		l := new(Ultra)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestUltraNegInvolutive(t *testing.T) {
	f := func(x *Ultra) bool {
		// t.Logf("x = %v", x)
		l := new(Ultra)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestUltraConjInvolutive(t *testing.T) {
	f := func(x *Ultra) bool {
		// t.Logf("x = %v", x)
		l := new(Ultra)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestUltraMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *Ultra) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Ultra), new(Ultra)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(Ultra).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestUltraMulInvAntiDistributive(t *testing.T) {
	f := func(x, y *Ultra) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Ultra), new(Ultra)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(Ultra).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestUltraAddConjDistributive(t *testing.T) {
	f := func(x, y *Ultra) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Ultra), new(Ultra)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(Ultra).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestUltraSubConjDistributive(t *testing.T) {
	f := func(x, y *Ultra) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Ultra), new(Ultra)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(Ultra).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestUltraAddScalDistributive(t *testing.T) {
	f := func(x, y *Ultra) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Ultra), new(Ultra)
		l.Scal(l.Add(x, y), a)
		r.Add(r.Scal(x, a), new(Ultra).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestUltraSubScalDistributive(t *testing.T) {
	f := func(x, y *Ultra) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Ultra), new(Ultra)
		l.Scal(l.Sub(x, y), a)
		r.Sub(r.Scal(x, a), new(Ultra).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestUltraAddMulDistributive(t *testing.T) {
	f := func(x, y, z *Ultra) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Ultra), new(Ultra)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(Ultra).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestUltraSubMulDistributive(t *testing.T) {
	f := func(x, y, z *Ultra) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Ultra), new(Ultra)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(Ultra).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Positivity

func TestUltraQuadPositive(t *testing.T) {
	f := func(x *Ultra) bool {
		// t.Logf("x = %v", x)
		return x.Quad().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Left-alternativity

func TestUltraLeftAlternative(t *testing.T) {
	f := func(x, y *Ultra) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Ultra)
		l.Associator(x, x, y)
		zero := new(Ultra)
		return l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Right-alternativity

func TestUltraRightAlternative(t *testing.T) {
	f := func(x, y *Ultra) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Ultra)
		l.Associator(x, y, y)
		zero := new(Ultra)
		return l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
