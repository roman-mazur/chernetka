package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	toolName := filepath.Base(os.Args[0])
	var outName string
	if len(os.Args) > 1 {
		outName = os.Args[1]
	}
	dir, _ := filepath.Abs(".")
	pkgName := filepath.Base(dir)

	out := os.Stdout
	if outName != "" {
		f, err := os.Create(outName)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		out = f
	}

	fmt.Fprintf(out, "// Code generated with %s %s. DO NOT EDIT.\n", toolName, outName)
	fmt.Fprintln(out, "package", pkgName)
	fmt.Fprintln(out)
	fmt.Fprintln(out, "var numbersLookup = [256]string{")

	for i := range 256 {
		fmt.Fprintf(out, "\t%q,\n", strconv.Itoa(i))
	}

	fmt.Fprintln(out, "}")
}
