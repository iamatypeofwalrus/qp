package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

var (
	param   = flag.String("param", "", "A comma seperated list of params to get values for")
	verbose = flag.Bool("verbose", false, "Print url and query string in addition to values matching params")
)

func main() {
	flag.Parse()

	if *param == "" {
		fmt.Println("You must provide a param with the -param option.")
		return
	}

	params := strings.Split(*param, ",")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		u, err := url.Parse(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		values, err := url.ParseQuery(u.RawQuery)

		if err != nil {
			log.Fatal(err)
		}

		// Create a slice of strings that is a fixed size since we will always be
		// looking for the same number of parameters.
		vals := make([]string, len(params), len(params))
		for i, p := range params {
			v := values.Get(p)

			if v != "" {
				vals[i] = v
			}
		}

		if printable(vals) {
			if *verbose {
				fmt.Printf("%v\t%v\t", u.Path, u.RawQuery)
			}

			printQueryValues(vals)
		}
	}
}

// printQueryValues prints the following if given a vals array of [xml, 3, derp]:
// xml\t3\tderp\n
func printQueryValues(vals []string) {
	numVals := len(vals)

	for i, val := range vals {

		if val == "" {
			fmt.Print(".")
		} else {
			fmt.Print(val)
		}

		var sep string
		if i != numVals {
			sep = "\t"
		}

		fmt.Print(sep)
	}

	fmt.Println()
}

func printable(vals []string) bool {
	for _, v := range vals {
		if v != "" {
			return true
		}
	}

	return false
}
