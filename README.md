# Novapi

**The Desktop-Native AI Gateway.**

[![Build and Test](https://github.com/shnodj/novapi/actions/workflows/ci.yml/badge.svg)](https://github.com/shnodj/novapi/actions/workflows/ci.yml)

Novapi is a lightweight, single-user alternative to [One-API](https://github.com/songquanpeng/one-api), designed to run locally on your desktop. It removes the need for Docker, Redis, and complex server configurations, giving you an OpenAI-compatible API gateway for all your LLM needs (Cloud + Local).

## 🚀 Features (Planned)

*   **Zero-Dependency**: No Docker, no Redis. Just a single executable.
*   **Local & Cloud**: Unified interface for OpenAI, Claude, Azure, and **local Ollama/Llama.cpp**.
*   **Privacy-First**: Data stays on your machine (SQLite).
*   **Dev-Friendly**: Built with Go + Wails + React.

## 📖 Documentation

*   [**Architecture Overview**](./docs/ARCHITECTURE.md): Learn how Novapi works under the hood.
*   [**Developer Guide**](./docs/DEV_GUIDE.md): Setup your local environment.
*   [**Contributing Guidelines**](./docs/CONTRIBUTING.md): How to contribute code.
*   [**CI/CD Plan**](./docs/CICD_PLAN.md): Our automation strategy.

## 🛠️ Quick Start (Development)

1.  **Prerequisites**: Go 1.21+, Node.js 18+, Wails CLI.
2.  **Clone**: `git clone https://github.com/shnodj/novapi`
3.  **Run**:
    ```bash
    wails dev
    ```

## 🏗️ Architecture

Novapi uses a "Dual-Process" model:
*   **GUI**: Wails (React) for configuration and management.
*   **Server**: Embedded Gin HTTP server (port 3000) for API compatibility.

See [Architecture](./docs/ARCHITECTURE.md) for details.

## 📄 License

MIT License.
