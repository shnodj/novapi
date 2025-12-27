# Novapi 架构重构方案：Host-Hosted Pattern

## 核心理念
将 `new-api` 视为**核心引擎（Engine）而非简单的代码库，采用“宿主-寄生”模式（Host-Hosted Pattern）**。
桌面程序（Wails App）作为**宿主**，将 `new-api` 的核心逻辑（Router, Relay, Model）作为**依赖包**引入。

## 1. 核心架构设计

```mermaid
classDiagram
    namespace DesktopHost {
        class WailsApp {
            +ManageWindow()
            +ManageLocalProcesses()
        }
        class ConfigInjector {
            +OverrideEnv()
            +InitSQLite()
        }
        class LocalProcessManager {
            +StartOllama()
            +StartLlamaCpp()
            +RegisterToNewAPI()
        }
    }

    namespace NewAPIEngine {
        class Router {
            +ServeHTTP()
        }
        class RelayService {
            +ForwardRequest()
        }
        class ChannelDB {
            +GetChannel()
        }
    }

    WailsApp --> ConfigInjector : 1. 初始化环境
    ConfigInjector ..> NewAPIEngine : 注入配置(SQLite路径/关闭Redis)
    WailsApp --> NewAPIEngine : 2. 启动 Gin Server (Goroutine)
    WailsApp --> LocalProcessManager : 3. 管理本地模型
    LocalProcessManager ..> ChannelDB : 4. 动态注册/更新本地渠道
```

## 2. 技术实现：引用 New-API

### A. go.mod 依赖管理

在 Wails 项目中引入 `new-api`。

```bash
go get github.com/Calcium-Ion/new-api
```
*(注：根据实际情况可能需要 replace 指向特定 fork)*

### B. 启动引导 (Bootstrap) 与 配置注入

在 `backend/bootstrap.go` 中：

1.  **强制配置注入**:
    *   `common.SQLitePath` 指向用户数据目录。
    *   `common.RedisEnabled = false`。
    *   `common.MemoryCacheEnabled = true`。
2.  **初始化 DB**: 调用 `common.InitSQLite()` 和 `model.InitDB()`。
3.  **Root Token**: 确保数据库存在一个用于前端调用的 Root Token。
4.  **启动 Gin Router**:
    *   `router.SetRouter`。
    *   在 Goroutine 中运行 Server。

### 3. 本地渠道兼容 (Manager Logic)

不修改表结构，将本地进程抽象为 HTTP 接口。

1.  **进程管理**: Wails 后端启动/停止本地模型（Ollama/llama-server），分配端口。
2.  **动态注册**:
    *   直接操作 `new-api` 的 `model.DB`。
    *   插入 `Channel` 记录 (Type=Ollama, BaseURL=localhost:port, Key=sk-ignore)。
    *   刷新 `common.MemoryCache`。

### 4. 功能裁剪

1.  **Router 层**: 屏蔽注册、支付、日志等不需要的接口。
2.  **前端**:
    *   完全重写，不使用 `new-api/web`。
    *   只保留设置、渠道管理、Token 显示。

### 5. 版本追踪

*   使用 `go mod replace` 可以在需要 fix bug 时指向自己的 fork。
*   主要依赖 `Relay` 逻辑，UI 和 支付逻辑被抛弃，兼容性风险低。
