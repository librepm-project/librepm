# Contributing & Branching Strategy

## Branch model

```
master   ←── develop ←── feature/...
                    ←── fix/...
         ←── hotfix/...
```

| Branch | Purpose | Direct commits |
|--------|---------|----------------|
| `master` | Production-ready code. Every push triggers the full CI pipeline and creates a release. | No — PRs only |
| `develop` | Main development branch. All features and fixes land here first. | No — PRs only |
| `feature/*` | New functionality. Branched from `develop`. | Yes |
| `fix/*` | Non-urgent bug fixes. Branched from `develop`. | Yes |
| `hotfix/*` | Urgent production fixes. Branched from `master`, merged back into **both** `master` and `develop`. | Yes |

---

## Day-to-day workflow

### Feature or fix

```bash
git checkout develop
git pull
git checkout -b feature/my-feature   # or fix/my-bug

# ... work ...

git push -u origin feature/my-feature
# → open PR targeting develop
```

### Release

1. Bump the version in `package.json` on `develop` (or in a dedicated `release/x.y.z` branch).
2. Open a PR from `develop` → `master`.
3. Merge after review and green CI.
4. The pipeline automatically:
   - Builds and pushes the Docker image tagged with the version (`ghcr.io/librepm-project/librepm:x.y.z`)
   - Creates a GitHub Release tagged `vx.y.z` with auto-generated notes and the image tarball attached.

### Hotfix

```bash
git checkout master
git pull
git checkout -b hotfix/critical-bug

# ... fix ...

# Merge into master
git checkout master
git merge --no-ff hotfix/critical-bug

# Merge into develop too — keeps branches in sync
git checkout develop
git merge --no-ff hotfix/critical-bug
```

---

## CI pipeline (`.github/workflows/ci.yml`)

Triggered on every **push to `master`** and every **pull request**.

```
Lint ──→ Test ──→ E2E ──→ Build & Push ──→ Release
                                 ↑                ↑
                         master push only  master push only
```

| Job | Runs on | What it does |
|-----|---------|-------------|
| `lint` | all triggers | Go + frontend ESLint |
| `test` | all triggers | Go unit tests + frontend tests |
| `e2e` | all triggers | Playwright tests in Docker |
| `build-and-push` | all triggers (push to GHCR only on `master`) | Builds Docker image; pushes only when merging to `master` |
| `release` | `master` push only | Creates GitHub Release `vx.y.z` with Docker image tarball |

PRs get the full lint → test → e2e → build cycle but **never push an image or create a release**.

---

## Versioning

The version is the single source of truth in `package.json`. The pipeline reads it with:

```bash
jq -r '.version' package.json
```

Bump it before merging to `master`. Use [Semantic Versioning](https://semver.org/):

| Change | Version bump |
|--------|-------------|
| Bug fix | `x.y.Z` patch |
| New feature, backwards-compatible | `x.Y.0` minor |
| Breaking change | `X.0.0` major |

The release tag (`v1.2.3`) and Docker image tag (`1.2.3`) are derived from this value automatically by CI. Never create tags manually.

---

## PR rules

- Target branch is always `develop` (or `master` for releases and hotfixes).
- CI must be green before merging.
- Squash or merge commit — no preference enforced, but keep the history readable.
- Delete the branch after merging.
