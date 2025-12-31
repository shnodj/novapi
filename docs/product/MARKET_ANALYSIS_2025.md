# Market Analysis & Technology Radar (2025)

**Date:** December 2025
**Author:** Novapi Product Team
**Last Updated:** 2025-12-31

## 1. Executive Summary

Novapi aims to be the premier **Desktop-Native AI Gateway**, solving "Local Token Anxiety" by unifying local LLMs (Ollama, Llama.cpp) and cloud APIs (OpenAI, Anthropic) into a single, standardized interface.

As of late 2025, the battleground has shifted from "who can run the model" to **"who can run the Agent"**. With Jan.ai v0.7.x releasing stable MCP and Browser support, and LM Studio maturing its host capabilities, Novapi must leverage its open architecture to provide the most flexible **Agent Runtime**.

## 2. Technology Radar

### 2.1 Model Context Protocol (MCP) - *Standardized*
**Status:** Core Requirement
**Description:** Open standard for connecting LLMs to external data and tools.
**Market Update:**
*   **Jan.ai (v0.7.5):** Full stable MCP support, including "Jan Browser MCP".
*   **LM Studio:** Established MCP Host support.
*   **Novapi Strategy:**
    1.  **Transport Versatility:** Support both Stdio (simple tools) and HTTP/SSE (complex agents like Browser Use).
    2.  **Config Portability:** seamless `mcp.json` import/export.

### 2.2 Local Inference
*   **Ollama:** The de-facto standard runtime.
*   **DeepSeek R1:** The dominant open-weights model for reasoning.
    *   *Action:* Novapi must provide specific presets for R1 distilled variants (7b/8b/14b/32b) to cater to different hardware tiers.

### 2.3 Agentic Capabilities (New)
*   **Browser Use:** The "Killer App" for local agents.
    *   *Implementation:* `mcp-browser-use` (Python/Playwright) via HTTP transport.
    *   *Differentiation:* Novapi can run this as a background service, exposing it to *any* client (e.g., a simple curl request to Novapi can trigger a browser agent task).

## 3. Competitor Analysis

| Product | Type | Strengths | Weaknesses | Novapi Differentiation |
| :--- | :--- | :--- | :--- | :--- |
| **Open WebUI** | Web App (Docker) | Feature-rich (RAG, Pipelines), polished UI. | Heavy (Docker), complex setup. | **Native Desktop App**. No Docker. Light footprint. |
| **LM Studio** | Desktop App | **MCP Support**, Great UI, Model Discovery. | **Closed Source**. | **Open Source**. Novapi acts as a *Server/Gateway* for other apps, not just a Chat UI. |
| **Jan.ai** | Desktop App | Open Source, **Browser MCP**, Projects. | Electron heaviness? | Focus on **Protocol Standardization** (MCP, OpenAI API) vs just Chat. We are the "Gateway". |
| **AnythingLLM** | Desktop/Cloud | Easy RAG setup. | focused on RAG/Docs. | Novapi focuses on **Developer/Power User** control and API unification. |

## 4. Strategic Opportunities

### 4.1 "The Universal Local Adapter"
The problem isn't just "running models" anymore (Ollama won that). The problem is **orchestration**.
*   Novapi must be the "Router" that sits between the User (or their IDE/Agent) and the Models.
*   **Killer Feature:** "Bring your own MCP Server".

### 4.2 "DeepSeek Ready"
*   Users want to try the latest models without reading docs.
*   Novapi should have a "Trending Models" section that automates the `ollama pull` command for hits like DeepSeek R1.

## 5. Roadmap Recommendations (2026 Q1)

1.  **MCP Host Implementation (P0):**
    *   Support `mcp.json` config file.
    *   Support stdio-based MCP servers (Node/Python).
    *   **NEW:** Support HTTP/SSE transport for robust Agent hosting.

2.  **Browser Agent Bundle (P1):**
    *   Detect if Python/Playwright is available.
    *   One-click spin-up of `mcp-browser-use` linked to the local LLM.

3.  **Featured Models UI (P1):**
    *   "One-click DeepSeek R1" button.
    *   Visual indicator of model status (downloading, running).
