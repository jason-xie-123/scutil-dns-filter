# AGENTS.md

Guidance for AI coding assistants (Claude Code, Codex, etc.) working in this repository.

## This tool functionally depends on macOS, but the code itself doesn't

`main()` calls `scutil.ReadMacOSDNS()`, which shells out to the macOS-only `scutil --dns` command — the tool is only useful when run on macOS. However, the source code has no `//go:build` constraint and compiles/tests fine on any host. Releases are intentionally limited to `darwin/amd64` and `darwin/arm64` — do not add Windows/Linux targets without a specific reason; it wouldn't be useful there.

## Project layout

- `cmd/scutil-dns-filter/main.go` — CLI entrypoint; thin wrapper around `filterDNSIPs`
- `cmd/scutil-dns-filter/main_test.go` — unit tests for `filterDNSIPs`, using hand-built `scutil.Config` fixtures (no real `scutil` invocation, no macOS required)
- `internal/version/version.go` — single `Version` constant, bumped manually before each release

## Build, test, lint

```sh
go build ./...
go test ./...
gofmt -l .              # must produce no output
golangci-lint run ./...  # must report 0 issues
```

Run all four before considering any change complete. Unlike some of the other CLI tools in this account, this one runs fine on any CI runner — no cross-platform test execution issues.

## Commit messages

Write commit messages in English. Keep them short and describe the actual change — avoid placeholder messages like `init` or `update`.

## Release process

Releases are tag-triggered, not push-triggered:

1. Draft `release_notes.md` locally by reading the diff since the last tag (`git diff <last-tag>..HEAD`) — an AI assistant can draft this, but a human must review it before tagging.
2. Bump the `Version` constant in `internal/version/version.go` to match the new tag.
3. `git tag vX.Y.Z && git push origin vX.Y.Z` — this triggers `.github/workflows/release.yml`, which builds `darwin/amd64` and `darwin/arm64` binaries and creates the GitHub Release using the committed `release_notes.md`.

Do not call any LLM API from within CI to generate release notes — that step happens locally, before tagging, to avoid paying per-run API costs in the pipeline.
