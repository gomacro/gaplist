// Copyright 2015 The GOMACRO Authors. All rights reserved.
// Use of this source code is governed by a GPLv2-style
// license that can be found in the LICENSE file.

package util

import (
	"testing"
	"fmt"
)

func round(t *testing.T, j, k int) {
	max := -1
	min := 0
	iters := 0
	i := 0
	for {
		if i >= j {
			t.Errorf("array [0] overrun i=%v", i)
		}
		if -i > k {
			t.Errorf("array [1] overrun i=%v", i)
		}

		fmt.Println(i, max, min)

		iters++

		if i > max {
			max = i
		} else if iters <= j  {
			t.Errorf("dwn %v", i)
		}

		if i < min {
			min = i
		} else if iters > j {
			t.Errorf("up %v", i)
		}

		// NEXT

		i = Nxt(i, j, k)
		if i == 0 {
			break
		}
	}
	if iters != j + k {
		t.Errorf("iters %v", iters)
	}
}

func TestCustom0(t *testing.T) {
//	round(t, 10, 10)


	round(t, 0, 0)
}
