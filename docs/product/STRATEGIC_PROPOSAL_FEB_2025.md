# Strategic Proposal: "MCP First" & Desktop Native (Feb 2025)

**Status:** Draft
**Owner:** Product Team
**Date:** 2025-02-24

## 1. Context & Motivation

The local AI landscape has shifted. "Running models" is no longer the primary differentiator; **orchestration** is.
*   **Ollama** has commoditized local inference.
*   **DeepSeek R1** has reset expectations for open-source model capability.
*   **Model Context Protocol (MCP)** has emerged as the standard for connecting LLMs to tools/data.

Our competitor **LM Studio** has recently released MCP Host support (v0.3.17+), positioning themselves as the central hub for local AI. To remain relevant, Novapi must pivot from being a "better OneAPI" to being the **"Open Source MCP Gateway"**.

## 2. The Core Proposal

### 2.1 Pivot to "MCP Host"
**Objective:** Novapi should be able to load any standard MCP Server (e.g., Filesystem, GitHub, Postgres) and expose its tools to *any* model it manages.

**User Story:**
> "As a developer, I want to install the 'Postgres MCP Server' in Novapi once, and then access my database from Claude (via API), Llama 3 (via Ollama), or DeepSeek R1, all through the Novapi chat interface or API."

**Requirements:**
*   Support `mcp.json` configuration standard (interoperable with Cursor/LM Studio).
*   Implement an MCP Client in the Go backend.
*   Map MCP Tools to OpenAI-compatible Function Calling definitions.

### 2.2 Capitalize on "DeepSeek Mania"
**Objective:** Reduce friction for users wanting to try the latest trending model.

**Feature:** **"Featured Models" Widget**
*   Instead of a blank list, show "Trending Now: DeepSeek R1".
*   Single button: **"Download & Run"** (wraps `ollama pull deepseek-r1`).
*   Visual status tracking.

### 2.3 The "De-Redis" Refactor (Pre-requisite)
To achieve the above on a desktop environment, we must remove the server-side bloat.
*   **Problem:** `new-api` defaults to requiring Redis, which is heavy for a desktop app.
*   **Solution:** Enforce SQLite + In-Memory Cache (Go-Cache) for the desktop build.

## 3. Competitive differentiation

| Feature | LM Studio | Open WebUI | Novapi (Proposed) |
| :--- | :--- | :--- | :--- |
| **License** | Closed | MIT | **MIT (Open Source)** |
| **Architecture** | Desktop Native | Docker Container | **Desktop Native (Wails)** |
| **Role** | Runner + Host | Chat App | **Gateway / Router** |
| **MCP Support** | Yes (Host) | WIP | **Yes (Host + Gateway)** |

## 4. Next Steps

1.  **Engineering:** Address the Go version dependency blocker (See `docs/technical/BACKEND_FEASIBILITY_REPORT.md`).
2.  **Product:** Design the "MCP Settings" UI in the frontend.
3.  **Marketing:** Prepare messaging around "Open Source Alternative to LM Studio".
