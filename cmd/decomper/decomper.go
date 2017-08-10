package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/stts-se/decomp"
)

// decomper is a command line utility for running the decomp.Decomp decompounder.

func main() {

	var wds []string

	// FLags: -help:
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "decomper <DEOMCP FILE> <words...>\n")
		os.Exit(0)
	}

	fn := os.Args[1]
	decomp, err := decomp.NewDecompounderFromFile(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "decomper: failed to load file '%s' : %v\n", fn, err)
		os.Exit(1)
	}
	_ = decomp

	// words as command line args
	if len(os.Args) > 2 {
		wds = os.Args[2:]
	} else {

		// words from stdin
		s := bufio.NewScanner(os.Stdin)
		r, err := regexp.Compile(`\s+`)
		if err != nil {
			fmt.Fprintf(os.Stderr, "decomper: split regexp failure : %v\n", err)
			os.Exit(1)
		}
		for s.Scan() {
			l := s.Text()
			lWds := r.Split(l, -1)
			wds = append(wds, lWds...)
		}
	}

	for _, w := range wds {
		for _, ds := range decomp.Decomp(w) {
			fmt.Printf("%s\t%v\n", w, strings.Join(ds, "+"))
		}
	}

}
