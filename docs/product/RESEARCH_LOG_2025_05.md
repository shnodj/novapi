# Research Log: May 2025

**Status:** Draft
**Date:** 2025-05-15
**Focus:** Competitor Feature Analysis & Strategic Integrations

## 1. Executive Summary

This research session focused on "DeepSeek R1" integration patterns in competitor tools (Jan.ai), the emergence of standardized MCP Registries (Smithery.ai, Glama.ai), and the latest capabilities of the "Browser Use" MCP server.

**Key Findings:**
1.  **DeepSeek R1 UI Patterns:** Competitors like Jan.ai and the official DeepSeek web UI separate "Reasoning Content" (the `<think>` block) into a collapsible, distinct UI element. This is becoming a standard expectation for "Reasoning Models".
2.  **MCP Registry Standardization:** Two major registries have emerged: `smithery.ai` and `glama.ai`. Both offer searchable catalogs of MCP servers. Deep integration (e.g., search registry within Novapi) is a competitive necessity.
3.  **Browser Use "Daemon" Mode:** The `mcp-browser-use` project now supports a "Streamable HTTP" transport mode (running as a daemon on port 8383) to solve timeout issues inherent in Stdio transport for long-running tasks.
4.  **LM Studio Deep Linking:** LM Studio v0.3.17+ supports `lmstudio://` deep links for one-click MCP server installation, a feature we should mirror (`novapi://`).

## 2. Detailed Findings

### 2.1 DeepSeek R1 & Reasoning UI

**Observation:**
DeepSeek R1 (and distilled versions) output a `<think>` tag containing the chain-of-thought process before the final answer.

**Competitor Implementation (Jan.ai / Cherry Studio):**
-   **Visuals:** The thinking process is not shown as raw text. Instead, a "Thought" or "Reasoning" accordion/dropdown is displayed at the top of the message.
-   **Interaction:** Users can expand this to see the step-by-step logic, but it defaults to collapsed to keep the chat clean.
-   **Technical:** The frontend parses the streaming response. Content between `<think>` and `</think>` is routed to a separate state variable `reasoning_content` vs `content`.

**Recommendation for Novapi:**
-   Update `new-api` to optionally parse these tags if the model `family` is "reasoning".
-   Update Frontend to render a "Thinking..." indicator that expands into a grayed-out code block or dedicated reasoning view.

### 2.2 MCP Registry Integration

**Observation:**
Users currently struggle to find and configure MCP servers. "Registry" tools are emerging to solve this.

-   **Smithery.ai:** positioned as the "Package Manager for MCP".
-   **Glama.ai:** offers a "Marketplace" feel with verified servers.

**Strategic Opportunity:**
Novapi should implement a "Marketplace" tab that pulls from these open registries (via their APIs or public indices) to allow "One-Click Install".
-   *Action:* Add a `Marketplace` view in Settings -> MCP.
-   *Source:* `https://github.com/smithery-ai/directory` or similar.

### 2.3 Browser Use "Daemon" Integration

**Observation:**
The `saik0s/mcp-browser-use` repo explicitly advises *against* Stdio for complex tasks due to timeouts.
-   **New Mode:** Runs a Python web server (FastAPI/FastMCP) on `localhost:8383`.
-   **Benefits:** Persistent connection, better observability (dashboard at `/dashboard`), and "Skills" persistence.

**Recommendation for Novapi:**
-   Instead of just spawning the process via stdio, Novapi should be able to *manage* this daemon.
-   *Option A:* Spawn the python script in background and connect via SSE/HTTP.
-   *Option B:* Docker container (heavier, but cleaner).
-   *Decision:* Stick to Option A (Local Process) for now to avoid Docker requirement, but use the HTTP transport config in `mcp.json`.

### 2.4 Deep Linking (`novapi://`)

**Observation:**
LM Studio uses `lmstudio://install-mcp?url=...` to allow web-based directories to trigger local installs.

**Recommendation:**
-   Register `novapi://` protocol handler in Wails.
-   Support `novapi://mcp/install?url=<github_url>` to auto-fill the "Add Server" modal.

## 3. Gap Analysis

| Feature | Jan.ai | LM Studio | Novapi (Current) | Action Item |
| :--- | :--- | :--- | :--- | :--- |
| **Reasoning UI** | Yes (Accordion) | Yes | Raw Text | **P1: Implement Parsing** |
| **MCP Registry** | Planned | Manual | Manual | **P1: Add Marketplace Tab** |
| **Browser Agent** | Via Extension | Via MCP | Manual Config | **P0: 1-Click Bundle** |
| **Deep Links** | No | Yes | No | **P2: Protocol Handler** |

## 4. Next Steps (Drafting Requirements)

1.  **Design "Thinking" UI Component:** Needs to handle streaming tokens and "finished" state.
2.  **Define Marketplace Source:** Choose one primary registry (Smithery seems most developer-friendly) to index initially.
3.  **Prototype Protocol Handler:** Investigate Wails `OnDomReady` or deep link handling capabilities for macOS/Windows.
