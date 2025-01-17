package marsContext

import (
	"context"
	"testing"
)

func TestWithCustomValue_ParentContextNil_ReturnsError(t *testing.T) {
	ctx, err := WithCustomValue(nil, "testValue")
	if ctx != nil {
		t.Errorf("Expected context to be nil, got %v", ctx)
	}
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestWithCustomValue_ValidParentContext_ReturnsContext(t *testing.T) {
	parentCtx := context.Background()
	ctx, err := WithCustomValue(parentCtx, "testValue")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if ctx == nil {
		t.Errorf("Expected context to be non-nil")
	}
}

func TestValue_CustomKey_ReturnsCustomValue(t *testing.T) {
	parentCtx := context.Background()
	ctx, _ := WithCustomValue(parentCtx, "testValue")
	value := ctx.Value("customKey")
	if value != "testValue" {
		t.Errorf("Expected customValue 'testValue', got %v", value)
	}
}

func TestValue_NonCustomKey_ReturnsParentValue(t *testing.T) {
	parentCtx := context.WithValue(context.Background(), "key", "parentValue")
	ctx, _ := WithCustomValue(parentCtx, "testValue")
	value := ctx.Value("key")
	if value != "parentValue" {
		t.Errorf("Expected parent value 'parentValue', got %v", value)
	}
}
