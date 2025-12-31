# Market Scan & Ecosystem Update (Q4 2025)

**Date:** 2025-12-31
**Status:** DRAFT
**Context:** Year-End Update to Strategy

## 1. Jan.ai Updates (v0.7.5+)
Jan.ai has moved aggressively into the MCP space.
*   **Nov 2025 (v0.7.3):** Introduced **"Jan Browser MCP"** for powerful browser automation.
*   **Aug 2025 (v0.6.9):** "Stable MCP" support (moved out of experimental).
*   **Key Insight:** They are positioning the Browser as a first-class citizen via MCP. This validates Novapi's strategy to bundle "Browser Use".
*   **Projects (v0.7.0):** Added "Projects" for workspace management, similar to VS Code workspaces.

## 2. Browser Use MCP
The ecosystem has consolidated around a few key implementations for browser automation via MCP.
*   **`mcp-browser-use` (Saik0s):** Popular option.
    *   **Agentic:** Uses Playwright + Python.
    *   **Transport:** Supports HTTP transport (`http://localhost:8383/mcp`) which is more reliable than stdio for long-running browser tasks.
    *   **Config:** `{"mcpServers": {"browser-use": {"type": "streamable-http", "url": "..."}}}`.
*   **`browser-mcp` (djyde):** Extension-based approach.
    *   **Copilot style:** Interacts with the *current* browser window (Chrome Extension) rather than a headless instance.
    *   **Usage:** Good for "Summarize this page" or "Fill this form", less for "Go research X".

**Recommendation for Novapi:**
*   **Agent Mode:** Use `Saik0s/mcp-browser-use` (via Docker or Python venv) for autonomous tasks.
*   **Copilot Mode:** Investigate bundling an extension if we want to control the user's *actual* browser, but "Agent Mode" is the primary differentiator for a background runner like Novapi.

## 3. Trending MCP Servers (2025 Winners)
The "Must Have" list for a Desktop Host:
1.  **Filesystem (Core):** Essential for coding/file manipulation.
2.  **GitHub / GitLab:** For PR summaries and issue tracking.
3.  **Chrome DevTools / Browser Use:** For web research.
4.  **Postgres / Supabase:** For database interactions.
5.  **Context7:** For up-to-date documentation (popular in Reddit threads).

## 4. DeepSeek R1 Status
*   **Availability:** DeepSeek R1 remains a top local model.
*   **Strategy:** Novapi should offer "Presets" in the `new-api` channel configuration that map to standard Ollama tags (e.g., `deepseek-r1:7b`, `deepseek-r1:14b`).

## 5. Strategic Pivots for Novapi
*   **HTTP Transport for MCP:** Stdio is flaky for heavy tools like Browser Use. Novapi must support **SSE (Server-Sent Events) / HTTP** transport for MCP clients, not just Stdio.
*   **"Projects" Concept:** Users want to scope their MCP tools to specific folders/projects (copying Jan.ai's v0.7.0 feature).
