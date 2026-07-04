## Changelog for v0.2.1

Bug-fix release, no CLI behavior changes.

- **Fixed `go install` not working**: `go.mod` declared its module path as the bare name `scutil-dns-filter`, which conflicted with the `go install github.com/jason-xie-123/scutil-dns-filter/cmd/scutil-dns-filter@latest` command documented in the README. The module path is now `github.com/jason-xie-123/scutil-dns-filter`, matching the actual import path.
