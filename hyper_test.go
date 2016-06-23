// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestHyperAddCommutative(t *testing.T) {
	f := func(x, y *Hyper) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Hyper).Add(x, y)
		r := new(Hyper).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHyperMulCommutative(t *testing.T) {
	f := func(x, y *Hyper) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Hyper).Mul(x, y)
		r := new(Hyper).Mul(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHyperNegConjCommutative(t *testing.T) {
	f := func(x *Hyper) bool {
		// t.Logf("x = %v", x)
		l, r := new(Hyper), new(Hyper)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestHyperSubAntiCommutative(t *testing.T) {
	f := func(x, y *Hyper) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Hyper), new(Hyper)
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

func TestHyperAddAssociative(t *testing.T) {
	f := func(x, y, z *Hyper) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Hyper), new(Hyper)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHyperMulAssociative(t *testing.T) {
	f := func(x, y, z *Hyper) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Hyper), new(Hyper)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestHyperAddZero(t *testing.T) {
	zero := new(Hyper)
	f := func(x *Hyper) bool {
		// t.Logf("x = %v", x)
		l := new(Hyper).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHyperMulOne(t *testing.T) {
	one := &Infra{
		l: *big.NewRat(1, 1),
	}
	zero := new(Infra)
	f := func(x *Hyper) bool {
		// t.Logf("x = %v", x)
		l := new(Hyper).Mul(x, &Hyper{*one, *zero})
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHyperMulInvOne(t *testing.T) {
	one := &Infra{
		l: *big.NewRat(1, 1),
	}
	zero := new(Infra)
	f := func(x *Hyper) bool {
		// t.Logf("x = %v", x)
		l := new(Hyper)
		l.Mul(x, l.Inv(x))
		return l.Equals(&Hyper{*one, *zero})
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHyperAddNegSub(t *testing.T) {
	f := func(x, y *Hyper) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Hyper), new(Hyper)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHyperAddScalDouble(t *testing.T) {
	f := func(x *Hyper) bool {
		// t.Logf("x = %v", x)
		l, r := new(Hyper), new(Hyper)
		l.Add(x, x)
		r.Scal(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestHyperInvInvolutive(t *testing.T) {
	f := func(x *Hyper) bool {
		// t.Logf("x = %v", x)
		l := new(Hyper)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHyperNegInvolutive(t *testing.T) {
	f := func(x *Hyper) bool {
		// t.Logf("x = %v", x)
		l := new(Hyper)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHyperConjInvolutive(t *testing.T) {
	f := func(x *Hyper) bool {
		// t.Logf("x = %v", x)
		l := new(Hyper)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestHyperMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *Hyper) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Hyper), new(Hyper)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(Hyper).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHyperMulInvAntiDistributive(t *testing.T) {
	f := func(x, y *Hyper) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Hyper), new(Hyper)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(Hyper).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestHyperAddConjDistributive(t *testing.T) {
	f := func(x, y *Hyper) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Hyper), new(Hyper)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(Hyper).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHyperSubConjDistributive(t *testing.T) {
	f := func(x, y *Hyper) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Hyper), new(Hyper)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(Hyper).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHyperAddScalDistributive(t *testing.T) {
	f := func(x, y *Hyper) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Hyper), new(Hyper)
		l.Scal(l.Add(x, y), a)
		r.Add(r.Scal(x, a), new(Hyper).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHyperSubScalDistributive(t *testing.T) {
	f := func(x, y *Hyper) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Hyper), new(Hyper)
		l.Scal(l.Sub(x, y), a)
		r.Sub(r.Scal(x, a), new(Hyper).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHyperAddMulDistributive(t *testing.T) {
	f := func(x, y, z *Hyper) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Hyper), new(Hyper)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(Hyper).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHyperSubMulDistributive(t *testing.T) {
	f := func(x, y, z *Hyper) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Hyper), new(Hyper)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(Hyper).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Positivity

func TestHyperNormPositive(t *testing.T) {
	f := func(x *Hyper) bool {
		// t.Logf("x = %v", x)
		return x.Norm().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func TestHyperComposition(t *testing.T) {
	f := func(x, y *Hyper) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(Hyper)
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
