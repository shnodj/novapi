# Product Requirements List

**Status:** Active
**Last Updated:** 2025-03-14

## 1. Core Platform (P0)

*   [ ] **De-Redis Refactor:** Complete removal of Redis dependency. Ensure `common.RedisEnabled` is false and `go-cache` is used for token/user caching.
*   [ ] **Single User Mode:** Auto-generate/inject Root Token on startup. Bypass login screen if configured for local-only mode.
*   [ ] **SQLite Enforcement:** Ensure SQLite is the only database option for desktop builds.

## 2. Model Context Protocol (MCP) (P0)

*   [ ] **MCP Host Core:** Implement a client in Go that can spawn and communicate with MCP Servers via Stdio.
*   [ ] **Config Compatibility:** Support loading/saving `mcp.json` (compatible with Cursor/LM Studio).
*   [ ] **MCP Permission Manager:** UI to Approve/Deny tool execution requests (Security).
*   [ ] **Essential MCP Servers (Bundled/1-Click):**
    *   **GitHub:** `github/github-mcp-server` (Codebase management).
    *   **Browser:** `microsoft/playwright-mcp` OR `mcp-browser-use` (Web automation).
*   [ ] **Tool Mapping:** Translate MCP "Resources" and "Tools" into OpenAI-compatible Function Definitions.

## 3. User Experience (P1)

*   [ ] **DeepSeek R1 Support (Critical):**
    *   **Thinking UI:** Parse `<think>...</think>` tags and render as a collapsible "Thinking Process" accordion.
    *   **One-Click Run:** "Featured Models" widget specifically for DeepSeek R1 (distilled versions).
*   [ ] **Knowledge Base (Lite RAG):**
    *   Local file indexing (PDF, MD, TXT).
    *   "Add to Context" button in Chat UI.
*   [ ] **Featured Models Widget:**
    *   Curated list (DeepSeek, Llama 3, Qwen).
    *   Visual status indicators (Downloading, Running).
*   [ ] **Tray Icon:** Run in background.

## 4. Documentation & onboarding (P2)

*   [ ] **MCP Tutorial:** "How to use the GitHub MCP with Novapi".
*   [ ] **DeepSeek Guide:** "Running DeepSeek R1 locally with Novapi".

## 5. Engineering / Tech Debt

*   [ ] **CI/CD:** Ensure Linux builds work (GTK dependencies).
*   [ ] **Release:** Automate signing for macOS (notarization) - *Long term*.
