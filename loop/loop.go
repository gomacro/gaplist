package loop

import (
	u "example.com/repo.git/gaplist/util2"
	"reflect"
)

func Begin(list interface{}) int {
	return u.Bgn(reflect.ValueOf(list).Index(0).Len())
}
func In(i int, list interface{}) bool {
	return u.I(i, reflect.ValueOf(list).Index(1).Len())
}
func Next(i int, list interface{}) int {
	return u.Nxt(i, reflect.ValueOf(list).Index(0).Len())
}
func Off(i int) int {
	return u.Off(i)
}
func Pos(i int) int {
	return u.Pos(i)
}
func Break() int {
	return u.Brk
}
