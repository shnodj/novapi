# Product Research Log (March 2025)

**Date:** 2025-03-11
**Focus:** Competitor Analysis (Jan, LM Studio) & MCP Ecosystem

## 1. Competitor Analysis

### Jan.ai (v0.7.5)
*   **Key Feature:** "Jan Browser MCP" (v0.7.3) allows the model to control a browser. This validates our "Browser Use" requirement.
*   **UX:** "Jan Projects" (v0.7.0) organizes chats/models by context.
*   **Stability:** Recent releases focus on "Hotfixes" and "Privacy" (Analytics opt-in).
*   **Takeaway:** We should prioritize the "Browser Use" MCP integration to match their capability.

### LM Studio (v0.3.17+)
*   **MCP Support:** Full "Host" support using `mcp.json`.
*   **Deep Linking:** Introduced `lmstudio://install-mcp?config=...` protocol. This allows one-click installation of MCP servers from websites (like Glama or Smithery).
*   **Compatibility:** They explicitly mention compatibility with Cursor's `mcp.json` format.
*   **Takeaway:** **Critical.** We must implement `novapi://install-mcp` to match this "One-Click" experience. It significantly lowers the barrier to entry for users.

## 2. MCP Ecosystem Updates
*   **Registries:** `smithery.ai` and `glama.ai/mcp` are the de-facto public registries.
*   **Protocol:** "Async operations" are coming to the protocol (long-running tasks).
*   **Servers:**
    *   `hf-mcp-server`: Official Hugging Face server for model/dataset search.
    *   `browser-use`: The leading open-source browser automation agent (Python).

## 3. Recommendations for Novapi

### 3.1 "One-Click" MCP Installation (Priority: High)
*   **Feature:** Register `novapi://` protocol handler.
*   **Action:** Support `novapi://install-mcp?url=...` or `?config=...`.
*   **Benefit:** Parity with LM Studio; ease of use.

### 3.2 "Browser Use" Integration (Priority: High)
*   **Strategy:** Instead of building our own, wrap the `Saik0s/mcp-browser-use` server.
*   **Deployment:** Since it requires Python/Playwright, we might need a "Docker-less" way to bundle it, or guide the user to install it via `uv` / `pip`.

### 3.3 Thinking Process UI (DeepSeek R1)
*   **Observation:** While no specific "standard" was found, the pattern of "Collapsible Thinking Blocks" (parsing `<think>...</think>`) is standard in web UIs (OpenWebUI).
*   **Requirement:** Ensure the Chat UI parses these tags and renders them in a disclosed state (collapsed by default).
