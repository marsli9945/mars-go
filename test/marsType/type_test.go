package marsType

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

	arrayInitForMap := marsType.ArrayInitForMap(map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true})
	marsLog.Logger().InfoF("arrayInitForMap: %v", arrayInitForMap)

	str := arrayInitForMap.Join(",")
	marsLog.Logger().InfoF("arrayInitForMap: %s", str)
}

func TestQueue(t *testing.T) {
	queue := make(marsType.Queue[int], 0, 5)
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	queue.Push(5)

	//for !queue.IsEmpty() {
	//	marsLog.Logger().InfoF("queue: %v", queue.Pop())
	//}

	for i := 0; i < 6; i++ {
		marsLog.Logger().InfoF("queue: %v", queue.Pop())
	}
}

func TestSet(t *testing.T) {
	set := marsType.NewSet[string]()
	elements := []string{"1", "2", "3", "4", "5", "6"}
	set.AddAll(elements)
	//set.Add("1")
	//set.Add("2")
	//set.Add("3")
	//set.Add("3")
	//set.Add("4")
	//set.Add("4")
	//set.Add("4")
	//set.Add("4")
	//set.Add("5")
	//set.Add("6")

	marsLog.Logger().InfoF("list: %v", set.ToList())

	set.Remove("4")
	marsLog.Logger().InfoF("list: %v", set.ToList())
}
