# Product Research Log - 2026-01-10

**Status:** Draft
**Author:** Product Team
**Focus:** MCP Ecosystem, Competitor Analysis (Jan.ai/LM Studio), DeepSeek R1 Integration

## 1. Executive Summary

This research cycle focused on bridging the gap between Novapi and leading competitors (Jan.ai, LM Studio) specifically in the area of MCP (Model Context Protocol) and local model integration.

Key takeaways:
1.  **Deep Linking is Critical:** LM Studio's `lmstudio://` protocol allows one-click MCP installation. Novapi needs `novapi://` equivalent.
2.  **Browser Use requires HTTP:** The standard Stdio transport is insufficient for long-running agents like "Browser Use". We must support HTTP/SSE transport for local MCP servers.
3.  **DeepSeek R1 UI:** The standard for R1 is to visualize the "Thinking Process" (content within `<think>` tags) as a collapsible section separate from the markdown response.
4.  **Local Channels:** We can provide high-value "free" intelligence by wrapping the GitHub Copilot CLI into a local OpenAI endpoint using tools like `copilot-api`.

## 2. Competitor Analysis

### Jan.ai (v0.7.5+)
*   **Jan Browser MCP:** A dedicated MCP server for browser automation.
*   **Jan Projects:** A new way to organize chats and context, likely similar to Claude Projects.
*   **Lesson:** Browser automation is becoming a table-stakes feature for local AI clients.

### LM Studio (v0.3.17+)
*   **MCP Host:** Now fully supports `mcp.json` (Cursor compatible).
*   **Deep Links:** Implemented `lmstudio://install-mcp?url=...` to reduce friction.
*   **Lesson:** Compatibility with `mcp.json` is the de-facto standard. We should not invent our own config format.

## 3. Technology Scans

### DeepSeek R1 Integration
*   **Challenge:** R1 models output a "Reasoning Trace" before the final answer.
*   **Pattern:** Parse the stream for `<think>...</think>` tags.
*   **UI Requirement:** Render this trace in a "grayed out" or collapsible "Thought Process" box. Do not treat it as part of the final Markdown.

### "Browser Use" Agent (Saik0s/mcp-browser-use)
*   **Problem:** Stdio transport has timeouts (often 60s) which breaks complex browser tasks.
*   **Solution:** This implementation uses a persistent HTTP daemon (default port 8383) and SSE (Server-Sent Events).
*   **Requirement:** Novapi's MCP Host implementation must support connecting to HTTP/SSE endpoints, not just spawning Stdio processes.

### Local Channels: GitHub Copilot Wrapper
*   **Tool:** `ericc-ch/copilot-api`
*   **Function:** Reverse engineers the Copilot plugin API to expose an OpenAI-compatible endpoint.
*   **Opportunity:** "Login with GitHub" in Novapi could spin up this bridge, giving users access to GPT-4o/Claude 3.5 Sonnet (via their Copilot sub) without paying per-token API fees.

## 4. Recommendations for Requirements

1.  **Update MCP Host (P0):** Add HTTP/SSE transport support.
2.  **New Feature (P1):** Deep Link Handler (`novapi://`).
3.  **UI Update (P1):** DeepSeek R1 "Thinking" visualization.
4.  **New Channel (P2):** Integration with `copilot-api` (or similar) for "Free" high-end models.
