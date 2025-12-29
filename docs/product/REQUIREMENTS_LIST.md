# Product Requirements List

**Status:** Active
**Last Updated:** 2025-02-24

## 1. Core Platform (P0)

*   [ ] **De-Redis Refactor:** Complete removal of Redis dependency. Ensure `common.RedisEnabled` is false and `go-cache` is used for token/user caching.
*   [ ] **Single User Mode:** Auto-generate/inject Root Token on startup. Bypass login screen if configured for local-only mode.
*   [ ] **SQLite Enforcement:** Ensure SQLite is the only database option for desktop builds.

## 2. Model Context Protocol (MCP) (P0)

*   [ ] **MCP Host Core:** Implement a client in Go that can spawn and communicate with MCP Servers via Stdio.
*   [ ] **Config Compatibility:** Support loading/saving `mcp.json` (compatible with Cursor/LM Studio).
*   [ ] **Tool Mapping:** Translate MCP "Resources" and "Tools" into OpenAI-compatible Function Definitions for the chat API.
*   [ ] **Ollama Bridge:** When an MCP tool is used, translate the request to an Ollama "tool call" if the backend model supports it.
*   [ ] **Browser Use Integration:** Bundle or provide 1-click setup for the "Browser Use" MCP server (web automation agent).

## 3. User Experience (P1)

*   [ ] **Knowledge Base (RAG):**
    *   Local file indexing (PDF, MD, TXT).
    *   "Add to Context" button in Chat UI.
*   [ ] **Featured Models:** Home screen widget showing "DeepSeek R1", "Llama 3", etc.
    *   One-click "Install" (triggers `ollama pull`).
    *   One-click "Run".
*   [ ] **Local Server Discovery:**
    *   Auto-detect running Ollama instance.
    *   Auto-detect LM Studio server (port 1234).
    *   Auto-detect Jan.ai / Cherry Studio APIs.
*   [ ] **Tray Icon:** Run in background (system tray) to serve API requests even when window is closed.

## 4. Documentation & onboarding (P2)

*   [ ] **MCP Tutorial:** "How to use the Filesystem MCP with Novapi".
*   [ ] **DeepSeek Guide:** "Running DeepSeek R1 locally with Novapi".

## 5. Engineering / Tech Debt

*   [ ] **CI/CD:** Ensure Linux builds work (GTK dependencies).
*   [ ] **Release:** Automate signing for macOS (notarization) - *Long term*.
