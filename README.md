# Mars Go 工具库使用说明

## 1. 项目简介
Mars Go 是基于 Go 语言开发的工具类集合，采用模块化设计，提供以下核心功能模块：
```text
├── marsContext    // 上下文增强工具
├── marsGin        // Gin 框架增强组件
├── marsHttp       // HTTP 客户端工具
├── marsJwt        // JWT 鉴权工具  
├── marsLog        // 日志处理模块
├── marsRepo       // 数据库操作抽象
│   ├── MongoDB 操作
│   └── SQL 操作
├── marsSql        // SQL 执行工具
├── marsType       // 数据结构扩展
│   ├── array 数组增强
│   ├── queue 队列实现  
│   └── set 集合实现
└── test           // 单元测试示例

```
## 2. 安装与配置
### 2.1 安装依赖
确保你的 Go 环境已经正确安装，并且可以正常工作。使用以下命令安装 mars-go：
```shell
GOPROXY=direct go get github.com/marsli9945/mars-go@v0.0.5
```
### 2.2 配置环境变量（可选）
部分模块可能需要配置环境变量，例如日志模块中的调试模式：
- DEBUG_ENABLE=true：启用调试日志
## 3. 模块介绍
### 3.1 日志模块 (marsLog)
功能特性：
- 多日志级别：DEBUG/INFO/WARN/ERROR/FATAL
- 上下文关联日志
- JSON 格式输出
#### 示例：
```text
import "github.com/marsli9945/mars-go/marsLog"

func main() {
   ctx := context.Background()
   
   Logger().Debug("Debug")
   Logger().Info("Info")
   Logger().Warn("Warn")
   Logger().Error("Error")
   Logger().Fatal("Debug")
   
   Logger().DebugF("F: %s", "DebugF")
   Logger().InfoF("F: %s", "InfoF")
   Logger().WarnF("F: %s", "WarnF")
   Logger().ErrorF("F: %s", "ErrorF")
   Logger().FatalF("F: %s", "FatalF")
   
   Logger().DebugFX(ctx, "FX: %s", "DebugFX")
   Logger().InfoFX(ctx, "FX: %s", "InfoFX")
   Logger().WarnFX(ctx, "FX: %s", "WarnFX")
   Logger().ErrorFX(ctx, "FX: %s", "ErrorFX")
   Logger().FatalFX(ctx, "FX: %s", "FatalFX")
}
```
#### 输出结果：
```text
[MARS] 2025-02-17 14:14:15 | DEBUG | log_test.go:21  Debug
[MARS] 2025-02-17 14:14:15 | INFO  | log_test.go:22  Info
[MARS] 2025-02-17 14:14:15 | WARN  | log_test.go:23  Warn
[MARS] 2025-02-17 14:14:15 | ERROR | log_test.go:24  Error
[MARS] 2025-02-17 14:14:15 | FATAL | log_test.go:25  Debug
[MARS] 2025-02-17 14:14:15 | DEBUG | log_test.go:27  F: DebugF
[MARS] 2025-02-17 14:14:15 | INFO  | log_test.go:28  F: InfoF
[MARS] 2025-02-17 14:14:15 | WARN  | log_test.go:29  F: WarnF
[MARS] 2025-02-17 14:14:15 | ERROR | log_test.go:30  F: ErrorF
[MARS] 2025-02-17 14:14:15 | FATAL | log_test.go:31  F: FatalF
[MARS] 2025-02-17 14:14:15 | DEBUG | log_test.go:33  context.Background
[MARS] 2025-02-17 14:14:15 | DEBUG | log_test.go:33  FX: DebugFX
[MARS] 2025-02-17 14:14:15 | INFO  | log_test.go:34  context.Background
[MARS] 2025-02-17 14:14:15 | INFO  | log_test.go:34  FX: InfoFX
[MARS] 2025-02-17 14:14:15 | WARN  | log_test.go:35  context.Background
[MARS] 2025-02-17 14:14:15 | WARN  | log_test.go:35  FX: WarnFX
[MARS] 2025-02-17 14:14:15 | ERROR | log_test.go:36  context.Background
[MARS] 2025-02-17 14:14:15 | ERROR | log_test.go:36  FX: ErrorFX
[MARS] 2025-02-17 14:14:15 | FATAL | log_test.go:37  context.Background
[MARS] 2025-02-17 14:14:15 | FATAL | log_test.go:37  FX: FatalFX
[MARS] 2025-02-17 14:14:15 | JSON  | log_test.go:43  {"cookie_code":"123","feed_id":"456"}
[MARS] 2025-02-17 14:14:15 | JSON  | log_test.go:44  
{
 "cookie_code": "123",
 "feed_id": "456"
}
```
### 3.2 数据库操作模块 (marsRepo)
封装了 MongoDB 和 SQL 数据库的操作，提供了便捷的增删改查接口。
#### MongoDB 操作：
InsertOneContext(ctx context.Context, model any)：插入单条数据
FindAndPageContext(ctx context.Context, results any, filter any, skip int64, limit int64)：分页查询
#### SQL 操作：
PrepareBatchContext(ctx context.Context, query string, rows [][]any)：批量执行 SQL 语句
示例：