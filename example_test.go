// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list_test

import (
	"fmt"
	. "example.com/repo.git/gaplist/util"
	"testing"
)
/*
func Append(list *[2][][]int, elem ... int) {
	(*list)[1][0] = append((*list)[1][0], elem...)
}
*/

/*
func Prepend(list *[2][][]int, elem ... int) {
	(*list)[0][0] = append(elem, (*list)[0][0]...)
}
*/


/*
func TestCustom0(t *testing.T) {
	var list [2][][]byte

	list[0] = [][]byte{{1,1,1,1,1,7,1,1,1,7,1,1,7,1,1}}
	list[1] = [][]byte{{1,1,1,1,1,7,1,1,1,7,1,1,7,1,1}}

	fmt.Println(list)

	for i := Bgn(len(list[0])); Tst(i, len(list[1])); i = Nxt(i, len(list[0]), len(list[1])) {
		chunk := &list[Neg(i)][Off(i, len(list[1]))]

		fmt.Println("Looping on ", *chunk)

		for j, val := range *chunk {
			if val == 7 {
	fmt.Println(list)
				fmt.Println("I am at", i, j, "gaps here here")



				if Neg(i) == 0 {

					list[1] = append([][]int{(*chunk)[j:]}, list[1]...)
//					*chunk = (*chunk)[:j]
					list[0] = append(list[0], []int{8})
					i = 999
					break
				} else {
				fmt.Println("chunk", *chunk)
//					list[0] = append(list[0], (*chunk)[:j])
					list[0] = append(list[0], []int{8})
					*chunk = (*chunk)[j:]
					break
				}

				// now i will insert digit 8 before

//				*chunk = (*chunk)[:j]

//				list[Neg(i)] = append(list[Neg(i)], []int{9})

			}
		}
//		fmt.Println(chunk)
	}

	fmt.Println(list)

	if false {
		t.Errorf("")
	}
}

*/
func Append(src [2][][]byte, a ...byte) [2][][]byte {
	return AppendS(src, a)
}
func AppendS(src [2][][]byte, a []byte) (dst [2][][]byte) {
	dst = src
	dst[1] = append(dst[1], a)
	return dst
}

func From(dst *[2][][]byte, src []byte) {
	var d [2][][]byte
	d[0] = append(d[0], src)
	*dst = d
}

func Collapse(dst *[]byte, src [2][][]byte) {
	for _, o := range src {
		for _, p := range o {
			*dst = append(*dst, p...)
		}
	}
}
func Slice(dst *[2][][]byte, src [2][][]byte, i, p, j, q int) {
	(*dst)[0] = src[0][i:]
	(*dst)[0][0] = (*dst)[0][0][p:]
	(*dst)[1] = src[1][:j-1+len(src[1])]
	l := len((*dst)[1])-1
	(*dst)[1][l] = (*dst)[1][l][:q]
}

func TestCustom0(t *testing.T) {
	var list [2][][]byte

	list[0] = [][]byte{{1,1,1,1,1,7,1,1,1,7,1,1,7,1,1}}
	list[1] = [][]byte{{1,1,1,1,1,7,1,1,1,7,1,1,7,1,1}}

	for i := Bgn(len(list[0])); Tst(i, len(list[1])); i = Nxt(i, len(list[0]), len(list[1])) {

	}


//	fmt.Println(list)
	list = AppendS(list, []byte{1,3,3,7})
	list = Append(list, 2, 4, 4, 8)


//	fmt.Println(list)

	var lst []byte

	Collapse(&lst, list)

//	fmt.Println(lst)

	From(&list, lst)
//	fmt.Println(list)
}

func TestSlice0(t *testing.T) {
	var list [2][][]byte
	list[0] = [][]byte{{1,2,3},{4,5,6,7,8},{9,10,11},{12,13},{14,15}}
	list[1] = [][]byte{{16,17},{18,19},{20,21,22},{23,24},{25,26,27}}
	fmt.Println(list)

	var list2 [2][][]byte

	Slice(&list2, list, 1, 3, -2, 1)

	fmt.Println(list2)

}

/*
func Example() {
	var list [2][][]int

	fmt.Println("hello")

	Append(&list, 1, 3, 3, 7)
	Prepend(&list, 8, 2, 2, 6)


	// Create a new list and put some numbers in it.
//	l := list.New()
//	e4 := l.PushBack(4)
//	e1 := l.PushFront(1)
//	l.InsertBefore(3, e4)
//	l.InsertAfter(2, e1)

	fmt.Println(list)

	for i := Bgn(len(list[0])); Tst(i, len(list[1])); i = Nxt(i, len(list[0]), len(list[1])) {
		fmt.Println(list[Neg(i)][Off(i, len(list[1]))])
	}

	// Iterate through list and print its contents.
//	for e := l.Front(); e != nil; e = e.Next() {
//		fmt.Println(e.Value)
//	}

	// Output:
	// 1
	// 2
	// 3
	// 4
}
*/
