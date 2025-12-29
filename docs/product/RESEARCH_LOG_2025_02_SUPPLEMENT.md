# Product Research Log - February 2025 (Supplement)

**Date:** 2025-02-24
**Topic:** Emerging Competitors & MCP Opportunities
**Researcher:** Jules (Product Manager Agent)

## 1. Competitive Landscape Analysis

### 1.1 Jan.ai (v0.6.3+)
*   **Overview:** A direct open-source competitor offering a "ChatGPT alternative" that runs 100% offline.
*   **Key Features:**
    *   **Experimental MCP Support:** Users can enable MCP Servers via settings.
    *   **Localhost Connectivity:** "Connects cleanly" to Ollama (:11434) and LM Studio (:1234).
    *   **Jan v1 Model:** A custom 4B parameter model optimized for tool calling and reasoning.
    *   **Local RAG:** Built-in "Knowledge Base" (files/web).
*   **Threat Assessment:** Jan is moving faster on "Client-side MCP" and RAG.
*   **Novapi Opportunity:** Novapi differentiates by being a **Gateway** (API Host), not just a Client. We should focus on *serving* these local tools to other apps (like Cursor or VS Code) via our API, rather than just consuming them in our UI.

### 1.2 Cherry Studio
*   **Overview:** "All-in-One AI Assistant" with aggressive feature updates.
*   **Key Features:**
    *   **Multi-Model Management:** Unified interface for OpenAI, Anthropic, and Local models.
    *   **Knowledge Base:** Strong file/web import capabilities.
    *   **Pre-configured Assistants:** 300+ "smart bodies" (agents).
    *   **MCP Support:** Listed as a core feature.
*   **Takeaway:** The market expectation for "Desktop AI" now includes **RAG** and **Multi-Provider Routing** as standard features.

### 1.3 Open WebUI Desktop
*   **Overview:** Now has an Alpha desktop build (Tauri v2).
*   **Impact:** Moves Open WebUI from "Docker/Server" territory into "Desktop App" territory.
*   **Weakness:** Still feels like a web-wrapper; Novapi/Wails can offer tighter system integration (Tray icon, global shortcuts).

## 2. Technology Trends & MCP

### 2.1 The "Browser Use" Phenomenon
*   **What is it?** An MCP server that allows an LLM to control a headless browser to perform tasks (click, type, extract).
*   **Relevance:** This is currently the most "viral" use case for local agents.
*   **Recommendation:** Novapi should bundle or provide a one-click setup for a "Browser Use" MCP server, allowing users to build "Web Agents" easily.

### 2.2 Trending MCP Servers (2025)
*   **Development:** GitHub, PostgreSQL, Docker.
*   **Knowledge:** Exa Search (Research), Context 7 (Docs).
*   **Utility:** Filesystem (Table stakes), Google Maps.

## 3. Strategic Recommendations

1.  **Pivot to "Gateway" Identity:** Don't just be another chat UI. Be the **Universal Local Host**. If a user runs Novapi, their Cursor/VS Code should instantly be able to access "Browser Use" or "Local RAG" via standard OpenAI-compatible endpoints provided by Novapi.
2.  **"Browser Use" Integration (P0):** This is a high-value feature that differentiates us from simple chat wrappers.
3.  **Knowledge Base (RAG) is Essential:** We cannot compete without "Add file to context".
