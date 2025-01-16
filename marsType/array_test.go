package marsType

import (
	"reflect"
	"sort"
	"testing"
)

func TestArrayInitForList(t *testing.T) {
	tests := []struct {
		input    []int
		expected Array[int]
	}{
		{nil, Array[int]{}},
		{[]int{1, 2, 3}, Array[int]{1, 2, 3}},
	}

	for _, test := range tests {
		actual := ArrayInitForList(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("ArrayInitForList(%v) = %v; expected %v", test.input, actual, test.expected)
		}
	}
}

func TestArrayInitForMap(t *testing.T) {
	tests := []struct {
		input    map[int]bool
		expected Array[int]
	}{
		{nil, Array[int]{}},
		{map[int]bool{1: true, 2: true}, Array[int]{1, 2}},
	}

	for _, test := range tests {
		actual := ArrayInitForMap(test.input)
		sort.Ints(actual)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("ArrayInitForMap(%v) = %v; expected %v", test.input, actual, test.expected)
		}
	}
}

func TestSplitArray(t *testing.T) {
	tests := []struct {
		input     Array[int]
		chunkSize int
		expected  []Array[int]
	}{
		{Array[int]{1, 2, 3, 4, 5}, 2, []Array[int]{{1, 2}, {3, 4}, {5}}},
		{Array[int]{1, 2, 3, 4, 5}, 0, []Array[int]{}},
		{Array[int]{}, 3, []Array[int]{}},
	}

	for _, test := range tests {
		actual := test.input.SplitArray(test.chunkSize)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("SplitArray(%v, %d) = %v; expected %v", test.input, test.chunkSize, actual, test.expected)
		}
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		input    Array[int]
		target   int
		expected bool
	}{
		{Array[int]{1, 2, 3}, 2, true},
		{Array[int]{1, 2, 3}, 4, false},
	}

	for _, test := range tests {
		actual := test.input.Contains(test.target)
		if actual != test.expected {
			t.Errorf("Contains(%v, %d) = %v; expected %v", test.input, test.target, actual, test.expected)
		}
	}
}

func TestNotContains(t *testing.T) {
	tests := []struct {
		input    Array[int]
		target   int
		expected bool
	}{
		{Array[int]{1, 2, 3}, 2, false},
		{Array[int]{1, 2, 3}, 4, true},
	}

	for _, test := range tests {
		actual := test.input.NotContains(test.target)
		if actual != test.expected {
			t.Errorf("NotContains(%v, %d) = %v; expected %v", test.input, test.target, actual, test.expected)
		}
	}
}

func TestJoin(t *testing.T) {
	tests := []struct {
		input    Array[int]
		sep      string
		expected string
	}{
		{Array[int]{}, ",", ""},
		{Array[int]{1, 2, 3}, ",", "1,2,3"},
		{Array[int]{1, 2, 3}, "", "123"},
	}

	for _, test := range tests {
		actual := test.input.Join(test.sep)
		if actual != test.expected {
			t.Errorf("Join(%v, %q) = %q; expected %q", test.input, test.sep, actual, test.expected)
		}
	}
}
