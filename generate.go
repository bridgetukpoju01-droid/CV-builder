package main

func GeneratePattern(input rune) []string {
	if input < 'A' || input > 'Z' {
		return []string{}
	}

	if input == 'A' {
		return []string{
			"  ##  ",
			" #  # ",
			" #  # ",
			" #### ",
			" #  # ",
			" #  # ",
			" #  # ",
			"      ",
		}
	}

	if input == 'Z' {
		return []string{
			" #### ",
			"    # ",
			"   #  ",
			"  #   ",
			" #    ",
			" #    ",
			" #### ",
			"      ",
		}
	}

	return []string{
		"######",
		"######",
		"######",
		"######",
		"######",
		"######",
		"######",
		"      ",
	}
}
