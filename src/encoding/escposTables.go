package encodings

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"

	iconv "github.com/djimenez/iconv-go"
)

var (
	CAPABILITES = "/home/esilva/Desktop/projetos/escpos-printer-db/dist/capabilities.json"
)

type PrinterEncodings struct {
	IconvName string   `json:"iconv"`
	Data      []string `json:"data"`
}

type PrinterProfile struct {
	CodePages map[string]string `json:"codePages"`
	Name      string            `json:"name"`
}

type Capabilites struct {
	Profiles  map[string]PrinterProfile   `json:"profiles"`
	Encodings map[string]PrinterEncodings `json:"encodings"`
}

type PrinterChar struct {
	Table int
	Hex   int
	Utf   string
}

func buildTables() map[int]PrinterChar {
	capabilities := loadCapabilities()

	data := make(map[int]PrinterChar)
	for n, enc := range capabilities.Profiles["HZ-8360"].CodePages {
		// Check if code page is present in the available encodings
		_, ok := capabilities.Encodings[enc]
		if ok && len(capabilities.Encodings[enc].IconvName) > 0 {
			// Open a new iconv converter for CP037 to UTF-8
			conv, err := iconv.NewConverter(enc, "UTF-8")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return nil
			}
			defer conv.Close()

			// Iterate over all byte values (128-255) and convert them to UTF-8
			for i := 128; i < 256; i++ {
				// Convert byte value to UTF-8 string
				in := []byte{byte(i)}
				out := make([]byte, 4)
				_, _, err := conv.Convert(in, out)

				// This usually happens when the table location is empty
				if err != nil {
					out = []byte(" ")
				}

				// Transform it into a rune
				char, _ := utf8.DecodeRune(out)
				tableNum, _ := strconv.Atoi(n)

				// Create the new char
				newChar := PrinterChar{
					Table: tableNum,
					Hex:   i,
					Utf:   string(out),
				}

				// If the char is already present we only attribute the current char
				// if the table num is lower
				_, ok := data[int(char)]
				if ok && data[int(char)].Table > tableNum {
					data[int(char)] = newChar
				}

				// If its not present we just add it
				if !ok {
					data[int(char)] = newChar
				}

			}
		}
	}

	return data
}

func loadCapabilities() Capabilites {
	file, err := os.ReadFile(CAPABILITES)
	if err != nil {
		fmt.Println("failed to open charmap: ", err)
		panic("")
	}

	var c Capabilites
	err = json.Unmarshal([]byte(file), &c)
	if err != nil {
		fmt.Println("failed to unmarshal: ", err)
		panic("")
	}

	return c
}
