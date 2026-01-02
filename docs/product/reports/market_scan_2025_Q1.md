# Market Scan & Technology Radar: 2025 Q1

**Date:** 2025-02-24
**Author:** Novapi Product Team
**Status:** Final

## 1. Executive Summary

This quarter's research confirms that the **Model Context Protocol (MCP)** has won the battle for the "Agent Interface" standard. The ecosystem has exploded with hundreds of community servers.

**Strategic Pivot:** Novapi should transition from just "supporting" MCP to being the **"Best MCP Host for Power Users"**. While competitors like Jan.ai are integrating basic MCP support, Novapi can differentiate by offering robust **Agentic Workflows** (via `browser-use`) and deep local system integration (File System, Git, Docker) that others are hesitant to enable by default due to safety concerns (which we can mitigate via a permission system).

## 2. Key Competitor Updates

### 2.1 Jan.ai (v0.7.5)
*   **Browser Control:** Launched "Jan Browser MCP" for browser automation. This directly validates our hypothesis that browser agents are the next killer app.
*   **DeepSeek R1 Support:** Native support for DeepSeek R1 models, presumably including reasoning trace visualization.
*   **Project Management:** Introduced "Projects" to organize chats and models, a feature Novapi currently lacks.
*   **Local API Server:** Strengthened their local server capabilities, moving closer to our "Gateway" value prop.

### 2.2 Open WebUI
*   **Feature Density:** continues to be the "kitchen sink" of local AI, with built-in RAG, pipelines, and multi-model orchestration.
*   **Weakness:** Still heavily Docker-reliant. The "Desktop App" is a work in progress. This remains our key wedge: **Native Binary Simplicity**.

### 2.3 LM Studio
*   **MCP Support:** Has successfully integrated MCP support via `mcp.json`.
*   **UX Leader:** Remains the benchmark for model discovery and downloading UX.

## 3. DeepSeek R1 & The "Reasoning" UX

The release of DeepSeek R1 has shifted user expectations. Users now expect to see *how* the model thinks, not just the final answer.

**Requirement:** The chat UI must support a collapsible `<think>` tag renderer.
*   **Raw:** `<think> ... reasoning steps ... </think> Final Answer`
*   **Rendered:** `[Thinking Process (Click to Expand)]` -> `Final Answer`

## 4. MCP Ecosystem Opportunities

We have identified high-value MCP servers to "Bundle" or "One-Click Install" in Novapi:

### Tier 1: Core Integrations (The "Must Haves")
1.  **Browser Use (`browser-use/browser-use-mcp-server`):**
    *   *Why:* Enables "Agentic" workflows (e.g., "Go to Amazon and find the price of X").
    *   *Action:* Bundle the Docker/Python setup for this to make it seamless.
2.  **Filesystem & Git:**
    *   *Why:* The basis of any coding assistant.
    *   *Action:* Built-in support, no configuration needed.
3.  **Brave Search / Google Search:**
    *   *Why:* Real-time internet access is the #1 requested feature for local LLMs.

### Tier 2: Developer Tools (The "Power User" Wedge)
1.  **PostgreSQL / SQLite:** Direct database querying.
2.  **Docker:** Managing containers via chat.
3.  **GitHub/GitLab:** Issue tracking and PR management.

## 5. Recommendations for Q2 Roadmap

1.  **MCP Marketplace/Registry UI:**
    *   Don't just edit `config.json`. Build a UI to browse and install these servers.
    *   Leverage `github.com/modelcontextprotocol/servers` as the source of truth.

2.  **"Agent Mode" Toggle:**
    *   A UI switch that enables the `browser-use` MCP server and switches the system prompt to be agentic.

3.  **DeepSeek R1 "Thought" Visualization:**
    *   Prioritize parsing the `<think>` tags in the React frontend.

4.  **Local RAG (Lite):**
    *   Instead of a heavy vector DB, start with a simple "Search in Files" MCP server that allows the LLM to grep the local project.
