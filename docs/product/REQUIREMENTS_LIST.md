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
*   [ ] **Browser Use Integration:**
    *   **Phase 1:** Support `npx` execution for `chrome-devtools-mcp` (Official DevTools protocol).
    *   **Phase 2:** Create a "Recipe" to install `Saik0s/mcp-browser-use` (requires Python/Playwright) and configure it to use Novapi's local API (`http://localhost:8080/v1`) as the LLM provider.

## 3. User Experience (P1)

*   [ ] **Knowledge Base (RAG) - *Updated*:**
    *   **Engine:** Integrate `philippgille/chromem-go` (Pure Go vector DB) for embedded, zero-dependency storage.
    *   **UI:** "Add to Context" button -> parses file -> embeds via Ollama -> stores in `chromem-go`.
*   [ ] **Featured Models - *Smart Recommendations*:**
    *   Auto-detect System RAM.
    *   Suggest `deepseek-r1:1.5b` for <8GB RAM, `deepseek-r1:8b` for 16GB, `deepseek-r1:32b` for 32GB+.
    *   One-click "Install & Run".
*   [ ] **Local Server Discovery:**
    *   Auto-detect running Ollama instance.
    *   Auto-detect LM Studio server (port 1234).
    *   Auto-detect Jan.ai (port 1337).
*   [ ] **Tray Icon:** Run in background (system tray) to serve API requests even when window is closed.

## 4. Documentation & onboarding (P2)

*   [ ] **MCP Tutorial:** "How to use the Filesystem MCP with Novapi".
*   [ ] **DeepSeek Guide:** "Running DeepSeek R1 locally with Novapi".
*   [ ] **Market Analysis:** Update `MARKET_ANALYSIS_2025.md` with Jan.ai v0.7.5 findings.

## 5. Engineering / Tech Debt

*   [ ] **CI/CD:** Ensure Linux builds work (GTK dependencies).
*   [ ] **Release:** Automate signing for macOS (notarization) - *Long term*.
