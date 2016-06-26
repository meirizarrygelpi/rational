// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestBiHamiltonAddCommutative(t *testing.T) {
	f := func(x, y *BiHamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(BiHamilton).Add(x, y)
		r := new(BiHamilton).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiHamiltonNegConjCommutative(t *testing.T) {
	f := func(x *BiHamilton) bool {
		// t.Logf("x = %v", x)
		l, r := new(BiHamilton), new(BiHamilton)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Non-commutativity

func TestBiHamiltonMulNonCommutative(t *testing.T) {
	f := func(x, y *BiHamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(BiHamilton).Commutator(x, y)
		zero := new(BiHamilton)
		return !l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestBiHamiltonSubAntiCommutative(t *testing.T) {
	f := func(x, y *BiHamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(BiHamilton), new(BiHamilton)
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

func TestBiHamiltonAddAssociative(t *testing.T) {
	f := func(x, y, z *BiHamilton) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(BiHamilton), new(BiHamilton)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiHamiltonMulAssociative(t *testing.T) {
	f := func(x, y, z *BiHamilton) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(BiHamilton), new(BiHamilton)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestBiHamiltonAddZero(t *testing.T) {
	zero := new(BiHamilton)
	f := func(x *BiHamilton) bool {
		// t.Logf("x = %v", x)
		l := new(BiHamilton).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiHamiltonMulOne(t *testing.T) {
	one := &Hamilton{
		l: Complex{
			l: *big.NewRat(1, 1),
		},
	}
	zero := new(Hamilton)
	f := func(x *BiHamilton) bool {
		// t.Logf("x = %v", x)
		l := new(BiHamilton).Mul(x, &BiHamilton{*one, *zero})
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiHamiltonMulInvOne(t *testing.T) {
	one := &Hamilton{
		l: Complex{
			l: *big.NewRat(1, 1),
		},
	}
	zero := new(Hamilton)
	f := func(x *BiHamilton) bool {
		// t.Logf("x = %v", x)
		l := new(BiHamilton)
		l.Mul(x, l.Inv(x))
		return l.Equals(&BiHamilton{*one, *zero})
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiHamiltonAddNegSub(t *testing.T) {
	f := func(x, y *BiHamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(BiHamilton), new(BiHamilton)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiHamiltonAddScalDouble(t *testing.T) {
	f := func(x *BiHamilton) bool {
		// t.Logf("x = %v", x)
		l, r := new(BiHamilton), new(BiHamilton)
		l.Add(x, x)
		r.Scal(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestBiHamiltonInvInvolutive(t *testing.T) {
	f := func(x *BiHamilton) bool {
		// t.Logf("x = %v", x)
		l := new(BiHamilton)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiHamiltonNegInvolutive(t *testing.T) {
	f := func(x *BiHamilton) bool {
		// t.Logf("x = %v", x)
		l := new(BiHamilton)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiHamiltonConjInvolutive(t *testing.T) {
	f := func(x *BiHamilton) bool {
		// t.Logf("x = %v", x)
		l := new(BiHamilton)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestBiHamiltonMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *BiHamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(BiHamilton), new(BiHamilton)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(BiHamilton).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiHamiltonMulInvAntiDistributive(t *testing.T) {
	f := func(x, y *BiHamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(BiHamilton), new(BiHamilton)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(BiHamilton).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestBiHamiltonAddConjDistributive(t *testing.T) {
	f := func(x, y *BiHamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(BiHamilton), new(BiHamilton)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(BiHamilton).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiHamiltonSubConjDistributive(t *testing.T) {
	f := func(x, y *BiHamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(BiHamilton), new(BiHamilton)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(BiHamilton).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiHamiltonAddScalDistributive(t *testing.T) {
	f := func(x, y *BiHamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(BiHamilton), new(BiHamilton)
		l.Scal(l.Add(x, y), a)
		r.Add(r.Scal(x, a), new(BiHamilton).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiHamiltonSubScalDistributive(t *testing.T) {
	f := func(x, y *BiHamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(BiHamilton), new(BiHamilton)
		l.Scal(l.Sub(x, y), a)
		r.Sub(r.Scal(x, a), new(BiHamilton).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiHamiltonAddMulDistributive(t *testing.T) {
	f := func(x, y, z *BiHamilton) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(BiHamilton), new(BiHamilton)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(BiHamilton).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBiHamiltonSubMulDistributive(t *testing.T) {
	f := func(x, y, z *BiHamilton) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(BiHamilton), new(BiHamilton)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(BiHamilton).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Positivity

func TestBiHamiltonNormPositive(t *testing.T) {
	f := func(x *BiHamilton) bool {
		// t.Logf("x = %v", x)
		return x.Norm().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func TestBiHamiltonComposition(t *testing.T) {
	f := func(x, y *BiHamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(BiHamilton)
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
