package main

import (
	"context"
	"errors"
	"fmt"
	packageVersion "github.com/jason-xie-123/scutil-dns-filter/internal/version"
	"os"

	"github.com/johnstarich/go/dns/scutil"
	"github.com/samber/lo"
	"github.com/urfave/cli/v2"
)

// filterDNSIPs returns the deduplicated set of DNS nameserver IPs for the
// resolvers matching interfaceName. Order is not guaranteed.
func filterDNSIPs(config scutil.Config, interfaceName string) []string {
	var dnsIPs = []string{}
	for _, resolver := range config.Resolvers {
		if resolver.InterfaceName != interfaceName {
			continue
		}

		dnsIPs = append(dnsIPs, resolver.Nameservers...)
	}

	dnsIPs = lo.MapToSlice(lo.GroupBy(dnsIPs, func(ipTmp string) string {
		return ipTmp
	}), func(ipTmp string, ipsTmp []string) string {
		return ipsTmp[0]
	})

	return dnsIPs
}

func main() {
	app := &cli.App{
		Name:    "scutil-dns-filter",
		Usage:   "CLI tool to format scutil-dns-filter scripts",
		Version: packageVersion.Version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "InterfaceName",
				Usage: "interface name to use for DNS filtering, such as 'en0'",
			},
		},
		Action: func(c *cli.Context) error {
			interfaceName := c.String("InterfaceName")
			if len(interfaceName) == 0 {
				return errors.New("interfaceName is required")
			}

			config, _ := scutil.ReadMacOSDNS(context.Background())

			for _, ip := range filterDNSIPs(config, interfaceName) {
				fmt.Printf("%s\n", ip)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
