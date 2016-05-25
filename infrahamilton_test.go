// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestInfraHamiltonAddCommutative(t *testing.T) {
	f := func(x, y *InfraHamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(InfraHamilton).Add(x, y)
		r := new(InfraHamilton).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraHamiltonNegConjCommutative(t *testing.T) {
	f := func(x *InfraHamilton) bool {
		// t.Logf("x = %v", x)
		l, r := new(InfraHamilton), new(InfraHamilton)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Non-commutativity

func TestInfraHamiltonMulNonCommutative(t *testing.T) {
	f := func(x, y *InfraHamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(InfraHamilton).Commutator(x, y)
		zero := new(InfraHamilton)
		return !l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestInfraHamiltonSubAntiCommutative(t *testing.T) {
	f := func(x, y *InfraHamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(InfraHamilton), new(InfraHamilton)
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

func TestInfraHamiltonAddAssociative(t *testing.T) {
	f := func(x, y, z *InfraHamilton) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(InfraHamilton), new(InfraHamilton)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Non-associativity

func TestInfraHamiltonMulNonAssociative(t *testing.T) {
	f := func(x, y, z *InfraHamilton) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l := new(InfraHamilton).Associator(x, y, z)
		zero := new(InfraHamilton)
		return !l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestInfraHamiltonAddZero(t *testing.T) {
	zero := new(InfraHamilton)
	f := func(x *InfraHamilton) bool {
		// t.Logf("x = %v", x)
		l := new(InfraHamilton).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraHamiltonMulOne(t *testing.T) {
	one := new(Hamilton)
	one.SetL(NewComplex(
		big.NewRat(1, 1),
		big.NewRat(0, 1),
	))
	zero := new(Hamilton)
	f := func(x *InfraHamilton) bool {
		// t.Logf("x = %v", x)
		l := new(InfraHamilton).Mul(x, &InfraHamilton{*one, *zero})
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraHamiltonMulInvOne(t *testing.T) {
	one := new(Hamilton)
	one.SetL(NewComplex(
		big.NewRat(1, 1),
		big.NewRat(0, 1),
	))
	zero := new(Hamilton)
	f := func(x *InfraHamilton) bool {
		// t.Logf("x = %v", x)
		l := new(InfraHamilton)
		l.Mul(x, l.Inv(x))
		return l.Equals(&InfraHamilton{*one, *zero})
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraHamiltonAddNegSub(t *testing.T) {
	f := func(x, y *InfraHamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(InfraHamilton), new(InfraHamilton)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraHamiltonAddScalDouble(t *testing.T) {
	f := func(x *InfraHamilton) bool {
		// t.Logf("x = %v", x)
		l, r := new(InfraHamilton), new(InfraHamilton)
		l.Add(x, x)
		r.Scal(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestInfraHamiltonInvInvolutive(t *testing.T) {
	f := func(x *InfraHamilton) bool {
		// t.Logf("x = %v", x)
		l := new(InfraHamilton)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraHamiltonNegInvolutive(t *testing.T) {
	f := func(x *InfraHamilton) bool {
		// t.Logf("x = %v", x)
		l := new(InfraHamilton)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraHamiltonConjInvolutive(t *testing.T) {
	f := func(x *InfraHamilton) bool {
		// t.Logf("x = %v", x)
		l := new(InfraHamilton)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestInfraHamiltonMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *InfraHamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(InfraHamilton), new(InfraHamilton)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(InfraHamilton).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraHamiltonMulInvAntiDistributive(t *testing.T) {
	f := func(x, y *InfraHamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(InfraHamilton), new(InfraHamilton)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(InfraHamilton).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestInfraHamiltonAddConjDistributive(t *testing.T) {
	f := func(x, y *InfraHamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(InfraHamilton), new(InfraHamilton)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(InfraHamilton).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraHamiltonSubConjDistributive(t *testing.T) {
	f := func(x, y *InfraHamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(InfraHamilton), new(InfraHamilton)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(InfraHamilton).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraHamiltonAddScalDistributive(t *testing.T) {
	f := func(x, y *InfraHamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(InfraHamilton), new(InfraHamilton)
		l.Scal(l.Add(x, y), a)
		r.Add(r.Scal(x, a), new(InfraHamilton).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraHamiltonSubScalDistributive(t *testing.T) {
	f := func(x, y *InfraHamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(InfraHamilton), new(InfraHamilton)
		l.Scal(l.Sub(x, y), a)
		r.Sub(r.Scal(x, a), new(InfraHamilton).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraHamiltonAddMulDistributive(t *testing.T) {
	f := func(x, y, z *InfraHamilton) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(InfraHamilton), new(InfraHamilton)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(InfraHamilton).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInfraHamiltonSubMulDistributive(t *testing.T) {
	f := func(x, y, z *InfraHamilton) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(InfraHamilton), new(InfraHamilton)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(InfraHamilton).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Positivity

func TestInfraHamiltonQuadPositive(t *testing.T) {
	f := func(x *InfraHamilton) bool {
		// t.Logf("x = %v", x)
		return x.Quad().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
