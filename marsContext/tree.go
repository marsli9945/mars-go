package marsContext

import (
	"context"
	"sync"
)

type ContextNode struct {
	Ctx        context.Context
	Cancel     context.CancelFunc
	Children   []*ContextNode
	Identifier string
	mu         sync.Mutex // 添加互斥锁以保证并发安全
}

func NewContextTree(root context.Context, identifier string) *ContextNode {
	ctx, cancel := context.WithCancel(root)
	return &ContextNode{
		Ctx:        ctx,
		Cancel:     cancel,
		Children:   make([]*ContextNode, 0),
		Identifier: identifier,
	}
}

func (n *ContextNode) AddChild(identifier string) *ContextNode {
	n.mu.Lock()
	defer n.mu.Unlock()

	child := NewContextTree(n.Ctx, identifier)
	n.Children = append(n.Children, child)
	return child
}

func (n *ContextNode) CancelBranch() {
	defer func() {
		if r := recover(); r != nil {
			// 处理 panic，确保所有子节点都能被取消
			n.CancelBranch()
		}
	}()

	n.Cancel()

	n.mu.Lock()
	defer n.mu.Unlock()

	if len(n.Children) == 0 {
		return
	}

	for _, child := range n.Children {
		child.CancelBranch()
	}
}
