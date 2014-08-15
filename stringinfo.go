package stringinfo

import (
	"unicode"
)

const (
	// 0xd800-0xdc00 encodes the high 10 bits of a pair.
	// 0xdc00-0xe000 encodes the low 10 bits of a pair..
	surr1 = 0xd800
	surr2 = 0xdc00
	surr3 = 0xe000
)

type StringInfo struct {
	runes  []rune
	length int
}

func New(s string) *StringInfo {
	return &StringInfo{runes: []rune(s), length: -1}
}

func (si *StringInfo) String() string {
	return string(si.runes)
}

// Gets the number of text elements in the current StringInfo object.
func (si *StringInfo) LengthInTextElements() int {
	if si.length < 0 {
		si.length = 0
		for idx := 0; idx < len(si.runes); si.length++ {
			idx += nextTextElementLength(si.runes, idx)
		}
	}
	return si.length
}

func nextTextElementLength(runes []rune, index int) int {
	r := runes[index]

	if unicode.Is(unicode.Cs, r) {
		// Check that it's a high surrogate followed by a low surrogate
		if surr1 <= r && r < surr2 {
			if index+1 < len(runes) && surr2 <= runes[index+1] && runes[index+1] < surr3 {
				// A valid surrogate pair
				return 2
			} else {
				// High surrogate on its own
				return 1
			}
		} else {
			// Low surrogate on its own
			return 1
		}
	} else {
		// Look for a base character, which may or may not be followed by a
		// series of combining characters
		if unicode.IsMark(r) {
			// Not a base character
			return 1
		}

		count := 1

		for index+count < len(runes) {
			if !unicode.IsMark(runes[index+count]) {
				// Finish the sequence
				break
			}
			count++
		}
		return count
	}
}
