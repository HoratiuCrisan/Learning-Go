package iteration

/* A function that adds a character 5 times to a string variable and returns it */

func Repeat(character rune, value int) (text string) {
	for i := 0; i < value; i++ {
		text += "a"
	}

	return text
}
