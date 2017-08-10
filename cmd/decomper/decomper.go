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
// Sample invocation: go run decomper.go ../../server/decomp_files/sv_nst.txt hundparkering
func main() {

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "decomper <DECOMP FILE> <words...>|<STDIN>\n")
		os.Exit(0)
	}

	fn := os.Args[1]
	decomp, err := decomp.NewDecompounderFromFile(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "decomper: failed to load file '%s' : %v\n", fn, err)
		os.Exit(1)
	}

	var wds []string
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

		ds := decomp.Decomp(w)
		if len(ds) == 0 {
			fmt.Printf("%s\t?\n", w)
		}
		for _, d := range ds {
			fmt.Printf("%s\t%v\n", w, strings.Join(d, "+"))
		}
	}

}
