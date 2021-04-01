package main

import (
	"flag"
)

func main() {
	seed := flag.String("seed", "", "Seed from seed.properties. 16 chars")
	input := flag.String("input", "datasources.yml", "Input yaml data sources")
	flag.Parse()
	if *seed == "" || *input == "" {
		flag.Usage()
	}
}
