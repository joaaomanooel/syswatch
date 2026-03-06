# CI/CD Release Pipeline

This document describes the continuous integration and continuous deployment pipeline for SysWatch.

## Overview

SysWatch uses GitHub Actions for automated building, testing, and releasing. The pipeline is designed to be reliable, fast, and produce consistent releases across multiple platforms.

## Workflows

### 1. Test Workflow (`test.yml`)

**Trigger**: Push to `main` branch or Pull Request

**Purpose**: Run tests and linting on every code change

**Steps**:
1. Checkout code
2. Set up Go 1.24
3. Run unit tests
4. Run linter (golangci-lint)

**Status**: Must pass before merging to main

---

### 2. Release Workflow (`release.yml`)

**Trigger**: Push of version tag (e.g., `v0.0.1`)

**Purpose**: Build binaries and create GitHub Release

**Steps**:
1. Checkout code
2. Set up Go 1.24
3. Run GoReleaser
   - Builds for all platforms:
     - macOS (amd64, arm64)
     - Linux (amd64, arm64)
     - Windows (amd64)
   - Generates .deb package (Debian/Ubuntu)
   - Generates .rpm package (Fedora/RHEL)
   - Creates checksums
   - Creates GitHub Release

**Artifacts**:
```
syswatch/
├── syswatch-darwin-amd64.zip
├── syswatch-darwin-arm64.zip
├── syswatch-linux-amd64.zip
├── syswatch-linux-arm64.zip
├── syswatch-windows-amd64.zip
├── syswatch_0.0.1_amd64.deb
├── syswatch-0.0.1-1.x86_64.rpm
├── checksums.txt
└── Source code (tarball)
```

---

### 3. Snap Store Workflow (`snap.yml`)

**Trigger**: When a GitHub Release is published

**Purpose**: Build and publish to Snap Store

**Steps**:
1. Checkout code
2. Build snap package using `snapcore/action-build`
3. Publish to Snap Store using `snapcore/action-publish`

**Requirements**:
- Snapcraft account
- `SNAPCRAFT_LOGIN_JSON` secret in GitHub repository

**Installation**:
```bash
sudo snap install syswatch --classic
```

---

### 4. Homebrew Workflow (`homebrew.yml`)

**Trigger**: When a GitHub Release is published

**Purpose**: Update Homebrew tap with new version

**Steps**:
1. Checkout code
2. Get latest tag
3. Create Pull Request to homebrew-syswatch tap

**Requirements**:
- GitHub repository: `joaaomanooel/homebrew-syswatch`

**Installation**:
```bash
brew install joaaomanooel/homebrew-syswatch/syswatch
```

---

## Creating a Release

### Step 1: Prepare Your Code

```bash
# Make your changes
git add .
git commit -m "Your changes"
```

### Step 2: Push to Main

```bash
git push origin main
```

### Step 3: Wait for Tests

The `test.yml` workflow will run. Ensure all tests pass.

### Step 4: Create Version Tag

```bash
git tag v0.0.1
git push origin v0.0.1
```

### Step 5: Release Pipeline Runs

1. `release.yml` triggers → Creates GitHub Release with binaries
2. `snap.yml` triggers → Publishes to Snap Store (if configured)
3. `homebrew.yml` triggers → Updates Homebrew tap (if configured)

---

## Secrets Configuration

### Required Secrets

| Secret | Where to Get | Purpose |
|--------|--------------|---------|
| `SNAPCRAFT_LOGIN_JSON` | Snapcraft account | Publish to Snap Store |

### Optional Secrets

| Secret | Purpose |
|--------|---------|
| `GITHUB_TOKEN` | Auto-generated, used for releases |

---

## Manual Release (Without CI)

If you need to create a release manually:

```bash
# Install GoReleaser
brew install goreleaser

# Create a tag
git tag v0.0.1

# Run GoReleaser
goreleaser release --clean
```

---

## Troubleshooting

### Release Failed

1. Check workflow logs in GitHub Actions
2. Verify GoReleaser configuration (`.goreleaser.yml`)
3. Ensure tests pass locally: `make test`

### Snap Store Publish Failed

1. Verify `SNAPCRAFT_LOGIN_JSON` is correct
2. Check Snapcraft account permissions
3. Ensure snapcraft.yml is valid

### Homebrew Update Failed

1. Verify `homebrew-syswatch` repository exists
2. Check that the version tag was created correctly

---

## Adding New Distribution Methods

### APT Repository (via GitHub Pages)

GoReleaser can generate APT repositories. Contact the maintainer for setup.

### Other Package Managers

1. Add configuration to `.goreleaser.yml`
2. Create corresponding workflow file
3. Update documentation

---

## Versioning Strategy

SysWatch follows [Semantic Versioning](https://semver.org/):

- **Major**: Breaking changes
- **Minor**: New features (backward compatible)
- **Patch**: Bug fixes

Example: `v1.2.3`
- 1 = Major version
- 2 = Minor version
- 3 = Patch version

---

## Support

For issues or questions:
- GitHub Issues: https://github.com/joaaomanooel/syswatch/issues
