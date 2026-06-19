package unidecode

import (
	"strings"

	"github.com/austinjp/unidecode/table"
)

// Version return version
func Version() string {
	return "0.2.1"
}

// Unidecode implements transliterate Unicode text into plain 7-bit ASCII.
// e.g. Unidecode("kožušček") => "kozuscek"
func Unidecode(s string) string {
	return unidecode(s)
}

func unidecode(s string) string {
	var ret strings.Builder
	ret.Grow(len(s)) // most chars map 1:1 or shrink, avoids reallocation
	for _, c := range s {

		// cast to uint32 to ensure compiler treats shifts as unsigned
		r := uint32(c)

		if r < 0x80 { // unicode.MaxASCII + 1, includes byte 127 which is DEL
			ret.WriteByte(byte(r)) // cheaper than WriteRune for known ASCII
			continue
		}
		if r > 0xeffff {
			continue
		}

		section := r >> 8    // Chop off the last two hex digits
		position := r & 0xff // Last two hex digits (use mask instead of mod, same result, marginally cheaper)

		if tb := table.Tables[section]; tb[position] != "" {
			ret.WriteString(tb[position])
		}
	}
	return ret.String()
}
