理解你的问题: doing
收集相关的信息: doing
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
│   ├── MongoDB 操作
│   └── SQL 操作
├── marsSql        // SQL 执行工具
├── marsType       // 数据结构扩展
│   ├── array 数组增强
│   ├── queue 队列实现  
│   └── set 集合实现
└── test           // 单元测试示例
```


## 2. 功能介绍

### 2.1 日志模块 (marsLog)
- **多日志级别**：支持 DEBUG/INFO/WARN/ERROR/FATAL 级别日志输出。
- **上下文关联日志**：支持在日志中打印上下文信息。
- **JSON 格式输出**：支持以 JSON 格式输出日志。

#### 示例代码
```go
import "github.com/marsli9945/mars-go/marsLog"

func main() {
    ctx := context.Background()
    marsLog.Logger().Debug("Debug")
    marsLog.Logger().Info("Info")
    marsLog.Logger().Warn("Warn")
    marsLog.Logger().Error("Error")
    marsLog.Logger().Fatal("Fatal")

    marsLog.Logger().DebugF("F: %s", "DebugF")
    marsLog.Logger().InfoFX(ctx, "FX: %s", "InfoFX")
}
```


### 2.2 数据库操作模块 (marsRepo)
封装了 MongoDB 和 SQL 数据库的操作，提供了便捷的增删改查接口。

#### MongoDB 操作
- `InsertOneContext`：插入单条数据。
- `FindAndPageContext`：分页查询。

#### SQL 操作
- `PrepareBatchContext`：批量执行 SQL 语句。

#### 示例代码
```go
import (
    "github.com/marsli9945/mars-go/marsRepo"
    "context"
)

func main() {
    repository := &marsRepo.MongoRepository{
        Database:   "testDB",
        Collection: "testCollection",
    }
    err := repository.InsertOneContext(context.Background(), map[string]string{"key": "value"})
    if err != nil {
        log.Fatal(err)
    }
}
```


### 2.3 HTTP 客户端工具 (marsHttp)
提供了简单的 HTTP 请求封装，支持 GET 和 POST 请求。

#### 示例代码
```go
import "github.com/marsli9945/mars-go/marsHttp"

func main() {
    response, err := marsHttp.Get("https://example.com")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(response)
}
```


### 2.4 数据结构扩展 (marsType)
- **Array**：数组增强功能，支持分割、查找等操作。
- **Queue**：队列实现。
- **Set**：集合实现。

#### 示例代码
```go
import (
    "github.com/marsli9945/mars-go/marsType"
    "fmt"
)

func main() {
    arr := marsType.ArrayInitForList([]int{1, 2, 3, 4, 5})
    fmt.Println(arr.Contains(3)) // true
    fmt.Println(arr.Join(","))  // "1,2,3,4,5"

    queue := make(marsType.Queue[int], 0, 5)
    queue.Push(1)
    queue.Push(2)
    fmt.Println(queue.Pop()) // 1

    set := marsType.NewSet[string]()
    set.AddAll([]string{"a", "b", "c"})
    fmt.Println(set.Contains("b")) // true
}
```


### 2.5 Gin 框架增强组件 (marsGin)
- **MiddlewareCors**：跨域设置。
- **MiddlewareErr**：错误处理中间件。
- **MiddlewareJWT**：JWT 鉴权中间件。

#### 示例代码
```go
import (
    "github.com/marsli9945/mars-go/marsGin"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.New()
    r.Use(gin.Logger(), marsGin.MiddlewareErr(), marsGin.MiddlewareCors())

    r.GET("/hello", marsGin.TransformHandle(func(g *marsGin.Gin) {
        g.Ok()
    }))

    r.Run(":8080")
}
```


### 2.6 上下文增强工具 (marsContext)
提供了上下文树结构，支持子上下文的创建和取消。

#### 示例代码
```go
import (
    "github.com/marsli9945/mars-go/marsContext"
    "context"
)

func main() {
    rootCtx := context.Background()
    rootNode := marsContext.NewContextTree(rootCtx, "root")
    childNode := rootNode.AddChild("child")
    rootNode.CancelBranch()
}
```


## 3. 总结
Mars Go 工具库通过模块化设计，为开发者提供了丰富的功能模块，简化了日常开发中的常见任务。无论是日志处理、数据库操作还是 HTTP 请求，Mars Go 都能帮助开发者快速构建高效的应用程序。