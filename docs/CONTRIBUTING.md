# Contributing to Novapi

First off, thank you for considering contributing to Novapi! It's people like you that make Novapi such a great tool.

## 1. Code of Conduct

By participating in this project, you are expected to uphold our Code of Conduct. Please be respectful and inclusive.

## 2. Development Workflow

We use the [Fork & Pull](https://github.com/susam/gitpr) workflow.

1.  **Fork** the repo on GitHub.
2.  **Clone** the project to your own machine.
3.  **Checkout** a new branch for your feature or fix.
    *   `git checkout -b feature/amazing-feature`
    *   `git checkout -b fix/annoying-bug`
4.  **Commit** changes to your own branch.
5.  **Push** your work back up to your fork.
6.  **Submit** a Pull Request so that we can review your changes.

## 3. Commit Message Convention

We follow the **Conventional Commits** specification. This allows us to automatically generate changelogs and version numbers.

**Format**: `<type>(<scope>): <subject>`

**Types**:
*   `feat`: A new feature
*   `fix`: A bug fix
*   `docs`: Documentation only changes
*   `style`: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
*   `refactor`: A code change that neither fixes a bug nor adds a feature
*   `perf`: A code change that improves performance
*   `test`: Adding missing tests or correcting existing tests
*   `chore`: Changes to the build process or auxiliary tools and libraries such as documentation generation

**Example**:
```
feat(channel): add support for local Ollama instance
fix(ui): resolve dark mode flickering on startup
docs(readme): update installation instructions
```

## 4. Coding Standards

### Backend (Go)
*   **Formatting**: Always run `go fmt` before committing.
*   **Linting**: We use `golangci-lint`. Ensure your code passes standard lint checks.
*   **Error Handling**: Don't ignore errors. Handle them or wrap them with context.
*   **Naming**: Follow standard Go naming conventions (CamelCase).

### Frontend (React)
*   **Formatting**: We use Prettier.
*   **Components**: Functional components with Hooks are preferred over Class components.
*   **State Management**: Use React Context or local state for simple cases.
*   **Structure**: Keep components small and focused.

## 5. Pull Request Guidelines

*   **Title**: Use the Conventional Commits format for the PR title.
*   **Description**: clearly describe the problem and the solution. Link to relevant issues (e.g., `Fixes #123`).
*   **Tests**: Ensure all tests pass. Add new tests for new features.
*   **Review**: Wait for at least one approval before merging (if you have write access).

## 6. Directory Structure

*   `backend/`: All Go code (business logic, database, API).
*   `frontend/`: React application (UI).
*   `docs/`: Project documentation.
*   `build/`: Wails build artifacts and assets.
