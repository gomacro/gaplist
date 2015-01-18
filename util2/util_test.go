// Copyright 2015 The GOMACRO Authors. All rights reserved.
// Use of this source code is governed by a GPLv2-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"testing"
)

const dEBUG = false

func round(t *testing.T, j, k int) {
	max := -1
	min := 0
	iters := 0

	for i := Bgn(j); Tst(i, k); i = Nxt(i, j) {

		if -1-i >= j {
			t.Errorf("array [0] overrun i=%v", -1-i)
		}
		if 0 > k {
			t.Errorf("array [1] overrun i=%v", -1-i)
		}
		if i >= 0 && Off(i) != i {
			t.Errorf("off [0]  i=%v, off=%v", i, Off(i))
		}
		if i < 0 && Off(i) != -i-1 {
			t.Errorf("off [1]  -i-1=%v, off=%v", -i-1, Off(i))
		}

		if dEBUG {

			fmt.Println(i, max, min, Pos(i), Off(i))

		}

		iters++

		if i > max {
			max = i
		} else if iters > j {
			t.Errorf("dwn %v", i)
		}

		if i < min {
			min = i
		} else if iters <= j {
			t.Errorf("up %v", i)
		}

	}

	if iters != j+k {
		t.Errorf("iters %v, need %v", iters, j+k)
	}
}

func TestCustom0(t *testing.T) {
	round(t, 10, 10)
}
func TestCustom1(t *testing.T) {
	round(t, 10, 0)
}
func TestCustom2(t *testing.T) {
	round(t, 0, 10)
}
func TestCustom3(t *testing.T) {
	round(t, 0, 0)
}
