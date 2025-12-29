# Product Research Log - February 2025

**Date:** 2025-02-24
**Topic:** Competitor Analysis & MCP Landscape
**Researcher:** Jules (Product Manager Agent)

## 1. Key Findings

### 1.1 LM Studio & MCP
*   **Update:** LM Studio v0.3.17+ now acts as an **MCP Host**.
*   **Mechanism:** Uses `mcp.json` configuration, compatible with Cursor IDE's format.
*   **Implication:** They have beaten Novapi to the punch on "MCP Host" for desktop. However, they are closed source. Novapi's differentiation remains "Open Source" and "Gateway/Router" focus (not just a runner).
*   **Feature:** They support a "Add to LM Studio" deeplink for installing MCP servers.

### 1.2 DeepSeek R1
*   **Status:** Extremely popular "Reasoning" model.
*   **Local Run:** Fully supported in Ollama (`ollama pull deepseek-r1`).
*   **Opportunity:** Novapi should have a "Featured Model" card on the home screen to one-click install/run this, catering to the current hype.

### 1.3 Ollama Tool Calling
*   **Status:** Ollama now supports native tool calling (single-shot).
*   **Relevance:** This is the technical primitive required for Novapi to translate "MCP Tool Calls" into "Ollama API Calls". We don't need complex prompt engineering; we can use the native API.

### 1.4 MCP Ecosystem
*   **Popular Servers:**
    *   **Filesystem:** Basic read/write access.
    *   **GitHub:** Manage PRs/Issues.
    *   **Postgres:** Database access.
    *   **Google Drive/Slack:** Enterprise data.
*   **Registry:** There is no single central registry yet, but "Awesome MCP" lists are growing.

### 1.5 Open WebUI
*   **Status:** Still heavy on Docker. The "Desktop" version is Alpha.
*   **Differentiation:** Novapi is still lighter and easier to install (Native binary vs Docker container).

## 2. Strategic Adjustments

*   **P0:** Support `mcp.json` standard (same as Cursor/LM Studio) to allow users to share config.
*   **P0:** "Featured Models" UI. Don't just list what's installed; suggest what *to* install (starting with DeepSeek R1).
*   **P1:** "Add to Novapi" Deep Link support (copying LM Studio's feature).

## 3. References
*   LM Studio MCP Docs: https://lmstudio.ai/docs/app/mcp
*   DeepSeek R1 on Ollama: `ollama run deepseek-r1`
*   Ollama Tool Calling: https://docs.ollama.com/capabilities/tool-calling
