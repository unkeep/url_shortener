package domain

import (
	"fmt"
	"math"
	"strings"
)

type urlIDEncoder interface {
	Encode(id uint64) (string, error)
	Decode(str string) (uint64, error)
}

const (
	base62Chars      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	base62CharsCount = 62
)

type base62URLIDEncoder struct{}

// Encode - encoded ID to a base62
func (b base62URLIDEncoder) Encode(id uint64) (string, error) {
	if id == 0 {
		return "", fmt.Errorf("invalid ID")
	}

	var encodedBuilder strings.Builder

	for ; id > 0; id = id / base62CharsCount {
		encodedBuilder.WriteByte(base62Chars[(id % base62CharsCount)])
	}

	return encodedBuilder.String(), nil
}

// Decode - decodes string to an ID
func (b base62URLIDEncoder) Decode(str string) (uint64, error) {
	var id uint64

	for i, symbol := range str {
		pos := strings.IndexRune(base62Chars, symbol)

		if pos == -1 {
			return 0, fmt.Errorf("invalid base62 character: %s", string(symbol))
		}
		id += uint64(pos) * uint64(math.Pow(float64(base62CharsCount), float64(i)))
	}

	return id, nil
}
