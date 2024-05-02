package encodings

import (
	"unicode/utf8"
)

type Parser struct {
	Charmap map[int]PrinterChar
}

// charSize returns the number of bits on given char in UTF-8
func charSize(b byte) int {
	toRet := -1

	if b>>7 == 0b0 {
		toRet = 1
	} else if b>>5 == 0b110 {
		toRet = 2
	} else if b>>4 == 0b1110 {
		toRet = 3
	} else if b>>3 == 0b11110 {
		toRet = 4
	} else if b>>2 == 0b111110 {
		toRet = 5
	} else if b>>1 == 0b1111110 {
		toRet = 6
	}

	return toRet
}

// Initialize the charmap
func (p *Parser) Initialize() {
	p.Charmap = buildTables()
}

// process will replace unicode characters of the message by the correct
// character of the correct code page for the thermal printer to print
func (p *Parser) Process(m string) []byte {
	message := make([]byte, 0)
	for k := 0; k < len([]byte(m)); k++ {
		size := charSize(m[k])

		toAppend := []byte{'X'}
		if size == 1 {
			// Check if its an escape sequence
			if int(m[k]) == '\\' {
				if int(m[k+1]) == 'n' {
					toAppend = []byte{0x0A}
				} else if int(m[k+1]) == 't' {
					toAppend = []byte{0x09}
				}
				k += 1
			} else {
				toAppend = []byte{m[k]}
			}
		} else {
			char, _ := utf8.DecodeRune([]byte(m[k : k+(size)]))
			_, ok := p.Charmap[int(char)]
			if ok {
				toAppend = []byte{
					0x1B,
					0x74,
					byte(p.Charmap[int(char)].Table),
					byte(p.Charmap[int(char)].Hex),
				}
			}

			k += (size - 1)
		}

		message = append(message, toAppend...)
	}

	return message
}
