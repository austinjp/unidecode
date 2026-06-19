package unidecode_test

import (
	"fmt"

	"github.com/austinjp/unidecode"
)

func ExampleUnidecode() {
	s := "北京kožušček"
	fmt.Println(unidecode.Unidecode(s))
	// Output: Bei Jing kozuscek
}
