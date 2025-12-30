# Product Research Log - February 2025 (Supplement)

**Date:** 2025-02-24 (Late Feb Update)
**Topic:** Browser Use Integration & RAG Feasibility
**Researcher:** Jules (Product Manager Agent)

## 1. Browser Use Integration Analysis

There are two distinct approaches to "Browser Use" for Novapi, catering to different user needs:

### 1.1 The "Agent" Approach (Autonomous)
*   **Tool:** `Saik0s/mcp-browser-use` (based on the `browser-use` Python library + Playwright).
*   **Mechanism:** Launches a *new*, clean browser instance controlled entirely by the LLM.
*   **Pros:** Good for complex, multi-step tasks (e.g., "Go to Amazon, find X, add to cart").
*   **Cons:** Heavy dependencies (Python, Playwright browsers). Hard to bundle in a single binary.
*   **Integration:** Can use `langchain_openai` with `base_url`. We can point this to Novapi's local Ollama endpoint (`http://localhost:8080/v1`).
*   **Recommendation:** Provide a guide or a "One-click Script" to install this side-by-side with Novapi, rather than embedding it.

### 1.2 The "Copilot" Approach (Interactive)
*   **Tool:** `djyde/browser-mcp` (Extension + Server) or `ChromeDevTools/chrome-devtools-mcp`.
*   **Mechanism:** Connects to the user's *running* browser via an extension or DevTools protocol.
*   **Pros:** Access to logged-in sessions (Gmail, GitHub, etc.). Lightweight (Node.js).
*   **Cons:** Privacy concerns (access to user's main browser). Requires user to install a Chrome Extension.
*   **Recommendation:** This fits the "Desktop App" vibe better. We should support `chrome-devtools-mcp` via `npx` (standard MCP) and investigate wrapping `djyde/browser-mcp` for a smoother UX.

## 2. Local RAG / Knowledge Base Feasibility

We need a lightweight, Go-native way to store vectors without requiring a heavy Python sidecar or Docker.

### 2.1 Solution: `chromem-go`
*   **Repo:** `philippgille/chromem-go`
*   **Type:** Embeddable vector database for Go.
*   **Features:**
    *   In-memory with optional persistence to disk.
    *   No CGO required (pure Go).
    *   Zero 3rd party dependencies.
    *   Supports OpenAI-compatible embedding APIs (which Novapi/Ollama provides).
*   **Verdict:** This is the perfect candidate for Novapi's "Knowledge Base" feature (P1). It eliminates the need for a separate vector DB process.

## 3. Jan.ai Updates (Feb 2025)

*   **Version 0.7.5:** Released recently.
*   **Key Features:**
    *   **MCP Support:** Jan is now an MCP Host (like LM Studio). They support `todoist` and `browser-mcp`.
    *   **DeepSeek R1:** Native support.
    *   **Extension Store:** They are moving towards an "Extension" model for features.
*   **Implication:** Jan.ai is moving fast on MCP. Novapi *must* implement `mcp.json` support immediately to remain competitive. Their "Jan Browser MCP" is a direct competitor to our planned "Browser Use" feature.

## 4. DeepSeek R1 Configuration

For users with "Local Token Anxiety", we should recommend specific quantized versions based on RAM:

*   **8GB RAM (MacBook Air):** `deepseek-r1:1.5b` (Fast, good for simple logic) or `deepseek-r1:7b` (Slower).
*   **16GB RAM (MacBook Pro):** `deepseek-r1:8b` (Sweet spot) or `deepseek-r1:14b` (If nothing else running).
*   **32GB+ RAM:** `deepseek-r1:32b`.

**Action:** The "Featured Models" UI should auto-detect system RAM (via Wails runtime) and suggest the correct tag.
