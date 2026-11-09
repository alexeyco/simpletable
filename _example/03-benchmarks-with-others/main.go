// Package main compares simpletable with other table libraries.
package main

import (
	"fmt"

	"github.com/alexeyco/simpletable/_example/03-benchmarks-with-others/bench"
)

func main() {
	fmt.Println("github.com/alexeyco/simpletable:")
	fmt.Println(bench.Simpletable())

	fmt.Println("github.com/apcera/termtables:")
	fmt.Println(bench.Termtables())

	fmt.Println("github.com/gosuri/uitable:")
	fmt.Println(bench.UITable())
}
