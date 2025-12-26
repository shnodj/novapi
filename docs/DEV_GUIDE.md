# Developer Guide

Welcome to the Novapi developer guide! This document will help you set up your environment and start coding.

## 1. Prerequisites

Before you begin, ensure you have the following installed:

*   **Go**: Version 1.21 or higher. [Download](https://go.dev/dl/)
*   **Node.js**: Version 18 or higher (LTS recommended). [Download](https://nodejs.org/)
*   **NPM**: Usually comes with Node.js.
*   **Wails CLI**: `go install github.com/wailsapp/wails/v2/cmd/wails@latest`
*   **GCC**: Required for CGO (SQLite).
    *   *Windows*: Install MinGW or TDM-GCC.
    *   *Linux*: `sudo apt install build-essential`
    *   *macOS*: `xcode-select --install`

## 2. Setup

1.  **Clone the repository**:
    ```bash
    git clone https://github.com/shnodj/novapi.git
    cd novapi
    ```

2.  **Install Frontend Dependencies**:
    ```bash
    cd frontend
    npm install
    cd ..
    ```

3.  **Download Go Modules**:
    ```bash
    go mod tidy
    ```

## 3. Running Locally

To run the application in "Dev Mode" (with hot reload):

```bash
wails dev
```

*   This will compile the backend, start the frontend dev server (Vite), and open the application window.
*   Backend changes will trigger a recompile.
*   Frontend changes will trigger a browser refresh.

## 4. Building for Production

To create a standalone binary:

```bash
wails build
```

*   The output binary will be located in the `build/bin/` directory.

## 5. Testing

### Backend Tests
Run standard Go tests:

```bash
go test -v ./...
```

### Frontend Verification
We use Playwright for E2E verification (see `docs/CICD_PLAN.md`).

## 6. Troubleshooting

*   **"missing go.sum entry"**: Run `go mod tidy`.
*   **"cgo: C compiler \"gcc\" not found"**: Ensure GCC is in your PATH.
*   **Ubuntu Build Failures**: Ensure you have the `webkit2gtk` libraries installed (see `CONTRIBUTING.md` or `README.md`).
