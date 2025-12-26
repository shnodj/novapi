# Novapi Architecture

Novapi is a desktop-native AI Gateway designed to be a lightweight, single-user alternative to One-API. It leverages **Wails** to combine a robust Go backend with a modern React frontend, running as a single executable.

## 1. High-Level Overview

```mermaid
graph TD
    User[User] --> GUI[Wails Frontend (React)]
    User --> API[External Apps (LangChain, Chatbox)]

    subgraph "Novapi Desktop App"
        GUI -- IPC (Wails Runtime) --> GoBackend[Go Backend]
        API -- HTTP (localhost:3000) --> GinServer[Gin HTTP Server]

        GoBackend -- "In-Process" --> GinServer

        GinServer --> Controller[Controllers]
        GoBackend --> Controller

        Controller --> Service[Services (Token, Channel)]
        Service --> Cache[Memory Cache]
        Service --> DB[(SQLite Database)]

        Service --> Upstream[Upstream APIs (OpenAI, Claude)]
        Service --> Local[Local Processes (Ollama, llama.cpp)]
    end
```

## 2. Core Components

### 2.1 The "Dual-Process" Model (Conceptual)

Although Novapi runs as a single binary process, functionally it operates in two modes simultaneously:

1.  **GUI Mode (Wails)**:
    *   **Role**: Manages the user interface, configuration, and dashboard.
    *   **Communication**: Uses Wails' IPC (Inter-Process Communication) to call Go methods from JavaScript.
    *   **Security**: Runs with the privileges of the desktop user.

2.  **Server Mode (Gin)**:
    *   **Role**: Provides the OpenAI-compatible API endpoints (e.g., `/v1/chat/completions`).
    *   **Communication**: Listens on a local TCP port (default: 3000).
    *   **Integration**: External tools connect to this port. It runs in a separate Goroutine managed by the main application.

### 2.2 Data Persistence

*   **Database**: **SQLite** (via GORM).
    *   chosen for its zero-configuration, single-file nature, perfect for desktop apps.
    *   File location: `~/.config/novapi/novapi.db` (or OS equivalent).
*   **Caching**: **In-Memory** (`go-cache`).
    *   **Decision**: We explicitly removed Redis to reduce dependencies.
    *   **Usage**: Stores short-lived tokens, rate limiting counters, and temporary session data.

### 2.3 Channel System

The channel system is the heart of Novapi, routing requests to various LLM providers.

*   **Upstream Channels**:
    *   Standard HTTP proxies to providers like OpenAI, Azure, Anthropic.
    *   Stored in DB with API Key and Base URL.
*   **Local Channels**:
    *   **Concept**: Novapi can manage or connect to local inference engines.
    *   **Implementation**: Can execute a shell command (e.g., `ollama run llama3`) or connect to a local port (e.g., `localhost:11434`).
    *   **Discovery**: Auto-scans common ports for active local services.

## 3. Key Design Decisions

### 3.1 "De-Redis" Strategy
One-API relies heavily on Redis for distributed locks and shared cache. Since Novapi is single-instance/desktop:
*   **Distributed Locks** -> **Go Mutexes**: We don't need distributed coordination.
*   **Redis Cache** -> **Memory Cache**: RAM is cheap and fast enough for a single user's state.

### 3.2 Single-User / Admin Mode
Novapi assumes a single owner.
*   **Auth**: No login screen. The app opens directly to the dashboard.
*   **API Auth**: "Master Key" generated on install. Used by external apps to connect.
*   **Multi-tenancy**: Removed. No "User" and "Group" complexity required for the core logic.

## 4. Technology Stack

*   **Frontend**: React, Vite, TailwindCSS (planned).
*   **Backend**: Go 1.21+.
*   **Framework**: Wails v2.
*   **Database**: SQLite.
*   **HTTP Server**: Gin Web Framework.
