opts
====

POSIX- and GNU-style options for Go, forked from https://code.google.com/p/opts-go

```go
package main

import (
	"github.com/webconnex/opts"
	"fmt"
)

func printVersion() {
	fmt.Printf("VERSION\n")
}

var showVersion = opts.Flag("", "--version", "Description")
var output = opts.Half("-o", "--output", "write output to file", "", "output")
var format = opts.Single("-f", "--format", "the output format to use", "csv")
var include = opts.Multi("-I", "--include", "files to include", "file")

func main() {
	opts.Xname = "main"
	opts.AddHelp("help screen")
	opts.Parse()
	if *showVersion {
		printVersion()
	}
	if *output != "" {
		fmt.Printf("Writing output to %s\n", *output)
	}
	fmt.Printf("Using format %s\n", *format)
	for _, file := range include {
		fmt.Printf("Including %s\n", file)
	}
}
```
