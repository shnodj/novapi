# Product Research Log - March 2025

**Date:** 2025-03-15
**Topic:** Competitor Roadmap & Agentic Workflows
**Researcher:** Jules (Product Manager Agent)

## 1. Executive Summary

A window of opportunity exists between now (March) and June 2025. Our primary competitor, **Jan.ai**, has scheduled their full MCP rollout for late June (v0.6.3). If Novapi ships a stable MCP Host with "Browser Use" capabilities in April/May, we can establish ourselves as the default "Agentic Gateway" before Jan.ai catches up.

## 2. Competitor Analysis

### 2.1 Jan.ai
*   **Current State (March 2025 - v0.5.16):** Focus is on stability, security, and cloud model support.
*   **Future Roadmap:**
    *   **June 25th, 2025 (v0.6.3):** "Unlocking MCP for everyone".
    *   **July/August:** Advanced MCP tutorials and Jupyter integration.
*   **Strategic Takeaway:** We have a ~3-month lead time to define what a "Local MCP Host" looks like.

### 2.2 LM Studio
*   **Current State:** v0.3.17+ supports `mcp.json`.
*   **Trend:** They are integrating heavily with LiteLLM for routing.
*   **Gap:** They are closed source. Power users want an open stack where they can inspect the "Agent's" code, especially for tools like "Browser Use" that control their screen.

## 3. DeepSeek R1 Integration

*   **Standard UX:** The industry standard (OpenWebUI, Jan.ai) is to parse the `<think>` ... `</think>` tags from DeepSeek R1 and render them in a **collapsible "Thinking Process" accordion**.
*   **Requirement:** Novapi's chat interface *must* implement this. Raw tags confuse users.

## 4. "Browser Use" & Agentic Infrastructure

*   **The Problem:** Standard MCP runs over `stdio` (standard input/output).
*   **The Finding:** Leading "Browser Use" implementations (e.g., `Saik0s/mcp-browser-use`) report that `stdio` has timeout issues for long-running browser automation tasks (30-120s+).
*   **The Solution:** They recommend or require **HTTP/SSE Transport** for the MCP connection to keep the daemon alive.
*   **Novapi Impact:** We cannot just implement `stdio` client support. We **must** support connecting to MCP servers via HTTP/SSE to support high-value agents like Browser Use.

## 5. Strategic Recommendations

1.  **Engineering P0:** Implement **HTTP/SSE Transport** for the MCP Host module.
2.  **UI P1:** Implement "Thinking Process" collapsible view for DeepSeek R1.
3.  **Marketing:** Position Novapi as the "First Open Source Desktop MCP Host" (beating Jan.ai to the June deadline).
