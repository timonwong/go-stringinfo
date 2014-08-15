package stringinfo

import (
	"unicode"
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
