package fun

import (
	"strconv"
)

type JsonASCII string

func (esc JsonASCII) MarshalJSON() ([]byte, error) {
	return []byte(strconv.QuoteToASCII(string(esc))), nil
}
