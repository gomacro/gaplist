package loop

import (
	. "example.com/repo.git/gaplist/util2"
	"reflect"
)

func Begin(list interface{}) int {
	j := reflect.ValueOf(list).Index(0).Len()
	return Bgn(j)
}
func In(i int, list interface{}) bool {
	k := reflect.ValueOf(list).Index(1).Len()
	return I(i, k)
}
func Next(i int, list interface{}) int {
	j := reflect.ValueOf(list).Index(0).Len()
	return Nxt(i, j)
}
