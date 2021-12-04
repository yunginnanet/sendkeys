package sendkeys

import (
	"errors"
	"strings"
)

func (kb *KBWrap) handleRunes(split []string) (keys []int) {
	for _, c := range split {
		d, dok := num[c]
		a, aok := alpha[c]
		sym, symok := Symbol[c]
		ca, caok := alpha[strings.ToLower(c)]

		switch {
		case aok:
			keys = append(keys, a)
		case dok:
			keys = append(keys, d)
		case caok:
			keys = append(keys, 0-ca)
		case symok:
			keys = append(keys, sym)
		default:
			kb.errors = append(
				kb.errors,
				errors.New(ErrKeyMappingNotFound.Error()+c),
			)
		}
	}
	return
}

func (kb *KBWrap) strToKeys(s string) (keys []int) {
	if !strings.Contains(s, " ") {
		return kb.handleRunes(strings.Split(s, ""))
	}
	splitspace := strings.Split(s, " ")
	for _, section := range splitspace {
		split := strings.Split(section, "")
		keys = append(keys, kb.handleRunes(split)...)
		keys = append(keys, 57)
	}
	return
}

var num = map[string]int{
	"1": 2, "2": 3, "3": 4, "4": 5, "5": 6,
	"6": 7, "7": 8, "8": 9, "9": 10, "0": 11,
}

var alpha = map[string]int{
	"q": 16, "w": 17, "e": 18, "r": 19, "t": 20,
	"y": 21, "u": 22, "i": 23, "o": 24, "p": 25,
	"a": 30, "s": 31, "d": 32, "f": 33, "g": 34,
	"h": 35, "j": 36, "k": 37, "l": 38, "z": 44,
	"x": 45, "c": 46, "v": 47, "b": 48, "n": 49,
	"m": 50,
}

// Symbol is a map of translations for sending non-alphanumeric key events.
var Symbol = map[string]int{
	"-": 12, "_": -12, "=": 13, "+": -13, "[": 26,
	"{": -26, "]": 27, "}": -27, "'": 40, "\"": -40,
	"`": 41, "~": -41, "\\": 43, "|": -43, ",": 51,
	"<": -51, ".": 52, ">": -52, "/": 53, "?": -53,
	" ": 57, "!": -2, "@": -3, "#": -4, "$": -5,
	"%": -6, "^": -7, "&": -8, "*": -9, "(": -10,
	")": -11, ";": 39, ":": -39,
}
