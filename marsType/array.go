package marsType

import (
	"fmt"
	"strconv"
	"strings"
)

type Array[T string | int] []T

func ArrayInitForList[T string | int](list []T) Array[T] {
	if list == nil {
		return Array[T]{}
	}

	// 预分配容量以提高性能
	result := make(Array[T], 0, len(list))
	for _, t := range list {
		result = append(result, t)
	}
	return result
}

func ArrayInitForMap[T string | int](m map[T]bool) Array[T] {
	if m == nil {
		return Array[T]{}
	}

	// 创建一个新的 Array[T] 实例，避免副作用
	result := make(Array[T], 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

func (arr Array[T]) SplitArray(chunkSize int) []Array[T] {
	n := len(arr)
	if chunkSize <= 0 || n == 0 {
		return []Array[T]{}
	}

	// 预分配内存
	numChunks := (n + chunkSize - 1) / chunkSize
	result := make([]Array[T], 0, numChunks)

	for i := 0; i < n; i += chunkSize {
		end := min(i+chunkSize, n)
		result = append(result, arr[i:end])
	}

	return result
}

func (arr Array[T]) Contains(target T) bool {
	for _, t := range arr {
		if target == t {
			return true
		}
	}
	return false
}

func (arr Array[T]) NotContains(target T) bool {
	return !arr.Contains(target)
}

func (arr Array[T]) Join(sep string) string {
	if len(arr) == 0 {
		return ""
	}

	var builder strings.Builder
	builder.Grow(len(arr) * (len(sep) + 1)) // 预估长度以提高性能

	for i, s := range arr {
		if i > 0 {
			builder.WriteString(sep)
		}
		builder.WriteString(convertToString(s))
	}

	return builder.String()
}

// convertToString 将传入的参数s，可能是string或者int，都转为string输出
func convertToString[T string | int](s T) string {
	switch v := any(s).(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}
