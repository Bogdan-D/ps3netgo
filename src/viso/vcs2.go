package viso

import (
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// UCS2 text codec.
type UCS2 []byte

// Encode to UCS2.
func (s UCS2) Encode() []byte {
	e := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	es, _, err := transform.Bytes(e.NewEncoder(), s)
	if err != nil {
		return s
	}
	return es
}

// Decode from UCS2.
func (s UCS2) Decode() []byte {
	e := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	es, _, err := transform.Bytes(e.NewDecoder(), s)
	if err != nil {
		return s
	}
	return es
}
