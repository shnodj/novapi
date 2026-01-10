# Product Requirements List

**Status:** Active
**Last Updated:** 2026-01-10

## 1. Core Platform (P0)

*   [ ] **De-Redis Refactor:** Complete removal of Redis dependency. Ensure `common.RedisEnabled` is false and `go-cache` is used for token/user caching.
*   [ ] **Single User Mode:** Auto-generate/inject Root Token on startup. Bypass login screen if configured for local-only mode.
*   [ ] **SQLite Enforcement:** Ensure SQLite is the only database option for desktop builds.

## 2. Model Context Protocol (MCP) (P0)

*   [ ] **MCP Host Core:** Implement a client in Go that can spawn (Stdio) and connect (HTTP/SSE) to MCP Servers.
    *   *Update:* Must support HTTP/SSE for long-running agents like `browser-use`.
*   [ ] **Config Compatibility:** Support loading/saving `mcp.json` (compatible with Cursor/LM Studio).
*   [ ] **Deep Link Handler:** Implement `novapi://mcp/install?url=...` for one-click server installation (Parity with LM Studio).
*   [ ] **Tool Mapping:** Translate MCP "Resources" and "Tools" into OpenAI-compatible Function Definitions.
*   [ ] **Browser Use Integration:** Bundle `Saik0s/mcp-browser-use` (via Docker or local binary) and connect via HTTP.

## 3. User Experience (P1)

*   [ ] **DeepSeek R1 UI:** Parse `<think>` tags in streaming responses. Visualize "Thinking Process" in a collapsible, distinct UI element.
*   [ ] **Knowledge Base (Lite RAG):**
    *   Use "Search Files" MCP instead of complex vector DBs.
    *   "Add to Context" button in Chat UI.
*   [ ] **Featured Models:** Home screen widget.
    *   One-click "Install" (triggers `ollama pull`).
*   [ ] **Local Server Discovery:**
    *   Auto-detect running Ollama, LM Studio, Jan.ai.

## 4. Local Channels & Integrations (P2)

*   [ ] **GitHub Copilot Bridge:** Explore integrating `copilot-api` to allow users to use their Copilot subscription as a local OpenAI endpoint.
*   [ ] **Tray Icon:** Background service mode.

## 5. Documentation & onboarding (P2)

*   [ ] **MCP Tutorial:** "How to use the Filesystem MCP with Novapi".
*   [ ] **DeepSeek Guide:** "Running DeepSeek R1 locally with Novapi".

## 6. Engineering / Tech Debt

*   [ ] **CI/CD:** Ensure Linux builds work (GTK dependencies).
*   [ ] **Release:** Automate signing for macOS.
