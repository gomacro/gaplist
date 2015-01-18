// Copyright 2015 The GOMACRO Authors. All rights reserved.
// Use of this source code is governed by a GPLv2-style
// license that can be found in the LICENSE file.

package util

// Pos is used to obtain the gap-buffer's half from interator
func Pos(i int) int {
	if i < 0 {
		return 0
	}
	return 1
}

// Off is ussed to obtain an offset within the half from interator
// The parameter Off(k) is the len([1])
func Off(i int) int {
	if i < 0 {
		return -1 - i
	}
	return i
}

// Nxt is used to increment the interator
// The parameter Nxt(j) is len([0])
func Nxt(i, j int) int {
	if i >= 0 {
		return i + 1
	}
	if j == -i {
		return 0
	}
	return i - 1
}

// Bgn is used to get the begin interator
// The parameter Bgn(j) is len([0])
func Bgn(j int) int {
	if j == 0 {
		return 0
	}
	return -1
}

// Bgn is used to get the begin interator
// The parameter Tst(k) is len([1])
func Tst(i, k int) bool {
	return i < k
}
