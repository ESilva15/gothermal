package encodings

import (
	"fmt"
	"unicode/utf8"

	iconv "github.com/djimenez/iconv-go"
)

func PrintTable(table string) {
	defChar := " "

	// Open a new iconv converter for CP037 to UTF-8
	conv, err := iconv.NewConverter(table, "UTF-8")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer conv.Close()

	// Iterate over all byte values (0-255) and convert them to UTF-8
	counter := 0
	for i := 128; i < 256; i++ {
		// Convert byte value to UTF-8 string

		in := []byte{byte(i)}
		out := make([]byte, 4)
		_, _, err := conv.Convert(in, out)

		if err != nil {
			out = []byte(defChar)
		}

		// We have to transform out into a rune
		char, _ := utf8.DecodeRune(out)

		fmt.Printf("[%c] ", char)
		counter++
		if counter == 16 {
			counter = 0
			fmt.Printf("\n")
		}
	}
}

func ShowAll() {
	tables := buildTables()
	fmt.Println(tables)
}
