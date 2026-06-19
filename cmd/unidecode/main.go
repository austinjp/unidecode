package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/austinjp/unidecode"
)

func main() {
	version := flag.Bool("V", false, "Output version info")
	input := flag.String("text", "", "Text to transliterate")

	flag.Parse()
	if *version {
		fmt.Printf("unidecode %s\n", unidecode.Version())
		os.Exit(0)
	}

	if *input == "" {
		fmt.Println("Usage: unidecode STRING")
		os.Exit(1)
	}

	fmt.Println(unidecode.Unidecode(*input))
}
