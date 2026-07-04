## Changelog for v0.2.2

Housekeeping release, no CLI behavior changes.

- Added unit test coverage for two edge cases in `filterDNSIPs`: an empty `Resolvers` slice, and a matching resolver with nil/empty `Nameservers`. Both were already handled correctly by the existing implementation; this just locks that behavior in.
