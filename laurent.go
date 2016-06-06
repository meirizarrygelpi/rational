// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package rational

import "math/big"

// Laurent represents a Laurent univariate polynomial with rational
// coefficients.
type Laurent map[int64]*big.Rat

// merge function.
func merge(a, b []int64) []int64 {
	la, lb := len(a), len(b)
	i, j := 0, 0
	n := la + lb
	c := make([]int64, n)
	for k := 0; k < n; k++ {
		if (i < la) && (j < lb) {
			if a[i] < b[j] {
				c[k] = a[i]
				i++
			} else {
				c[k] = b[j]
				j++
			}
		} else if i > la-1 {
			c[k] = b[j]
			j++
		} else {
			c[k] = a[i]
			i++
		}
	}
	return c
}

// sort function.
func sort(x []int64) []int64 {
	n := len(x)
	if (n == 0) || (n == 1) {
		return x
	}
	h := n / 2
	a := sort(x[:h])
	b := sort(x[h:])
	return merge(a, b)
}

// reverse function.
func reverse(x []int64) []int64 {
	n := len(x)
	y := make([]int64, n)
	for i := 0; i < n; i++ {
		y[i] = x[len(x)-i-1]
	}
	return y
}

// Degrees returns two sorted slices with the negative and non-negative degrees
// in p.
func (p Laurent) Degrees() (neg, nonneg []int64) {
	for n := range p {
		if n < 0 {
			neg = append(neg, n)
		} else {
			nonneg = append(nonneg, n)
		}
		neg = reverse(sort(neg))
		nonneg = sort(nonneg)
	}
	return
}
