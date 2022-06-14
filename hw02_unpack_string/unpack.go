package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

type runeKind string

const (
	runeKindDigit     = runeKind("digit")
	runeKindLetter    = runeKind("letter")
	runeKindBackslash = runeKind("backslash")
)

var ErrInvalidString = errors.New("invalid string")

type symbol struct {
	value  rune
	repeat int
}

func Unpack(source string) (string, error) {
	if source == "" {
		return source, nil
	}

	symbols, err := toSymbolSlice(source)
	if err != nil {
		return "", err
	}

	var result strings.Builder
	for _, s := range symbols {
		switch {
		case s.repeat > 0:
			result.WriteString(strings.Repeat(string(s.value), s.repeat))
		case s.repeat < 0:
			result.WriteRune(s.value)
		default:
			// skip
		}
	}

	return result.String(), nil
}

func toSymbolSlice(source string) ([]*symbol, error) {
	result := make([]*symbol, 0, len(source))
	if len(source) == 0 {
		return result, nil
	}

	runes := toRunesSlice(source)
	index := 0
	for index < len(runes) {
		currentRune := runes[index]
		currentRuneKind := kindOfRune(currentRune)

		var effectiveRune rune
		switch currentRuneKind {
		case runeKindBackslash:
			if index == len(runes)-1 {
				return nil, ErrInvalidString
			}
			nextRune := runes[index+1]
			if kindOfRune(nextRune) == runeKindLetter {
				return nil, ErrInvalidString
			}
			effectiveRune = nextRune
			index++
		case runeKindDigit:
			return nil, ErrInvalidString
		case runeKindLetter:
			effectiveRune = currentRune
		}

		index++
		if index > len(runes)-1 {
			result = append(result, &symbol{
				value:  effectiveRune,
				repeat: -1,
			})
			break
		}

		repeatCount := -1
		repeatCountRune := runes[index]
		if unicode.IsDigit(repeatCountRune) {
			value, err := strconv.Atoi(string(repeatCountRune))
			if err != nil {
				return nil, err
			}
			repeatCount = value
			index++
		}

		result = append(result, &symbol{
			value:  effectiveRune,
			repeat: repeatCount,
		})
	}

	return result, nil
}

func toRunesSlice(source string) []rune {
	result := make([]rune, 0, len(source))
	for _, r := range source {
		result = append(result, r)
	}
	return result
}

func kindOfRune(source rune) runeKind {
	switch {
	case string(source) == "\\":
		return runeKindBackslash
	case unicode.IsDigit(source):
		return runeKindDigit
	default:
		return runeKindLetter
	}
}
