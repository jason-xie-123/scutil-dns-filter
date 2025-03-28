package main

import (
	"fmt"
	packageVersion "scutil-dns-filter/version"
)

func main() {
	fmt.Printf("%v\n", packageVersion.Version)
}
