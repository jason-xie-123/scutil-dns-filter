## Changelog for v0.2.0

This release is a repository relaunch — the CLI's flags, output, and filtering behavior are unchanged. It focuses on making the project properly usable, testable, and maintainable by others:

- **Licensing**: added an MIT `LICENSE` (previously the repo had none).
- **Testability**: the DNS filter/dedup logic was extracted into `filterDNSIPs()` (previously inlined in `main()`, untestable without a real macOS DNS configuration) and now has unit test coverage.
- **CI/CD migrated from Azure DevOps to GitHub Actions**: runs entirely on `ubuntu-latest` (no platform-specific build constraints here). Releases are now triggered by pushing a `vX.Y.Z` tag instead of every push to `main`; the published binaries are still `darwin/amd64` and `darwin/arm64` only, no new platform targets were added.
- **Project layout**: moved to the standard `cmd/scutil-dns-filter/` + `internal/version/` Go layout.
- **Docs**: expanded `README.md` with install/usage/dev instructions, added `AGENTS.md` for AI coding assistants.
- Fixed a `.gitignore` bug where a bare `scutil-dns-filter` pattern unintentionally excluded the new `cmd/scutil-dns-filter/` directory from version control.
