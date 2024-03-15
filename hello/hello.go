package hello

const (
	french  = "French"
	spanish = "Spanish"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) string { /* keeping our domain from the outside world (side-effects) */
	if name == "" {
		name = "world"
	}

	return greetigsPrefix(language) + name + "!"
}

func greetigsPrefix(language string) (prefix string) /* A variable with the name `prefix` will be created in the function */ {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}
