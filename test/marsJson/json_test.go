package marsJson

import (
	"github.com/marsli9945/mars-go/marsJson"
	"testing"
)

// Marshal 方法的单元测试
func TestMarshal_Success(t *testing.T) {
	input := map[string]string{"key": "value"}
	expected := `{"key":"value"}`
	result := marsJson.Marshal(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestMarshal_Failure(t *testing.T) {
	// 创建一个循环引用
	type Cycle struct {
		Ptr *Cycle
	}
	cycle := Cycle{}
	cycle.Ptr = &cycle
	result := marsJson.Marshal(cycle)
	if result != "{}" {
		t.Errorf("Expected {}, got %s", result)
	}
}

// UnMarshal 方法的单元测试
func TestUnMarshal_Success(t *testing.T) {
	input := `{"key":"value"}`
	var result map[string]string
	err := marsJson.UnMarshal(input, &result)
	if err != nil {
		t.Errorf("UnMarshal failed: %v", err)
	}
	expected := map[string]string{"key": "value"}
	if !mapsEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func mapsEqual[T comparable, V comparable](m1, m2 map[T]V) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if w, ok := m2[k]; !ok || v != w {
			return false
		}
	}
	return true
}

func TestUnMarshal_Failure(t *testing.T) {
	input := `invalid json`
	var result map[string]string
	err := marsJson.UnMarshal(input, &result)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}

// PrettyString 方法的单元测试
func TestPrettyString_Success(t *testing.T) {
	input := `{"key":"value"}`
	expected := "{\n \"key\": \"value\"\n}"
	result := marsJson.PrettyString(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPrettyString_Failure(t *testing.T) {
	input := `invalid json`
	result := marsJson.PrettyString(input)
	if result != "" {
		t.Errorf("Expected empty string, got %s", result)
	}
}
