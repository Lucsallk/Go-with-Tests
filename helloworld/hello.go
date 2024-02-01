package main

import "fmt"

const (
	french = "French"
	spanish = "Spanish"
	
	englishPrefix = "Hello, "
	frenchPrefix = "Bonjour, "
	spanishPrefix = "Hola, "
)


// Funcoes privadas comecam com letras minúsculas
func definePrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchPrefix
	case "Spanish":
		prefix = spanishPrefix
	default :
		prefix = englishPrefix
	}
	return
}

func Hello(value string, language string) string {
	if value == "" {
		value = "World"
	}
	
	return DefinePrefix(language) + value
}

func main() {
	fmt.Println(Hello("", "French"))
}
