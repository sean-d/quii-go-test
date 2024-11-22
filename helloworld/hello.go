package main

const (
	EnglishPrefix = "Hello, "
	SpanishPrefix = "Hola, "
	FrenchPrefix  = "Bonjour, "
)

func Hello(language, s string) string {
	var prefix string

	prefix = greetingPrefix(language)

	if s == "" {
		s = "world"
	}
	return prefix + s
}

func greetingPrefix(language string) string {
	var prefix string

	switch language {
	case "spanish":
		prefix = SpanishPrefix
	case "french":
		prefix = FrenchPrefix
	default:
		prefix = EnglishPrefix
	}
	return prefix
}
func main() {
	//fmt.Println(Hello("spanish", "you"))
}
