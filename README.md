# rational

Package `rational` brings rational [complex](https://en.wikipedia.org/wiki/Complex_number), [split-complex](https://en.wikipedia.org/wiki/Split-complex_number), and [dual](https://en.wikipedia.org/wiki/Dual_number) numbers to Go. It borrows heavily from the `math`, `math/cmplx`, and `math/big` packages. Indeed, it is built on top of the `big.Rat` type.

[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/github.com/meirizarrygelpi/rational) [![GoDoc](https://godoc.org/github.com/meirizarrygelpi/rational?status.svg)](https://godoc.org/github.com/meirizarrygelpi/rational)

This package contains three two-dimensional types (e.g. complex numbers), five four-dimensional types (e.g. quaternions), and seven eight-dimensional types (e.g. octonions). Each type is either an elliptic, a parabolic, or a hyperbolic Cayley-Dickson construct.

## Cayley-Dickson Constructs

Given elements from what I will call a seed algebra, the Cayley-Dickson construction allows you to build elements of a higher-dimensional algebra with certain interesting properties. Let `a`, `b`, `c`, and `d` be elements in the seed algebra. Let `p = (a, b)` and `q = (c, d)` be elements in the construct algebra. Thus,
```go
	Add(p, q) = (Add(a, c), Add(b, d))
```
That is, addition is element-wise. The multiplication operation can be any of three kinds. Before we look into these, we should mention that the seed algebra must also include two involutions, `Conj` and `Neg`, such that the construct algebra also has two involutions `Conj` and `Neg` given by:
```go
	Neg(p) = (Neg(a), Neg(b))
	Conj(p) = (Conj(a), Neg(b))
```
With `Neg` you can define subtraction:
```go
	Sub(p, q) = Add(p, Neg(q))
```
The three multiplication operations are named after conic sections. Each has the form
```go
	Mul(p, q) = (F(a, b, c, d), G(a, b, c, d))
```
With `F` and `G` being a linear combination of bilinear terms.

### Elliptic Multiplication

The **elliptic** multiplication operation is:
```go
	F(a, b, c, d) = Sub(Mul(a, c), Mul(Conj(d), b))
	G(a, b, c, d) = Add(Mul(d, a), Mul(b, Conj(c)))
```
The order of the arguments in all `Mul` calls is very important.

### Parabolic Multiplication

The **parabolic** multiplication operation is:
```go
	F(a, b, c, d) = Mul(a, c)
	G(a, b, c, d) = Add(Mul(d, a), Mul(b, Conj(c)))
```
The order of the arguments in all `Mul` calls is very important. Note that if `a` and `c` are both zero, then `Mul(p, q) = (0, 0)` with `p` and `q` not necessarily being zero. That is, there are zero divisors with this multiplication operation.

### Hyperbolic Multiplication

The **hyperbolic** multiplication operation is:
```go
	F(a, b, c, d) = Add(Mul(a, c), Mul(Conj(d), b))
	G(a, b, c, d) = Add(Mul(d, a), Mul(b, Conj(c)))
```
The order of the arguments in all `Mul` calls is very important. Although it is not obvious, this multiplication operation also leads to zero divisors.

### Non-Sesquilinear Multiplication

The three "conic" multiplication operations above all use the `Conj` involution of the seed algebra. One can also consider non-sesquilinear multiplication rules. If the seed algebra is commutative and associative, then the resulting construct algebra will also be commutative and associative.

The **non-sesquilinear elliptic** multiplication operation is:
```go
	F(a, b, c, d) = Sub(Mul(a, c), Mul(d, b))
	G(a, b, c, d) = Add(Mul(d, a), Mul(b, c))
```
The **non-sesquilinear parabolic** multiplication operation is:
```go
	F(a, b, c, d) = Mul(a, c)
	G(a, b, c, d) = Add(Mul(d, a), Mul(b, c))
```
The **non-sesquilinear hyperbolic** multiplication operation is:
```go
	F(a, b, c, d) = Add(Mul(a, c), Mul(d, b))
	G(a, b, c, d) = Add(Mul(d, a), Mul(b, c))
```
The resulting construct algebras are very different from the familiar Cayley-Dickson constructs.

## Two-Dimensional Types

There are three two-dimensional types. The (binary) multiplication operation for all two-dimensional types is **commutative** and **associative**.

### rational.Complex

The `rational.Complex` type represents a rational complex number. It corresponds to an elliptic Cayley-Dickson construct with `big.Rat` values. The imaginary unit element is denoted `i`. The multiplication rule is:
```go
	Mul(i, i) = -1
```
This type can be used to study [Gaussian integers](https://en.wikipedia.org/wiki/Gaussian_integer).

### rational.Perplex

The `rational.Perplex` type represents a rational perplex number. It corresponds to a hyperbolic Cayley-Dickson construct with `big.Rat` values. The split unit element is denoted `s`. The multiplication rule is:
```go
	Mul(s, s) = +1
```
Perplex numbers are more commonly known as [split-complex numbers](https://en.wikipedia.org/wiki/Split-complex_number), but "perplex" is used here for brevity and symmetry with "complex".

### rational.Infra

The `rational.Infra` type represents a rational infra number. It corresponds to a parabolic Cayley-Dickson construct with `big.Rat` values. The dual unit element is denoted `α`. The multiplication rule is:
```go
	Mul(α, α) = 0
```
Infra numbers are useful for computing numerical first-order derivatives. If `z = a + 1α`, and `f` is a function, then
```
	f(z) = f(a) + f'(a)α
```
Infra numbers are more commonly known as [dual numbers](https://en.wikipedia.org/wiki/Dual_number), but "infra" is used here along with "supra" and "ultra" for the higher-dimensional analogs.

## Four-Dimensional Types

There are five four-dimensional types. The (binary) multiplication operation for all four-dimensional types is **noncommutative** but **associative**.

### rational.Hamilton

The `rational.Hamilton` type represents a rational Hamilton quaternion. It corresponds to an elliptic Cayley-Dickson construct with `rational.Complex` values. The imaginary unit elements are denoted `i`, `j`, and `k`. The multiplication rules are:
```go
	Mul(i, i) = Mul(j, j) = Mul(k, k) = -1
	Mul(i, j) = -Mul(j, i) = k
	Mul(j, k) = -Mul(k, j) = i
	Mul(k, i) = -Mul(i, k) = j
```
Hamilton quaternions are [traditional quaternions](https://en.wikipedia.org/wiki/Quaternion). The type is named after W.R. Hamilton, who discovered quaternions.

This type can be used to study [Hurwitz and Lipschitz integers](https://en.wikipedia.org/wiki/Hurwitz_quaternion).

### rational.Cockle

The `rational.Cockle` type represents a rational Cockle quaternion. It corresponds to a hyperbolic Cayley-Dickson construct with `rational.Complex` values. The imaginary unit element is denoted `i`, and the split unit elements are denoted `t` and `u`. The multiplication rules are:
```go
	Mul(i, i) = -1
	Mul(t, t) = Mul(u, u) = +1
	Mul(i, t) = -Mul(t, i) = u
	Mul(u, t) = -Mul(t, u) = i
	Mul(u, i) = -Mul(i, u) = t
```
Cockle quaternions are more commonly known as [split-quaternions](https://en.wikipedia.org/wiki/Split-quaternion). The type is named after J. Cockle, who discovered them.

Alternatively, you can obtain the Cockle quaternions from an *elliptic* Cayley-Dickson construct with `rational.Perplex` values; or from a *hyperbolic* Cayley-Dickson construct with `rational.Perplex` values.

### rational.Supra

The `rational.Supra` type represents a rational supra number. It corresponds to a parabolic Cayley-Dickson construct with `rational.Infra` values. The dual unit elements are denoted `α`, `β`, and `γ`. The multiplication rules are:
```go
	Mul(α, α) = Mul(β, β) = Mul(γ, γ) = 0
	Mul(α, β) = -Mul(β, α) = γ
	Mul(β, γ) = Mul(γ, β) = 0
	Mul(γ, α) = Mul(α, γ) = 0
```
Note that supra numbers are very different from [hyper-dual numbers](http://adl.stanford.edu/hyperdual/): the multiplication operation for supra numbers is noncommutative. In some ways, supra numbers are the dual analog of quaternions.

### rational.InfraComplex

The `rational.InfraComplex` type represents a rational infra complex number. It corresponds to a parabolic Cayley-Dickson construct with `rational.Complex` values. The imaginary unit element is denoted `i`, and the dual unit elements are denoted `β` and `γ`. The multiplication rules are:
```go
	Mul(i, i) = -1
	Mul(β, β) = Mul(γ, γ) = 0
	Mul(β, γ) = Mul(γ, β) = 0
	Mul(i, β) = -Mul(β, i) = γ
	Mul(γ, i) = -Mul(i, γ) = β
```
Note that `i` **does not commute** with either `β` or `γ`. This makes the infra complex numbers different from the more familiar dual complex numbers.

Alternatively, you can obtain the infra complex numbers from an *elliptic* Cayley-Dickson construct with `rational.Infra` values.

### rational.InfraPerplex

The `rational.InfraPerplex` type represents a rational infra perplex number. It corresponds to a parabolic Cayley-Dickson construct with `rational.Perplex` values. The split unit element is denoted `s`, and the dual unit elements are denoted `τ` and `υ`. The multiplication rules are:
```go
	Mul(s, s) = +1
	Mul(τ, τ) = Mul(υ, υ) = 0
	Mul(τ, υ) = Mul(υ, τ) = 0
	Mul(s, τ) = -Mul(τ, s) = υ
	Mul(s, υ) = -Mul(υ, s) = τ
```
Like `i` in the infra complex numbers, `s` **does not commute** with either `τ` or `υ`. This makes the infra perplex numbers different from the more familiar dual split-complex numbers.

Alternatively, you can obtain the infra perplex numbers from a *hyperbolic* Cayley-Dickson construct with `rational.Infra` values.

## Eight-Dimensional Types

There are seven eight-dimensional types. The (binary) multiplication operation for all eight-dimensional types is **noncommutative** and **nonassociative**.

### rational.Cayley

The `rational.Cayley` type represents a rational Cayley octonion. It corresponds to an elliptic Cayley-Dickson construct with `rational.Hamilton` values. The imaginary unit elements are denoted `i`, `j`, `k`, `m`, `n`, `p`, and `q`. The multiplication rules are:
```go
	Mul(i, i) = Mul(j, j) = Mul(k, k) = -1
	Mul(m, m) = Mul(n, n) = Mul(p, p) = Mul(q, q) = -1
	Mul(i, j) = -Mul(j, i) = +k
	Mul(i, k) = -Mul(k, i) = -j
	Mul(i, m) = -Mul(m, i) = +n
	Mul(i, n) = -Mul(n, i) = -m
	Mul(i, p) = -Mul(p, i) = -q
	Mul(i, q) = -Mul(q, i) = +p
	Mul(j, k) = -Mul(k, j) = +i
	Mul(j, m) = -Mul(m, j) = +p
	Mul(j, n) = -Mul(n, j) = +q
	Mul(j, p) = -Mul(p, j) = -m
	Mul(j, q) = -Mul(q, j) = -n
	Mul(k, m) = -Mul(m, k) = +q
	Mul(k, n) = -Mul(n, k) = -p
	Mul(k, p) = -Mul(p, k) = +n
	Mul(k, q) = -Mul(q, k) = -m
	Mul(m, n) = -Mul(n, m) = +i
	Mul(m, p) = -Mul(p, m) = +j
	Mul(m, q) = -Mul(q, m) = +k
	Mul(n, p) = -Mul(p, n) = -k
	Mul(n, q) = -Mul(q, n) = +j
	Mul(p, q) = -Mul(q, p) = -i
```
Cayley octonions are [traditional octonions](https://en.wikipedia.org/wiki/Octonion). The type is named after A. Cayley, who was **not** the first person to discover octonions. The first person to discover octonions was J.T. Graves.

This type can be used to study [Gravesian and Kleinian integers](https://en.wikipedia.org/wiki/Octonion#Integral_octonions), as well as other integral octonions.

### rational.Zorn

The `rational.Zorn` type represents a rational Zorn octonion. It corresponds to a hyperbolic Cayley-Dickson construct with `rational.Hamilton` values. The imaginary unit elements are denoted `i`, `j`, and `k`, and the split unit elements are `r`, `s`, `t`, and `u`. The multiplication rules are:
```go
	Mul(i, i) = Mul(j, j) = Mul(k, k) = -1
	Mul(r, r) = Mul(s, s) = Mul(t, t) = Mul(u, u) = +1
	Mul(i, j) = -Mul(j, i) = +k
	Mul(i, k) = -Mul(k, i) = -j
	Mul(i, r) = -Mul(r, i) = +s
	Mul(i, s) = -Mul(s, i) = -r
	Mul(i, t) = -Mul(t, i) = -u
	Mul(i, u) = -Mul(u, i) = +t
	Mul(j, k) = -Mul(k, j) = +i
	Mul(j, r) = -Mul(r, j) = +t
	Mul(j, s) = -Mul(s, j) = +u
	Mul(j, t) = -Mul(t, j) = -r
	Mul(j, u) = -Mul(u, j) = -s
	Mul(k, r) = -Mul(r, k) = +u
	Mul(k, s) = -Mul(s, k) = -t
	Mul(k, t) = -Mul(t, k) = +s
	Mul(k, u) = -Mul(u, k) = -r
	Mul(r, s) = -Mul(s, r) = -i
	Mul(r, t) = -Mul(t, r) = -j
	Mul(r, u) = -Mul(u, r) = -k
	Mul(s, t) = -Mul(t, s) = +k
	Mul(s, u) = -Mul(u, s) = -j
	Mul(t, u) = -Mul(u, t) = +i
```
Zorn octonions are commonly known as [split-octonions](https://en.wikipedia.org/wiki/Split-octonion). The type is named after M.A. Zorn, who developed a vector-matrix algebra for working with split-octonions.

Alternatively, you can obtain the Zorn octonions from an *elliptic* Cayley-Dickson construct with `rational.Cockle` values; or from a *hyperbolic* Cayley-Dickson construct with `rational.Cockle` values.

### rational.Ultra

The `rational.Ultra` type represents a rational ultra number. It corresponds to a parabolic Cayley-Dickson construct with `rational.Supra` values. The dual unit elements are denoted `α`, `β`, `γ`, `δ`, `ε`, `ζ`, and `η`. The multiplication rules are:
```go
	Mul(α, α) = Mul(β, β) = Mul(γ, γ) = 0
	Mul(δ, δ) = Mul(ε, ε) = Mul(ζ, ζ) = Mul(η, η) = 0
	Mul(α, β) = -Mul(β, α) = +γ
	Mul(α, γ) = Mul(γ, α) = 0
	Mul(α, δ) = -Mul(δ, α) = +ε
	Mul(α, ε) = Mul(ε, α) = 0
	Mul(α, ζ) = -Mul(ζ, α) = -η
	Mul(α, η) = -Mul(η, α) = +ζ
	Mul(β, γ) = Mul(γ, β) = 0
	Mul(β, δ) = -Mul(δ, β) = +ζ
	Mul(β, ε) = -Mul(ε, β) = +η
	Mul(β, ζ) = Mul(ζ, β) = 0
	Mul(β, η) = Mul(η, β) = 0
	Mul(γ, δ) = -Mul(δ, γ) = +η
	Mul(γ, ε) = Mul(ε, γ) = 0
	Mul(γ, ζ) = Mul(ζ, γ) = 0
	Mul(γ, η) = Mul(η, γ) = 0
	Mul(δ, ε) = Mul(ε, δ) = 0
	Mul(δ, ζ) = Mul(ζ, δ) = 0
	Mul(δ, η) = Mul(η, δ) = 0
	Mul(ε, ζ) = Mul(ζ, ε) = 0
	Mul(ε, η) = Mul(η, ε) = 0
	Mul(ζ, η) = Mul(η, ζ) = 0
```
In some ways, ultra numbers are the dual analog of octonions.

### rational.InfraHamilton

The `rational.InfraHamilton` type represents a rational infra Hamilton quaternion. It corresponds to a parabolic Cayley-Dickson construct with `rational.Hamilton` values. The imaginary unit elements are denoted `i`, `j` and `k`, and the dual unit elements are denoted `α`, `β`, `γ`, and `δ`. The multiplication rules are:
```go
	Mul(i, i) = Mul(j, j) = Mul(k, k) = -1
	Mul(α, α) = Mul(β, β) = Mul(γ, γ) = Mul(δ, δ) = 0
	Mul(i, j) = -Mul(j, i) = +k
	Mul(i, k) = -Mul(k, i) = -j
	Mul(i, α) = -Mul(α, i) = +β
	Mul(i, β) = -Mul(β, i) = -α
	Mul(i, γ) = -Mul(γ, i) = -δ
	Mul(i, δ) = -Mul(δ, i) = +γ
	Mul(j, k) = -Mul(k, j) = +i
	Mul(j, α) = -Mul(α, j) = +γ
	Mul(j, β) = -Mul(β, j) = +δ
	Mul(j, γ) = -Mul(γ, j) = -α
	Mul(j, δ) = -Mul(δ, j) = -β
	Mul(k, α) = -Mul(α, k) = +δ
	Mul(k, β) = -Mul(β, k) = -γ
	Mul(k, γ) = -Mul(γ, k) = +β
	Mul(k, δ) = -Mul(δ, k) = -α
	Mul(α, β) = Mul(β, α) = 0
	Mul(α, γ) = Mul(γ, α) = 0
	Mul(α, δ) = Mul(δ, α) = 0
	Mul(β, γ) = Mul(γ, β) = 0
	Mul(β, δ) = Mul(δ, β) = 0
	Mul(γ, δ) = Mul(δ, γ) = 0
```
The infra Hamilton quaternions are different from the more familiar dual quaternions.

Alternatively, you can obtain the infra Hamilton quaternions from an *elliptic* Cayley-Dickson construct with `rational.InfraComplex` values.

### rational.InfraCockle

The `rational.InfraCockle` type represents a rational infra Cockle quaternion. It corresponds to a parabolic Cayley-Dickson construct with `rational.Cockle` values. The imaginary unit element is denoted `i`, the split unit elements are denoted `t` and `u`, and the dual unit elements are denoted `ρ`, `σ`, `τ`, and `υ`. The multiplication rules are:
```go
	Mul(i, i) = -1
	Mul(t, t) = Mul(u, u) = +1
	Mul(ρ, ρ) = Mul(σ, σ) = Mul(τ, τ) = Mul(υ, υ) = 0
	Mul(i, t) = -Mul(t, i) = +u
	Mul(i, u) = -Mul(u, i) = -t
	Mul(i, ρ) = -Mul(ρ, i) = +σ
	Mul(i, σ) = -Mul(σ, i) = -ρ
	Mul(i, τ) = -Mul(τ, i) = -υ
	Mul(i, υ) = -Mul(υ, i) = +τ
	Mul(t, u) = -Mul(u, t) = -i
	Mul(t, ρ) = -Mul(ρ, t) = +τ
	Mul(t, σ) = -Mul(σ, t) = +υ
	Mul(t, τ) = -Mul(τ, t) = +ρ
	Mul(t, υ) = -Mul(υ, t) = +σ
	Mul(u, ρ) = -Mul(ρ, u) = +υ
	Mul(u, σ) = -Mul(σ, u) = -τ
	Mul(u, τ) = -Mul(τ, u) = -σ
	Mul(u, υ) = -Mul(υ, u) = +ρ
	Mul(ρ, σ) = Mul(σ, ρ) = 0
	Mul(ρ, τ) = Mul(τ, ρ) = 0
	Mul(ρ, υ) = Mul(υ, ρ) = 0
	Mul(σ, τ) = Mul(τ, σ) = 0
	Mul(σ, υ) = Mul(υ, σ) = 0
	Mul(τ, υ) = Mul(υ, τ) = 0
```
The infra Cockle quaternions are different from the more familiar dual split-quaternions.

Alternatively, you can obtain the infra Cockle quaternions from a *hyperbolic* Cayley-Dickson construct with `rational.InfraComplex` values; an *elliptic* Cayley-Dickson construct with `rational.InfraPerplex` values; or a *hyperbolic* Cayley-Dickson construct with `rational.InfraPerplex` values.

### rational.SupraComplex

The `rational.SupraComplex` type represents a rational supra complex number. It corresponds to a parabolic Cayley-Dickson construct with `rational.InfraComplex` values. The imaginary unit element is denoted `i`, and the dual unit elements are denoted `α`, `β`, `γ`, `δ`, `ε`, and `ζ`. The multiplication rules are:
```go
	Mul(i, i) = -1
	Mul(α, α) = Mul(β, β) = Mul(γ, γ) = 0
	Mul(δ, δ) = Mul(ε, ε) = Mul(ζ, ζ) = 0
	Mul(i, α) = -Mul(α, i) = +β
	Mul(i, β) = -Mul(β, i) = -α
	Mul(i, γ) = -Mul(γ, i) = +δ
	Mul(i, δ) = -Mul(δ, i) = -γ
	Mul(i, ε) = -Mul(ε, i) = -ζ
	Mul(i, ζ) = -Mul(ζ, i) = +ε
	Mul(α, β) = Mul(β, α) = 0
	Mul(α, γ) = -Mul(γ, α) = +ε
	Mul(α, δ) = -Mul(δ, α) = +ζ
	Mul(α, ε) = Mul(ε, α) = 0
	Mul(α, ζ) = Mul(ζ, α) = 0
	Mul(β, γ) = -Mul(γ, β) = +ζ
	Mul(β, δ) = -Mul(δ, β) = -ε
	Mul(β, ε) = Mul(ε, β) = 0
	Mul(β, ζ) = Mul(ζ, β) = 0
	Mul(γ, δ) = Mul(δ, γ) = 0
	Mul(γ, ε) = Mul(ε, γ) = 0
	Mul(γ, ζ) = Mul(ζ, γ) = 0
	Mul(δ, ε) = Mul(ε, δ) = 0
	Mul(δ, ζ) = Mul(ζ, δ) = 0
	Mul(ε, ζ) = Mul(ζ, ε) = 0
```
Alternatively, you can obtain the supra complex numbers from an *elliptic* Cayley-Dickson construct with `rational.Supra` values.

### rational.SupraPerplex

The `rational.SupraPerplex` type represents a rational supra perplex number. It corresponds to a parabolic Cayley-Dickson construct with `rational.InfraPerplex` values. The split unit element is denoted `s`, and the dual unit elements are denoted `ρ`, `σ`, `τ`, `υ`, `φ`, and `ψ`. The multiplication rules are:
```go
	Mul(s, s) = +1
	Mul(ρ, ρ) = Mul(σ, σ) = Mul(τ, τ) = 0
	Mul(υ, υ) = Mul(φ, φ) = Mul(ψ, ψ) = 0
	Mul(s, ρ) = -Mul(ρ, s) = +σ
	Mul(s, σ) = -Mul(σ, s) = +ρ
	Mul(s, τ) = -Mul(τ, s) = +υ
	Mul(s, υ) = -Mul(υ, s) = +τ
	Mul(s, φ) = -Mul(φ, s) = -ψ
	Mul(s, ψ) = -Mul(ψ, s) = -φ
	Mul(ρ, σ) = Mul(σ, ρ) = 0
	Mul(ρ, τ) = -Mul(τ, ρ) = +φ
	Mul(ρ, υ) = -Mul(υ, ρ) = +ψ
	Mul(ρ, φ) = Mul(φ, ρ) = 0
	Mul(ρ, ψ) = Mul(ψ, ρ) = 0
	Mul(σ, τ) = -Mul(τ, σ) = +ψ
	Mul(σ, υ) = -Mul(υ, σ) = +φ
	Mul(σ, φ) = Mul(φ, σ) = 0
	Mul(σ, ψ) = Mul(ψ, σ) = 0
	Mul(τ, υ) = Mul(υ, τ) = 0
	Mul(τ, φ) = Mul(φ, τ) = 0
	Mul(τ, ψ) = Mul(ψ, τ) = 0
	Mul(υ, φ) = Mul(φ, υ) = 0
	Mul(υ, ψ) = Mul(ψ, υ) = 0
	Mul(φ, ψ) = Mul(ψ, φ) = 0
```
Alternatively, you can obtain the supra perplex numbers from a *hyperbolic* Cayley-Dickson construct with `rational.Supra` values.

## And Beyond...

Using any of the Cayley-Dickson constructs on any of the eight-dimensional types would produce one of nine sixteen-dimensional types. These types include the [sedenions](https://en.wikipedia.org/wiki/Sedenion), which are infamous for containing zero divisors.

## Other Types

Besides all the types mentioned above, other types are included. These types have multiplication operations that do not involve conjugation.

### rational.BiComplex

The `rational.BiComplex` type represents a [bicomplex number](https://en.wikipedia.org/wiki/Bicomplex_number). It corresponds to a non-sesquilinear elliptic construct with `rational.Complex` values. The imaginary units are denoted `i` and `J`, and the split unit is denoted `S`. The multiplication rules are:
```go
	Mul(i, i) = Mul(J, J) = -1
	Mul(S, S) = +1
	Mul(i, J) = Mul(J, i) = +S
	Mul(J, S) = Mul(S, J) = -i
	Mul(S, i) = Mul(i, S) = -J
```
Note that this multiplication operation is **commutative** and **associative**.

### rational.BiPerplex

The `rational.BiPerplex` type represents a biperplex number. It corresponds to a non-sesquilinear hyperbolic construct with `rational.Perplex` values. The split units are denoted `s`, `T`, and `U`. The multiplication rules are:
```go
	Mul(s, s) = Mul(T, T) = Mul(U, U) = +1
	Mul(s, T) = Mul(T, s) = U
	Mul(T, U) = Mul(U, T) = s
	Mul(U, s) = Mul(s, U) = T
```
Note that this multiplication operation is **commutative** and **associative**.

### rational.Hyper

The `rational.Hyper` type represents a hyper-dual number. It corresponds to a non-sesquilinear parabolic construct with `rational.Infra` values. The dual units are denoted `α`, `κ`, and `λ`. The multiplication rules are:
```go
	Mul(α, α) = Mul(κ, κ) = Mul(λ, λ) = 0
	Mul(α, κ) = Mul(κ, α) = λ
	Mul(κ, λ) = Mul(λ, κ) = 0
	Mul(λ, α) = Mul(α, λ) = 0
```
Note that this multiplication operation is **commutative** and **associative**.

Hyper-dual numbers are useful for computing second-order derivatives. If `z = a + 1α + 1κ + 0λ`, and `f` is a function, then
```
	f(z) = f(a) + f'(a)α + f'(a)κ + f''(a)λ
```
Another name for the dual numbers could be the *nilplex* numbers. Thus, in analogy with the bicomplex numbers, another name for the hyper-dual numbers could be the *binilplex* numbers.

### rational.DualComplex

The `rational.DualComplex` type represents a dual-complex number. It corresponds to a non-sesquilinear parabolic construct with `rational.Complex` values. The imaginary unit is denoted `i`, and the dual units are denoted `κ` and `λ`. The multiplication rules are:
```go
	Mul(i, i) = -1
	Mul(κ, κ) = Mul(λ, λ) = 0
	Mul(i, κ) = Mul(κ, i) = +λ
	Mul(κ, λ) = Mul(λ, κ) = 0
	Mul(λ, i) = Mul(i, λ) = -κ
```
Note that this multiplication operation is **commutative** and **associative**.

Dual-complex numbers are useful for evaluating two-dimensional vector derivatives. Let `(x, y)` be the coordinate vector in a two-dimensional space. Also, let `f = (g, h)` be a function vector. Then, the **divergence** and the **curl** of `f` are:
```go
	Div(f) = (∂g / ∂x) + (∂h / ∂y)
	Curl(f) = (∂h / ∂x) - (∂g / ∂y)
```
A two-dimensional vector `(x, y)` can be associated with a complex number `z = x + yi`. Introduce the Wirtinger derivative
```
	(∂ / ∂z) = (1 / 2) * ((∂ / ∂x) - (∂ / ∂y)i)
```
The function vector `(g, h)` becomes the complex number `f = g + hi`. Then:
```
	2 * (∂f / ∂z) = Div(f) + (Curl(f))i
```
Thus, in analogy with the dual numbers, given a dual-complex number `p = a + bi + 2κ + 0λ`, and a function `f`, then
```
	f(p) = f(a + bi) + (∂f / ∂z) * (2κ + 0λ)
		 = f(a + bi) + (Div(f))κ + (Curl(f))λ
```
In this manner, `f`, `Div(f)`, and `Curl(f)` can be calculated at the point `a + bi` by just evaluating `f(a + bi + 2κ + 0λ)`. Note that this derivation assumes that `f` is holomorphic (i.e. that it only depends on `a` and `b` in the combination `a + bi`).

### rational.DualPerplex

The `rational.DualPerplex` type represents a dual-perplex number. It corresponds to a non-sesquilinear parabolic construct with `rational.Perplex` values. The split unit is denoted `s`, and the dual units are denoted `κ` and `λ`. The multiplication rules are:
```go
	Mul(s, s) = +1
	Mul(κ, κ) = Mul(λ, λ) = 0
	Mul(s, κ) = Mul(κ, s) = +λ
	Mul(κ, λ) = Mul(λ, κ) = 0
	Mul(λ, s) = Mul(s, λ) = +κ
```
Note that this multiplication operation is **commutative** and **associative**.

Just like dual-complex numbers, the dual-perplex numbers can be used to evaluate certain vector derivatives. Let `w = x + ys` and `f = g + hs`. The Wirtinger derivative is now
```
	(∂ / ∂w) = (1 / 2) * ((∂ / ∂x) + (∂ / ∂y)s)
```
Then:
```
	2 * (∂f / ∂w) = ((∂g / ∂x) + (∂h / ∂y)) + ((∂h / ∂x) + (∂g / ∂y))s
```
This suggests the introduction of the hyperbolic version of the curl:
```
	Hurl(f) = (∂h / ∂x) + (∂g / ∂y)
```
Now let `p = a + bs + 2κ + 0λ`. Thus,
```
	f(p) = f(a + bs) + (Div(f))κ + (Hurl(f))λ
```
Again, `f`, `Div(f)`, and now `Hurl(f)` can be calculated at the point `a + bs` by just evaluating `f(a + bs + 2κ + 0λ)`.

## To Do

1. Improve documentation
1. Tests
1. TriComplex type
1. TriPerplex type
1. TriNilplex type
1. BiHamilton type
1. BiCockle type
1. DualHamilton type
1. DualCockle type
1. Elementary and special functions via Padé approximants