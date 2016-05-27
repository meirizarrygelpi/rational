// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestSupraComplexAddCommutative(t *testing.T) {
	f := func(x, y *SupraComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(SupraComplex).Add(x, y)
		r := new(SupraComplex).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSupraComplexNegConjCommutative(t *testing.T) {
	f := func(x *SupraComplex) bool {
		// t.Logf("x = %v", x)
		l, r := new(SupraComplex), new(SupraComplex)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Non-commutativity

func TestSupraComplexMulNonCommutative(t *testing.T) {
	f := func(x, y *SupraComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(SupraComplex).Commutator(x, y)
		zero := new(SupraComplex)
		return !l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestSupraComplexSubAntiCommutative(t *testing.T) {
	f := func(x, y *SupraComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(SupraComplex), new(SupraComplex)
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

func TestSupraComplexAddAssociative(t *testing.T) {
	f := func(x, y, z *SupraComplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(SupraComplex), new(SupraComplex)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Non-associativity

func TestSupraComplexMulNonAssociative(t *testing.T) {
	f := func(x, y, z *SupraComplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l := new(SupraComplex).Associator(x, y, z)
		zero := new(SupraComplex)
		return !l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestSupraComplexAddZero(t *testing.T) {
	zero := new(SupraComplex)
	f := func(x *SupraComplex) bool {
		// t.Logf("x = %v", x)
		l := new(SupraComplex).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSupraComplexMulOne(t *testing.T) {
	one := &InfraComplex{
		l: Complex{
			l: *big.NewRat(1, 1),
		},
	}
	zero := new(InfraComplex)
	f := func(x *SupraComplex) bool {
		// t.Logf("x = %v", x)
		l := new(SupraComplex).Mul(x, &SupraComplex{*one, *zero})
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSupraComplexMulInvOne(t *testing.T) {
	one := &InfraComplex{
		l: Complex{
			l: *big.NewRat(1, 1),
		},
	}
	zero := new(InfraComplex)
	f := func(x *SupraComplex) bool {
		// t.Logf("x = %v", x)
		l := new(SupraComplex)
		l.Mul(x, l.Inv(x))
		return l.Equals(&SupraComplex{*one, *zero})
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSupraComplexAddNegSub(t *testing.T) {
	f := func(x, y *SupraComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(SupraComplex), new(SupraComplex)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSupraComplexAddScalDouble(t *testing.T) {
	f := func(x *SupraComplex) bool {
		// t.Logf("x = %v", x)
		l, r := new(SupraComplex), new(SupraComplex)
		l.Add(x, x)
		r.Scal(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestSupraComplexNegInvolutive(t *testing.T) {
	f := func(x *SupraComplex) bool {
		// t.Logf("x = %v", x)
		l := new(SupraComplex)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSupraComplexConjInvolutive(t *testing.T) {
	f := func(x *SupraComplex) bool {
		// t.Logf("x = %v", x)
		l := new(SupraComplex)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestSupraComplexMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *SupraComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(SupraComplex), new(SupraComplex)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(SupraComplex).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSupraComplexMulInvAntiDistributive(t *testing.T) {
	f := func(x, y *SupraComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(SupraComplex), new(SupraComplex)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(SupraComplex).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestSupraComplexAddConjDistributive(t *testing.T) {
	f := func(x, y *SupraComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(SupraComplex), new(SupraComplex)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(SupraComplex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSupraComplexSubConjDistributive(t *testing.T) {
	f := func(x, y *SupraComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(SupraComplex), new(SupraComplex)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(SupraComplex).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSupraComplexAddScalDistributive(t *testing.T) {
	f := func(x, y *SupraComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(SupraComplex), new(SupraComplex)
		l.Scal(l.Add(x, y), a)
		r.Add(r.Scal(x, a), new(SupraComplex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSupraComplexSubScalDistributive(t *testing.T) {
	f := func(x, y *SupraComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(SupraComplex), new(SupraComplex)
		l.Scal(l.Sub(x, y), a)
		r.Sub(r.Scal(x, a), new(SupraComplex).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSupraComplexAddMulDistributive(t *testing.T) {
	f := func(x, y, z *SupraComplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(SupraComplex), new(SupraComplex)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(SupraComplex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSupraComplexSubMulDistributive(t *testing.T) {
	f := func(x, y, z *SupraComplex) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(SupraComplex), new(SupraComplex)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(SupraComplex).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Positivity

func TestSupraComplexQuadPositive(t *testing.T) {
	f := func(x *SupraComplex) bool {
		// t.Logf("x = %v", x)
		return x.Quad().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Left-alternativity

func TestSupraComplexLeftAlternative(t *testing.T) {
	f := func(x, y *SupraComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(SupraComplex)
		l.Associator(x, x, y)
		zero := new(SupraComplex)
		return l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Right-alternativity

func TestSupraComplexRightAlternative(t *testing.T) {
	f := func(x, y *SupraComplex) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(SupraComplex)
		l.Associator(x, y, y)
		zero := new(SupraComplex)
		return l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
