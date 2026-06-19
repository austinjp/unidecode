# unidecode

ASCII transliterations of Unicode text, for Go. Forked from the original <github.com/mozillazg/go-unidecode>


## Installation

```
go get github.com/austinjp/unidecode
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/austinjp/unidecode"
)

func main() {
	s := "abc"
	fmt.Println(unidecode.Unidecode(s))
	// Output: abc

	s = "北京"
	fmt.Println(unidecode.Unidecode(s))
	// Output: Bei Jing

	s = "kožušček"
	fmt.Println(unidecode.Unidecode(s))
	// Output: kozuscek
}
```
