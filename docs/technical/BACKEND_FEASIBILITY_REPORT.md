# Backend Feasibility Report & Incident Log

**Date:** 2025-02-24
**Topic:** Build Blockers & Dependency Analysis
**Status:** Critical

## 1. Issue Description

During the initial build verification of the backend, a critical dependency version mismatch was identified.

**Error Log:**
```
go: github.com/QuantumNous/new-api@v0.9.18 requires go >= 1.25.1 (running go 1.24.3)
```

**Context:**
*   **Current Environment:** Go 1.24.3 (Linux/AMD64).
*   **Dependency:** `github.com/QuantumNous/new-api` (The core engine of Novapi).
*   **Anomaly:** Go 1.25 has not been released (stable is 1.23/1.24). This suggests the upstream library `new-api` has explicitly set `go 1.25.1` in its `go.mod` file, possibly by mistake or due to using a nightly/unstable toolchain.

## 2. Impact Assessment

*   **Blocker:** We cannot run `go test`, `go build`, or run the application locally without resolving this.
*   **CI/CD:** Any CI pipeline using standard Go images (e.g., `golang:1.23`) will fail.

## 3. Proposed Solutions

### Option A: Downgrade Dependency (Recommended)
Investigate previous versions of `new-api` (e.g., v0.9.17) to see if they have a valid Go version requirement.
*   *Pros:* Immediate fix.
*   *Cons:* Might miss features/fixes in v0.9.18.

### Option B: Fork & Patch
Fork `new-api`, edit `go.mod` to `go 1.23` or `go 1.24`, and verify if it actually uses any "future" features (unlikely). Then point our `go.mod` to the fork using `replace`.
*   *Pros:* Keeps us on the latest code.
*   *Cons:* Maintenance burden of a fork.

### Option C: Contact Upstream
Open an issue on `QuantumNous/new-api` reporting the invalid Go version.
*   *Pros:* Correct fix.
*   *Cons:* Slow (asynchronous).

## 4. Recommendation

**Execute Option B (Fork/Replace) locally to unblock development**, while pursuing Option C (Upstream Report).

1.  Add a `replace` directive in `go.mod` to a local version or a patched fork.
2.  Or, simply try to downgrade to verify if v0.9.17 works.

## 5. Other Findings

*   **Redis Dependency:** The codebase currently sets `common.RedisEnabled = false` in `bootstrap.go`, which is correct for our "De-Redis" goal. However, we must verify that `new-api` doesn't panic if Redis is disabled (some paths might strictly require it).
