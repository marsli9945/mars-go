package marsContext

import (
	"context"
	"errors"
	"github.com/marsli9945/mars-go/marsLog"
	"sync"
)

type Context struct {
	context.Context
	customValue string
	mu          sync.RWMutex
}

func (c *Context) Value(key any) any {
	if keyStr, ok := key.(string); ok && keyStr == "customKey" {
		c.mu.RLock()
		defer c.mu.RUnlock()
		return c.customValue
	}
	return c.Context.Value(key)
}

func WithCustomValue(parent context.Context, value string) (child context.Context, e error) {
	if parent == nil {
		marsLog.Logger().Error("parent context cannot be nil")
		return nil, errors.New("parent context cannot be nil")
	}
	return &Context{
		Context:     parent,
		customValue: value,
	}, nil
}
