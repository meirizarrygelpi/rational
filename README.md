# rational

Package `rational` brings rational [complex](https://en.wikipedia.org/wiki/Complex_number), [split-complex](https://en.wikipedia.org/wiki/Split-complex_number), and [dual](https://en.wikipedia.org/wiki/Dual_number) numbers to Go. It borrows heavily from the `math`, `math/cmplx`, and `math/big` packages. Indeed, it is built on top of the `big.Rat` type.

[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/github.com/meirizarrygelpi/rational) [![GoDoc](https://godoc.org/github.com/meirizarrygelpi/rational?status.svg)](https://godoc.org/github.com/meirizarrygelpi/rational)

## Two-Dimensional Types

There are three two-dimensional types. The (binary) multiplication operation for all two-dimensional types is commutative and associative.

### rational.Complex

The `rational.Complex` type represents a rational complex number. It corresponds to a complex Cayley-Dickson construct with `big.Rat` values.

This type can be used to study [Gaussian integers](https://en.wikipedia.org/wiki/Gaussian_integer).

### rational.Perplex

The `rational.Perplex` type represents a rational perplex number. It corresponds to a perplex Cayley-Dickson construct with `big.Rat` values. Perplex numbers are more commonly known as [split-complex numbers](https://en.wikipedia.org/wiki/Split-complex_number), but "perplex" is used here for brevity and symmetry with "complex".

### rational.Infra

The `rational.Infra` type represents a rational infra number. It corresponds to a null Cayley-Dickson construct with `big.Rat` values. Infra numbers are more commonly known as [dual numbers](https://en.wikipedia.org/wiki/Dual_number), but "infra" is used here along with "supra" and "ultra" for the higher-dimensional analogs.

## Four-Dimensional Types

There are five four-dimensional types: `rational.Hamilton`, `rational.Cockle`, `rational.Supra`, `rational.InfraComplex`, and `rational.InfraPerplex`. The (binary) multiplication operation for all four-dimensional types is noncommutative but associative.

## Eight-Dimensional Types

There are seven eight-dimensional types: `rational.Cayley`, `rational.Klein`, `rational.Ultra`, `rational.InfraHamilton`, `rational.InfraCockle`, `rational.SupraComplex`, and `rational.SupraPerplex`. The (binary) multiplication operation for all eight-dimensional types is noncommutative and nonassociative.

## To Do

1. Improve documentation
1. Tests
1. Improve README
1. Improve memory management
1. Add InfraPerplex type
1. Add InfraHamilton type
1. Add InfraCockle type
1. Add SupraComplex type
1. Add SupraPerplex type