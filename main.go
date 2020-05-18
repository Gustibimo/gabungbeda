package main

import (
"bufio"
"flag"
"fmt"
"io"
"os"
)

func main() {

	flag.Parse()
	fa := flag.Arg(0)

	lines := make(map[string]bool)

	var f io.WriteCloser

	if fa != "" {
		// read file covert into a map if it exists
		r, err := os.Open(fa)
		if err == nil {
			sc := bufio.NewScanner(r)

			for sc.Scan() {
				lines[sc.Text()] = true
			}
			r.Close()
		}

		// re-open the file for appending new stuff
		f, err = os.OpenFile(fa, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open file for writing: %s\n", err)
			return
		}
		defer f.Close()
	}

	// read, append and output them if they are new
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		line := sc.Text()
		if lines[line] {
			continue
		}

		// add the line to the map so we don't get any duplicates from stdin
		lines[line] = true


		if fa != "" {
			fmt.Fprintf(f, "%s\n", line)
		}
	}
}
