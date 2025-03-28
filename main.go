package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	packageVersion "scutil-dns-filter/version"

	"github.com/johnstarich/go/dns/scutil"
	"github.com/samber/lo"
	"github.com/urfave/cli/v2"
)

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

			for _, ip := range dnsIPs {
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
