package marsType

import (
	"github.com/marsli9945/mars-go/marsType"
	"reflect"
	"sort"
	"testing"
)

func TestSet_Add(t *testing.T) {
	set := marsType.NewSet[int]()
	set.Add(1)
	if !set.Contains(1) {
		t.Errorf("Expected set to contain 1")
	}
}

func TestSet_AddAll(t *testing.T) {
	set := marsType.NewSet[int]()
	set.AddAll([]int{1, 2, 3})
	if !set.Contains(1) || !set.Contains(2) || !set.Contains(3) {
		t.Errorf("Expected set to contain 1, 2, and 3")
	}
}

func TestSet_Remove(t *testing.T) {
	set := marsType.NewSet[int]()
	set.Add(1)
	set.Remove(1)
	if set.Contains(1) {
		t.Errorf("Expected set to not contain 1")
	}
}

func TestSet_Contains(t *testing.T) {
	set := marsType.NewSet[int]()
	set.Add(1)
	if !set.Contains(1) {
		t.Errorf("Expected set to contain 1")
	}
	if set.Contains(2) {
		t.Errorf("Expected set to not contain 2")
	}
}

func TestSet_ToList(t *testing.T) {
	set := marsType.NewSet[int]()
	set.AddAll([]int{1, 2, 3})
	list := set.ToList()
	sort.Ints(list)
	expected := marsType.Array[int]{1, 2, 3}
	if !reflect.DeepEqual(list, expected) {
		t.Errorf("Expected list to be %v, got %v", expected, list)
	}
}

func TestSet_EmptyToList(t *testing.T) {
	set := marsType.NewSet[int]()
	list := set.ToList()
	if len(list) != 0 {
		t.Errorf("Expected empty list, got %v", list)
	}
}
