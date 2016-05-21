# rational

Package `rational` brings rational [complex](https://en.wikipedia.org/wiki/Complex_number), [split-complex](https://en.wikipedia.org/wiki/Split-complex_number), and [dual](https://en.wikipedia.org/wiki/Dual_number) numbers to Go. It borrows heavily from the `math`, `math/cmplx`, and `math/big` packages. Indeed, it is built on top of the `big.Rat` type.

[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/github.com/meirizarrygelpi/rational) [![GoDoc](https://godoc.org/github.com/meirizarrygelpi/rational?status.svg)](https://godoc.org/github.com/meirizarrygelpi/rational)

## Two-Dimensional Types

There are three two-dimensional types. The (binary) multiplication operation for all two-dimensional types is **commutative** and **associative**.

### rational.Complex

The `rational.Complex` type represents a rational complex number. It corresponds to a complex Cayley-Dickson construct with `big.Rat` values. The imaginary unit element is denoted `i`.

This type can be used to study [Gaussian integers](https://en.wikipedia.org/wiki/Gaussian_integer).

### rational.Perplex

The `rational.Perplex` type represents a rational perplex number. It corresponds to a perplex Cayley-Dickson construct with `big.Rat` values. The split unit element is denoted `s`.

Perplex numbers are more commonly known as [split-complex numbers](https://en.wikipedia.org/wiki/Split-complex_number), but "perplex" is used here for brevity and symmetry with "complex".

### rational.Infra

The `rational.Infra` type represents a rational infra number. It corresponds to a null Cayley-Dickson construct with `big.Rat` values. The dual unit element is denoted `α`.

Infra numbers are more commonly known as [dual numbers](https://en.wikipedia.org/wiki/Dual_number), but "infra" is used here along with "supra" and "ultra" for the higher-dimensional analogs.

## Four-Dimensional Types

There are five four-dimensional types. The (binary) multiplication operation for all four-dimensional types is **noncommutative** but **associative**.

### rational.Hamilton

The `rational.Hamilton` type represents a rational Hamilton quaternion. It corresponds to a complex Cayley-Dickson construct with `rational.Complex` values. The imaginary unit elements are denoted `i`, `j`, and `k`.

Hamilton quaternions are [traditional quaternions](https://en.wikipedia.org/wiki/Quaternion). The type is named after W.R. Hamilton, who discovered them.

This type can be used to study [Hurwitz and Lipschitz integers](https://en.wikipedia.org/wiki/Hurwitz_quaternion).

### rational.Cockle

The `rational.Cockle` type represents a rational Cockle quaternion. It corresponds to a perplex Cayley-Dickson construct with `rational.Complex` values. The imaginary unit element is denoted `i`, and the split unit elements are denoted `t` and `u`.

Cockle quaternions are more commonly known as [split-quaternions](https://en.wikipedia.org/wiki/Split-quaternion). The type is named after J. Cockle, who discovered them.

### rational.Supra

The `rational.Supra` type represents a rational supra number. It corresponds to a null Cayley-Dickson construct with `rational.Infra` values. The dual unit elements are denoted `α`, `β`, and `γ`.

Note that supra numbers are very different from [hyper-dual numbers](http://adl.stanford.edu/hyperdual/): the multiplication operation for supra numbers is noncommutative.

### rational.InfraComplex

The `rational.InfraComplex` type represents a rational infra complex number. It corresponds to a null Cayley-Dickson construct with `rational.Complex` values. The imaginary unit element is denoted `i`, and the dual unit elements are denoted `α` and `β`.

### rational.InfraPerplex

The `rational.InfraPerplex` type represents a rational infra perplex number. It corresponds to a null Cayley-Dickson construct with `rational.Perplex` values. The split unit element is denoted `s`, and the dual unit elements are denoted `σ` and `τ`.

## Eight-Dimensional Types

There are seven eight-dimensional types. The (binary) multiplication operation for all eight-dimensional types is **noncommutative** and **nonassociative**.

### rational.Cayley

...

### rational.Klein

...

### rational.Ultra

...

### rational.InfraHamilton

...

### rational.InfraCockle

...

### rational.SupraComplex

...

### rational.SupraPerplex

...

## To Do

1. Improve documentation
1. Tests
1. Improve README
1. Improve memory management
1. Add InfraHamilton type
1. Add InfraCockle type
1. Add SupraComplex type
1. Add SupraPerplex type