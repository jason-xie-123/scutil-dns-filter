# scutil-dns-filter
Filter DNS by network interface using the scutil --dns command on Darwin.

You can verify the results by running:

```
/usr/sbin/scutil --dns
```

# Why Golang?
Golang allows us to compile directly into platform-specific binary files that can run independently without requiring additional environments.

# How to Use

```
./scutil-dns-filter -h                 
NAME:
   scutil-dns-filter - CLI tool to format scutil-dns-filter scripts

USAGE:
   scutil-dns-filter [global options] command [command options]

VERSION:
   0.1.2

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --InterfaceName value  interface name to use for DNS filtering, such as 'en0'
   --help, -h             show help
   --version, -v          print the version
```