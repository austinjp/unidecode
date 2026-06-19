package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/austinjp/unidecode"
	"github.com/mattn/go-isatty"
)

func main() {
	version := flag.Bool("V", false, "Output version info")
	flag.Parse()
	if *version {
		v := unidecode.Version()
		fmt.Printf("unidecode %s\n", v)
		os.Exit(0)
	}

	textSlice := flag.Args()
	stdin := []byte{}
	if !isatty.IsTerminal(os.Stdin.Fd()) {
		stdin, _ = io.ReadAll(os.Stdin)
	}
	if len(stdin) > 0 {
		textSlice = append(textSlice, string(stdin))
	}

	if len(textSlice) == 0 {
		fmt.Println("Usage: unidecode STRING")
		os.Exit(1)
	}

	s := strings.Join(textSlice, " ")
	ret := unidecode.Unidecode(s)
	fmt.Println(ret)
}
