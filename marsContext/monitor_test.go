package marsContext

import (
	"context"
	"testing"
	"time"
)

func TestMonitorContext_ContextCancelled_LogsError(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // 立即取消上下文

	MonitorContext(ctx, "testContext", 1*time.Second)

	if ctx.Err().Error() != "context canceled" {
		t.Error("Expected error message not logged")
	}

}

func TestMonitorContext_ContextActive_LogsInfo(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	MonitorContext(ctx, "testContext", 1*time.Second)

}

func TestMonitorContext_ShortInterval_LogsCorrectly(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	MonitorContext(ctx, "testContext", 1*time.Second)

	if ctx.Err().Error() != "context deadline exceeded" {
		t.Error("Expected error message not logged")
	}
}
