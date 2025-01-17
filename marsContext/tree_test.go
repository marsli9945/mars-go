package marsContext

import (
	"context"
	"testing"
	"time"
)

func TestNewContextTree(t *testing.T) {
	rootCtx := context.Background()
	node := NewContextTree(rootCtx, "root")
	if node == nil {
		t.Fatal("Expected a non-nil node")
	}
	if node.Ctx == nil {
		t.Fatal("Expected a non-nil context")
	}
	if node.Cancel == nil {
		t.Fatal("Expected a non-nil cancel function")
	}
	if len(node.Children) != 0 {
		t.Fatal("Expected no children initially")
	}
	if node.Identifier != "root" {
		t.Fatalf("Expected identifier to be 'root', got '%s'", node.Identifier)
	}
}

func TestAddChild(t *testing.T) {
	rootCtx := context.Background()
	rootNode := NewContextTree(rootCtx, "root")
	childNode := rootNode.AddChild("child")
	if childNode == nil {
		t.Fatal("Expected a non-nil child node")
	}
	if childNode.Identifier != "child" {
		t.Fatalf("Expected child identifier to be 'child', got '%s'", childNode.Identifier)
	}
	if len(rootNode.Children) != 1 {
		t.Fatal("Expected one child node")
	}
	if rootNode.Children[0] != childNode {
		t.Fatal("Expected the child node to be added to the root node's children")
	}
}

func TestCancelBranch(t *testing.T) {
	rootCtx := context.Background()
	rootNode := NewContextTree(rootCtx, "root")
	childNode := rootNode.AddChild("child")
	grandChildNode := childNode.AddChild("grandchild")

	rootNode.CancelBranch()

	select {
	case <-rootNode.Ctx.Done():
	case <-time.After(1 * time.Second):
		t.Fatal("Expected root context to be canceled")
	}

	select {
	case <-childNode.Ctx.Done():
	case <-time.After(1 * time.Second):
		t.Fatal("Expected child context to be canceled")
	}

	select {
	case <-grandChildNode.Ctx.Done():
	case <-time.After(1 * time.Second):
		t.Fatal("Expected grandchild context to be canceled")
	}
}
