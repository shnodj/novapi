# Market Analysis & Technology Radar (2025)

**Date:** February 2025
**Author:** Novapi Product Team

## 1. Executive Summary

Novapi aims to be the premier **Desktop-Native AI Gateway**, solving "Local Token Anxiety" by unifying local LLMs (Ollama, Llama.cpp) and cloud APIs (OpenAI, Anthropic) into a single, standardized interface.

This document analyzes the current landscape of Local AI tools, new protocols like MCP, and identifies strategic opportunities for Novapi.

## 2. Technology Radar

### 2.1 Model Context Protocol (MCP) - *CRITICAL*
**Status:** Adopt
**Description:** Introduced by Anthropic, MCP is an open standard that allows LLMs to connect to external data (databases, file systems) and tools in a standardized way.
**Relevance:**
*   **The "N x M" Problem:** Previously, connecting `Ollama` to `Postgres` required custom code. MCP standardizes this.
*   **Novapi's Opportunity:** Novapi should act as an **MCP Host**. This means Novapi can load "MCP Servers" (plugins) and expose their capabilities to any model it manages.
*   **Use Case:** A user installs the "Filesystem MCP Server" in Novapi. Now, *any* model (Claude via API, or Llama 3 via Ollama) routed through Novapi can read/write local files securely.

### 2.2 Local Inference Standards
*   **Ollama:** Has become the *de facto* standard for running models locally on Linux/Mac/Windows. It exposes a REST API.
*   **Llama.cpp:** The engine underneath most tools.
*   **Realtime API:** OpenAI's new Websocket-based API for low-latency voice/interaction.

### 2.3 AI Gateways
The market is shifting from simple proxies to "AI Operations Platforms".
*   **LiteLLM / Helicone:** Focus on observability, logging, and cost tracking.
*   **One-API / New-API:** The foundation of Novapi. Strong on channel management, weak on local desktop integration.

## 3. Competitor Analysis

| Product | Type | Strengths | Weaknesses | Novapi Differentiation |
| :--- | :--- | :--- | :--- | :--- |
| **Open WebUI** | Web App (Docker) | Feature-rich (RAG, Pipelines), polished UI. | Heavy (requires Docker), complex setup for non-techs. | **Native Desktop App (Wails)**. No Docker required. Single executable. |
| **LM Studio** | Desktop App | Excellent Model Discovery & UI. | Closed Source. Mainly a runner, not a gateway for other apps. | **Open Source**. Focus on being a *Gateway* (API) for other tools (IDE, Agents). |
| **Jan.ai** | Desktop App | Open Source, clean UI, local-first. | Extensions ecosystem is custom. | Focus on **Protocol Standardization** (MCP, OpenAI API) rather than just Chat UI. |
| **Ollama** | CLI / Service | The standard backend. | No GUI. No auth/management. | Novapi *wraps* Ollama, adding management, auth, and unified routing. |

## 4. Strategic Opportunities

### 4.1 "The Universal Local Adapter"
Users have fragmented AI access:
*   `Ollama` running on port 11434.
*   `LM Studio` on port 1234.
*   `GitHub Copilot CLI` authenticated in terminal.
*   `AWS Q` authenticated in terminal.

**Opportunity:** Novapi scans the machine, detects these services/CLIs, and wraps them all into `http://localhost:3000/v1/chat/completions`. One API Key to rule them all.

### 4.2 MCP Integration
By implementing MCP, Novapi moves from being a "dumb pipe" (Gateway) to a "smart agent host".
*   **Phase 1:** MCP Host. Allow models to call tools.
*   **Phase 2:** MCP Server. Expose Novapi-managed models *as* an MCP resource to other IDEs (like Cursor or Windsurf).

## 5. Recommendations
1.  **Prioritize Auto-Discovery:** The "Magic" moment for users is installing Novapi and immediately seeing their Ollama models ready to use.
2.  **Invest in MCP:** This is the future of tool usage. Being the first desktop gateway to support generic MCP servers will be a huge differentiator.
3.  **Keep it Lightweight:** Strict "No Docker" policy. Use SQLite and embedded processes.
