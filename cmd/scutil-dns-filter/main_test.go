package main

import (
	"sort"
	"testing"

	"github.com/johnstarich/go/dns/scutil"
)

func sortedStrings(s []string) []string {
	out := append([]string{}, s...)
	sort.Strings(out)
	return out
}

func TestFilterDNSIPsMatchingInterface(t *testing.T) {
	config := scutil.Config{
		Resolvers: []scutil.Resolver{
			{InterfaceName: "en0", Nameservers: []string{"1.1.1.1", "8.8.8.8"}},
			{InterfaceName: "en5", Nameservers: []string{"9.9.9.9"}},
		},
	}

	got := sortedStrings(filterDNSIPs(config, "en0"))
	want := []string{"1.1.1.1", "8.8.8.8"}

	if len(got) != len(want) {
		t.Fatalf("filterDNSIPs() = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("filterDNSIPs() = %v, want %v", got, want)
		}
	}
}

func TestFilterDNSIPsNoMatchingInterface(t *testing.T) {
	config := scutil.Config{
		Resolvers: []scutil.Resolver{
			{InterfaceName: "en0", Nameservers: []string{"1.1.1.1"}},
		},
	}

	got := filterDNSIPs(config, "en99")
	if len(got) != 0 {
		t.Fatalf("filterDNSIPs() = %v, want empty", got)
	}
}

func TestFilterDNSIPsDedupesAcrossResolvers(t *testing.T) {
	config := scutil.Config{
		Resolvers: []scutil.Resolver{
			{InterfaceName: "en0", Nameservers: []string{"1.1.1.1", "8.8.8.8"}},
			{InterfaceName: "en0", Nameservers: []string{"8.8.8.8", "9.9.9.9"}},
		},
	}

	got := sortedStrings(filterDNSIPs(config, "en0"))
	want := []string{"1.1.1.1", "8.8.8.8", "9.9.9.9"}

	if len(got) != len(want) {
		t.Fatalf("filterDNSIPs() = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("filterDNSIPs() = %v, want %v", got, want)
		}
	}
}
