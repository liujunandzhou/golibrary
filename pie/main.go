package main

import (
	"fmt"
	"strings"

	piev2 "github.com/elliotchance/pie/v2"
)

func main() {

	slices := []string{"bob", "sally", "john", "jane"}

	name := piev2.Of(slices).FilterNot(func(name string) bool {
		return strings.HasPrefix(name, "J")
	}).Map(strings.ToUpper).Last()

	fmt.Printf("name=%v\n", name)
}
