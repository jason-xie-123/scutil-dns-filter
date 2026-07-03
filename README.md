# scutil-dns-filter

[![CI](https://github.com/jason-xie-123/scutil-dns-filter/actions/workflows/ci.yml/badge.svg)](https://github.com/jason-xie-123/scutil-dns-filter/actions/workflows/ci.yml)
[![Release](https://img.shields.io/github/v/release/jason-xie-123/scutil-dns-filter)](https://github.com/jason-xie-123/scutil-dns-filter/releases/latest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

Filter DNS resolvers by network interface using the `scutil --dns` command on macOS.

You can verify the results by running:

```sh
/usr/sbin/scutil --dns
```

## Why Golang?

Golang allows us to compile directly into platform-specific binary files that can run independently without requiring additional environments.

## Install

Download the binary for your Mac's architecture from the [latest release](https://github.com/jason-xie-123/scutil-dns-filter/releases/latest) (darwin-amd64 or darwin-arm64).

Or build from source:

```sh
go install github.com/jason-xie-123/scutil-dns-filter/cmd/scutil-dns-filter@latest
```

## How to Use

```
scutil-dns-filter -h
NAME:
   scutil-dns-filter - CLI tool to format scutil-dns-filter scripts

USAGE:
   scutil-dns-filter [global options] command [command options]

GLOBAL OPTIONS:
   --InterfaceName value  interface name to use for DNS filtering, such as 'en0'
   --help, -h              show help
   --version, -v           print the version
```

Example:

```sh
scutil-dns-filter --InterfaceName en0
```

## Development

```sh
go build ./...
go test ./...
gofmt -l .
golangci-lint run ./...
```

The core filtering/dedup logic lives in `filterDNSIPs` (`cmd/scutil-dns-filter/main.go`) and is unit tested without needing a real macOS DNS configuration — see `main_test.go`.

Releases are cut by pushing a `vX.Y.Z` tag — see `.github/workflows/release.yml`. Release notes live in `release_notes.md` and are drafted locally before tagging (see `AGENTS.md`).

## License

[MIT](./LICENSE)
