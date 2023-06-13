package main

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

func main() {

	fmt.Printf("That file is %s.\n", humanize.Bytes(82854982))
	fmt.Printf("You owe $%s.\n", humanize.Comma(6582491))
	fmt.Printf("You're my %s best friend.", humanize.Ordinal(193))

	size, _ := humanize.ParseBytes("42 MB") // size is 42000000, nil
	fmt.Printf("42 MB is %d\n", size)
}
