# Research Log: March 2025

**Date:** 2025-03-14
**Author:** Product Team (Jules)
**Status:** Draft

## 1. Market Overview & Competitor Updates

### 1.1 Jan.ai (v0.5.16 - March 13, 2025)
*   **Focus:** Security patches, UI polish (Model Hub redesign), and performance.
*   **Key Features:**
    *   **Redesigned Model Hub:** Cleaner layout for model discovery.
    *   **Thread Settings:** More granular control per chat.
    *   **New Models:** Support for Claude 3.7 Sonnet, Gemma 3, GPT-4.5 Preview.
    *   **Local API Server:** continues to be a core feature.
*   **Takeaway for Novapi:** The "Model Hub" redesign suggests users find current discovery interfaces overwhelming. Novapi's "Featured Models" widget is the right direction but needs to be visually distinct and curated.

### 1.2 LM Studio (v0.3.x)
*   **MCP Support:** Fully integrated as a Host. Users configure via `mcp.json`.
*   **DeepSeek R1:** Official support with blog posts. They handle the "Reasoning" models well, likely by parsing the `<think>` tags.
*   **Differentiation:** LM Studio is closed source. Novapi's open-source nature remains its biggest leverage point for developer adoption (custom forks, audits).

### 1.3 DeepSeek R1 & "Reasoning" UX
*   **Trend:** "Reasoning" models (CoT) are the standard for coding/logic tasks.
*   **UX Pattern:** The standard interaction pattern is to *hide* the raw `<think>` tokens by default but allow expanding them for inspection.
*   **Requirement:** Novapi's Chat UI must parse `<think>...</think>` tags and render them in a collapsible "Thinking Process" accordion.

## 2. Model Context Protocol (MCP) Ecosystem

### 2.1 Top Servers (Candidates for "Featured" List)
1.  **`github/github-mcp-server` (15k+ stars):** The gold standard for engineering agents. Allows managing issues, PRs, and files. **Must Include.**
2.  **`microsoft/playwright-mcp` (11k+ stars):** A potential alternative to `browser-use`. It's backed by Microsoft and might be more stable for "Browser Automation" tasks.
    *   *Action:* Evaluate replacing `mcp-browser-use` with `playwright-mcp` or offering both.
3.  **`awslabs/mcp` & `hashicorp/terraform-mcp-server`:** DevOps focus. Good for "Power User" persona.

### 2.2 Security & Governance
*   **Zero Trust for Agents:** Articles emphasize the need for security layers between Agents and Tools.
*   **Novapi Implementation:** We must implement an "Approval" modal for MCP tool execution (e.g., "Allow this agent to read /etc/passwd?").

## 3. Strategic Recommendations (additions to Backlog)

1.  **UI Upgrade for DeepSeek R1:** Implement `<think>` tag parsing.
2.  **MCP Registry/Discovery:** Instead of just a raw list, we should pull from `github.com/modelcontextprotocol/servers` or a curated JSON list.
3.  **Playwright Integration:** Investigate `microsoft/playwright-mcp` as the default "Browser" capability.
4.  **GitHub Integration:** Add a specific "Connect GitHub" flow that configures the `github-mcp-server` automatically (Auth handling).

## 4. Updates to Roadmap
*   **Priority Shift:** DeepSeek R1 UX is now P0 (Critical) to match market hype.
*   **New Feature:** "MCP Permission Manager" (User Approval UI).
