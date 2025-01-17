package marsContext

import (
	"context"
	"github.com/marsli9945/mars-go/marsLog"
	"sync"
	"time"
)

func MonitorContext(ctx context.Context, name string, interval time.Duration) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		lastLogTime := time.Now()
		for {
			select {
			case <-ctx.Done():
				marsLog.Logger().WarnF("Context %s 已取消，原因：%v", name, ctx.Err())
				// 添加清理逻辑
				// cleanup()
				return
			default:
				if time.Since(lastLogTime) >= interval {
					marsLog.Logger().InfoF("Context %s 仍在运行", name)
					lastLogTime = time.Now()
				}
				time.Sleep(time.Second)
			}
		}
	}()
	wg.Wait()
}
