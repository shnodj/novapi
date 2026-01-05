# Product Research Log - March 2025

**Date:** 2025-03-01
**Topic:** Browser Agents, Deep Research, and MCP Marketplaces
**Researcher:** Jules (Product Manager Agent)

## 1. Key Findings

### 1.1 Browser Automation (The "Killer App" for MCP)
*   **Leading Solution:** `Saik0s/mcp-browser-use` (860+ stars) is the most mature implementation.
*   **Architecture Shift:** It uses a **long-running HTTP Daemon** instead of just Stdio.
    *   *Why?* Browser tasks take 30-120 seconds. Standard Stdio transports often time out or are hard to debug.
    *   *Novapi Implication:* Our MCP Host implementation must support **HTTP/SSE Transport** (Server-Sent Events) to handle these long-running agents reliably, not just Stdio.
*   **Features:**
    *   "Deep Research" mode: Plan → Search → Synthesize.
    *   "Skills": Record and replay API calls (faster than UI clicking).
    *   Web Dashboard: It comes with its own UI at `localhost:8383`.

### 1.2 Jan.ai Updates (v0.7.x)
*   **Browser MCP:** They shipped a native "Jan Browser MCP" in v0.7.3.
*   **Proactive Mode:** Anticipates user needs by capturing context.
*   **Status:** They are moving fast. Novapi needs to match the "One-click Browser Agent" capability to remain competitive.

### 1.3 LM Studio Updates (v0.3.36)
*   **Stability:** Fully supports `mcp.json` config.
*   **Remote MCP:** Opt-in support for remote servers, validating the HTTP transport requirement.

### 1.4 DeepSeek R1 & Reasoning UX
*   **Standard Pattern:** No specific "Thinking UI" standard has emerged in open source beyond the `<think>` tag usage.
*   **Observation:** Most tools just stream the output.
*   **Requirement:** Novapi should parse `<think>` tags and render them in a collapsible "Thinking Process" block to clean up the chat interface.

### 1.5 MCP Ecosystem & Marketplaces
*   **Emerging Registries:**
    *   `smithery.ai`: Curated registry.
    *   `glama.ai/mcp`: Another popular list.
    *   `github.com/modelcontextprotocol/servers`: The "official" reference implementation list.
*   **Strategy:** Novapi should pull from `smithery.ai` or similar APIs if available, or just parse the "Awesome MCP" list to populate a "Marketplace" tab.

## 2. Strategic Recommendations

### 2.1 Architecture Upgrade: HTTP/SSE Support (Critical)
We cannot rely solely on Stdio for agents like `browser-use`.
*   **Action:** Update `MCP Host Core` requirement to include HTTP transport support.
*   **Action:** Ensure the UI can render a "Connect to Remote MCP" dialog.

### 2.2 The "Agent Bundle"
Instead of just "installing a server", we should offer a **"Browser Agent Bundle"**:
1.  Downloads `chromium` (via Playwright).
2.  Starts `mcp-browser-use` as a background daemon.
3.  Connects Novapi to it via HTTP.

### 2.3 UX Refinement: "Thinking" State
*   **Action:** Implement a specific parser for `<think>...</think>` tags in the chat stream.
*   **Visual:** Show a pulsing brain icon or "Reasoning..." label that expands to show the raw trace.

## 3. References
*   **mcp-browser-use:** https://github.com/Saik0s/mcp-browser-use
*   **Jan.ai Changelog:** https://www.jan.ai/changelog
*   **Smithery:** https://smithery.ai
