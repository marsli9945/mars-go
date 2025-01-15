package test

import (
	"github.com/marsli9945/mars-go/marsLog"
	"github.com/marsli9945/mars-go/marsType"
	"testing"
)

func TestArray(t *testing.T) {
	arr := marsType.ArrayInitForList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	//arr := marsType.ArrayInitForList([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"})
	arrayStr := arr.Join(",")
	marsLog.Logger().InfoF("arrayStr: %s", arrayStr)

	contains2 := arr.Contains(2)
	contains11 := arr.Contains(11)

	marsLog.Logger().InfoF("contains2: %v", contains2)
	marsLog.Logger().InfoF("contains11: %v", contains11)

	notContains2 := arr.NotContains(2)
	notContains11 := arr.NotContains(11)

	marsLog.Logger().InfoF("notContains2: %v", notContains2)
	marsLog.Logger().InfoF("notContains11: %v", notContains11)
}
