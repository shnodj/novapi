# CICD Plan for Novapi

This document outlines the plan for Continuous Integration and Continuous Deployment (CICD) for the Novapi project.

## 1. Continuous Integration (CI)

We utilize GitHub Actions for our CI pipeline. The workflow is defined in `.github/workflows/ci.yml`.

### Triggers
- **Push** to `main` branch.
- **Pull Request** targeting `main` branch.

### Jobs
- **Build & Test**:
    - Runs on `ubuntu-latest`.
    - Sets up Go and Node.js environments.
    - Installs Linux dependencies required for Wails (GTK3, WebKit2GTK).
    - Installs Wails CLI.
    - Builds the Frontend (React/Vite).
    - Builds the Backend (Go).
    - Runs Go unit tests.

## 2. Future Improvements (CD)

For Continuous Deployment, we plan to add the following:

- **Cross-Platform Building**: Use matrix builds to compile for Windows, macOS, and Linux.
- **Release Automation**: Automatically create GitHub Releases with binaries attached when a tag is pushed.
- **Signing**: Code signing for Windows and macOS binaries.

## 3. Testing Strategy

- **Unit Tests**: Go `testing` package for backend logic (especially `backend/channel` and `backend/model`).
- **Integration Tests**: Test the Gin server endpoints using `httptest`.
- **E2E Tests**: Use Playwright to test the Wails application frontend and its interaction with the backend (this requires a more complex setup with xvfb on CI).
