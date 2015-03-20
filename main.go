package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"
)

var (
	verbose    = flag.Bool("v", false, "Print url and query string in addition to values matching params")
	nullString = flag.String("null", "NULL", "Value to print when a param is not present.")
)

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		usage()
		os.Exit(1)
	}

	param := flag.Args()[0]

	// TODO: it may also be nice to read in from a file if it is specified
	queryParamParser(os.Stdin, os.Stdout, getParamsArray(param))
}

func usage() {
	fmt.Printf("Usage: qp [options] <comma delimited list of param keys>\n")
	flag.PrintDefaults()

	fmt.Print("\nExample: qp -v -null 'derp' 'format, sort'\n")
}

func getParamsArray(p string) []string {
	p = strings.Replace(p, " ", "", -1)

	return strings.Split(p, ",")
}

func queryParamParser(r io.Reader, w io.Writer, params []string) {
	scanner := bufio.NewScanner(r)

	// Allocate an array for params values once
	vals := make([]string, len(params), len(params))

	for scanner.Scan() {
		u, err := url.Parse(scanner.Text())

		// Ignore all errors due to badly formatted URLs
		if err != nil {
			continue
		}

		values, err := url.ParseQuery(u.RawQuery)

		if err != nil {
			continue
		}

		// Fill values array
		for i, p := range params {
			v := values.Get(p)

			if v != "" {
				vals[i] = v
			}
		}

		// Check for at least one non-blank string
		if printable(vals) {
			if *verbose {
				fmt.Fprintf(w, "%v\t%v\t", u.Path, u.RawQuery)
			}

			printQueryValues(w, vals)
		}

		// Zero out memory
		for i := range vals {
			vals[i] = ""
		}
	}
}

// printQueryValues prints the following if given a vals array of ["xml", "3", "", "derp"]:
// xml\t3\tNULL\tderp\n
func printQueryValues(w io.Writer, vals []string) {
	numVals := len(vals)

	for i, val := range vals {

		var printStr string
		if val == "" {
			printStr = *nullString
		} else {
			printStr = val
		}

		fmt.Fprint(w, printStr)

		// Initializes to blank string ("")
		var sep string
		if i != numVals {
			sep = "\t"
		}

		fmt.Fprint(w, sep)
	}

	fmt.Fprintln(w)
}

// printable checks to see if there is at least one non blank string in the
// values array
func printable(vals []string) bool {
	for _, v := range vals {
		if v != "" {
			return true
		}
	}

	return false
}
