// Copyright 2015 The GOMACRO Authors. All rights reserved.
// Use of this source code is governed by a GPLv2-style
// license that can be found in the LICENSE file.

package util

// Neg is used to obtain the gap-buffer's half from interator
func Neg(i int) int {
	if i < 0 {
		return 1
	}
	return 0
}

// Off is ussed to obtain an offset within the half from interator
// The parameter Off(j) is the len([1])
func Off(i, k int) int {
	if i < 0 {
		return k + i
	}
	return i
}

// Nxt is used to increment the interator
// The parameter Nxt(j,k) are len([0]) len([1])
func Nxt(i, j, k int) int {
	if i >= 0 {
		i++
		if i >= j {
			return -1
		}
		return i
	}
	return i-1
}
