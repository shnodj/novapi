# Product Research Log - March 2025

**Date:** 2025-03-01
**Topic:** DeepSeek R1 UI, MCP Ecosystem Expansion, and Competitor Benchmarking
**Author:** Jules (Product Engineer)

## 1. Market & Competitor Analysis

### 1.1 DeepSeek R1 & "Reasoning" Models
The release of DeepSeek R1 has shifted user expectations. Users now expect to see the "Thinking Process" separated from the final answer.

*   **Jan.ai:** Has optimized their UI to handle DeepSeek R1's `<think>` tags. They collapse this by default but allow users to expand it to inspect the chain of thought.
*   **LM Studio:** Version 0.3.29 introduced explicit "Reasoning support" in their API and UI, allowing for parsing of reasoning output and controlling effort levels.
*   **Novapi Implication:** We **must** implement a UI parser for `<think>` tags (or the `reasoning_content` API field) immediately. Showing raw `<think>` tags is now considered a poor UX.

### 1.2 MCP Host Maturity
*   **LM Studio (v0.3.17+):** Full MCP Host support using `mcp.json`. They have added a "Deep Link" feature (`lmstudio://`) to install MCP servers with one click.
*   **Cursor:** Remains the gold standard for `mcp.json` configuration.
*   **Novapi Implication:** We need to support `mcp.json` natively to allow users to copy-paste configs from Cursor/LM Studio. We should also investigate handling `novapi://` deep links for easier installation.

## 2. MCP Ecosystem Scouting

The MCP server ecosystem has exploded. Based on the `modelcontextprotocol/servers` registry and community trends, here are the key categories for Novapi to support or bundle.

### 2.1 "Must Have" Bundles (Core)
These should be recommended immediately upon installation.
1.  **Filesystem:** The absolute baseline.
2.  **Git/GitHub:** Essential for developers.
3.  **Brave Search / Tavily:** For "Lite RAG" (Web Search).

### 2.2 The "Agentic" Frontier (High Value)
*   **Browser Use (`mcp-browser-use`):** This is the hottest agentic capability. It uses Playwright (often via Docker or local python) to let the LLM control a browser.
    *   *Implementation Note:* Novapi needs to decide if we bundle a Python environment to run this, or rely on a Docker container. `co-browser/browser-use-mcp-server` is a leading implementation.
*   **Sequential Thinking:** A server that allows the model to "think" in steps, even if it's not a reasoning model like R1.

### 2.3 Enterprise/Utility
*   **PostgreSQL / SQLite:** For interacting with local data.
*   **Google Drive / Slack:** Common enterprise request.

## 3. Strategic Recommendations

### 3.1 Feature: "Thinking" UI (P0)
*   **Requirement:** Parse streaming responses. If `<think>` or `reasoning_content` is detected, render it in a collapsible "Thinking..." accordion component.
*   **Goal:** Parity with Jan.ai and LM Studio for DeepSeek R1.

### 3.2 Feature: MCP Marketplace (P1)
*   **Requirement:** A UI to browse and install MCP servers.
*   **Source:** We can pull metadata from the `modelcontextprotocol/servers` repository `README.md` or their `registry.json` if available.
*   **One-Click Install:** For Node-based servers (`npx ...`), we can automate the installation. For Python (`uvx ...`), we need to ensure `uv` is present.

### 3.3 Feature: Deep Link Support (P2)
*   **Requirement:** Register `novapi://` protocol handler.
*   **Use Case:** Clicking "Install in Novapi" from a web directory adds the server to `mcp.json`.

## 4. Technical Feasibility Notes
*   **Browser Use:** Running `browser-use` locally is heavy. It requires Playwright + Browsers. We might want to start by supporting the *connection* to an existing server, rather than bundling the entire browser stack in the installer.
*   **MCP Config:** We should watch the `~/.config/novapi/mcp.json` file (or equivalent) for changes to support external editing.
