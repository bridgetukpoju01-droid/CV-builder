package main

import "fmt"

func GenerateFont() map[rune][8]string {
	font := make(map[rune][8]string)

	for c := rune(32); c <= rune(126); c++ {
		var lines [8]string

		if c == ' ' {
			for i := 0; i < 8; i++ {
				lines[i] = "        "
			}
		} else {
			for i := 0; i < 8; i++ {
				lines[i] = fmt.Sprintf("%02d-%05d", i, c)
			}
		}

		font[c] = lines
	}

	return font
}
