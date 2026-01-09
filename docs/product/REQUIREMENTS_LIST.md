# Product Requirements List

**Status:** Active
**Last Updated:** 2025-05-15

## 1. Core Platform (P0)

*   [ ] **De-Redis Refactor:** Complete removal of Redis dependency. Ensure `common.RedisEnabled` is false and `go-cache` is used for token/user caching.
*   [ ] **Single User Mode:** Auto-generate/inject Root Token on startup. Bypass login screen if configured for local-only mode.
*   [ ] **SQLite Enforcement:** Ensure SQLite is the only database option for desktop builds.
*   [ ] **Deep Link Support:** Register `novapi://` protocol handler to support web-to-local actions (e.g., `novapi://mcp/install?url=...`).

## 2. Model Context Protocol (MCP) (P0)

*   [ ] **MCP Host Core:** Implement a client in Go that can spawn and communicate with MCP Servers via Stdio and **SSE/HTTP**.
*   [ ] **MCP Marketplace:** Integrate with **Smithery.ai** or **Glama.ai** to list and install community servers directly from the UI.
*   [ ] **Config Compatibility:** Support loading/saving `mcp.json` (compatible with Cursor/LM Studio).
*   [ ] **Browser Use Integration:** Provide a 1-click setup for `mcp-browser-use` (Daemon mode on port 8383) to enable stable, long-running web agents.
*   [ ] **Ollama Bridge:** When an MCP tool is used, translate the request to an Ollama "tool call" if the backend model supports it.

## 3. User Experience (P1)

*   [ ] **Reasoning UI (DeepSeek R1):**
    *   Parse `<think>` tags from model output.
    *   Display "Thinking Process" in a collapsible, distinct UI element (gray text, accordion style).
*   [ ] **Knowledge Base (Lite RAG):**
    *   Local file indexing (PDF, MD, TXT) using the "Filesystem MCP" or a dedicated "Search" MCP.
    *   "Add to Context" button in Chat UI.
*   [ ] **Featured Models:** Home screen widget showing "DeepSeek R1", "Llama 3", etc.
    *   One-click "Install" (triggers `ollama pull`).
    *   One-click "Run".
*   [ ] **Local Server Discovery:**
    *   Auto-detect running Ollama instance.
    *   Auto-detect LM Studio server (port 1234).
*   [ ] **Tray Icon:** Run in background (system tray) to serve API requests even when window is closed.

## 4. Documentation & onboarding (P2)

*   [ ] **MCP Tutorial:** "How to use the Filesystem MCP with Novapi".
*   [ ] **DeepSeek Guide:** "Running DeepSeek R1 locally with Novapi".
*   [ ] **Browser Agent Guide:** "Automating the Web with Novapi & Browser Use".

## 5. Engineering / Tech Debt

*   [ ] **CI/CD:** Ensure Linux builds work (GTK dependencies).
*   [ ] **Release:** Automate signing for macOS (notarization) - *Long term*.
