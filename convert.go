package main

import "strings"

func StringToArt(input string) string {
	if input == "" {
		return ""
	}

	artMap := map[rune][5]string{
		'1': {
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
		},
		'2': {
			" ___ ",
			"    |",
			" ___|",
			"|    ",
			"|___ ",
		},
	}

	var result strings.Builder
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		var rowBuilders [5]strings.Builder
		for _, char := range line {
			art, exists := artMap[char]
			if !exists {
				return ""
			}
			for r := 0; r < 5; r++ {
				rowBuilders[r].WriteString(art[r])
			}
		}

		for r := 0; r < 5; r++ {
			result.WriteString(rowBuilders[r].String() + "\n")
		}
	}

	return result.String()
}
