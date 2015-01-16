// Copyright 2015 The GOMACRO Authors. All rights reserved.
// Use of this source code is governed by a GPLv2-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"testing"
)

const dEBUG = true

func round(t *testing.T, j, k int) {
	max := -1
	min := 0
	iters := 0
	i := 0
	if j == 0 {
		i = -1
	}
	if j+k != 0 {



		for i >= -k {

			if i >= j {
				t.Errorf("array [0] overrun i=%v", i)
			}
			if -i > k {
				t.Errorf("array [1] overrun i=%v", i)
			}
			if i >= 0 && Off(i, k) != i {
				t.Errorf("off [0]  i=%v, off=%v", i, Off(i, k))
			}
			if i < 0 && Off(i, k) != i+k {
				t.Errorf("off [1]  i+k=%v, off=%v", i+k, Off(i, k))
			}

			if dEBUG {

				fmt.Println(i, max, min, Neg(i), Off(i, k))

			}

			iters++

			if i > max {
				max = i
			} else if iters <= j {
				t.Errorf("dwn %v", i)
			}

			if i < min {
				min = i
			} else if iters > j {
				t.Errorf("up %v", i)
			}

			// NEXT

			i = Nxt(i, j, k)

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
