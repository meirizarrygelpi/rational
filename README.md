# rational

Package `rational` brings rational [complex](https://en.wikipedia.org/wiki/Complex_number), [split-complex](https://en.wikipedia.org/wiki/Split-complex_number), and [dual](https://en.wikipedia.org/wiki/Dual_number) numbers to Go. It borrows heavily from the `math`, `math/cmplx`, and `math/big` packages. Indeed, it is built on top of the `big.Rat` type.

[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/github.com/meirizarrygelpi/rational) [![GoDoc](https://godoc.org/github.com/meirizarrygelpi/rational?status.svg)](https://godoc.org/github.com/meirizarrygelpi/rational)

## Two-Dimensional Types

`rational` contains three two-dimensional types: `rational.Complex`, `rational.Perplex`, and `rational.Infra`.

## Four-Dimensional Types

`rational` contains five four-dimensional types: `rational.Hamilton`, `rational.Cockle`, `rational.Supra`, `rational.InfraComplex`, and `rational.InfraPerplex`.

## Eight-Dimensional Types

`rational` contains seven eight-dimensional types: `rational.Cayley`, `rational.Klein`, `rational.Ultra`, `rational.SupraComplex`, `rational.SupraPerplex`, `rational.InfraHamilton`, and `rational.InfraCockle`.

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