package marsType

import (
	"fmt"
	"testing"
	"time"
)

func Test_Property(t *testing.T) {
	// 标准库 map
	start := time.Now()
	m := make(map[string]interface{})
	for i := 0; i < 1000000; i++ {
		m[fmt.Sprintf("key%d", i)] = i
	}
	for i := 0; i < 1000000; i++ {
		m[fmt.Sprintf("key%d", i)] = i + 2
	}
	for i := 0; i < 1000000; i++ {
		delete(m, fmt.Sprintf("key%d", i))
	}
	fmt.Println("Standard map insert time:", time.Since(start))

	// Swiss Table
	start = time.Now()
	st := NewSwissTable()
	for i := 0; i < 1000000; i++ {
		st.Insert(fmt.Sprintf("key%d", i), i)
	}
	for i := 0; i < 1000000; i++ {
		st.Insert(fmt.Sprintf("key%d", i), i+2)
	}
	for i := 0; i < 1000000; i++ {
		st.Delete(fmt.Sprintf("key%d", i))
	}
	fmt.Println("Swiss Table insert time:", time.Since(start))
}
