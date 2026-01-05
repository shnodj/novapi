# Product Requirements List

**Status:** Active
**Last Updated:** 2025-03-01

## 1. Core Platform (P0)

*   [ ] **De-Redis Refactor:** Complete removal of Redis dependency. Ensure `common.RedisEnabled` is false and `go-cache` is used for token/user caching.
*   [ ] **Single User Mode:** Auto-generate/inject Root Token on startup. Bypass login screen if configured for local-only mode.
*   [ ] **SQLite Enforcement:** Ensure SQLite is the only database option for desktop builds.

## 2. Model Context Protocol (MCP) (P0)

*   [ ] **MCP Host Core (Stdio):** Implement client for spawning/communicating with local Stdio MCP Servers.
*   [ ] **MCP Host Core (HTTP/SSE):** Support connecting to long-running MCP agents (like `browser-use`) via HTTP/SSE. *[New - Critical for Agents]*
*   [ ] **Config Compatibility:** Support loading/saving `mcp.json` (compatible with Cursor/LM Studio).
*   [ ] **Tool Mapping:** Translate MCP "Resources" and "Tools" into OpenAI-compatible Function Definitions.
*   [ ] **Ollama Bridge:** Translate MCP tool calls to Ollama native tool calls.

## 3. High-Value Integrations (P1)

*   [ ] **Browser Agent Bundle:**
    *   One-click setup for `Saik0s/mcp-browser-use`.
    *   Manage the background daemon process.
    *   UI for "Deep Research" (Plan -> Search -> Synthesize).
*   [ ] **MCP Marketplace UI:**
    *   Browse/Install servers from `smithery.ai` or `modelcontextprotocol/servers`.
    *   Visual management of installed servers.

## 4. User Experience (P1)

*   [ ] **Reasoning UI (DeepSeek R1):**
    *   Parse `<think>` tags in streaming responses.
    *   Render as collapsible "Thinking Process" section (default collapsed).
*   [ ] **Knowledge Base (Lite RAG):**
    *   Local file indexing using "Search Files" MCP server (vs embedded vector DB).
    *   "Add to Context" drag-and-drop.
*   [ ] **Local Server Discovery:**
    *   Auto-detect running Ollama, LM Studio, Jan.ai instances.

## 5. Documentation & Onboarding (P2)

*   [ ] **MCP Tutorial:** "How to use the Filesystem MCP with Novapi".
*   [ ] **DeepSeek Guide:** "Running DeepSeek R1 locally with Novapi".

## 6. Engineering / Tech Debt

*   [ ] **CI/CD:** Ensure Linux builds work (GTK dependencies).
*   [ ] **Release:** Automate signing for macOS (notarization) - *Long term*.
