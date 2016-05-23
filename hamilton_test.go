package rational

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestHamiltonAddCommutative(t *testing.T) {
	f := func(x, y *Hamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Hamilton).Add(x, y)
		r := new(Hamilton).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHamiltonNegConjCommutative(t *testing.T) {
	f := func(x *Hamilton) bool {
		// t.Logf("x = %v", x)
		l, r := new(Hamilton), new(Hamilton)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Non-commutativity

func TestHamiltonMulNonCommutative(t *testing.T) {
	f := func(x, y *Hamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Hamilton).Commutator(x, y)
		zero := new(Complex)
		return !l.Equals(&Hamilton{*zero, *zero})
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestHamiltonSubAntiCommutative(t *testing.T) {
	f := func(x, y *Hamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Hamilton), new(Hamilton)
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

func TestHamiltonAddAssociative(t *testing.T) {
	f := func(x, y, z *Hamilton) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Hamilton), new(Hamilton)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHamiltonMulAssociative(t *testing.T) {
	f := func(x, y, z *Hamilton) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Hamilton), new(Hamilton)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestHamiltonAddZero(t *testing.T) {
	zero := new(Hamilton)
	f := func(x *Hamilton) bool {
		// t.Logf("x = %v", x)
		l := new(Hamilton).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHamiltonMulOne(t *testing.T) {
	one := new(Complex)
	one.SetL(big.NewRat(1, 1))
	one.SetR(big.NewRat(0, 1))
	zero := new(Complex)
	f := func(x *Hamilton) bool {
		// t.Logf("x = %v", x)
		l := new(Hamilton).Mul(x, &Hamilton{*one, *zero})
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHamiltonMulInvOne(t *testing.T) {
	one := new(Complex)
	one.SetL(big.NewRat(1, 1))
	one.SetR(big.NewRat(0, 1))
	zero := new(Complex)
	f := func(x *Hamilton) bool {
		// t.Logf("x = %v", x)
		l := new(Hamilton)
		l.Mul(x, l.Inv(x))
		return l.Equals(&Hamilton{*one, *zero})
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHamiltonAddNegSub(t *testing.T) {
	f := func(x, y *Hamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Hamilton), new(Hamilton)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHamiltonAddScalDouble(t *testing.T) {
	f := func(x *Hamilton) bool {
		// t.Logf("x = %v", x)
		l, r := new(Hamilton), new(Hamilton)
		l.Add(x, x)
		r.Scal(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestHamiltonInvInvolutive(t *testing.T) {
	f := func(x *Hamilton) bool {
		// t.Logf("x = %v", x)
		l := new(Hamilton)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHamiltonNegInvolutive(t *testing.T) {
	f := func(x *Hamilton) bool {
		// t.Logf("x = %v", x)
		l := new(Hamilton)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHamiltonConjInvolutive(t *testing.T) {
	f := func(x *Hamilton) bool {
		// t.Logf("x = %v", x)
		l := new(Hamilton)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestHamiltonMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *Hamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Hamilton), new(Hamilton)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(Hamilton).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHamiltonMulInvAntiDistributive(t *testing.T) {
	f := func(x, y *Hamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Hamilton), new(Hamilton)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(Hamilton).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestHamiltonAddConjDistributive(t *testing.T) {
	f := func(x, y *Hamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Hamilton), new(Hamilton)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(Hamilton).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHamiltonSubConjDistributive(t *testing.T) {
	f := func(x, y *Hamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Hamilton), new(Hamilton)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(Hamilton).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHamiltonAddScalDistributive(t *testing.T) {
	f := func(x, y *Hamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Hamilton), new(Hamilton)
		l.Scal(l.Add(x, y), a)
		r.Add(r.Scal(x, a), new(Hamilton).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHamiltonSubScalDistributive(t *testing.T) {
	f := func(x, y *Hamilton) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Hamilton), new(Hamilton)
		l.Scal(l.Sub(x, y), a)
		r.Sub(r.Scal(x, a), new(Hamilton).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHamiltonAddMulDistributive(t *testing.T) {
	f := func(x, y, z *Hamilton) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Hamilton), new(Hamilton)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(Hamilton).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHamiltonSubMulDistributive(t *testing.T) {
	f := func(x, y, z *Hamilton) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Hamilton), new(Hamilton)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(Hamilton).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Positivity

func TestHamiltonQuadPositive(t *testing.T) {
	f := func(x *Hamilton) bool {
		// t.Logf("x = %v", x)
		return x.Quad().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
